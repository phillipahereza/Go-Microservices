package dbclient

import (
	"encoding/json"
	"fmt"
	"github.com/phillipahereza/go_microservices/model"
	bolt "go.etcd.io/bbolt"
	"log"
	"strconv"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(accountId string) (model.Account, error)
	Seed()
	Check() bool
}

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BoltClient) QueryAccount(accountId string) (model.Account, error) {
	account := model.Account{}

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("AccountBucket"))

		accountBytes := b.Get([]byte(accountId))

		if accountBytes == nil {
			return fmt.Errorf("No account found for " + accountId)
		}

		json.Unmarshal(accountBytes, &account)
		return nil
	})

	if err != nil {
		return model.Account{}, err
	}
	return account, nil
}

func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts()
	
}

func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

func (bc *BoltClient) seedAccounts() {
	total := 100
	for i := 0; i< total; i++ {
		key := strconv.Itoa(10000 + i)
		
		acc := model.Account{
			Id:   key,
			Name: "Person_" + strconv.Itoa(i),
		}

		jsonBytes, _ := json.Marshal(acc)

		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake accounts...\n", total)
}

func (bc *BoltClient) Check() bool {
	return bc.boltDB != nil
}