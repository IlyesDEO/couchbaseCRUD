package couchbase

import (
	"log"
	"time"

	"github.com/couchbase/gocb/v2"
	"github.com/kilianp07/CassandraCRUD/utils/structs"
)

type Couchbase struct {
	host       string
	bucketname string
	username   string
	password   string
	cluster    *gocb.Cluster
	Bucket     *gocb.Bucket
}

func (c *Couchbase) Connect() {
	var err error
	c.cluster, err = gocb.Connect("couchbase://"+c.host, gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: c.username,
			Password: c.password,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	c.Bucket = c.cluster.Bucket(c.bucketname)
	err = c.Bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func NewCouchbase(host string, bucketname string, username string, password string) *Couchbase {
	couchConnexion := &Couchbase{
		host:       host,
		bucketname: bucketname,
		username:   username,
		password:   password,
	}
	couchConnexion.Connect()
	return couchConnexion
}

func (c *Couchbase) MigrateData(data structs.Contact) error {
	var err error
	col := c.Bucket.DefaultCollection()
	_, err = col.Upsert(data.Id, data, nil)
	if err != nil {
		return err
	}
	return nil
}
