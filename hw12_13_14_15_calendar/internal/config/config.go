package config

type Config struct {
	Logger  LoggerConf
	Storage StorageConf
}

type LoggerConf struct {
	Level string
}

type StorageConf struct {
	Type     string
	Database DatabaseConf
}

type DatabaseConf struct {
	DirMigrate string
	Dialect    string
	Host       string
	Port       int
	Name       string
	User       string
	Password   string
}

func NewConfig() Config {
	return Config{}
}
