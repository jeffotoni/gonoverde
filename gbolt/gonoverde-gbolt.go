/*
* Go Library (C) 2018 Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
* @project     Noverde
* @package     main
* @author      @jeffotoni
* @size        15/04/2018
 */

// We have created some methods below to help us abstract some
// functions of boltdb.
// We realize how full the pkg bolt is, so many possibilities.
//
// We have discovered that there is no way to do a singleton of the
// DB object, but we are still testing to try to figure out some more
// flexible way to connect to buckets.
//
// The pkg bolt is very extensive and complete, so we had to implement the
// basics, and gradually deepening as much as possible and as needed to
// mature enough to propose improvements.
//
// The interesting thing is that we have several doubts that will
// be solved with tests, we love to test, as we saw the possibility
// of creating several databases and several buckets, we did not test yet
// but realized that there is possibility to leave the environment
// even more robust as the need.
package gbolt

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/boltdb/bolt"
)

// constante do banco
const (
	BDContas = "DBGonoverde"
	BDTrans  = "DBGonoverdeTransaction"
)

// Objects that define the database, directory where it will
// contain the file that is all in the no-sql storage.
//
// We can create a hash of databases and storage or buckets
// according to the need of the application and its complexity.
// Multiple files can be created each with their own needs,
// this further extends our possibilities for more
// robust deployments with no-sql.
var (
	DatabaseC = []byte(BDContas)
	DatabaseT = []byte(BDTrans)
	DirDb     = "db"
	PathDb    = "db/gbolt.db"
)

// Our struct for boltdb connection,
// it is with it that we will instantiate
// and do all necessary manipulation
// operations in our no-sql database.
type DB struct {
	*bolt.DB
}

// Here we define the variables
// that will manipulate and participate
// in our entire program
var (
	dbbolt *bolt.DB
	err    error
)

// We create a struct of our structure we can call
// the no-sql table, where we will use json to serve
// as a standard for our data recording in our bucket,
// remembering that our no-sql only accepts key and
// value, then our value will be a json Composed
// of several other fields.
type JsonDataDb struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Size    int64  `json:"size"`
	Path    string `json:"path"`
	Created string `json:"key"`
}

// We created a json date type of our structure
var djson JsonDataDb
var db2 *DB

// removendo a base de dados
func DropDatabase() {

	// get path local
	pwd, err := os.Getwd()

	// tratando erro
	if err != nil {
		log.Println(err)
		//os.Exit(0)
	}

	// constuindo path banco
	pwd = pwd + "/" + PathDb

	//removendo banco
	RemoveFile(pwd)
}

// using Connect
// This method is what does
// and returns our instance for access
// to our no-sql database.
func Connect() *DB {

	if db2 == nil {

		// Singleton the bank has to close every call,
		// save, update, get etc ..
		//
		// Testing and verifying if there is a directory
		// and file of our bucket, if it does not exist
		// it creates the directory and the file so that
		// we can manipulate all our bucket.
		//
		// Remember that boltdb with the open function also creates.
		if err := DataBaseTest(PathDb); err != nil {

			log.Fatal("Error Test database", err)
		}

		// Here is the object responsible for
		// allowing calls to the methods, such as Get, Save, etc.
		dbbolt, err = bolt.Open(PathDb, 0600, &bolt.Options{Timeout: 1 * time.Second})

		if err != nil {

			log.Fatal("connect error: ", err)
		}

		// We create a new reference
		// just to facilitate
		// understanding and syntax
		db2 = &DB{dbbolt}
	}

	return db2
}

//  DataBaseTest This method is if there is a directory
//  of the database
//
// if it does not exist it creates the directory
// and the file that we can call the bucket.
func DataBaseTest(PathDb string) error {

	// if file exist
	if !ExistDb(DirDb) {

		os.MkdirAll(DirDb, 0755)
	}

	// detect if file exists
	if !ExistDb(PathDb) {

		// create path
		var file, err = os.Create(PathDb)
		checkError(err)
		defer file.Close()

		// create file not exist
		w, errx := os.OpenFile(PathDb, os.O_WRONLY|os.O_CREATE, 0644)
		checkError(errx)
		defer w.Close()
	}

	return nil
}

