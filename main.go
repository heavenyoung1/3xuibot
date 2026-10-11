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
		log.Printf("предупреждение: .env на загружен %v", err)
	}

	baseURL := os.Getenv("URL")
	token := os.Getenv("TOKEN")
	if baseURL == "" || token == "" {
		log.Fatal("ошибка загрузки констант из .env")
	}

	api := xui.NewClientAPI(baseURL, token)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clients, err := api.ListClients(ctx)
	if err != nil {
		log.Fatalf("ошибка получения клиентов %v", err)
	}
	for _, c := range clients {
		fmt.Println(c.Email, c.UUID)
	}
}
