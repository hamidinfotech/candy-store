package candystore

import (
	"candystore/entity"
	"candystore/repository"
)

// Sqlite implementation repository
var candystoreRepository repository.CandystoreRepository = repository.CandystoreRepositorySqlite{}

// CSV implementation repository
// var candystoreRepository repository.CandystoreRepository = repository.CandystoreRepositoryCSV{}

func TopCustomers() []entity.CustomerStat {
	return candystoreRepository.GetTopCustomers()
}
