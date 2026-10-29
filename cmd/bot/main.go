package main

import (
	"3xuibot/internal/app"
	"3xuibot/internal/config"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := run(); err != nil {
		slog.Error("application stopped with error", "error", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	application, err := app.New(cfg)
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	slog.Info("starting application")

	application.Run(ctx)

	slog.Info("application stopped")

	return nil
}

// func main() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Printf(".env на загружен %v", err)
// 	}

// 	baseURL := os.Getenv("URL")
// 	token := os.Getenv("TOKEN")
// 	if baseURL == "" || token == "" {
// 		log.Fatal("загрузка констант из .env")
// 	}

// 	api := xui.NewClientAPI(baseURL, token)

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	clients, err := api.ListClients(ctx)
// 	if err != nil {
// 		log.Fatalf("получение клиентов: %v", err)
// 	}
// 	for _, c := range clients {
// 		fmt.Println(c.Email, c.UUID)
// 	}

// 	email := "AntonY"
// 	client, err := api.ClientByEmail(ctx, email)
// 	switch {
// 	case errors.Is(err, xui.ErrClientNotFound):
// 		log.Printf("клиент %q не найден", email)
// 	case err != nil:
// 		log.Fatalf("получение клиента %v", err)
// 	default:
// 		log.Printf("клиент: %+v", client)
// 	}

// TODO пока что поиск по TgID не работает
// clientsTgId, err := api.ClientByTgId(ctx, 1234567891)
// if err != nil {
// 	log.Fatalf("получение клиента по tgid: %v", err)
// }
// fmt.Println(clientsTgId)
// }
