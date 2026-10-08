# Eventhub - Backend Service API <image src="https://raw.githubusercontent.com/marwin1991/profile-technology-icons/refs/heads/main/icons/go.png" width="25px">


[![License MIT](https://img.shields.io/badge/LICENSE-MIT-05a815?style=flat&logo=opensourceinitiative&logoColor=8fcf95)](https://opensource.org/license/mit)
[![Status](https://img.shields.io/badge/Status-In%20Development-orange)](#)

> Welcome to the **EventHub Backend API** repository. This service powers the EventHub platform, providing robust RESTful APIs designed for seamless integration with the EventHub Frontend application.

> The core service manages authentication, events management, community interactions, notifications, and user/organizer administration.

## Tech Stack
[![Golang](https://img.shields.io/badge/GO-1.27.1-blue?logo=go)](https://go.dev/)
[![Gin Gonic](https://img.shields.io/badge/Gin_Gonic-1.12.0-blue?logo=gin&logoColor=orange)](https://gin-gonic.com/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16.15-blue?logo=postgresql)](https://hub.docker.com/_/postgres)
[![Redis](https://img.shields.io/badge/Redis-latest-blue?logo=redis)](https://hub.docker.com/_/redis)
[![Swagger](https://img.shields.io/badge/Swagger-latest-blue?logo=swagger)](https://github.com/swaggo/swag)

<br/>

## 🚀 Features
> * Authentication (`/api/auth`)
> * User Profiles (`/api/users`)
> * Events Management (`/api/events`)
> * Community Management (`/api/community`)
> * Testimonials & Feedback (`/api/testimonies`)
> * Notifications (`/api/notifications`)
> * Dashboards (`/api/dashboard`)

## ⚡ Getting Started
### Installation & Run

1. **Clone The Repository :**
   ```bash
   $ git clone https://github.com/nico-hondo/koda-b9-backend.git
   cd koda-b9-backend
   ```
2. **Configure The Environment :**
   > Copy `.env.example` to `.env` and fill in your DB credentials and JWT secret.
   ```bash
   $ cp .env.example .env
   ```
3. **Install Dependencies :**
   ```bash
   $ go mod download
   ```
4. **Run The Application :**
   ```bash
   $ go run main.go

### Documentation
> For Complete documentation visit `/swagger/index.html`

<br/>

## 📄 License
>This project is licensed under the **MIT License** and part of the **Koda Academy** curriculum and is intended for demonstration and portfolio purposes.

---
## 🔗 Frontend Integration
>This repository serves as the backend engine for the **EventHub Frontend**. Ensure your frontend configuration points its base API URL to this backend server [Frontend - Eventhub](https://github.com/nico-hondo/eventhub-app.git)