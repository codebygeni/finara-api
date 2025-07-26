# Geni Firestore API

A simple REST API built with Go and Gin framework to retrieve data from Firebase Firestore.

## Project Structure

```
geni-firestore-api/
├── main.go                                    # Main application entry point
├── go.mod                                     # Go module dependencies
├── config/
│   ├── firebase.go                           # Firebase configuration
│   └── firebase-service-account.json        # Firebase service account key
├── handlers/
│   └── user_handler.go                       # User API handlers
└── router/
    └── routes.go                             # API route definitions
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check endpoint |
| GET | `/users` | Retrieve all users |
| GET | `/users/:userId` | Retrieve specific user by ID |
| GET | `/users/:userId/goal_info` | Retrieve all goals for a specific user |
| GET | `/users/:userId/goal_info/:goalId` | Retrieve a specific goal by ID |

## Quick Start

1. **Install dependencies**:
   ```bash
   go mod download
   ```
1. **Install dependencies**:
   ```bash
   go mod tidy
   ```

2. **Add your Firebase service account key**:
   - Place your `firebase-service-account.json` file in the `config/` directory

3. **Run the application**:
   ```bash
   go run main.go
   ```

4. **Test the endpoints**:
   ```bash
   curl http://localhost:8080/health
   curl http://localhost:8080/users/6666666666
   curl http://localhost:8080/users/6666666666/goal_info/goal_01
   ```

## Sample API Responses

### GET /users/6666666666/goal_info/goal_01
```json
{
  "id": "goal_01",
  "goal_amount": 1000000,
  "goal_description": "Save for retirement",
  "goal_line": "Retirement Planning",
  "goal_timeline": 25
}
```