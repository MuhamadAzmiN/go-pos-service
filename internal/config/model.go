package config

type Config struct {
	App        App
	Server     Server
	Jwt        Jwt
	Log        Log
	MongoDB    *MongoDB   // Optional
	PostgresDB *PostgresDB // Optional
	Redis      *Redis     // Optional
}

type App struct {
	Name    string
	Env     string // development | production | testing
	Version string
	Debug   bool
}

type Server struct {
	Host         string
	Port         string
	ReadTimeout  int // in seconds
	WriteTimeout int
}

type Log struct {
	Level    string // debug | info | warn | error
	FilePath string
}


type MongoDB struct {
	URI      string
	Database string
}

type PostgresDB struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
	SSLMode  string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	DB       int
}


type Jwt struct {
	Key string
	Exp int
}



