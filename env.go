package konggo

import "os"

func GetEnvOrDefault(key string, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}

	return result
}
