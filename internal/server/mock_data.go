package server

import (
	"fmt"
	"os"

	"github.com/taalhach/aroundhome-challennge/internal/server/database"
	"github.com/taalhach/aroundhome-challennge/internal/server/manager"

	"github.com/spf13/cobra"
	_ "github.com/taalhach/aroundhome-challennge/docs"
)

const (
	dataSourceEnvKey = "DATA_SOURCE"
)

var mockData = &cobra.Command{
	Use:                   "mock_data",
	Short:                 fmt.Sprintf("servers api on %v port", port),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		// check if data source path env variable is set
		file := os.Getenv(dataSourceEnvKey)
		if file == "" {
			fmt.Printf("Missing env variable: %v", dataSourceEnvKey)
			os.Exit(1)
		}

		if _, err := os.Stat(file); err != nil {
			fmt.Printf("failed to file file: %s", file)
			os.Exit(1)
		}

		_, err := loadConfigs()
		if err != nil {
			fmt.Printf("failed to load configs: %v", err)
			os.Exit(1)
		}

		manager.InitMockDataBuilder(&database.Db, file)
	},
}

func init() {
	rootCommand.AddCommand(mockData)
	//	add flags
	rootCommand.PersistentFlags().Bool("reset.db", true, "reset database")
}
