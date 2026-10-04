package service

import (
	"context"
	"errors"

	"github.com/nico-hondo/internal/dto"
	"github.com/nico-hondo/internal/repo"
)

type CommunitiesService struct {
	cr *repo.CommunityRepo
	nr *repo.NotifRepo
}

func NewCommunityService(cr *repo.CommunityRepo, nr *repo.NotifRepo) *CommunitiesService {
	return &CommunitiesService{
		cr: cr,
		nr: nr,
	}
}

func (cs *CommunitiesService) GetCommunityService(ctx context.Context, filter dto.CommunityFilterParam, userId int) ([]dto.PopularCommunityResponse, error) {
	communities, err := cs.cr.GetCommunityRepo(ctx, filter, userId)
	if err != nil {
		return nil, err
	}

	return communities, nil
}

func (cs *CommunitiesService) GetCommunityDetailService(ctx context.Context, communityID, userID int) (dto.CommunityDetailResponse, error) {
	if communityID <= 0 {
		return dto.CommunityDetailResponse{}, errors.New("ID komunitas tidak valid")
	}

	// 1. Ambil data detail & statistik komunitas
	detail, err := cs.cr.GetCommunityByIDRepo(ctx, communityID, userID)
	if err != nil {
		return dto.CommunityDetailResponse{}, err
	}

	// 2. Ambil list anggotanya
	members, err := cs.cr.GetCommunityMembersRepo(ctx, communityID)
	if err != nil {
		return dto.CommunityDetailResponse{}, err
	}

	if members == nil {
		members = []dto.MemberResponse{}
	}

	detail.Members = members
	return detail, nil
}

func (cs *CommunitiesService) GetPopularCommunitiesService(ctx context.Context) ([]dto.PopularCommunityResponse, error) {
	community, err := cs.cr.GetPopularCommunity(ctx)
	if err != nil {
		return nil, err
	}

	if community == nil {
		community = []dto.PopularCommunityResponse{}
	}

	return community, nil
}
