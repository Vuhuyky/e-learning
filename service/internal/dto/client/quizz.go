package requestdata

import (
	"app/internal/entity"
)

type CreateQuizzRequest struct {
	Ask        string             `json:"ask"`
	ResultType entity.RESULT_TYPE `json:"resultType"`
	Result     []string           `json:"result"`
	Option     []string           `json:"option"`
	Time       int                `json:"time"`
	EntityType entity.ENTITY_TYPE `json:"entityType"`
	EntityId   uint               `json:"entityId"`
}

type UpdateQuizzRequest struct {
	Id         uint64             `json:"id"`
	Ask        string             `json:"ask"`
	ResultType entity.RESULT_TYPE `json:"resultType"`
	Result     []string           `json:"result"`
	Option     []string           `json:"option"`
	Time       int                `json:"time"`
}

type DeleteQuizzRequest struct {
	Id uint64 `json:"id"`
}
