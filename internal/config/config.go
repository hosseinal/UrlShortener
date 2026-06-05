package config

type SQLDatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type RedisCacheConfig struct {
	Host     string
	Port     int
	Password string
}

func NewSQLDatabaseConfig(host string, port int, user string, password string, dbName string) SQLDatabaseConfig {
	return SQLDatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}
}

func NewRedisCacheConfig(host string, port int, password string) RedisCacheConfig {
	return RedisCacheConfig{
		Host:     host,
		Port:     port,
		Password: password,
	}
}
