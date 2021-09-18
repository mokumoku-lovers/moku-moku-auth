package cassandra

import (
	"github.com/joho/godotenv"
)
var (
	host    string
)

func loadEnvironment() {
	err := godotenv.Load("./clients/cassandra/.env")
	if err != nil {
		log.Println("Could not load environment variables")
		panic(err)
	}
}

func init() {
	loadEnvironment()
	host = os.Getenv(cHost)
}
}
