package config

import "fmt"

type DBConfig struct {
	host     string
	port     int
	name     string
	user     string
	password string
}

func (dc DBConfig) ConnString() string {
	//return fmt.Sprintf("dbname=%s user=%s password='%s' host=%s sslmode=disable", dc.name, dc.user, dc.password, dc.host)
	return fmt.Sprintf("%s://%s%s:%d/%s?sslmode=disable", dc.user, dc.password, dc.host, dc.port, dc.name)
}
