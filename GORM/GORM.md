# Using GORM for Database Interactions

## Overview

**GORM** is a powerful and widely-used ORM (Object-Relational Mapping) library for Go. It simplifies database interactions by allowing developers to work with databases using Go's native data structures. Including GORM in your advanced Go topics will enhance your repository and provide valuable insights into effective database management in Go applications.

### Why Use GORM?

- **Productivity**: Reduces boilerplate code for common database operations.
- **Maintainability**: Provides a consistent and readable codebase.
- **Community Support**: Active development and a large community offering plugins and help.

---

## Installation

Install GORM and the database driver you need:

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql # For MySQL
```
```
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

func main() {
    dsn := "username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    // Use 'db' for further operations
}
```

## Defining Models
Define your database models as Go structs. Use GORM's struct tags to map fields to database columns.

```
type User struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:100;not null"`
    Email     string    `gorm:"unique;not null"`
    Age       int       `gorm:"default:18"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

gorm:"primaryKey" indicates the primary key field.

Other tags like size, unique, and default define column constraints.

## Auto Migration
Automatically migrate your schema:

```
db.AutoMigrate(&User{})
```

# CRUD Operations

## CREATE

```
user := User{Name: "Alice", Email: "alice@example.com"}
result := db.Create(&user) // Pass pointer of data to Create

if result.Error != nil {
    log.Printf("Error creating user: %v", result.Error)
}

fmt.Println(user.ID) // New ID assigned by the database
```

## Read

```
var user User

// Find by primary key
db.First(&user, 1) // SELECT * FROM users WHERE id = 1;

// Find by condition
db.First(&user, "email = ?", "alice@example.com")
```

## UPDATE

```
// Update single column
db.Model(&user).Update("Age", 20)

// Update multiple columns
db.Model(&user).Updates(User{Age: 21, Name: "Alice Smith"})

```

## Delete

```
db.Delete(&user, user.ID)
```