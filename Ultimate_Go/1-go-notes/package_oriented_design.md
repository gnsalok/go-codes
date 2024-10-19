# Package Oriented Design


When creating an API in Go with a **package-oriented design**, it’s important to follow a structured and scalable folder layout. This layout should promote maintainability, clarity, and clean separation of concerns. A common approach is to adopt a layout that uses `internal`, `cmd`, and other key directories, helping you organize different aspects of your application. Below is a best-practice folder structure, along with explanations of each part.

### Example Folder Structure

```
/your-project
├── /cmd
│   ├── /your-api
│   │   └── main.go
├── /internal
│   ├── /api
│   │   ├── handlers.go
│   │   ├── routes.go
│   ├── /services
│   │   └── user_service.go
│   ├── /repository
│   │   └── user_repo.go
├── /pkg
│   ├── /models
│   │   └── user.go
│   ├── /utils
│   │   └── logger.go
│   └── /middleware
│       └── auth_middleware.go
├── /configs
│   └── config.yaml
├── /scripts
│   └── some-script.sh
├── /test
│   ├── integration_test.go
│   └── unit_test.go
├── /docs
│   └── api_documentation.md
├── go.mod
├── go.sum
└── README.md
```

### Key Directories and When to Use Them

1. **`cmd/` (Command Directory)**:
   - **Purpose**: This directory is used to hold the main application entry points for different binaries. In Go, the `cmd` folder is a convention used to organize your executable code.
   - **Why/When**: If you have multiple applications or binaries, you can organize them under `cmd`. For instance, you might have multiple commands like `cmd/your-api` for your API server, and another like `cmd/migrate` for a database migration tool.
   - **Example**: Inside `cmd/your-api/main.go`, you would initialize the server, set up routes, middlewares, and so on.

   ```go
   // cmd/your-api/main.go
   package main

   import (
       "fmt"
       "yourproject/internal/api"
   )

   func main() {
       fmt.Println("Starting API server...")
       api.StartServer()  // You might call your API entry point here
   }
   ```

2. **`internal/` (Private Application Logic)**:
   - **Purpose**: The `internal` directory is for code that is private to your application or service. Go enforces that code inside `internal/` cannot be imported by code outside your project, providing a way to explicitly limit the scope of your internal packages.
   - **Why/When**: Anything inside `internal/` is considered implementation detail and not for public use, which is especially helpful in larger applications or libraries. This can be where most of your API-specific logic and business domain code lives.
   - **Subdirectories**:
     - `/api`: Contains the API handler logic, such as route definitions and request/response handling.
     - `/services`: Encapsulates business logic (e.g., user services, payment processing).
     - `/repository`: Data access logic (e.g., interaction with databases or external services).
   
   ```go
   // internal/api/routes.go
   package api

   import (
       "github.com/gin-gonic/gin"
   )

   func SetupRouter() *gin.Engine {
       r := gin.Default()
       r.GET("/users", GetUserHandler)
       return r
   }
   ```

   ```go
   // internal/api/handlers.go
   package api

   import "github.com/gin-gonic/gin"

   func GetUserHandler(c *gin.Context) {
       // Logic to get user data
   }
   ```

3. **`pkg/` (Reusable Libraries)**:
   - **Purpose**: `pkg` is for packages that could potentially be shared across different projects or could be reusable across different parts of your application. This directory is useful for code that’s not specific to your application but is instead a general utility.
   - **Why/When**: Anything inside `pkg/` is public and reusable. For example, utility functions, logging, or models can go here. This code can be imported into your other projects if needed.
   - **Subdirectories**:
     - `/models`: Holds the structs and data models that define your domain objects (e.g., `User`, `Order`).
     - `/utils`: General helper functions (e.g., string utilities, logger).
     - `/middleware`: Middleware code for things like authentication, logging, etc.

   ```go
   // pkg/models/user.go
   package models

   type User struct {
       ID    int
       Name  string
       Email string
   }
   ```

   ```go
   // pkg/utils/logger.go
   package utils

   import "log"

   func LogInfo(message string) {
       log.Println("INFO: ", message)
   }
   ```

4. **`configs/` (Configuration Files)**:
   - **Purpose**: Contains configuration files for your project (e.g., YAML, JSON, or TOML files that contain environment variables, database configurations, etc.).
   - **Why/When**: Keeping configuration separate makes it easier to adjust environment-specific settings without altering your codebase. This might include database credentials, API tokens, etc.

   ```yaml
   # configs/config.yaml
   server:
     port: 8080
   database:
     host: localhost
     port: 5432
   ```

5. **`scripts/` (Automation Scripts)**:
   - **Purpose**: Contains utility scripts for automating tasks such as database migrations, deployment scripts, or CI/CD pipelines.
   - **Why/When**: This is useful when you want to automate development tasks or deployment processes.

6. **`test/` (Testing)**:
   - **Purpose**: Stores integration and unit tests, separate from the main application code.
   - **Why/When**: To separate your testing logic from your production code. This directory can also contain testing utilities, mock data, and more.

   ```go
   // test/unit_test.go
   package test

   import (
       "testing"
   )

   func TestSomething(t *testing.T) {
       t.Log("Test passed")
   }
   ```

7. **`docs/` (Documentation)**:
   - **Purpose**: Stores documentation for your API, such as Swagger/OpenAPI specs, markdown documentation, or developer guides.
   - **Why/When**: Proper documentation ensures that developers (including yourself) can quickly understand and contribute to the project. This is essential for open-source projects or teams.

   ```markdown
   # API Documentation

   This API does XYZ...

   ## Endpoints

   - `GET /users`: Retrieves all users.
   - `POST /users`: Creates a new user.
   ```

### Benefits of This Structure

1. **Separation of Concerns**:
   - `cmd` separates the entry point from the business logic, allowing multiple binaries or applications within the same project.
   - `internal` keeps your internal logic encapsulated and prevents accidental usage outside of your project.
   - `pkg` contains reusable code that can be shared across projects or other teams.

2. **Encapsulation**:
   - By using `internal`, you enforce encapsulation, ensuring that only what needs to be exposed is available to the outside world. This allows for better maintainability.

3. **Modularity and Scalability**:
   - This structure is scalable for large projects because it allows you to organize code into distinct modules. As your API grows, adding more services, repositories, or utility functions becomes easier.

4. **Testing**:
   - The `test` directory ensures that your tests are kept separate from your business logic, making it easier to maintain and run tests independently.

5. **Clean Imports**:
   - By organizing your code into packages, imports remain clear and explicit. For example:
     - `yourproject/internal/api`
     - `yourproject/pkg/models`

### Conclusion

This folder structure gives you a clear separation of concerns, improves encapsulation, and ensures maintainability for medium to large-scale Go projects. Using `internal` protects implementation details, `cmd` handles various entry points, and `pkg` offers a way to reuse and share code.

Let me know if you want more details or examples!