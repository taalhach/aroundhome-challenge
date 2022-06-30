package server

import (
	"fmt"
	"os"

	"github.com/taalhach/aroundhome-challennge/internal/server/models"

	ini "github.com/nanitor/goini"
	"github.com/taalhach/aroundhome-challennge/internal/server/configs"
	"github.com/taalhach/aroundhome-challennge/internal/server/database"
)

const (
	envKey = "SETTINGS"
)

var (
	MainConfigs *configs.MainConfig
)

func loadConfigs() (*configs.MainConfig, error) {
	file := os.Getenv(envKey)
	if file == "" {
		fmt.Printf("Missing env variable: %v", envKey)
		os.Exit(1)
	}

	dict, err := ini.Load(file)
	if err != nil {
		return nil, err
	}

	MainConfigs, err = configs.LoadMainConfig(dict)
	if err != nil {
		return nil, err
	}

	// must make database connection or panic
	database.MustConnectDB(MainConfigs.Database)

	// now apply migrations
	if err := database.Migrate(models.Material{}, models.Partner{}, models.PartnerMaterial{}); err != nil {
		fmt.Printf("db migrations failed, err: %v", err)
		os.Exit(1)
	}

	return MainConfigs, err
}
