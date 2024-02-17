package config

type AppConfig struct {
	Config GeneralConfig
}

type GeneralConfig struct {
	AppPort     string
	Environment string
}
