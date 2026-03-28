package configs

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Env struct {
	Env            string
	Port           string
	DatabasePath   string
	RedisHost      string
	RedisPassword  string
	RedisDatabase  int
	JWTExpTime     int
	JWTSecret      string
	UserSessionTTL int
}

func LoadEnv() Env {
	err := loadEnvFromFile("./configs/.env")
	if err != nil {
		slog.Error("failed to read env", "error", err)
	}

	redisDb, _ := strconv.Atoi(getEnv("REDIS_DATABASE", "0"))
	jwtExpTime, _ := strconv.Atoi(getEnv("JWT_EXP_TIME_SECOND", "900"))
	userSessionTTL, _ := strconv.Atoi(getEnv("USER_SESSION_TTL", "900"))

	return Env{
		Env:            getEnv("ENV", "dev"),
		Port:           getEnv("PORT", "8080"),
		DatabasePath:   getEnv("DB_PATH", ".linkvault.db"),
		RedisHost:      getEnv("REDIS_HOST", "localhots:6379"),
		RedisDatabase:  redisDb,
		RedisPassword:  getEnv("REDIS_PASSWORD", ""),
		JWTExpTime:     jwtExpTime,
		JWTSecret:      getEnv("JWT_SECRET_KEY", "Js0nW3bT0K3n"),
		UserSessionTTL: userSessionTTL,
	}
}

// Function to get env value
// this function accept two argument, which key for env key
// and the fallback value if the env is not exist
func getEnv(key string, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}

// Function to read .env file, this function accept file path as argument
// and the env key format should be like this:
// SERVER_PORT=8080
// API_CONTEXT=/api
func loadEnvFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, val, found := strings.Cut(line, "=")
		if !found {
			continue
		}

		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)

		val = strings.Trim(val, `"'`)

		if os.Getenv(key) == "" {
			os.Setenv(key, val)
		}
	}

	return scanner.Err()
}
