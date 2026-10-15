package xui

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var ErrClientNotFound = errors.New("клиент не найден")

const msgRecordNotFound = "record not found"

// Обертка для работы с клиентами по API 3X-UI
type ClientAPI struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

// Создание нового клиента API
func NewClientAPI(baseURL, token string) *ClientAPI {
	return &ClientAPI{
		baseURL:    baseURL,
		token:      token,
		httpClient: &http.Client{Timeout: 15 * time.Second},
	}
}

// Общий шаблон запроса, используемый в вызове эндпоинтов
func (c *ClientAPI) do(ctx context.Context, method, path string, payload []byte) ([]byte, error) {
	fullURL := c.baseURL + path

	var payloadReader io.Reader
	if payload != nil {
		payloadReader = bytes.NewReader(payload)
	}
	request, err := http.NewRequestWithContext(ctx, method, fullURL, payloadReader)
	if err != nil {
		return nil, fmt.Errorf("создание запроса %s: %w", fullURL, err)
	}
	request.Header.Set("Authorization", "Bearer "+c.token)
	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("выполнение запроса %s: %w", fullURL, err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("чтение ответа %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: HTTP %d: %s", path, response.StatusCode, body)
	}

	return body, nil
}

// Получить список клиентов
func (c *ClientAPI) ListClients(ctx context.Context) ([]Client, error) {
	path := "/panel/api/clients/list"

	body, err := c.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var resp ClientsResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("разбор JSON %w, body=%s", err, body)
	}
	if !resp.Success {
		return nil, fmt.Errorf("панель вернула ошибку: %s", resp.Message)
	}
	return resp.Object, nil
}

// Получить конкретного клиента по <email>
func (c *ClientAPI) ClientByEmail(ctx context.Context, email string) (*Client, error) {
	path := "/panel/api/clients/get/" + url.PathEscape(email)

	body, err := c.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var resp ClientResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("разбор JSON выполнен некорректно %w, body=%s", err, body)
	}
	//log.Println(string(body))
	if !resp.Success {
		if strings.Contains(resp.Message, msgRecordNotFound) {
			return nil, ErrClientNotFound
		}
		return nil, fmt.Errorf("панель вернула ошибку: %s", resp.Message)
	}
	return &resp.Object.Client, nil
}
