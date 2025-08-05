package config

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

var env ConfigDto

// Go default function to run when program runs
func init() {
	if env.port == ""{
		LoadEnvironment()
	}
	ConfigEnv()
}

type ConfigDto struct {
	port         string
	secret_key   string
	database_url string
	database_name string
}

func ConfigEnv() {
	env = ConfigDto{
		port:         os.Getenv("PORT"),
		secret_key:   os.Getenv("SECRET_KEY"),
		database_url: os.Getenv("MONGO_DB_URL"),
		database_name: os.Getenv("DATABASE_NAME"),
	}
}

func LoadEnvironment() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env variable")
	}
}

func accessEnv(key string) (string, error) {
	v := reflect.ValueOf(env)
	t := v.Type()
	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("expected struct in Env")
	}
	_, ok := t.FieldByName(key)

	if !ok {
		return "", fmt.Errorf("%s key not found", key)
	}
	return v.FieldByName(key).String(), nil
}

func GetEnvProperty(key string) string {
	if env.port == "" {
		ConfigEnv()
	}

	value, err := accessEnv(key)

	if err != nil {
		fmt.Println("get env error:", err)
		return value
	}
	return value
}
