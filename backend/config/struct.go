package config

type config struct {
	App      App
	Server   Server
	Database Database
}

type App struct {
	Page      int
	PageSize  int
	JwtSecret string
}

type Server struct {
	AppMode      string
	Port         int
	ReadTimeout  int
	WriteTimeout int
	CorsOrigin   string
}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Dbname      string
	TablePrefix string
}
