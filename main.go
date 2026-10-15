package main

import (
	"context"
	"errors"
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

	email := "AntonY"
	client, err := api.ClientByEmail(ctx, email)
	switch {
	case errors.Is(err, xui.ErrClientNotFound):
		log.Printf("клиент %q не найден", email)
	case err != nil:
		log.Fatalf("получение клиента %v", err)
	default:
		log.Printf("клиент: %+v", client)
	}

	// TODO пока что поиск по TgID не работает
	// clientsTgId, err := api.ClientByTgId(ctx, 1234567891)
	// if err != nil {
	// 	log.Fatalf("получение клиента по tgid: %v", err)
	// }
	// fmt.Println(clientsTgId)
}
