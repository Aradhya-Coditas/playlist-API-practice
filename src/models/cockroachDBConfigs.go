package models

type CockroachDBConfig struct {
	Host             string
	Port             string
	User             string
	Password         string
	DBName           string
	SSLMode          string
	MaxIdleConns     int
	MaxOpenConns     int
	ConnMaxLifetime  uint16
	TimeZone         string
	IsMockConnection bool
}
