package configs

import (
	"bufio"
	"os"
	"strings"
)

type Config struct {
	Env  string
	Port string
}

func Load() Config {
	_ = loadEnvFromFile(".env")
	return Config{
		Env:  getEnv("ENV", "dev"),
		Port: getEnv("PORT", "8080"),
	}
}

func getEnv(key string, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}

/*
 * Function to read .env file
 * the env key format should be like this:
 * SERVER_PORT=8080
 * API_CONTEXT=/api
 */
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
