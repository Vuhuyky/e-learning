package queuepayload

type QueueFileDeleteMp4 struct {
	Uuid string `json:"uuid"`
}

type QueueFileM3U8Payload struct {
	Path     string `json:"path"`
	IpServer string `json:"ipServer"`
	Uuid     string `json:"uuid"`
	Quantity string `json:"quantity"`
}

type QueueUrlQuantityPayload struct {
	Url      string `json:"url"`
	Quantity string `json:"quantity"`
	Uuid     string `json:"uuid"`
}
