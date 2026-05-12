package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort     string `mapstructure:"APP_PORT"`
	AppEnv      string `mapstructure:"APP_ENV"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	RedisURL    string `mapstructure:"REDIS_URL"`
	AppPrefork  bool   `mapstructure:"APP_PREFORK"`
	RateLimitMax        int           `mapstructure:"RATE_LIMIT_MAX"`
	RateLimitExpiration time.Duration `mapstructure:"RATE_LIMIT_EXPIRATION"`
	
	JWTSecret      string        `mapstructure:"JWT_SECRET"`
	AccessTokenTTL time.Duration `mapstructure:"ACCESS_TOKEN_TTL"`
	RefreshTokenTTL time.Duration `mapstructure:"REFRESH_TOKEN_TTL"`

	MinIOEndpoint  string `mapstructure:"MINIO_ENDPOINT"`
	MinIOAccessKey string `mapstructure:"MINIO_ACCESS_KEY"`
	MinIOSecretKey string `mapstructure:"MINIO_SECRET_KEY"`
	MinIOUseSSL    bool   `mapstructure:"MINIO_USE_SSL"`
	MinIOBucket    string `mapstructure:"MINIO_BUCKET"`
	StorageEncryptionKey string `mapstructure:"STORAGE_ENCRYPTION_KEY"`

	ElasticsearchURL string `mapstructure:"ELASTICSEARCH_URL"`

	OCRServiceURL string `mapstructure:"OCR_SERVICE_URL"`
	OCRServiceUser string `mapstructure:"OCR_SERVICE_USER"`
	OCRServicePass string `mapstructure:"OCR_SERVICE_PASS"`

	GeminiAPIKey string `mapstructure:"GEMINI_API_KEY"`
	OpenAIAPIKey string `mapstructure:"OPENAI_API_KEY"`
	OllamaURL    string `mapstructure:"OLLAMA_URL"`
	AIRefinementEnabled bool `mapstructure:"AI_REFINEMENT_ENABLED"`
	AIMetadataEnabled   bool `mapstructure:"AI_METADATA_ENABLED"`

	SMTPHost string `mapstructure:"SMTP_HOST"`
	SMTPPort int    `mapstructure:"SMTP_PORT"`
	SMTPUser string `mapstructure:"SMTP_USER"`
	SMTPPass string `mapstructure:"SMTP_PASS"`

	WAGatewayURL string `mapstructure:"WA_GATEWAY_URL"`
	WAApiKey     string `mapstructure:"WA_API_KEY"`

	LDAPEnabled  bool   `mapstructure:"LDAP_ENABLED"`
	LDAPURL      string `mapstructure:"LDAP_URL"`
	LDAPBaseDN   string `mapstructure:"LDAP_BASE_DN"`
	LDAPBindDN   string `mapstructure:"LDAP_BIND_DN"`
	LDAPBindPass string `mapstructure:"LDAP_BIND_PASS"`
	LDAPUserFilter string `mapstructure:"LDAP_USER_FILTER"`
	LDAPSimulation bool   `mapstructure:"LDAP_SIMULATION"`

	HTTPSEnabled  bool   `mapstructure:"HTTPS_ENABLED"`
	HTTPSCertFile string `mapstructure:"HTTPS_CERT_FILE"`
	HTTPSKeyFile  string `mapstructure:"HTTPS_KEY_FILE"`

	WorkerConcurrency int `mapstructure:"WORKER_CONCURRENCY"`
	CORSAllowedOrigins string `mapstructure:"CORS_ALLOWED_ORIGINS"`
	AppURL             string `mapstructure:"APP_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	
	// Set defaults
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("REDIS_URL", "127.0.0.1:6379")
	viper.SetDefault("WORKER_CONCURRENCY", 5)
	viper.SetDefault("CORS_ALLOWED_ORIGINS", "http://localhost:3000,http://127.0.0.1:3000")
	
	// Default to .env
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// Try to read the default .env file
	_ = viper.ReadInConfig()

	// If APP_ENV is set, try to load that specific file (e.g., .env.staging)
	env := viper.GetString("APP_ENV")
	if env != "" && env != "development" {
		viper.SetConfigName(".env." + env)
		err = viper.MergeInConfig()
		if err != nil {
			log.Printf("Warning: Could not find .env.%s file, using defaults\n", env)
		}
	}

	err = viper.Unmarshal(&config)
	return
}
