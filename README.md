# Go + Fiber + SQLite CRUD Application

A beginner-friendly RESTful CRUD application built to learn **Go**, the **Fiber v3** web framework, and **GORM** with **SQLite**.

---

## 📁 Project Structure

```text
.
├── controllers/       # HTTP request handlers & CRUD business logic
│   └── book.go
├── db/                # Database connection & auto-migration setup
│   └── db.go
├── docs/              # Auto-generated Swagger/OpenAPI specifications
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── models/            # Database schemas & request DTO structs
│   └── book.go
├── routes/            # Route registrations & grouping
│   └── routes.go
├── main.go            # Application entry point
├── books.db           # SQLite database file
├── go.mod
└── go.sum
```

---

## 🚀 How to Run

1. **Run the server:**
   ```bash
   go run main.go
   ```
   *(The server will start at `http://localhost:8080` and automatically create the SQLite database file `books.db`)*

---

## 📖 Interactive Swagger / OpenAPI Documentation

Once the server is running, open your browser and navigate to:
👉 **[http://localhost:8080/swagger/](http://localhost:8080/swagger/)**

You can inspect all routes, read request/response schemas, and execute test requests directly from the UI.

### Regenerating Docs after Code Changes:
Whenever you add or change annotations in `controllers/` or `main.go`, re-run:
```bash
swag init
```

---

## 📡 API Endpoints & Testing with cURL

### 1. Create a Book (POST)
```bash
curl -X POST http://localhost:8080/api/books \
  -H "Content-Type: application/json" \
  -d '{"title": "The Go Programming Language", "author": "Alan Donovan & Brian Kernighan", "year": 2015}'
```

Create a second book:
```bash
curl -X POST http://localhost:8080/api/books \
  -H "Content-Type: application/json" \
  -d '{"title": "Clean Code", "author": "Robert C. Martin", "year": 2008}'
```

---

### 2. Get All Books (GET)
```bash
curl http://localhost:8080/api/books
```

---

### 3. Get Single Book by ID (GET)
```bash
curl http://localhost:8080/api/books/1
```

---

### 4. Update a Book (PUT)
```bash
curl -X PUT http://localhost:8080/api/books/1 \
  -H "Content-Type: application/json" \
  -d '{"title": "The Go Programming Language (2nd Edition)", "year": 2024}'
```

---

### 5. Delete a Book (DELETE)
```bash
curl -X DELETE http://localhost:8080/api/books/1
```

---

## 🧠 Key Go & Fiber Concepts Demonstrated

1. **GORM Struct Tags**:
   - `json:"title"` defines the field name when serializing/deserializing JSON.
   - `gorm:"primaryKey"` tells GORM to treat the column as the primary key.
   - `gorm:"not null"` enforces database constraints.

2. **Auto-Migration**:
   - `DB.AutoMigrate(&Book{})` creates the SQLite table automatically without writing raw SQL.

3. **Fiber v3 Request Body Binding**:
   - `c.Bind().JSON(&input)` parses incoming JSON bodies safely into Go structs.

4. **Error Handling**:
   - `errors.Is(err, gorm.ErrRecordNotFound)` checks if a query failed because the record doesn't exist (returns `404 Not Found`).

