package db

import (
	"context"

	"github.com/vsivarajah/RiotStatistics/producer"
)

type db struct {
	ConnStr string
	Port    string
}

func New(...string) producer.Sender {
	return &db{
		ConnStr: "",
		Port:    "",
	}
}

func (db *db) Init(ctx context.Context, config interface{}) error {
	return nil
}

func (db *db) Send(ctx context.Context, message interface{}) error {

	// err := sql.Query(ctx , "query")

	return nil
}
