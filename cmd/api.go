/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/pstpmn/my-golang-hexagonal-template/conf"
	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/port"
	"github.com/pstpmn/my-golang-hexagonal-template/internal/core/usecase"
	httpHandler "github.com/pstpmn/my-golang-hexagonal-template/pkg/handlers/http"
	userRepository "github.com/pstpmn/my-golang-hexagonal-template/pkg/storage/mongo/repositories/user"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	timezone = "Asia/Bangkok"
)

// apiCmd represents the API server command.
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the Order Autoplatf API server",
	Long: `Launches the Order Autoplatf API server, providing endpoints for user management
and order processing. Requires a MongoDB connection and environment configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runAPI(); err != nil {
			log.Fatalf("API startup failed: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

// runAPI initializes and starts the API server.
func runAPI() error {
	fmt.Println("Starting API server...")

	if _, err := time.LoadLocation(timezone); err != nil {
		return fmt.Errorf("failed to load timezone %s: %w", timezone, err)
	}

	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Failed to load .env file: %v", err)
	}

	cfg := conf.LoadConfigFromEnv()

	// mongoClient := mongo.NewMongo()
	// conn, err := mongoClient.Connect(cfg.Mongo.Uri)
	// if err != nil {
	// 	return fmt.Errorf("failed to connect to MongoDB: %w", err)
	// }
	// defer conn.Disconnect(context.Background())

	// Optional: Test MongoDB connection
	// if err := mongoClient.Ping(*conn); err != nil {
	// 	return fmt.Errorf("MongoDB ping failed: %w", err)
	// }

	// Initialize dependencies
	var cache port.ICache // Placeholder for cache implementation
	middleware := httpHandler.NewMiddlewareHandler()
	userRepo := userRepository.NewUserRepo(&mongo.Client{}, cache)
	userUseCase := usecase.NewUserUseCase(userRepo)
	handler := httpHandler.NewUserHandler(userUseCase)

	fmt.Println(cfg.App)
	// Start the server
	router := httpHandler.NewRouter(cfg.App, handler, middleware)
	router.Serve()
	return nil
}
