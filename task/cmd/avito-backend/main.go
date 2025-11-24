package main

import (
	_ "avito-test/docs"
	"avito-test/internal/app"
)

// @title           Avito-Test-Task API
// @version         1.0
// @description     ## Overview
// @description     This API provides functionality for managing pull requests and automatically assigning random reviewers from teams.
// @description
// @description     ### Key Features:
// @description     - Create and manage pull requests
// @description     - Automatically assign 2 random reviewers from team members
// @description     - Manage team structures and user assignments
// @description     - Track pull request status and merge operations
// @description
// @description     ### Authentication
// @description     This API uses JWT Bearer token authentication. Admin tokens have elevated privileges.

// @host      localhost:8080
// @BasePath  /
// @schemes   http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	app := app.NewApp()
	app.Start()
}
