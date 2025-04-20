package main

import (
	"context"
	"grpc/handler"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	ctx := context.TODO()
	if len(args) > 0 {
		switch args[0] {
		case "http":
			handler.RunHttp(ctx)
		case "grpc":
			handler.RunGrpc(ctx)
		}
	} else {
		log.Println("Usage: go run main.go [http|grpc]")
	}
}
