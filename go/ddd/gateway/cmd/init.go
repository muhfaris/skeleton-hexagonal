package cmd

import (
	"fmt"
	"time"

	"github.com/muhfaris/lib-go/psql"
	"github.com/muhfaris/redigo/redis"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	cfgFile   = ""
	cachePool *redis.Pool
)

// splash print plain text message to console
func splash() {
	fmt.Println(`
  __    ___   __   ___   ___   ___   ___  _____
 / /\  | | \ ( ( | |_) / / \ | |_) / / \  | |
/_/--\ |_|_/ _)_) |_| \ \_\_/ |_|_) \_\_/  |_|
		`)
}

func initconfig() {
	viper.SetConfigType("toml")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// search config in home directory with name "config" (without extension)
		viper.AddConfigPath("./configs")
		viper.SetConfigName("config")
	}

	//read env
	viper.AutomaticEnv()

	// if a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Config application:", err)
	}

	log.Println("using config file:", viper.ConfigFileUsed())
}

func initDatabase() {
	dbOptions := psql.DBOptions{
		Host:     viper.GetString("persistence.database.host"),
		Port:     viper.GetInt("persistence.database.port"),
		Username: viper.GetString("persistence.database.username"),
		Password: viper.GetString("persistence.database.password"),
		DBName:   viper.GetString("persistence.database.name"),
		SSLMode:  viper.GetString("persistence.database.ssl_mode"),
	}

	conn, err := psql.Connect(&dbOptions)
	if err != nil {
		log.Fatalln("Database:", err)
	}

	log.Println("Database connected ...")
	dbPool = conn
}

// initCache initiate cache pool object from previously read config.
func initCache() {
	host := viper.GetString("persistence.redis.host")
	port := viper.GetInt("persistence.redis.port")
	password := viper.GetString("persistence.redis.password")
	maxIdle := viper.GetInt("persistence.redis.max_idle")
	idleTimeout := viper.GetInt("persistence.redis.idle_timeout")
	database := viper.GetInt("persistence.redis.database.index")
	maxConnLifetime := viper.GetInt("persistence.redis.max_conn_lifetime")

	if cachePool == nil {
		cachePool = &redis.Pool{
			MaxIdle:     maxIdle,
			IdleTimeout: time.Duration(idleTimeout) * time.Second,
			Dial: func() (redis.Conn, error) {
				address := fmt.Sprintf("%s:%d", host, port)
				c, err := redis.Dial(
					"tcp",
					address,
					redis.DialPassword(password),
				)
				if err != nil {
					return nil, err
				}

				// Do authentication process if password not empty.
				if password != "" {
					if _, err := c.Do("AUTH", password); err != nil {
						c.Close()
						return nil, err
					}
				}

				if database != 0 {
					if _, err := c.Do("SELECT", database); err != nil {
						c.Close()
						return nil, err
					}
				}

				return c, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}

				_, err := c.Do("PING")
				return err
			},
			Wait:            true,
			MaxConnLifetime: time.Duration(maxConnLifetime) * time.Minute,
		}

		return
	}
}
