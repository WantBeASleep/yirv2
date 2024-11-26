package config

type Config struct {
	App App `yaml:"app"`
	DB  DB  `yaml:"db"`
}

type App struct {
	Url string `yaml:"url" env:"url" env-default:"localhost:50060"`
	// Host string `yaml:"host" env:"HOST" env-default:"localhost"`
	// Port string `yaml:"grpc" env:"PORT" env-default:"50060"`
}

type DB struct {
	Dsn string `yaml:"dsn" env:"DB_DSN" env-required:"true"`
}
