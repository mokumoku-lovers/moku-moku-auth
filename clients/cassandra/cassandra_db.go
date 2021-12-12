package cassandra

import (
	"log"
	"os"

	"github.com/gocql/gocql"
	"github.com/joho/godotenv"
)

const (
	cHost = "C_HOST"
)

var (
	cluster *gocql.ClusterConfig
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
	cluster = gocql.NewCluster(host)
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
