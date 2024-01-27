package configs

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

type AuthConfig struct {
	SHASalt       string
	JWTSigningKey string
	JWTTokenTTL   int
}
