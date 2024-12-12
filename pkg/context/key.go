package appcontext

type AppContextKey string

const (
	// DatastoreKey is the key used to store the datastore in the context.
	DatastoreKey AppContextKey = "datastore"
	// LoggerKey is the key used to store the logger in the context.
	LoggerKey AppContextKey = "logger"
	// RequestIDKey is the key used to store the request ID in the context.
	RequestIDKey AppContextKey = "request_id"
	// UserKey is the key used to store the user in the context.
	UserKey AppContextKey = "user"
	// ConfigKey is the key used to store the config in the context.
	ConfigKey AppContextKey = "config"
)
