package rawquery

import (
	"app/internal/connection"
	requestdata "app/internal/dto/client"

	"gorm.io/gorm"
)

type queryRawService[T any] struct {
	psql *gorm.DB
}

type QueryRawService[T any] interface {
	Query(payload requestdata.QueryRawReq[T]) (*T, error)
	QueryAll(payload requestdata.QueryRawReq[T]) ([]T, error)
}

func Register[T any]() QueryRawService[T] {
	return &queryRawService[T]{
		psql: connection.GetPsql(),
	}
}
