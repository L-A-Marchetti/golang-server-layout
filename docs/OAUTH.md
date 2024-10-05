# OAuth Integration for Google and GitHub

## Overview
This document outlines the implementation of OAuth authentication for Google and GitHub.

## 13. Integrate Google authentication

### Configure Google API
1. Go to the [Google Cloud Console](https://console.cloud.google.com/).
2. Create a new project or select an existing one.
3. Enable the Google+ API.
4. Create OAuth 2.0 credentials (client ID and client secret).
5. Set up the authorized redirect URIs.

### Implement server-side OAuth flow

```go
package handlers

import (
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "net/http"
)

var googleOauthConfig = &oauth2.Config{
    ClientID:     "YOUR_GOOGLE_CLIENT_ID",
    ClientSecret: "YOUR_GOOGLE_CLIENT_SECRET",
    RedirectURL:  "http://localhost:8080/auth/google/callback",
    Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
    Endpoint:     google.Endpoint,
}

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
    url := googleOauthConfig.AuthCodeURL("state")
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
    // Handle the OAuth2 callback
    // Exchange the authorization code for a token
    // Use the token to get user information
    // Create or update user in your database
    // Set up a session for the user
}
```

### Create Google login button on client-side

```html
{{define "google-login"}}
<a href="/auth/google/login" class="google-btn">
    <img src="/static/img/google-logo.png" alt="Google Logo">
    Sign in with Google
</a>
{{end}}
```

## 14. Integrate GitHub authentication

### Configure GitHub API
1. Go to [GitHub Developer Settings](https://github.com/settings/developers).
2. Create a new OAuth App.
3. Get the client ID and client secret.
4. Set up the authorization callback URL.

### Implement server-side OAuth flow

```go
package handlers

import (
    "golang.org/x/oauth2"
    "golang.org/x/oauth2/github"
    "net/http"
)

var githubOauthConfig = &oauth2.Config{
    ClientID:     "YOUR_GITHUB_CLIENT_ID",
    ClientSecret: "YOUR_GITHUB_CLIENT_SECRET",
    RedirectURL:  "http://localhost:8080/auth/github/callback",
    Scopes:       []string{"user:email"},
    Endpoint:     github.Endpoint,
}

func GitHubLoginHandler(w http.ResponseWriter, r *http.Request) {
    url := githubOauthConfig.AuthCodeURL("state")
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GitHubCallbackHandler(w http.ResponseWriter, r *http.Request) {
    // Handle the OAuth2 callback
    // Exchange the authorization code for a token
    // Use the token to get user information
    // Create or update user in your database
    // Set up a session for the user
}
```

### Create GitHub login button on client-side

```html
{{define "github-login"}}
<a href="/auth/github/login" class="github-btn">
    <img src="/static/img/github-logo.png" alt="GitHub Logo">
    Sign in with GitHub
</a>
{{end}}
```