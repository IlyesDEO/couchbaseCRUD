package contactRepo

import (
	"fmt"
	"log"

	"github.com/IlyesDEO/goCrud/pkg/couchbase"
	"github.com/couchbase/gocb/v2"
	"github.com/kilianp07/CassandraCRUD/utils/structs"
)

func GetById(id string) (*structs.Contact, error) {
	c := couchbase.NewCouchbase("localhost", "contact", "Administrator", "root123")
	col := c.Bucket.DefaultCollection()

	getResult, err := col.Get(id, nil)
	if err != nil {
		return nil, err
	}
	var contact structs.Contact
	err = getResult.Content(&contact)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func GetAll() ([]*structs.Contact, error) {
	contacts := []*structs.Contact{}

	c := couchbase.NewCouchbase("127.0.0.1", "contact", "Administrator", "root123")

	contactScope := c.Bucket.Scope("_default")
	queryResult, err := contactScope.Query(
		fmt.Sprintf("SELECT * FROM _default"),
		&gocb.QueryOptions{Adhoc: true},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Print each found Row
	for queryResult.Next() {
		var result interface{}
		err := queryResult.Row(&result)
		if err != nil {
			log.Fatal(err)
		}
		a := result.(map[string]interface{})
		b := a["_default"].(map[string]interface{})

		contact := structs.Contact{
			Id:          b["id"].(string),
			Title:       b["title"].(string),
			Name:        b["name"].(string),
			Address:     b["address"].(string),
			RealAddress: b["realAddress"].(string),
			Departement: b["department"].(string),
			Country:     b["country"].(string),
			Tel:         b["tel"].(string),
			Email:       b["email"].(string),
		}

		contacts = append(contacts, &contact)
	}
	return contacts, nil
}

func Create(contact *structs.Contact) error {

	c := couchbase.NewCouchbase("localhost", "contact", "Administrator", "root123")
	var err error
	col := c.Bucket.DefaultCollection()
	_, err = col.Upsert(contact.Id, contact, nil)
	if err != nil {
		return err
	}
	return nil
}

func Update(contact *structs.Contact) error {

	c := couchbase.NewCouchbase("localhost", "contact", "Administrator", "root123")
	var err error

	col := c.Bucket.DefaultCollection()

	updateGetResult, err := col.Get(contact.Id, nil)
	if err != nil {
		panic(err)
	}

	err = updateGetResult.Content(&contact)
	if err != nil {
		panic(err)
	}

	_, err = col.Replace(contact.Id, contact, &gocb.ReplaceOptions{
		Cas: updateGetResult.Cas(),
	})
	if err != nil {
		return err
	}
	return nil
}

func Delete(id string) error {
	var ctx *gocb.TransactionAttemptContext
	c := couchbase.NewCouchbase("localhost", "contact", "Administrator", "root123")
	col := c.Bucket.DefaultCollection()

	var err error
	docC, err := ctx.Get(col, id)
	if err != nil {
		return err
	}

	err = ctx.Remove(docC)
	if err != nil {
		return err
	}
	return nil
}
