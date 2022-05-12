package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func insertOne(usersCollection *mongo.Collection) error {
	user := bson.D{{"fullName", "User 1"}, {"age", 30}}
	result, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	fmt.Println(result.InsertedID)
	return nil
}

func insertMany(usersCollection *mongo.Collection) error {
	users := []interface{}{
		bson.D{{"fullName", "User 2"}, {"age", 25}},
		bson.D{{"fullName", "User 3"}, {"age", 20}},
		bson.D{{"fullName", "User 4"}, {"age", 28}},
	}

	results, err := usersCollection.InsertMany(context.TODO(), users)
	if err != nil {
		return err
	}
	fmt.Println(results.InsertedIDs)
	return nil
}

func readAll(usersCollection *mongo.Collection) error {
	cursor, err := usersCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return err
	}

	for _, result := range results {
		fmt.Println(result)
	}
	return nil
}

func readMany(usersCollection *mongo.Collection) error {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}}},
				},
			},
		},
	}
	cursor, err := usersCollection.Find(context.TODO(), filter)
	if err != nil {
		return err
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return err
	}

	for _, result := range results {
		fmt.Println(result)
	}
	return nil
}

func readOne(usersCollection *mongo.Collection) error {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}}},
				},
			},
		},
	}

	// retrieving the first document that match the filter
	var result bson.M
	if err := usersCollection.FindOne(context.TODO(), filter).Decode(&result); err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}

func update(usersCollection *mongo.Collection) error {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 25}}},
				},
			},
		},
	}

	update := bson.D{
		{"$set",
			bson.D{
				{"age", 40},
			},
		},
	}

	result, err := usersCollection.UpdateOne(context.TODO(), filter, update)
	// result, err := usersCollection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println("Number of documents updated:", result.ModifiedCount)
	return nil
}

func replaceOne(usersCollection *mongo.Collection) error {
	filter := bson.D{{"fullName", "User 1"}}
	replacement := bson.D{
		{"firstName", "John"},
		{"lastName", "Doe"},
		{"age", 30},
		{"emailAddress", "johndoe@email.com"},
	}

	result, err := usersCollection.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		return err
	}
	fmt.Println("Number of documents updated:", result.ModifiedCount)
	return nil
}

func deleteOne(usersCollection *mongo.Collection) error {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 28}}},
				},
			},
		},
	}
	result, err := usersCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	fmt.Println("Number of documents deleted:", result.DeletedCount)

	// // delete every document that match the filter
	// results, err := usersCollection.DeleteMany(context.TODO(), filter)
	// // check for errors in the deleting
	// if err != nil {
	// 	panic(err)
	// }
	// // display the number of documents deleted
	// fmt.Println("deleting every result from the search filter")
	// fmt.Println("Number of documents deleted:", results.DeletedCount)
	return nil
}

func DeleteMany(usersCollection *mongo.Collection) error {
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{
					{"age", bson.D{{"$gt", 23}}},
				},
			},
		},
	}
	results, err := usersCollection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}

	fmt.Println("Number of documents deleted:", results.DeletedCount)
	return nil
}

func main() {
	// connect
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	// ping
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	// use the collection of the database
	usersCollection := client.Database("db01").Collection("coll01")

	// if err := insertOne(usersCollection); err != nil {
	// 	fmt.Println("failed to insert data:", err)
	// }

	// if err := insertMany(usersCollection); err != nil {
	// 	fmt.Println("failed to insert data:", err)
	// }

	// if err := readAll(usersCollection); err != nil {
	// 	fmt.Println("failed to read:", err)
	// }

	// if err := readOne(usersCollection); err != nil {
	// 	fmt.Println("failed to read:", err)
	// }

	// if err := update(usersCollection); err != nil {
	// 	fmt.Println("failed to update:", err)
	// }

	// if err := replaceOne(usersCollection); err != nil {
	// 	fmt.Println("failed to replaceOne:", err)
	// }

	// if err := deleteOne(usersCollection); err != nil {
	// 	fmt.Println("failed to deleteOne:", err)
	// }

	if err := DeleteMany(usersCollection); err != nil {
		fmt.Println("failed to deleteOne:", err)
	}

}
