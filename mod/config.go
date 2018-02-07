package mod

type InfluxDB struct {
	Addr     string
	Username string
	Password string
}

type Mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}
