package requestdata

type InfoVideo struct {
	LessionId uint   `json:"lessionId"`
	Uuid      string `json:"uuid"`
}

type GetListVideoResponse struct {
	Files []string `json:"files"`
}
