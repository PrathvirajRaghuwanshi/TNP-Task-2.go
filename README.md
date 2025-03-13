# TNP-Task-2.go
# TnpTask2 - Certificate Management System

##  Project Overview
This project is a **Certificate Management System** that provides functionalities for:
- Fetching certificates from a PostgreSQL database.
- Sending individual certificates via email.
- Sending bulk messages to recipients.

It is built using **Go**, **Gorilla Mux**, **PostgreSQL**, and **Go-Mailer (gomail)** for email notifications.

---

##  Approach to the Solution
The project is structured using the **MVC (Model-View-Controller) pattern** with separate layers for:
- **cmd/** → Contains the main entry point (`main.go`).
- **handlers/** → Handles HTTP requests.
- **database/** → Manages database connections.
- **utils/** → Utility functions like JSON responses.

### Key Steps Taken:
1. **Database Connection Handling** → Used PostgreSQL with a centralized connection (`database/database.go`).
2. **Routing & Middleware** → Used `gorilla/mux` for routing and added authentication middleware.
3. **Email Sending** → Implemented email functionality using `gomail` with validation checks.
4. **Modularization** → Split the code into `cmd/`, `handlers/`, and `database/` to keep it maintainable.
5. **Validation** → Used `go-playground/validator` to ensure correct input data.

---

##  Problems Encountered & Solutions
###  1. **Broken Imports ("Could not import module")**
- **Problem:** Initially, Go modules were not set up correctly.
- **Solution:** Initialized Go modules with `go mod init TnpTask2` and fixed import paths.

###  2. **Database Connection Issue**
- **Problem:** PostgreSQL connection refused due to incorrect configurations.
- **Solution:** Fixed credentials in `database.go` and ensured PostgreSQL was running.

###  3. **Validation Errors in JSON Requests**
- **Problem:** Some API requests failed due to missing required fields.
- **Solution:** Integrated `go-playground/validator` for request validation.

###  4. **Emails Not Sending**
- **Problem:** SMTP configuration was incorrect.
- **Solution:** Used correct SMTP credentials and switched to a test SMTP server.

---

##  What Makes This Solution Unique?
✅ **Modular & Scalable:** The project is divided into reusable modules, making it easy to extend.
✅ **Error Handling & Validation:** Uses structured error messages and request validation.
✅ **Bulk Email Support:** Allows mass communication via SMTP with a loop-based email sender.
✅ **Database-Driven:** Fetches certificates dynamically from PostgreSQL.
✅ **Middleware Integration:** Uses authentication middleware for secure endpoints.

---

## How to Run the Project
### 1️⃣ Clone the Repository
```sh
$ git clone https://github.com/yourusername/TnpTask2.git
$ cd TnpTask2
```

### 2️⃣ Setup PostgreSQL Database
- Create a PostgreSQL database and update `database/database.go` with your credentials.
- Run migration scripts (if needed).

### 3️⃣ Install Dependencies
```sh
$ go mod tidy
```

### 4️⃣ Run the Server
```sh
$ go run cmd/main.go
```

### 5️⃣ API Endpoints
| Method | Endpoint | Description |
|--------|---------------------------|-------------------------------|
| GET | `/certificates` | Fetch all certificates |
| POST | `/certificates/{id}/send` | Send a certificate via email |
| POST | `/bulk-messages` | Send bulk email notifications |

---

##  Technologies Used
- **Golang** (Backend)
- **PostgreSQL** (Database)
- **Gorilla Mux** (Router)
- **Go-Mailer (gomail)** (Emailing)
- **Go-Playground Validator** (Input validation)

---




