package main

import (
	"context"
	"fmt"

	"github.com/SoumyadipBhowmik/go-backend/driver"
)

var ctx = context.Background()

func init() {
	driver.LoadEnv()
}

func main() {

	conn := driver.ConnectToPostgresDB(ctx)
	err := conn.Ping(ctx)
	if err != nil {
		fmt.Print("couldn't establish connection to database ", err)
	} else {
		fmt.Print("connection established")
	}
	defer conn.Close()
}
