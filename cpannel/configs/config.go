package configs

// ApplicationConfig 程序配置
type ApplicationConfig struct {
	ENV  string       `json:"env"`
	Host string       `json:"host"`
	Port string       `json:"port"`
	DB   *MySQLConfig `json:"db"`
}
