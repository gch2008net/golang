package main

import (
	"data-access/mongodb"
	"fmt"
)

func main() {
	fmt.Println("this is a data-access example  for mysql.. ")
	// mysqldb.MysqlCURD()

	fmt.Println("this is a data-access example for mongodb.. ")
	mongodb.MongodbCURD()
}
