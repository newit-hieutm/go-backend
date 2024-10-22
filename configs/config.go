package configs

type Config struct {
	Server
	Db
	Security
}

type Server struct {
	Host string
	Port int
}

type Db struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Security struct {
	Jwt
}

type Jwt struct {
	TokenSecret string `mapstructure:"token_secret"`
}
