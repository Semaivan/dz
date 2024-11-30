package configmail

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Mail MailConfig
	Pass PassConfig
	Addr AddrConfig
}
type MailConfig struct {
	Mail string
}
type PassConfig struct {
	Pass string
}
type AddrConfig struct {
	Address string
}

func LoadConfig() *Config {
	currentWorkDirectory, _ := os.Getwd()
	rootPath := []byte(currentWorkDirectory)
	err := godotenv.Load(string(rootPath) + `\.env`)
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		Mail: MailConfig{
			Mail: os.Getenv("MAIL"),
		},
		Pass: PassConfig{
			Pass: os.Getenv("PASSWORD"),
		},
		Addr: AddrConfig{
			Address: os.Getenv("ADDR"),
		},
	}
}
