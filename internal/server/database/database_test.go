package database

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taalhach/aroundhome-challennge/internal/server/configs"
)

func TestMustConnectDB(t *testing.T) {
	t.Run("must db connect", func(t *testing.T) {
		cfg := configs.DatabaseConfig{
			Name:     "postgres",
			Host:     "localhost",
			Port:     5434,
			Password: "postgres",
			User:     "postgres",
		}
		err := ConnectDB(&cfg)
		require.Nil(t, err)
	})

	t.Run("db connect must failed", func(t *testing.T) {
		cfg := configs.DatabaseConfig{
			Name:     "postgres",
			Host:     "localhost",
			Port:     5434,
			Password: "pqstgres",
			User:     "postgres",
		}
		err := ConnectDB(&cfg)
		require.NotNil(t, err)
	})

}
