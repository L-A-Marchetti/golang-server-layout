# Login and Registration Implementation

## Overview
This document outlines the implementation of login and registration functionalities, including HTML pages and server-side functions.

## 9. Create HTML login page

```html
{{define "login"}}
<h2>Login</h2>
<form action="/login" method="post">
    <label for="email">Email:</label><br>
    <input type="email" id="email" name="email" required><br><br>
    <label for="password">Password:</label><br>
    <input type="password" id="password" name="password" required><br><br>
    <input type="submit" value="Login">
</form>
{{end}}
```

## 10. Implement server-side login function

```go
package handlers

import (
    "net/http"
    "db"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    email := r.FormValue("email")
    password := r.FormValue("password")

    user, err := db.GetUserByEmail(email)
    if err != nil {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    if !db.CheckPasswordHash(password, user.Password) {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    // TODO: Implement session management here

    http.Redirect(w, r, "/", http.StatusSeeOther)
}
```

## 11. Create HTML registration page

```html
{{define "register"}}
<h2>Register</h2>
<form action="/register" method="post">
    <label for="email">Email:</label><br>
    <input type="email" id="email" name="email" required><br><br>
    <label for="username">Username:</label><br>
    <input type="text" id="username" name="username" required><br><br>
    <label for="password">Password:</label><br>
    <input type="password" id="password" name="password" required><br><br>
    <input type="submit" value="Register">
</form>
{{end}}
```

## 12. Implement server-side registration function

```go
package handlers

import (
    "net/http"
    "db"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    email := r.FormValue("email")
    username := r.FormValue("username")
    password := r.FormValue("password")

    // Check if email is already taken
    if _, err := db.GetUserByEmail(email); err == nil {
        http.Error(w, "Email already in use", http.StatusConflict)
        return
    }

    // Encrypt password before storing
    hashedPassword, err := db.HashPassword(password)
    if err != nil {
        http.Error(w, "Error processing registration", http.StatusInternalServerError)
        return
    }

    err = db.CreateUser(email, username, hashedPassword)
    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/login", http.StatusSeeOther)
}
```