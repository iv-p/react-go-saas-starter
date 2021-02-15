package config

// Config of the app
type Config struct {
	AppName    string `env:"APP_NAME"`
	AppVersion string `env:"APP_VERSION" envDefault:"dev"`
	Port       int    `env:"APP_PORT" envDefault:"8080"`

	Auth0Domain string `env:"AUTH0_DOMAIN"`
}
