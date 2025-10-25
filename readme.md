# CRUD TASK
This is fullstack task CRUD application built with Go (Gin) for the backend, MongoDB for persistence, and Vue.js for the frontend.

## What I built
1. Backend Go REST endpoints for CRUD (create, read, update, delete) tasks and login authentication.
2. Frontend with Vue.js that consumes the API and provides interactive UI for managing tasks (filtering, pagination, sorting).

## Folder Structure
```
/backend  
   ├─ model/            → domain models (Task, Auth)  
   ├─ handler/           → HTTP handlers and interface definitions  
   ├─ usecase/           → business logic  
   ├─ repository/        → persistence implementations (Mongo)  
   ├─ middleware/        → JWT auth middleware  
/frontend  
   ├─ src/              → Vue app source  
   ├─ components/       → UI components  
   ├─ views/            → page views  
```

## Design decisions & architecture
### Backend
1. Using Gin for fast HTTP routing for golang
2. Clean architecture (or layered architecture) style: packages for model, handler/interfaces, usecase, repository, etc. This separation helps maintainability, testability, and future extensions.
3. Interfaces are mocked in tests using Mockery so that business logic (use-cases) can be tested independent of database layer.
4. JWT token for authentication

### Frontend
1. Used Vue.js for simplicity and reactive UI.
2. Tailwind CSS for styling responsive UI.
3. Axios for REST API calls.

### MongoDB Indexes
1. Index on status, because listing tasks often filters for status to know in backlog, progress, or completed.
2. Index on title, to support for search by title
3. Index on created_at with sort descending, because listing is often sort by most recently created.
4. Compound index status and created_at, for covering both fields the filter and the sort

## Strength of this Module
1. Separation of concern. Business logic in use-cases, repository abstracts persistence, handlers only deal with HTTP. This makes the code easier to maintain, extend, and test.
2. Test coverage by mocking interfaces, we can test core logic without needing a live database or HTTP server.
3. Task listing (filter + sort + pagination) is efficient, even with large datasets.

## Setup and Run
### Backend
1. `cd backend`
2. Copy `.env.example` into `.env`
3. `go mod tidy`
4. `go run main.go`

If want to test run `make test`

### Frontend
1. `cd frontend`
2. Copy `.env.example` into `.env`
3. `npm install`
4. `npm run dev`