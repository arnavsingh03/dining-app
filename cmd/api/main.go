package main

import (
    "log"
    "github.com/arnavsingh03/dining-app/internal/config"
    "github.com/gin-gonic/gin"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Initialize Gin router
    router := gin.Default()

    // Start server
    if err := router.Run(":" + cfg.Server.Port); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}