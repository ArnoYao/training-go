package dao

type DBConfig struct {
	DBPath string
}

func DefaultDBConfig() *DBConfig {
	return &DBConfig{
		DBPath: "./02week/db/01.db",
	}
}
