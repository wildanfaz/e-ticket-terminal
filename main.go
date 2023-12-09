package main

import (
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/wildanfaz/e-ticket-terminal/cmd"
)

func main() {
	cmd.InitCmd(context.Background())
}
