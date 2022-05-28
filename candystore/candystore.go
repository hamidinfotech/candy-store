package candystore

import (
	repository_sql "candystore/repository/sqlite"
	//"candystore/repository/csv"
	"candystore/entity"
)

func TopCustomers() []entity.CustomerStat {
	return repository_sql.GetTopCustomers()
}
