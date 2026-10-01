# EventHub - Backend API Services

[![Go Standard Library](https://img.shields.io/badge/Language-Go-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Architecture](https://img.shields.io/badge/Architecture-Modular%20%2F%20Clean%20Architecture-blue)](#project-architecture)
[![Status](https://img.shields.io/badge/Status-In%20Development-orange)](#)

Welcome to the **EventHub Backend API** repository. This service powers the EventHub platform, providing robust RESTful APIs designed for seamless integration with the EventHub Frontend application.

The core service manages authentication, events management, community interactions, notifications, and user/organizer administration.

---

## 🛠 Project Architecture

This project strictly adheres to a **Modular / Layered Architecture** pattern to ensure maintainability, scalability, and loose coupling between components.

```text
internal/
├── dto/         # Data Transfer Objects & Request/Response Validation
├── handler/     # HTTP Request Handlers / Controllers
├── model/       # Database Entities & Models
├── repo/        # Database Access Layer (Repositories)
├── router/      # API Route definitions & Middleware wiring
└── service/     # Core Business Logic Layer
```

---

## 🚀 Features & API Endpoints

All API endpoints are documented below grouped by functional domain.

### 🔑 Authentication & Security (`/api/v1/auth`)
* `POST /auth/register` — Register a new user account.
* `POST /auth/login` — Authenticate user and issue session token/JWT.
* `POST /auth/logout` — Revoke active token/session (Blacklist / Whitelist mechanism).
* `POST /auth/forgot-password` — Request a password reset link or OTP.
* `POST /auth/reset-password` — Create a new password following a reset request.

### 🎟️ Events Management (`/api/v1/events`)
* `GET /events` — Retrieve event list with support for search keywords and category filters.
* `GET /events/:id` — Get detailed information about a specific event.
* `GET /events/upcoming` — Fetch upcoming featured/trending events.
* `GET /events/my-events` — Get events created or managed by the authenticated user.
* `POST /events` — Create a new event (Organizer/Admin).
* `PUT /events/:id` — Edit an existing event details.
* `POST /events/:id/join` — Join/Register for an event.
* `DELETE /events/:id/leave` — Cancel participation / Leave an event.

### 👥 Communities (`/api/v1/communities`)
* `GET /communities` — Retrieve community list with search and filter parameters.
* `GET /communities/popular` — Get top/popular communities based on activity or member count.
* `GET /communities/:id` — Get community detail and info.
* `GET /communities/:id/members` — Fetch member list of a specific community.
* `POST /communities/:id/join` — Join a community.
* `DELETE /communities/:id/leave` — Leave a community.

### 👤 User Profile & Settings (`/api/v1/users`)
* `GET /users/profile` — Fetch current user's profile details.
* `PUT /users/profile` — Update user profile details (Name, Bio, Avatar, etc.).
* `PUT /users/change-password` — Change password for authenticated users.

### 💬 Testimonials & Feedback (`/api/v1/testimonies`)
* `GET /testimonies` — Get testimonials submitted by users.
* `POST /testimonies` — Submit or update a user testimonial.

### 🔔 Notifications (`/api/v1/notifications`)
* `GET /notifications` — Fetch all notifications for the authenticated user.

### 📊 Dashboards (`/api/v1/dashboard`)
* `GET /dashboard/organizer` — Get organizer metrics, statistics, and information.
* `GET /dashboard/admin` — Get platform-wide administrative metrics and control information.

---

## 💻 Tech Stack & Standards

* **Language:** Go (Golang)
* **Naming Conventions:** All functions, methods, and variables strictly follow standard **English** naming conventions.
* **Architecture:** Modular Layered Architecture (Handler - Service - Repository Pattern).

---

## ⚡ Getting Started

### Prerequisites
* Go 1.20 or higher installed.
* PostgreSQL (or configured database instance).

### Installation & Run

1. **Clone the repository:**
   ```bash
   git clone https://github.com/your-username/eventhub-backend.git
   cd eventhub-backend
   ```

2. **Configure Environment Variables:**
   Copy `.env.example` to `.env` and fill in your DB credentials and JWT secret.
   ```bash
   cp .env.example .env
   ```

3. **Install dependencies:**
   ```bash
   go mod download
   ```

4. **Run the Application:**
   ```bash
   go run main.go
   ```

---

## 🔗 Frontend Integration

This repository serves as the backend engine for the **EventHub Frontend**. Ensure your frontend configuration points its base API URL to this backend server (`http://localhost:8000/api/v1` or configured port).

---

## 📄 License
This project is part of the **Koda Academy** curriculum and is intended for demonstration and portfolio purposes.