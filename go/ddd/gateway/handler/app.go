package handler

import (
	"database/sql"

	"github.com/muhfaris/redigo/redis"
	"github.com/spf13/viper"

	mw "github.com/muhfaris/adsrobot/gateway/middleware"
	"github.com/muhfaris/adsrobot/internal/logger"
	"github.com/muhfaris/request"
)

// App is
type App struct {
	Config     Config
	Middleware mw.Middleware

	Connections ConnectionApp
}

// ConnectionApp  is connection source
type ConnectionApp struct {
	DB    *sql.DB
	Redis *redis.Pool
}

// Config application
type Config struct {
	Name        string
	Port        int
	Environment string

	APIs  APIsConfig
	Auth  AuthConfig
	Cache CacheConfig
	HTTP  HTTPConfig
	User  UserConfig

	// another config
	standardLog *logger.StandardLog
	request     *request.ReqApp
}

// APIsConfig is all api config
type APIsConfig struct {
	Facebook string
}

// AuthConfig is config for auth
type AuthConfig struct {
	SecretKey string
}

// CacheConfig is config for caching
type CacheConfig struct {
	UserKey string
}

// ChangeUserKey is change worker delete key
func (cc *CacheConfig) ChangeUserKey(key string) {
	if key != "" {
		cc.UserKey = key
	}
}

// UserConfig is user data
type UserConfig struct {
	Username string
	Role     string
	Token    string
	FBToken  string
}

type HTTPConfig struct {
	ReadTimeout int
}

// NewApp create handler
func NewApp(db *sql.DB, redis *redis.Pool) *App {
	// logger
	appName := viper.GetString("app.name")
	stdLogger := logger.NewLogger(appName)
	requestApp := &request.ReqApp{
		ContentType: request.MimeTypeJSON,
	}

	app := &App{
		Config: Config{
			Name:        viper.GetString("app.name"),
			Port:        viper.GetInt("app.port"),
			Environment: viper.GetString("app.env"),
			APIs: APIsConfig{
				Facebook: viper.GetString("app.apis.facebook"),
			},
			Auth: AuthConfig{
				SecretKey: viper.GetString("app.auth.secret_key"),
			},
			HTTP: HTTPConfig{
				ReadTimeout: viper.GetInt("app.http.read_timeout"),
			},
			standardLog: stdLogger,
			request:     requestApp,
		},
		Connections: ConnectionApp{
			DB:    db,
			Redis: redis,
		},
		Middleware: mw.Middleware{
			APIs: mw.APIsConfig{
				User: viper.GetString("app.apis.user"),
				Cron: viper.GetString("app.apis.cron"),
			},
			ManagixInternalKey: viper.GetString("app.apis.internal_key"),
		},
	}

	return app
}
