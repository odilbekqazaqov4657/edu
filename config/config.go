package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	DebugMode   string = "debug"
	TestMode    string = "test"
	ReleaseMode string = "release"
)

type GeneralConfig struct {
	AppName     string
	Environment string //debug
	Version     string
	HTTPPort    string
	HTTPScheme  string
	SignInKey   string
}

type PgConfig struct {
	Username     string
	Password     string
	Host         string
	Port         int
	DatabaseName string
}

type RedisConfig struct {
	Username     string
	Password     string
	Host         string
	Port         int
	DatabaseName int
}

type Config struct {
	GeneralConfig GeneralConfig
	PgConfig      PgConfig
	RedisConfig   RedisConfig
}

func NewConfig() Config {
	return Config{GeneralConfig: GeneralConfig{AppName: "my_blogs_site_back_end"}}
}

func Load() Config {

	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found !")
	}

	var config = NewConfig()

	// general configs
	config.GeneralConfig.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.GeneralConfig.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))
	config.GeneralConfig.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8080"))
	config.GeneralConfig.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))
	config.GeneralConfig.SignInKey = cast.ToString(getOrReturnDefaultValue("SIGN_IN_KEY", "sdfghjkdfghjk4567uhgf"))

	// postgres config
	config.PgConfig.Username = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "superuser"))
	config.PgConfig.Password = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "superuser"))
	config.PgConfig.Host = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PgConfig.Port = cast.ToInt(getOrReturnDefaultValue("pOSTGRES_PORT", 5432))
	config.PgConfig.DatabaseName = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", config.GeneralConfig.AppName))

	// redis config
	config.RedisConfig.Host = cast.ToString(getOrReturnDefaultValue("REDIS_HOST", "localhost"))
	config.RedisConfig.Port = cast.ToInt(getOrReturnDefaultValue("REDIS_PORT", 6379))

	return config

}

// URL example:  "postgres://username:password@localhost:5432/database_name"
func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {

	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
