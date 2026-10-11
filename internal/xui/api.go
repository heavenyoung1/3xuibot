package xui

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Обертка для работы с клиентами по API 3X-UI
type ClientAPI struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

func NewClientAPI(baseURL, token string) *ClientAPI {
	return &ClientAPI{
		baseURL:    baseURL,
		token:      token,
		httpClient: &http.Client{Timeout: 15 * time.Second},
	}
}

func (c *ClientAPI) do(ctx context.Context, method, path string) ([]byte, error) {
	url := c.baseURL + path
	request, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания запроса %s: %w", url, err)
	}
	request.Header.Set("Authorization", "Bearer "+c.token)
	//request.Header.Set("Accept", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("Ошибка выполнения запроса %s: %w", url, err)
	}
	log.Printf("запрос %v, ошибка %d", response.Body, response.StatusCode)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения ответа %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: HTTP %d: %s", path, response.StatusCode, body)
	}

	return body, nil
}

func (c *ClientAPI) ListClients(ctx context.Context) ([]Client, error) {
	body, err := c.do(ctx, http.MethodGet, "/panel/api/clients/list")
	if err != nil {
		return nil, err
	}

	var clients Clients
	if err := json.Unmarshal(body, &clients); err != nil {
		return nil, fmt.Errorf("разбор JSON выполнен некорректно %w, %s", err, body)
	}
	if clients.Success != true {
		return nil, fmt.Errorf("панель вернула ошибку: %s", clients.Message)
	}
	return clients.Objects, nil

}
