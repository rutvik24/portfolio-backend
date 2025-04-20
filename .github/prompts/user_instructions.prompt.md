## 🧠 Project Goal

Build a scalable, clean, and modular backend service in **Go** for a terminal-style portfolio frontend. The backend should:

- Use **GORM** as the ORM
- Support **PostgreSQL** and **SQLite**, switchable via environment variables
- Use **Gorilla Mux** or **net/http** for routing
- Provide complete **CRUD operations** for portfolio content:
  - Projects
  - Experience
  - Skills/Tech Stack
  - Education
  - Contact (submissions/messages)
- Have **CORS configuration** controllable via environment variables
- Follow clean architecture principles (separate handlers, services, models, and database layers)
- Use **dotenv or environment variable loading**
- Be developer-friendly and production-ready

## ✅ Best Practices to Follow

1. **Project Structure**
    
    /cmd
      main.go
    /config
      config.go
    /models
      <entity>.go
    /db
      db.go
    /routes
      routes.go
    /handlers
      <entity>_handler.go
    /services
      <entity>_service.go
    /middleware
      cors.go
    .env
    go.mod
    

2. **Environment Switching**
   - Use `.env` with `DB_TYPE=postgres` or `sqlite`
   - Load `.env` using `github.com/joho/godotenv`

3. **CORS Configuration**
   - Use Gorilla handlers or custom middleware
   - Load allowed origins from environment variable: `CORS_ORIGINS=*` (comma-separated)

4. **Database**
   - Initialize DB connection in `db/db.go` using GORM
   - Support auto-migration

5. **CRUD API Design (REST)**
   - Follow RESTful route naming:
     ```
     GET    /api/projects
     POST   /api/projects
     GET    /api/projects/{id}
     PUT    /api/projects/{id}
     DELETE /api/projects/{id}
     ```

6. **Code Style**
   - Use meaningful function names
   - Handle errors properly (return structured responses)
   - Prefer interfaces for services

7. **JSON Serialization**
   - Use `json` struct tags in models
   - Follow snake_case in API responses

8. **Logging & Error Handling**
   - Use `log` package or a logging library like `zap`
   - Return standard JSON error format from handlers

## 🤖 How Copilot Can Help

- Generate GORM model structs for each entity
- Scaffold CRUD handlers using Gorilla Mux
- Implement service interfaces and logic
- Add CORS middleware
- Suggest dotenv config loading
- Set up GORM DB connection switchable by env
- Provide .env templates and docker support if needed

## 🧪 Entities to Scaffold

Each of the following should include:
- Model
- Handler
- Service logic
- CRUD routes

### 📁 Project
type Project struct {
  ID          uint   `gorm:"primaryKey" json:"id"`
  Title       string `json:"title"`
  Description string `json:"description"`
  RepoURL     string `json:"repo_url"`
  DemoURL     string `json:"demo_url"`
  TechStack   string `json:"tech_stack"`
}

### 👨‍💼 Experience

### 🛠️ Skills / Tech Stack

### 🎓 Education

### 📩 Contact / Messages

---

Use this instruction file as guidance for all Copilot completions, agent generation and code suggestions to match this backend architecture.
