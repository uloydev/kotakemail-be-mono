package appcontext

type AppContextKey string

const (
	// databasesKey is the key used to store the databases in the context.
	DatabasesKey AppContextKey = "databases"
	// LoggerKey is the key used to store the logger in the context.
	LoggerKey AppContextKey = "logger"
	// RequestIDKey is the key used to store the request ID in the context.
	RequestIDKey AppContextKey = "request_id"
	// UserKey is the key used to store the user in the context.
	UserKey AppContextKey = "user"
	// ConfigKey is the key used to store the config in the context.
	ConfigKey AppContextKey = "config"
	// TracerKey is the key used to store the tracer in the context.
	TracerKey AppContextKey = "tracer"
	// EnvinronmentKey is the key used to store the environment in the context.
	EnvironmentKey AppContextKey = "environment"
	// AppNameKey is the key used to store the app name in the context.
	AppNameKey AppContextKey = "app_name"
)
