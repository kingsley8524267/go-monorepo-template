package config

type MyService struct {
	Common `mapstructure:",squash"`
}
