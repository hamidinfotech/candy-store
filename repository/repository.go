package repository

import "candystore/entity"

type CandystoreRepository interface {
	GetTopCustomers() []entity.CustomerStat
}