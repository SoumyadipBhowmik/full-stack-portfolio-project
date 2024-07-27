package utils

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
