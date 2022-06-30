package database

//func TestMustConnectDB(t *testing.T) {
//	cfg := configs.DatabaseConfig{
//		Name:     "postgres",
//		Host:     "localhost",
//		Port:     5432,
//		Password: "temp123",
//	}
//
//	MustConnectDB(&cfg)
//
//	t.Run("db connect must failed", func(t *testing.T) {
//		cfg := configs.DatabaseConfig{
//			Name:     "postgres",
//			Host:     "ocalhost",
//			Port:     5432,
//			Password: "temp123",
//		}
//		err := MustConnectDB(&cfg)
//		require.NotNil(t, err)
//	})
//
//}
