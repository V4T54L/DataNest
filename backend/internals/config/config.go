package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerAddr     string
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
	config := Config{
		ServerAddr:     getStr("SERVER_ADDR", ":8080"),
		DBURI:          getStr("DB_URI", "postgres://user:[email protected]:port/dbname?sslmode=disable"),
		DBMaxOpenConns: getInt("DBMaxOpenConns", 30),
		DBMaxIdleConns: getInt("DBMaxIdleConns", 30),
		DBMaxIdleTime:  getStr("DBMaxIdleTime", "15m"),
	}

	return &config
}