//  [ExistDb Method only tests whether directory or file exists]
func ExistDb(name string) bool {

	if _, err := os.Stat(name); err != nil {

		if os.IsNotExist(err) {

			return false
		}
	}

	return true
}

// Save This method is responsible for saving on boltdb
func Save(keyS, valueS string, opcional ...string) error {

	DataBaseTmp := DatabaseC
	// if len(opcional) > 0 {

	// 	if opcional[0] == BDTrans {

	// 		DataBaseTmp = DatabaseT

	// 	} else {

	// 		fmt.Println("erro try save, database nao encontrada!", err)
	// 		os.Exit(0)
	// 	}
	// }

	db := Connect()

	//defer db.Close()

	key := []byte(keyS)
	value := []byte(valueS)

	err := db.Update(func(tx *bolt.Tx) error {

		bucket, err := tx.CreateBucketIfNotExists(DataBaseTmp)

		if err != nil {

			return err
		}

		err = bucket.Put(key, value)

		if err != nil {

			return err

		} else {

			//fmt.Println("save sucess")
			return nil
		}
	})

	if err != nil {

		fmt.Println("erro try save ", err)
		os.Exit(0)
	}

	return nil
}

// Get This method returns a string result as the last key
func Get(keyS string, opcional ...string) string {

	DataBaseTmp := DatabaseC
	// if len(opcional) > 0 {

	// 	if opcional[0] == BDTrans {

	// 		DataBaseTmp = DatabaseT

	// 	} else {

	// 		fmt.Println("erro try get, database nao encontrada!", err)
	// 		os.Exit(0)
	// 	}
	// }

	db := Connect()

	//defer db.Close()

	key := []byte(keyS)

	var valbyte []byte

	err = db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(DataBaseTmp)

		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", DataBaseTmp)
		}

		valbyte = bucket.Get(key)

		return nil
	})

	if err != nil {

		log.Fatal("Error open db, ", err)
	}

	return string(valbyte)
}

// using ListAllKeys
// Method responsible for listing all key parts
// and their respective values, in the bucket
// that is configured.
func ListAllKeys() error {

	db := Connect()

	if ExistDb(PathDb) {

		fmt.Println("Exist", db)

	} else {

		fmt.Println("Not exist!")
		os.Exit(0)
	}

	db.View(func(tx *bolt.Tx) error {

		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(DatabaseC))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {

			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})

	return nil
}

// SaveDb This method prepares the whole json string to save in boltdb
func SaveDb(keyfile string, namefile string, sizefile int64, pathFile string) error {

	times := fmt.Sprintf("%s", time.Now())

	stringJson := JsonDataDb{keyfile, namefile, sizefile, pathFile, times}

	respJson, err := json.Marshal(stringJson)

	respJsonX := string(respJson)

	err = Save(keyfile, respJsonX)

	if err == nil {

		//fmt.Println("save sucess..")
		return nil

	} else {

		//fmt.Println("Error", err)
		return err
	}
}

// JsonGet This method is responsible for returning the
// content in json format]
func JsonGet(keyS string) string {

	db := Connect()

	//defer db.Close()

	key := []byte(keyS)

	var valbyte []byte

	err = db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket(DatabaseC)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", DatabaseC)
		}

		valbyte = bucket.Get(key)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	byt := []byte(string(valbyte))

	///interface

	errjs := json.Unmarshal(byt, &djson)

	fmt.Println("here: ", djson)

	if errjs != nil {

		log.Fatal(fmt.Println("Error Json: ", errjs))
	}

	return string(valbyte)
}

// using checkError Test the errors
func checkError(err error) {

	if err != nil {
		fmt.Println("Error Database: ", err.Error())
		os.Exit(0)
	}
}

// removendo files
func RemoveFile(File string) (err error) {

	// if o arquivo existir
	// apagar para gerar
	// uma nova versao
	if ExistsFile(File) {

		//matando o arquivo
		err = os.Remove(File)
		// tratando o erro
		if err != nil {
			log.Println(err)
			return
		}
	}

	return
}

// Exists file in disck
func ExistsFile(file string) bool {

	if _, err := os.Stat(file); err != nil {

		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}
