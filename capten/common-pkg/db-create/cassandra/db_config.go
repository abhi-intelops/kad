// Package cassandra contains ...
package cassandra

import "github.com/gocql/gocql"

type DBConfig struct {
	DbAddresses         []string `envconfig:"DB_ADDRESSES" required:"true" default:"localhost:9042"`
	DbAdminUsername     string   `envconfig:"DB_ADMIN_USERNAME" required:"true" default:"cassandra"`
	DbServiceUsername   string   `envconfig:"DB_SERVICE_USERNAME" required:"false"`
	DbName              string   `envconfig:"CASSANDRA_DB_NAME" required:"false"`
	DbReplicationFactor string   `envconfig:"DB_REPLICATION_FACTOR" required:"true" default:"1"`
	DbAdminPassword     string   `envconfig:"DB_ADMIN_PASSWD" required:"true" default:"cassandra"`
	DbServicePassword   string   `envconfig:"DB_SERVICE_PASSWD" required:"false"`
}

type Store interface {
	Connect(dbAddrs []string, dbAdminUsername string, dbAdminPassword string) (err error)
	Session() *gocql.Session
	Close()
	CreateDbUser(serviceUsername string, servicePassword string) (err error)
	GrantPermission(serviceUsername string, dbName string) (err error)
	CreateDb(dbName string, replicationFactor string) (err error)
	CreateLockSchemaDb(replicationFactor string) (err error)
}
