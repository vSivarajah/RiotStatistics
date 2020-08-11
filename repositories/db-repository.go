package producer

import (
	"github.com/vsivarajah/RiotStatistics/repositories/db"
)

type DbRepository struct {
	db.MongoRepository
	db.RedisRepository
}
