package repo

import "github.com/MdZunaed/bookshop/config"

type Repository struct {
	UserRepository MongoRepoInterface
	BookRepository MongoRepoInterface
}

func GetRepository() *Repository {
	dbName:= config.GetEnvProperty("database_name")
	return &Repository{
		UserRepository: GetMongoRepository(dbName, "user"),
		BookRepository: GetMongoRepository(dbName, "book"),
	}
}
