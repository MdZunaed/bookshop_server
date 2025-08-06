package repository

import "github.com/MdZunaed/bookshop/config"

type Repository struct {
	UserRepository MongoRepoInterface
}

func GetRepository() *Repository {
	return &Repository{
		UserRepository: GetMongoRepository(config.GetEnvProperty("database_name"), "user"),
	}
}
