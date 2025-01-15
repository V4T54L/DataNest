package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	ServerPort     string
	DBURI          string
	DBMaxOpenConns int
	DBMaxIdleConns int
	DBMaxIdleTime  string
}

func getStr(key, fallback string) string {
	val := os.Getenv(key)

	if len(val) > 0 {
		return val
	}

	return fallback
}
func getInt(key string, fallback int) int {
	valStr := os.Getenv(key)

	if len(valStr) > 0 {
		val, err := strconv.Atoi(valStr)
		if err != nil {
			return val
		}
	}

	return fallback
}

func GetConfig() *Config {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		getStr("DB_USER", "postgres"), getStr("DB_PASSWORD", "8080"), getStr("DB_HOST", "localhost"), getStr("DB_PORT", "5432"), getStr("DB_DATABASE_NAME", "db"),
	)

	config := Config{
		ServerPort:     getStr("SERVER_PORT", "8080"),
		DBURI:          connStr,
		DBMaxOpenConns: getInt("DBMaxOpenConns", 30),
		DBMaxIdleConns: getInt("DBMaxIdleConns", 30),
		DBMaxIdleTime:  getStr("DBMaxIdleTime", "15m"),
	}

	return &config
}
