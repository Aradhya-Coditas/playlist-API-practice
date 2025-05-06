package models

type OpenSearchConfig struct {
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	User             string `yaml:"user"`
	Password         string `yaml:"password"`
	SSLMode          string `yaml:"sslMode"`
	TimeZone         string `yaml:"TimeZone"`
	IsMockConnection bool   `yaml:"isMockConnection"`
	CACertPath       string `yaml:"caCertPath"`
}
