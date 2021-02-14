package config

type Config struct {
	AppName    string `env:"APP_NAME"`
	AppVersion string `env:"APP_VERSION" envDefault:"dev"`
	Port       int    `env:"port" envDefault:"80"`

	Auth0Domain   string `env:"AUTH0_DOMAIN"`
	Auth0Audience string `env:"AUTH0_AUDIENCE"`
}
