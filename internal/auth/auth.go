package auth

import (   
    "net/http"
)
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    username := r.FormValue("username")
    password := r.FormValue("password")
    //Just a basic logic which would work for only one user
    if username != "admin" || password != "password" {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    tokenString, err := GenerateJWT(username)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(tokenString))
}