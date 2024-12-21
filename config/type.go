package config

type DatabaseType string

const (
	DB_MONGO    DatabaseType = "MONGODB"
	DB_POSTGRES DatabaseType = "POSTGRES"
	DB_MYSQL    DatabaseType = "MYSQL"
)

type DatabaseConfig struct {
	Host     string       `mapstructure:"host"`
	Port     int          `mapstructure:"port"`
	Username string       `mapstructure:"username"`
	Password string       `mapstructure:"password"`
	Name     string       `mapstructure:"name"`
	Type     DatabaseType `mapstructure:"type"`
	Database string       `mapstructure:"database"`
}

type (
	LogLevel  string
	LogOutput string
)

const (
	LOG_TRACE   LogLevel = "TRACE"
	LOG_DEBUG   LogLevel = "DEBUG"
	LOG_INFO    LogLevel = "INFO"
	LOG_WARNING LogLevel = "WARNING"
	LOG_ERROR   LogLevel = "ERROR"
	LOG_FATAL   LogLevel = "FATAL"

	LOG_OUTPUT_CONSOLE LogOutput = "CONSOLE"
	LOG_OUTPUT_FILE    LogOutput = "FILE"
)

type LoggingConfig struct {
	Level      LogLevel  `mapstructure:"level"`
	Output     LogOutput `mapstructure:"output"`
	File       string    `mapstructure:"file"`
	MaxSize    int       `mapstructure:"max_size"`
	MaxAge     int       `mapstructure:"max_age"`
	MaxBackups int       `mapstructure:"max_backups"`
	Compress   bool      `mapstructure:"compress"`
}

type StorageType string

const (
	STORAGE_LOCAL = "LOCAL"
	STORAGE_S3    = "S3"
)

type StorageConfig struct {
	Type      string `mapstructure:"type"`
	BasePath  string `mapstructure:"base_path, omitempty"`
	EndPoint  string `mapstructure:"end_point,omitempty"`
	AccessKey string `mapstructure:"access_key,omitempty"`
	SecretKey string `mapstructure:"secret_key,omitempty"`
	Bucket    string `mapstructure:"bucket,omitempty"`
}
