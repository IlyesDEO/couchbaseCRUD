package main

import (
	"github.com/IlyesDEO/goCrud/pkg/api"
	"github.com/IlyesDEO/goCrud/pkg/couchbase"
)

func main() {

	couchbase.NewCouchbase("localhost", "contact", "Administrator", "root123")

	api.Start()
}
