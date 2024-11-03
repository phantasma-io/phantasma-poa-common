package configuration

import "strconv"

type DbConfiguration struct {
	DbName   string
	Host     string
	Port     int
	SslMode  string
	User     string
	Password string
}

func (c DbConfiguration) GetConnectionString() string {
	r := "dbname=" + c.DbName
	r += " host=" + c.Host
	r += " port=" + strconv.Itoa(c.Port)
	r += " sslmode=" + c.SslMode
	r += " user=" + c.User
	r += " password=" + c.Password

	return r
}
