package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	xui "3xuibot/internal/xui"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf(".env на загружен %v", err)
	}

	baseURL := os.Getenv("URL")
	token := os.Getenv("TOKEN")
	if baseURL == "" || token == "" {
		log.Fatal("загрузка констант из .env")
	}

	api := xui.NewClientAPI(baseURL, token)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clients, err := api.ListClients(ctx)
	if err != nil {
		log.Fatalf("получение клиентов: %v", err)
	}
	for _, c := range clients {
		fmt.Println(c.Email, c.UUID)
	}

	client, err := api.ClientByEmail(ctx, "Test")
	if err != nil {
		log.Fatalf("получение клиента: %v", err)
	}
	fmt.Println(client)
}
