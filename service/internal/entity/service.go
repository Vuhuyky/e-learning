package entity

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	Type TYPE_SERVICE `json:"type"`
	Ip   string       `json:"ip"`
}

type TYPE_SERVICE string

const (
	MERGE_BLOB_SERVICE TYPE_SERVICE = "MERGE_BLOB_SERVICE"
	QUANTITY_SERVICE   TYPE_SERVICE = "QUANTITY_SERVICE"
)
