package main

import (
	"context"
	"fmt"
	"golang-backend/database"
	"golang-backend/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var itemCollection *mongo.Collection = database.OpenCollection(database.Client, "item")

	data := []interface{}{
		bson.D{
			{Key: "name", Value: "fallon1"},
			{Key: "image", Value: "https://images.unsplash.com/photo-1508296695146-257a814070b4?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8c3VuJTIwZ2xhc3N8ZW58MHx8MHx8&auto=format&fit=crop&w=500&q=60"},
			{Key: "des", Value: "The range of lifestyle products will be launched under a new sub-brand Qubo Go. The launch of the sub-brand comes at a time when the need for smart lifestyle devices is emerging, particularly to bring convenience while matching the fast-paced on-the-go lifestyle."},
			{Key: "quantity", Value: 0},
			{Key: "price", Value: 1200},
		},
		bson.D{
			{Key: "name", Value: "Clear glass"},
			{Key: "image", Value: "https://images.unsplash.com/photo-1515613813261-5cd015bcd184?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8c3VuJTIwZ2xhc3N8ZW58MHx8MHx8&auto=format&fit=crop&w=500&q=60"},
			{Key: "des", Value: "The range of lifestyle products will be launched under a new sub-brand Qubo Go. The launch of the sub-brand comes at a time when the need for smart lifestyle devices is emerging, particularly to bring convenience while matching the fast-paced on-the-go lifestyle."},
			{Key: "quantity", Value: 0},
			{Key: "price", Value: 900},
		}, bson.D{
			{Key: "name", Value: "transparent"},
			{Key: "image", Value: "https://images.unsplash.com/photo-1511499767150-a48a237f0083?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8c3VuZ2xhc3N8ZW58MHx8MHx8&auto=format&fit=crop&w=500&q=60"},
			{Key: "des", Value: "The range of lifestyle products will be launched under a new sub-brand Qubo Go. The launch of the sub-brand comes at a time when the need for smart lifestyle devices is emerging, particularly to bring convenience while matching the fast-paced on-the-go lifestyle."},
			{Key: "quantity", Value: 0},
			{Key: "price", Value: 600},
		}, bson.D{
			{Key: "name", Value: "cooling glass"},
			{Key: "image", Value: "https://images.unsplash.com/photo-1502767089025-6572583495f9?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8c3VuZ2xhc3N8ZW58MHx8MHx8&auto=format&fit=crop&w=500&q=60"},
			{Key: "des", Value: "The range of lifestyle products will be launched under a new sub-brand Qubo Go. The launch of the sub-brand comes at a time when the need for smart lifestyle devices is emerging, particularly to bring convenience while matching the fast-paced on-the-go lifestyle."},
			{Key: "quantity", Value: 0},
			{Key: "price", Value: 1500},
		}, bson.D{
			{Key: "name", Value: "lens"},
			{Key: "image", Value: "https://images.unsplash.com/photo-1625591340248-6d289000f96a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8c3VuZ2xhc3N8ZW58MHx8MHx8&auto=format&fit=crop&w=500&q=60"},
			{Key: "des", Value: "The range of lifestyle products will be launched under a new sub-brand Qubo Go. The launch of the sub-brand comes at a time when the need for smart lifestyle devices is emerging, particularly to bring convenience while matching the fast-paced on-the-go lifestyle."},
			{Key: "quantity", Value: 0},
			{Key: "price", Value: 800},
		}, bson.D{
			{Key: "name", Value: "agarwal"},
			{Key: "image", Value: "https://images.unsplash.com/photo-1625591339971-4c9a87a66871?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTF8fHN1bmdsYXNzfGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60"},
			{Key: "des", Value: "The range of lifestyle products will be launched under a new sub-brand Qubo Go. The launch of the sub-brand comes at a time when the need for smart lifestyle devices is emerging, particularly to bring convenience while matching the fast-paced on-the-go lifestyle."},
			{Key: "quantity", Value: 0},
			{Key: "price", Value: 1000},
		}, bson.D{
			{Key: "name", Value: "aravind"},
			{Key: "image", Value: "https://images.unsplash.com/photo-1572635196237-14b3f281503f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8c3VuZ2xhc3N8ZW58MHx8MHx8&auto=format&fit=crop&w=500&q=60"},
			{Key: "des", Value: "The range of lifestyle products will be launched under a new sub-brand Qubo Go. The launch of the sub-brand comes at a time when the need for smart lifestyle devices is emerging, particularly to bring convenience while matching the fast-paced on-the-go lifestyle."},
			{Key: "quantity", Value: 0},
			{Key: "price", Value: 1400},
		}, bson.D{
			{Key: "name", Value: "optics"},
			{Key: "image", Value: "https://images.unsplash.com/photo-1625591342274-013866180475?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8N3x8c3VuZ2xhc3N8ZW58MHx8MHx8&auto=format&fit=crop&w=500&q=60"},
			{Key: "des", Value: "The range of lifestyle products will be launched under a new sub-brand Qubo Go. The launch of the sub-brand comes at a time when the need for smart lifestyle devices is emerging, particularly to bring convenience while matching the fast-paced on-the-go lifestyle."},
			{Key: "quantity", Value: 0},
			{Key: "price", Value: 1800},
		},
		bson.D{
			{Key: "name", Value: "anand"},
			{Key: "image", Value: "https://images.unsplash.com/photo-1566421966482-ad8076104d8e?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MjB8fHN1bmdsYXNzfGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60"},
			{Key: "des", Value: "The range of lifestyle products will be launched under a new sub-brand Qubo Go. The launch of the sub-brand comes at a time when the need for smart lifestyle devices is emerging, particularly to bring convenience while matching the fast-paced on-the-go lifestyle."},
			{Key: "quantity", Value: 0},
			{Key: "price", Value: 1650},
		},
	}

	var mani model.Item
	fmt.Print(mani)

	itemResult, err := itemCollection.InsertMany(context.Background(), data)
	if err != nil {
		log.Fatal((err))
	}
	fmt.Printf("Inserted %v documents into items collection! \n", len(itemResult.InsertedIDs))
}
