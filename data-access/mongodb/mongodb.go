package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongodbCURD() {

	// 设置MongoDB连接选项
	clientOptions := options.Client().ApplyURI("mongodb://root:openIM123@192.168.8.82:37017")

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()

	// 连接到MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// 获取数据库和集合
	db := client.Database("test")
	collection := db.Collection("user")

	// 检查连接
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("已成功连接到MongoDB!")

	//插入文档
	user := map[string]interface{}{"name": "Alice", "age": 29}
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}

	users := []interface{}{
		map[string]interface{}{"name": "Bob", "age": 31},
		map[string]interface{}{"name": "Charlie", "age": 25},
	}
	_, err = collection.InsertMany(ctx, users)
	if err != nil {
		log.Fatal(err)
	}

	//查询文档
	//查询单个文档
	var result map[string]interface{}
	filter := map[string]interface{}{"name": "Alice"}
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	//查询多个文档
	filter = map[string]interface{}{"age": map[string]interface{}{"$gte": 30}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var results []map[string]interface{}
	if err = cursor.All(ctx, &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		log.Println(result)
	}

	//更新文档
	//更新单个文档
	filter = map[string]interface{}{"name": "Bob"}
	update := map[string]interface{}{"$set": map[string]interface{}{"age": 32}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	//更新多个文档
	filter = map[string]interface{}{"age": map[string]interface{}{"$lt": 30}}
	update = map[string]interface{}{"$set": map[string]interface{}{"valid": true}}
	_, err = collection.UpdateMany(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	//删除文档
	//删除单个文档
	filter = map[string]interface{}{"name": "Charlie"}
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	//删除多个文档
	filter = map[string]interface{}{"valid": true}
	_, err = collection.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	// 当程序结束时关闭连接
	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("已断开与MongoDB的连接!")
}

type user struct {
	Id   int64  //字段首字母要大写要大写要大写，重要的事说三遍
	Name string //字段首字母要大写要大写要大写，重要的事说三遍
}

type Trainer struct {
	Name string
	Age  int
	City string
}
