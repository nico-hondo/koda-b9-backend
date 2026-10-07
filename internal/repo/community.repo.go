package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nico-hondo/internal/dto"
)

type CommunityRepo struct {
	db *pgxpool.Pool
}

func NewCommunityRepo(db *pgxpool.Pool) *CommunityRepo {
	return &CommunityRepo{
		db: db,
	}
}

func (cr *CommunityRepo) GetCommunityRepo(ctx context.Context, filter dto.CommunityFilterParam) ([]dto.PopularCommunityResponse, error) {
	sql := `
		SELECT 
			c.id, 
			c.name, 
			c.slug, 
			c.description, 
			c.category, 
			c.image_url, 
			c.location, 
			c.is_active, 
			c.created_at,
			COALESCE(cm.total_members, 0) AS total_members,
			COALESCE(ue.total_upcoming_events, 0) AS total_upcoming_events
		FROM communities c
		-- Subquery 1: Total Anggota Komunitas
		LEFT JOIN (
			SELECT community_id, COUNT(user_id) AS total_members 
			FROM community_members 
			GROUP BY community_id
		) cm ON c.id = cm.community_id
		-- Subquery 2: Total Event yang Akan Datang (Upcoming)
		LEFT JOIN (
			SELECT community_id, COUNT(id) AS total_upcoming_events 
			FROM events 
			WHERE start_time >= NOW() 
			GROUP BY community_id
		) ue ON c.id = ue.community_id
	`

	var args []interface{}
	argIdx := 1

	//Filter search keyword(nama atau deskripsi)
	if filter.Search != "" {
		sql += fmt.Sprintf(" AND (c.name ILIKE $%d OR c.description ILIKE $%d)", argIdx, argIdx)
		args = append(args, "%"+filter.Search+"%")
		argIdx++
	}

	if filter.Category != "" && strings.ToLower(filter.Category) != "all" {
		sql += fmt.Sprintf(" AND c.category ILIKE $%d", argIdx)
		args = append(args, filter.Category)
		argIdx++
	}

	// if filter.IsJoined != "" && strings.ToLower(filter.IsJoined) != "all" {
	// 	sql += fmt.Sprintf(" AND cm.user_id = $%d", argIdx)
	// 	args = append(args, filter.Category)
	// 	argIdx++
	// }

	sql += " ORDER BY c.id ASC"
	communities, err := cr.db.Query(ctx, sql, args...)

	if err != nil {
		return nil, err
	}

	defer communities.Close()

	var data []dto.PopularCommunityResponse
	for communities.Next() {
		var community dto.PopularCommunityResponse
		if err := communities.Scan(
			&community.Id,
			&community.Name,
			&community.Slug,
			&community.Description,
			&community.Category,
			&community.ImageUrl,
			&community.Location,
			&community.Is_Active,
			&community.Created_At,
			&community.Total_Members,
			&community.TotalUpcomingEvents,
		); err != nil {
			return nil, err
		}

		data = append(data, community)
	}
	return data, nil
}

func (cr *CommunityRepo) GetCommunityByIDRepo(ctx context.Context, communityID int) (dto.CommunityDetailResponse, error) {
	query := `
		SELECT 
			c.id, c.name, c.slug, c.description, c.category, c.image_url, c.location, c.is_active, c.created_at,
			COALESCE(cm.total_members, 0) AS total_members,
			COALESCE(ue.total_upcoming_events, 0) AS total_upcoming_events
		FROM communities c
		LEFT JOIN (
			SELECT community_id, COUNT(user_id) AS total_members 
			FROM community_members 
			GROUP BY community_id
		) cm ON c.id = cm.community_id
		LEFT JOIN (
			SELECT community_id, COUNT(id) AS total_upcoming_events 
			FROM events 
			WHERE start_time >= NOW() 
			GROUP BY community_id
		) ue ON c.id = ue.community_id
		WHERE c.id = $1
	`

	var detail dto.CommunityDetailResponse
	err := cr.db.QueryRow(ctx, query, communityID).Scan(
		&detail.ID,
		&detail.Name,
		&detail.Slug,
		&detail.Description,
		&detail.Category,
		&detail.ImageUrl,
		&detail.Location,
		&detail.IsActive,
		&detail.CreatedAt,
		&detail.TotalMembers,
		&detail.TotalUpcomingEvents,
	)

	if err != nil {
		return dto.CommunityDetailResponse{}, err
	}

	return detail, nil
}

func (cr *CommunityRepo) GetCommunityMembersRepo(ctx context.Context, communityID int) ([]dto.MemberResponse, error) {
	query := `
		SELECT 
			u.id AS user_id, 
			u.name, 
			COALESCE(u.job, '') AS job, 
			COALESCE(u.avatar_url, '') AS avatar_url
		FROM community_members cm
		JOIN users u ON cm.user_id = u.id
		WHERE cm.community_id = $1
		ORDER BY cm.joined_at ASC
		LIMIT 10
	`

	rows, err := cr.db.Query(ctx, query, communityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []dto.MemberResponse
	for rows.Next() {
		var m dto.MemberResponse
		if err := rows.Scan(&m.UserID, &m.Name, &m.Job, &m.AvatarUrl); err != nil {
			return nil, err
		}
		members = append(members, m)
	}

	return members, nil
}

func (cr *CommunityRepo) GetPopularCommunity(ctx context.Context) ([]dto.PopularCommunityResponse, error) {
	sql := `
		SELECT
			c.id,
			c.name,
			c.slug,
			c.description,
			c.category,
			c.image_url,
			c.location,
			c.is_active,
			c.created_at,
			COALESCE(cm.total_members, 0) AS total_members,
			COALESCE(ue.total_upcoming_events, 0) AS total_upcoming_events
		FROM communities c
		LEFT JOIN (
			SELECT community_id, COUNT(user_id) AS total_members
			FROM community_members
			GROUP BY community_id
		) cm ON c.id = cm.community_id
		LEFT JOIN (
			SELECT community_id, COUNT(id) AS total_upcoming_events
			FROM events
			WHERE start_time >= NOW()
			GROUP BY community_id
		) ue ON c.id = ue.community_id
		ORDER BY total_members DESC
		LIMIT 8;
	`

	rows, err := cr.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []dto.PopularCommunityResponse
	for rows.Next() {
		var item dto.PopularCommunityResponse
		err := rows.Scan(
			&item.Id,
			&item.Name,
			&item.Slug,
			&item.Description,
			&item.Category,
			&item.ImageUrl,
			&item.Location,
			&item.Is_Active,
			&item.Created_At,
			&item.Total_Members,
			&item.TotalUpcomingEvents,
		)
		if err != nil {
			return nil, err
		}
		data = append(data, item)
	}

	if data == nil {
		data = []dto.PopularCommunityResponse{}
	}

	return data, nil
}

// func (cr *CommunityRepo) JoinCommunity(ctx context.Context, userId, community_id int) (dto.JoinCommunityResponse, error){
// 	checkJoin := `
// 		SELECT
// 			c.name,
// 			c.description,
// 			(cm.user_id IS NOT NULL) AS is_joined
// 		FROM communities c
// 		LEFT JOIN community_members cm
// 			ON c.id = cm.community_id AND cm.user_id = $2
// 		WHERE c.id = $1
// 	`

// 	var result dto.JoinCommunityResponse

// 	cr.db.QueryRow(ctx, checkJoin, community_id, userId).Scan(
// 		&result.Community
// 	)
// }
