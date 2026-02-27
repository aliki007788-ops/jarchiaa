package main

import (
    "log"
    "net/http"
    "user-service/handlers"
    "user-service/repository"
    "user-service/service"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
    // اتصال به دیتابیس
    db, err := repository.Connect("postgres://jarchia:jarchia123@localhost:5432/jarchia?sslmode=disable")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    
    // ایجاد repositoryها
    userRepo := repository.NewUserRepository(db)
    sessionRepo := repository.NewSessionRepository(db)
    
    // ایجاد سرویس‌ها
    authService := service.NewAuthService(userRepo, sessionRepo, "your-jwt-secret-key-change-this")
    
    // ایجاد هندلرها
    authHandler := handlers.NewAuthHandler(authService)
    
    // تنظیم روتر
    r := mux.NewRouter()
    api := r.PathPrefix("/api/auth").Subrouter()
    
    api.HandleFunc("/login", authHandler.Login).Methods("POST")
    api.HandleFunc("/verify", authHandler.Verify).Methods("POST")
    api.HandleFunc("/register", authHandler.Register).Methods("POST")
    api.HandleFunc("/me", authHandler.GetMe).Methods("GET")
    
    // CORS
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    })
    
    handler := c.Handler(r)
    
    log.Println("🚀 User service running on :8081")
    log.Fatal(http.ListenAndServe(":8081", handler))
}
