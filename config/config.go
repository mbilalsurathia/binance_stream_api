package config

type Config struct {
	Name     string   `yaml:"name"`
	Version  string   `yaml:"version"`
	Mode     string   `yaml:"mode"`
	Host     string   `yaml:"host"`
	Port     string   `yaml:"port"`
	Domain   string   `yaml:"domain"`
	Logger   Logger   `yaml:"logger"`
	Postgres Postgres `yaml:"postgres"`
}

type Logger struct {
	LogEnvironment string `yaml:"logEnvironment"`
	Path           string `yaml:"path"`
}

type Postgres struct {
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Name     string `yaml:"Name"`
	SslMode  string `yaml:"SslMode"`
}
