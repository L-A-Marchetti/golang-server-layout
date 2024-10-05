# Session Management with UUID Cookies

## Overview
This document outlines the implementation of session management using UUID cookies.

## 15. Implement session management with UUID cookies

### Generate UUIDs for user sessions

```go
package session

import (
    "github.com/google/uuid"
    "net/http"
    "time"
)

func GenerateSessionID() string {
    return uuid.New().String()
}

func CreateSession(w http.ResponseWriter, userID int) {
    sessionID := GenerateSessionID()
    
    // Set cookie
    http.SetCookie(w, &http.Cookie{
        Name:     "session_id",
        Value:    sessionID,
        Expires:  time.Now().Add(24 * time.Hour),
        HttpOnly: true,
        Secure:   true, // Set to true if using HTTPS
        SameSite: http.SameSiteStrictMode,
    })

    // Store session in server (implement this function)
    StoreSession(sessionID, userID)
}
```

### Store session state securely on the server

```go
package session

import (
    "sync"
    "time"
)

type Session struct {
    UserID    int
    CreatedAt time.Time
}

var (
    sessions = make(map[string]Session)
    mutex    sync.RWMutex
)

func StoreSession(sessionID string, userID int) {
    mutex.Lock()
    defer mutex.Unlock()
    sessions[sessionID] = Session{
        UserID:    userID,
        CreatedAt: time.Now(),
    }
}

func GetSession(sessionID string) (Session, bool) {
    mutex.RLock()
    defer mutex.RUnlock()
    session, exists := sessions[sessionID]
    return session, exists
}

func DeleteSession(sessionID string) {
    mutex.Lock()
    defer mutex.Unlock()
    delete(sessions, sessionID)
}
```

### Middleware for session validation

```go
package middleware

import (
    "net/http"
    "yourproject/session"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("session_id")
        if err != nil {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        sessionID := cookie.Value
        _, exists := session.GetSession(sessionID)
        if !exists {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }

        next.ServeHTTP(w, r)
    }
}
```