
## **🔹 Example: Mocking a Database in Unit Tests**
### **Scenario**
- We have a `Database` interface with a `GetUser(id int) User` method.
- A `RealDatabase` struct implements it.
- During testing, we replace `RealDatabase` with a `MockDatabase`.

---

### **1️⃣ Define the Interface**
```go
package main

import "fmt"

// User model
type User struct {
    ID   int
    Name string
}

// Define the Database interface
type Database interface {
    GetUser(id int) User
}
```
✅ **Why use an interface?**
- It allows different implementations (real database or mock).
- Makes the code testable and flexible.

---

### **2️⃣ Implement the Real Database**
```go
// RealDatabase struct (Actual implementation)
type RealDatabase struct{}

// Real method that interacts with the database
func (db RealDatabase) GetUser(id int) User {
    fmt.Println("Fetching from real database...")
    return User{ID: id, Name: "Alice"}
}

// Function that depends on the interface
func FetchUser(db Database, id int) {
    user := db.GetUser(id)
    fmt.Println("User found:", user.Name)
}
```
✅ **Why pass `Database` instead of `RealDatabase`?**
- Allows swapping the real database with a **mock** in tests.
- Decouples the function from a specific implementation.

---

### **3️⃣ Create a Mock Database for Testing**
```go
// MockDatabase struct (Used for testing)
type MockDatabase struct{}

// Mock method returning fake data (no real DB calls)
func (db MockDatabase) GetUser(id int) User {
    fmt.Println("Fetching from mock database...")
    return User{ID: id, Name: "TestUser"}
}
```
✅ **Why use a mock?**
- **No need for a real database** during tests.
- **Fast execution** without external dependencies.
- **Controlled output** for predictable test results.

---

### **4️⃣ Run with Real and Mock Implementations**
```go
func main() {
    realDB := RealDatabase{}
    mockDB := MockDatabase{}

    fmt.Println("Using Real Database:")
    FetchUser(realDB, 1)

    fmt.Println("\nUsing Mock Database:")
    FetchUser(mockDB, 1)
}
```
---

### **🔹 Output**
```
Using Real Database:
Fetching from real database...
User found: Alice

Using Mock Database:
Fetching from mock database...
User found: TestUser
```
---

## **🔹 Writing a Unit Test with Mocking (Using `testing` Package)**
Now, let's write a **unit test** using Go’s `testing` package.

### **5️⃣ Unit Test with Mock Database**
```go
package main

import (
    "testing"
)

// Test function using mock database
func TestFetchUser(t *testing.T) {
    mockDB := MockDatabase{} // Using mock

    user := mockDB.GetUser(1)

    if user.Name != "TestUser" {
        t.Errorf("Expected TestUser, got %s", user.Name)
    }
}
```
✅ **Why is this useful?**
- No need to **connect to a real database**.
- The test is **fast and independent**.
- Predictable **mock output** ensures reliability.

---

## **🔹 Mocking External APIs**
Instead of hitting a real API in tests, we can replace it with a mock.

### **Example: Mocking an HTTP API Call**
```go
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

// Interface for API clients
type APIClient interface {
    FetchData(url string) string
}

// Real API client making HTTP calls
type RealClient struct{}

func (r RealClient) FetchData(url string) string {
    resp, _ := http.Get(url)
    body, _ := ioutil.ReadAll(resp.Body)
    return string(body)
}

// Mock API client for testing
type MockClient struct{}

func (m MockClient) FetchData(url string) string {
    return "Mock API Response"
}

func main() {
    realClient := RealClient{}
    mockClient := MockClient{}

    fmt.Println("Real API Call:", realClient.FetchData("http://example.com"))
    fmt.Println("Mock API Call:", mockClient.FetchData("http://example.com"))
}
```
✅ **Why use a mock API?**
- Avoid **slow network calls** in tests.
- Prevent **API rate limits** from affecting tests.
- Ensure **consistent test results**.

---

## **🔹 Final Summary**
| **Scenario** | **Use Real Struct** | **Use Interface & Mock** |
|-------------|-----------------|----------------|
| Need actual data & logic | ✅ Yes | ❌ No |
| Need fast & independent unit tests | ❌ No | ✅ Yes |
| Need to replace implementation dynamically | ❌ No | ✅ Yes |
| Dependency on an external service (DB, API) | ❌ No | ✅ Yes |

---

### **🔹 Key Takeaways**
✔ Use **interfaces** for defining behavior.  
✔ Implement **real structs** for actual functionality.  
✔ Replace **real implementations with mocks** in unit tests.  
✔ Improve **testability, flexibility, and speed** using mocks.