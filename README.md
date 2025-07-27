# 🚀 Finara API - Financial Goal Management System

A robust REST API built with **Go** and **Gin framework** for managing financial goals and serving dynamic HTML dashboards from Firebase Firestore.

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Firebase](https://img.shields.io/badge/Firebase-Firestore-FFCA28?style=flat&logo=firebase)](https://firebase.google.com/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## 📋 Table of Contents

- [Features](#-features)
- [Project Structure](#-project-structure)
- [API Endpoints](#-api-endpoints)
- [Quick Start](#-quick-start)
- [Configuration](#-configuration)
- [Usage Examples](#-usage-examples)
- [Goal Dashboard](#-goal-dashboard)
- [Sample Responses](#-sample-responses)
- [Development](#-development)
- [Deployment](#-deployment)
- [Contributing](#-contributing)

## 🌟 Features

- **User Management**: Complete CRUD operations for user profiles
- **Goal Tracking**: Financial goal creation, retrieval, and management
- **Dynamic Dashboards**: HTML dashboard generation with real-time data
- **Firebase Integration**: Seamless Firestore database connectivity
- **CORS Support**: Ready for frontend integration
- **Health Monitoring**: Built-in health check endpoints
- **RESTful Design**: Clean, intuitive API design

## 🏗️ Project Structure

```
finara-api/
├── main.go                                    # 🎯 Application entry point
├── go.mod                                     # 📦 Go module dependencies
├── go.sum                                     # 🔒 Dependency checksums
├── goal_dashboard.html                        # 📊 Sample dashboard template
├── config/
│   ├── firebase.go                           # 🔧 Firebase configuration
│   └── firebase-service-account.json        # 🔑 Firebase service account
├── handlers/
│   └── user_handler.go                       # 🎮 API request handlers
└── router/
    └── routes.go                             # 🛣️  API route definitions
```

## 🔗 API Endpoints

### Core Endpoints

| Method | Endpoint | Description | Status |
|--------|----------|-------------|--------|
| `GET` | `/health` | Health check | ✅ Active |
| `GET` | `/` | API information | ✅ Active |

### User Management

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| `GET` | `/users` | Get all users | - |
| `GET` | `/users/:userId` | Get user by ID | - |
| `POST` | `/users/:userId` | Create/Update user | [UserRequest](#user-request) |

### Goal Management

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| `GET` | `/users/:userId/goal_info` | Get user's goals | - |
| `GET` | `/users/:userId/goal_info/:goalId` | Get specific goal | - |
| `POST` | `/users/:userId/goal_info/:goalId` | Create/Update goal | [GoalRequest](#goal-request) |

### Dashboard Endpoints

| Method | Endpoint | Description | Output |
|--------|----------|-------------|--------|
| `GET` | `/dashboard` | Static dashboard | HTML |
| `GET` | `/dashboard/:userId` | User dashboard | HTML |
| `GET` | `/users/:userId/goal_info/:goalId/status/:statusId/dashboard` | Goal dashboard | HTML |
| `GET` | `/users/:userId/goal_info/:goalId/status/:statusId/view` | Goal dashboard (alt) | HTML |

## 🚀 Quick Start

### Prerequisites

- **Go 1.23+** installed
- **Firebase project** with Firestore enabled
- **Service account key** for Firebase

### Installation

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd finara-api
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   go mod tidy
   ```

3. **Configure Firebase**:
   - Place your `firebase-service-account.json` in the `config/` directory
   - Update Firebase project settings if needed

4. **Run the application**:
   ```bash
   go run main.go
   ```

5. **Verify installation**:
   ```bash
   curl http://localhost:8080/health
   ```

The server will start on port `8080` by default (configurable via `PORT` environment variable).

## ⚙️ Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8080` |
| `GOOGLE_APPLICATION_CREDENTIALS` | Path to service account | `config/firebase-service-account.json` |

### Firebase Setup

1. Create a Firebase project
2. Enable Firestore Database
3. Generate a service account key
4. Download and place in `config/firebase-service-account.json`

## 💡 Usage Examples

### User Operations

**Create a User:**
```bash
curl -X POST http://localhost:8080/users/12345 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "age": 30,
    "email": "john@example.com",
    "mobile_no": "1234567890",
    "preferred_language": "English",
    "marrital_status": "Single",
    "city": "New York",
    "career_stage": "Mid-level"
  }'
```

**Get User:**
```bash
curl http://localhost:8080/users/12345
```

### Goal Operations

**Create a Goal:**
```bash
curl -X POST http://localhost:8080/users/12345/goal_info/goal_01 \
  -H "Content-Type: application/json" \
  -d '{
    "goal_amount": "200000",
    "goal_description": "Buy a car",
    "goal_line": "Transportation",
    "goal_timeline": "3 years",
    "goal_set_date": "2024-01-01"
  }'
```

**Get Goal:**
```bash
curl http://localhost:8080/users/12345/goal_info/goal_01
```

### Dashboard Access

**View Goal Dashboard:**
```bash
# Open in browser
http://localhost:8080/users/6666666666/goal_info/goal_01/status/status_01/dashboard
```

## 📊 Goal Dashboard

The Goal Dashboard feature provides dynamic, interactive HTML dashboards with:

- **Financial Metrics**: Current assets, goal progress, success probability
- **Visual Charts**: Progress bars, line charts, and allocation charts
- **Strategic Insights**: Personalized financial advice and recommendations
- **Responsive Design**: Mobile-friendly interface

### Dashboard Components

- 📈 **Progress Tracking**: Visual goal completion status
- 💡 **Smart Insights**: AI-powered financial recommendations  
- 📊 **Interactive Charts**: Built with Chart.js
- 🎯 **Action Items**: Prioritized improvement suggestions

## 📄 Sample Responses

### User Request
```json
{
  "name": "John Doe",
  "age": 30,
  "email": "john@example.com",
  "mobile_no": "1234567890",
  "preferred_language": "English",
  "marrital_status": "Single",
  "city": "New York",
  "career_stage": "Mid-level"
}
```

### Goal Request
```json
{
  "goal_amount": "200000",
  "goal_description": "Buy a car",
  "goal_line": "Transportation",
  "goal_timeline": "3 years",
  "goal_set_date": "2024-01-01"
}
```

### User Response
```json
{
  "id": "12345",
  "name": "John Doe",
  "age": 30,
  "email": "john@example.com",
  "mobile_no": "1234567890",
  "preferred_language": "English",
  "marrital_status": "Single",
  "city": "New York",
  "career_stage": "Mid-level"
}
```

### Goal Response
```json
{
  "id": "goal_01",
  "goal_amount": "200000",
  "goal_description": "Buy a car",
  "goal_line": "Transportation",
  "goal_timeline": "3 years",
  "goal_set_date": "2024-01-01"
}
```

## 🛠️ Development

### Project Dependencies

```go
require (
    cloud.google.com/go/firestore v1.18.0
    firebase.google.com/go/v4 v4.17.0
    github.com/gin-contrib/cors v1.7.6
    github.com/gin-gonic/gin v1.10.1
    google.golang.org/api v0.243.0
    google.golang.org/grpc v1.73.0
)
```
