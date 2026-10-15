package xui

// ――― Клиент ―――
type Client struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	SubID        string `json:"subId"`
	UUID         string `json:"uuid"`
	Password     string `json:"password"`
	Auth         string `json:"auth"`
	Flow         string `json:"flow"`
	Security     string `json:"security"`
	PrivateKey   string `json:"privateKey"`
	PublicKey    string `json:"publicKey"`
	AllowedIPs   string `json:"allowedIPs"`
	PreSharedKey string `json:"preSharedKey"`
	KeepAlive    int    `json:"keepAlive"`
	Secret       string `json:"secret"`
	AdTag        string `json:"adTag"`
	LimitIP      int    `json:"limitIp"`
	TotalGB      int64  `json:"totalGB"`
	ExpiryTime   int64  `json:"expiryTime"`
	Enable       bool   `json:"enable"`
	TgID         int64  `json:"tgId"`
	Group        string `json:"group"`
	Comment      string `json:"comment"`
	Reset        int    `json:"reset"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	Reverse      any    `json:"reverse"`
	InboundIDs   []int  `json:"inboundIds"`
}

// Трафик клиента
type Traffic struct {
	ID     int   `json:"id"`
	Enable bool  `json:"enable"`
	Up     int64 `json:"up"`
	Down   int64 `json:"down"`
}

// ――― Ответ сервера - Получить всех клиентов ―――
type ClientsResponse struct {
	Success bool     `json:"success"`
	Message string   `json:"msg"`
	Object  []Client `json:"obj"`
}

// ――――――――――――――――――――――――――――――――――――――――――――――

// ――― Структуры для получения клиента ―――
type ExternalLink struct {
	ID        int    `json:"id"`
	ClientID  int    `json:"clientId"`
	Kind      string `json:"kind"`
	Value     string `json:"value"`
	Remark    string `json:"remark"`
	SortIndex int    `json:"sortIndex"`
	CreatedAt int64  `json:"createdAt"`
}

type ClientDetails struct {
	Client        Client         `json:"client"`
	ExternalLinks []ExternalLink `json:"externalLinks"`
	InboundIDs    []int          `json:"inboundIds"`
	UsedTraffic   int64          `json:"usedTraffic"`
}

// ――― Ответ сервера - Получить одного клиента ―――
type ClientResponse struct {
	Success bool          `json:"success"`
	Message string        `json:"msg"`
	Object  ClientDetails `json:"obj"`
}

// Описание структуры входящего соединения
// type Inbound struct {
// 	ID int `json:"success"`
// 	InboundIDs int `json:"inboundIds"`
// 	Enable bool `json:"success"`
// 	Email string `json:"success"`
// }
