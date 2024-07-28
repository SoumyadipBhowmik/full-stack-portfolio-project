package utils

import "github.com/joho/godotenv"

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func CommonErrorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func Checkport(port string) string {
	if port == "" {
		port = "3000"
	}
	return port
}
