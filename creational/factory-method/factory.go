package main

type FactoryDatabase interface {
	CreateDatabase() Database
}

type FactorySql struct{}

func (*FactorySql) CreateDatabase() Database {
	return &Sql{}
}

type FactoryMongo struct{}

func (*FactoryMongo) CreateDatabase() Database {
	return &Mongo{}
}
