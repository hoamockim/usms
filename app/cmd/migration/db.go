package migration

import (
	"fmt"
	"time"
)

type MySQL struct {
	Username string `envconfig:"MYSQL_USER" default:"root"`
	Password string `envconfig:"MYSQL_PASS" default:"Admin@1234"`
	Host     string `envconfig:"MYSQL_HOST" default:"localhost"`
	Port     int    `envconfig:"MYSQL_PORT" default:"3306"`
	Database string `envconfig:"MYSQL_DB" default:"reward_intelligent"`
}

func (c *MySQL) ConnectionString() string {
	format := "%v:%v@tcp(%v:%v)/%v?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	connection := fmt.Sprintf(format, c.Username, c.Password, c.Host, c.Port, c.Database)
	return connection
}

type JobStatement struct {
	Name   string `json:"name"`
	Sql    string `json:"sql"`
	Status string `json:"status"`
}

//Read from to json
type JobMigrate struct {
	JobId    string          `json:"job_id,omitempty"`
	Migrates []*JobStatement `json:"migrates"`
}

//run sql & save into table migration
type Migration struct {
	JobId       string
	Name        string
	IdMigration int
	Statements  string // sql query
	Status      string
	StartTime   time.Time
	EndTime     time.Time
}
