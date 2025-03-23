package dtoclientservice

type InitProcessRequest struct {
}

type AddData struct {
	Mess        string `json:"mess"`
	UuidProcess string `json:"uuidProcess"`
}
