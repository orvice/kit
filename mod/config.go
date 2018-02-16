package mod

type Mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

type InfluxDB struct {
	Addr     string
	Username string
	Password string
}

type Telegram struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}
