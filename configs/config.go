package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)
import _ "github.com/go-chi/jwtauth"

type conf struct {
	Driver          string `mapstructure:"DB_DRIVER"`
	Host            string `mapstructure:"DB_HOST"`
	Port            string `mapstructure:"DB_PORT"`
	User            string `mapstructure:"DB_USER"`
	Password        string `mapstructure:"DB_PASSWORD"`
	Name            string `mapstructure:"DB_NAME"`
	MigrationTag    bool   `mapstructure:"MIGRATION_TAG"`
	MigrationSource string `mapstructure:"MIGRATION_SOURCE"`
	WebServerPort   string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret       string `mapstructure:"JWT_SECRET"`
	JWTExpiration   int    `mapstructure:"JWT_EXPIRATION"`
	TokenAuth       *jwtauth.JWTAuth
	Test            bool `mapstructure:"TEST"`
}

type Config interface {
	GetDriver() string
	GetHost() string
	GetPort() string
	GetUser() string
	GetPassword() string
	GetName() string
	GetMigrationTag() bool
	GetMigrationSource() string
	GetWebServerPort() string
	GetJWTSecret() string
	GetJWTExpiration() int
	GetTokenAuth() *jwtauth.JWTAuth
	GetTest() bool
}

func LoadConfig(filePath string) (*Config, error) {
	conf, err := loadConfig(filePath)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}

func loadConfig(path string) (Config, error) {
	var cfg *conf

	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}

func (c *conf) GetDriver() string {
	return c.Driver
}

func (c *conf) GetHost() string {
	return c.Host
}

func (c *conf) GetPort() string {
	return c.Port
}

func (c *conf) GetUser() string {
	return c.User
}

func (c *conf) GetPassword() string {
	return c.Password
}

func (c *conf) GetName() string {
	return c.Name
}

func (c *conf) GetMigrationTag() bool {
	return c.MigrationTag
}

func (c *conf) GetMigrationSource() string {
	return c.MigrationSource
}

func (c *conf) GetWebServerPort() string {
	return c.WebServerPort
}

func (c *conf) GetJWTSecret() string {
	return c.JWTSecret
}

func (c *conf) GetJWTExpiration() int {
	return c.JWTExpiration
}

func (c *conf) GetTokenAuth() *jwtauth.JWTAuth {
	return c.TokenAuth
}

func (c *conf) GetTest() bool {
	return c.Test
}
