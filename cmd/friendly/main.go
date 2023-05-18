package main

import (
	"flag"
	"friendly/internal/app"
)

// @title My Gin API
// @description This is a sample API for Gin framework.
// @version 1.0
// @host localhost:8080
// @BasePath /api/v1
func main() {
	configPath := flag.String("configPath", "configs/friendlyConfig-dev.yaml", "Path to config file")
	firebaseSecretPath := flag.String("firebaseSecretPath", "configs/firebaseSecret.json", "Path to firebase secret file")
	pathsConfigFile := flag.String("csvConfigPath", "configs/pathsConfig.csv", "csv config for auth middleware")

	app.Start(*configPath, *firebaseSecretPath, *pathsConfigFile)
}
