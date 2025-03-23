package entity

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Quizz struct {
	gorm.Model
	Ask        string         `json:"ask"`
	ResultType RESULT_TYPE    `json:"resultType"`
	Result     pq.StringArray `json:"result" gorm:"type:text[]"`
	Option     pq.StringArray `json:"option" gorm:"type:text[]"`
	Time       int            `json:"time"`

	EntityType ENTITY_TYPE `json:"entityType"`
	EntityId   uint        `json:"entityId"`
}

type ENTITY_TYPE string

const (
	QUIZZ_VIDEO_LESSION ENTITY_TYPE = "QUIZZ_VIDEO_LESSION"
	QUIZZ_LESSION       ENTITY_TYPE = "QUIZZ_LESSION"
)

type RESULT_TYPE string

const (
	QUIZZ_SINGLE_RESULT RESULT_TYPE = "QUIZZ_SINGLE_RESULT"
	QUIZZ_MULTI_RESULT  RESULT_TYPE = "QUIZZ_MULTI_RESULT"
)
