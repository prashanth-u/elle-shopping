package config

type DatabaseConfig struct {
    DBUser     string `json:"db_user"`
    DBPassword string `json:"db_password"`
    DBHost     string `json:"db_host"`
    DBName     string `json:"db_name"`
    DBPort int `json:"db_port"`
}