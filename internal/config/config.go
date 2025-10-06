// sentiric-vertical-public-service/internal/config/config.go
package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	GRPCPort string
	HttpPort string
	CertPath string
	KeyPath  string
	CaPath   string
	LogLevel string
	Env      string

	// Kamu servisi bağımlılıkları (Placeholder)
	PublicSectorEndpoint string
	PublicSectorAPIKey   string
}

func Load() (*Config, error) {
	godotenv.Load()

	// Harmonik Mimari Portlar (Dikey Servisler, 205XX bloğu atandı)
	return &Config{
		GRPCPort: GetEnv("VERTICAL_PUBLIC_SERVICE_GRPC_PORT", "20511"),
		HttpPort: GetEnv("VERTICAL_PUBLIC_SERVICE_HTTP_PORT", "20510"),

		CertPath: GetEnvOrFail("VERTICAL_PUBLIC_SERVICE_CERT_PATH"),
		KeyPath:  GetEnvOrFail("VERTICAL_PUBLIC_SERVICE_KEY_PATH"),
		CaPath:   GetEnvOrFail("GRPC_TLS_CA_PATH"),
		LogLevel: GetEnv("LOG_LEVEL", "info"),
		Env:      GetEnv("ENV", "production"),

		PublicSectorEndpoint: GetEnv("PUBLIC_SECTOR_ENDPOINT", "mock_gov_api"),
		PublicSectorAPIKey:   GetEnv("PUBLIC_SECTOR_API_KEY", ""),
	}, nil
}

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func GetEnvOrFail(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal().Str("variable", key).Msg("Gerekli ortam değişkeni tanımlı değil")
	}
	return value
}
