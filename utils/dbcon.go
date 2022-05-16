package utils

type DBConfig struct {
	Host     string `yaml:"mysql_host"`
	Port     int    `yaml:"mysql_port"`
	User     string `yaml:"mysql_user"`
	Password string `yaml:"mysql_pwd"`
	DBName   string `yaml:"mysql_db"`
}
