package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	Name    string     `json:"name" bson:"name"`
	Age     int        `json:"age" bson:"age"`
	Address *Address   `json:"address" bson:"address,omitempty"`
	Salary  int        `json:"salary" bson:"salary,omitempty"`
	Date    *time.Time `json:"date" bson:"date"`
}

type Address struct {
	Country string `json:"country" bson:"country"`
	City    string `json:"city" bson:"city"`
}

func insertOne(usersCollection *mongo.Collection) error {
	user := User{
		Name: "chen4",
		Age:  4,
	}
	result, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	fmt.Println(result.InsertedID)
	return nil
}

func insertMany(usersCollection *mongo.Collection) error {
	users := []interface{}{
		User{Name: "chen4", Age: 4},
		User{Name: "chen5", Age: 5},
		User{Name: "chen6", Age: 6},
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
	defer cursor.Close(context.TODO())

	users := []User{}
	if err = cursor.All(context.TODO(), &users); err != nil {
		return err
	}

	for _, user := range users {
		fmt.Println(user)
	}
	return nil
}

func readMany(usersCollection *mongo.Collection) error {
	filter := primitive.M{"age": primitive.M{"$gt": 2}}
	cursor, err := usersCollection.Find(context.TODO(), filter)
	if err != nil {
		return err
	}

	users := []User{}
	// if err = cursor.All(context.TODO(), &users); err != nil {
	// 	return err
	// }

	for cursor.Next(context.TODO()) {
		user := User{}
		if err := cursor.Decode(&user); err != nil {
			return err
		}
		users = append(users, user)
	}
	fmt.Println(users)
	return nil
}

func readOne(usersCollection *mongo.Collection) error {
	filter := primitive.M{"name": "chen1"}
	user := User{}
	if err := usersCollection.FindOne(context.TODO(), filter).Decode(&user); err != nil {
		return err
	}
	fmt.Println(user)
	return nil
}

func update(usersCollection *mongo.Collection) error {
	filter := primitive.M{"name": "chen4", "age": 4}
	// address := Address{Country: "USA", City: "New York"}
	// update := primitive.M{"$set": primitive.M{"salary": 1000}} // It is working, even though it is empty or nil
	// update := primitive.M{"$set": primitive.M{"address.city": "USA"}} // It is working, even though it is empty
	update := primitive.M{"$set": primitive.M{"address": nil}} // It is working, even though it is empty

	result, err := usersCollection.UpdateOne(context.TODO(), filter, update)
	// result, err := usersCollection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println("Number of documents updated:", result.ModifiedCount)
	return nil
}

func replaceOne(usersCollection *mongo.Collection) error {
	filter := primitive.M{"name": "chen"}
	replacement := User{Name: "chen", Age: 30}
	result, err := usersCollection.ReplaceOne(context.TODO(), filter, replacement)
	if err != nil {
		return err
	}
	fmt.Println("Number of documents updated:", result.ModifiedCount)
	return nil
}

func deleteOne(usersCollection *mongo.Collection) error {
	filter := primitive.M{"age": 5}
	result, err := usersCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	fmt.Println("Number of documents deleted:", result.DeletedCount)
	return nil
}

func DeleteMany(usersCollection *mongo.Collection) error {
	filter := primitive.M{"age": primitive.M{"$gt": 3}}
	results, err := usersCollection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	fmt.Println("Number of documents deleted:", results.DeletedCount)
	return nil
}

func Count(usersCollection *mongo.Collection) (count int64, err error) {
	filter := primitive.M{"age": primitive.M{"$gt": 2}}
	return usersCollection.CountDocuments(context.TODO(), filter)
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
	usersCollection := client.Database("db01").Collection("users")

	// if err := insertOne(usersCollection); err != nil {
	// 	fmt.Println("failed to insert data:", err)
	// }

	// if err := insertMany(usersCollection); err != nil {
	// 	fmt.Println("failed to insert data:", err)
	// }

	// if err := readAll(usersCollection); err != nil {
	// 	fmt.Println("failed to read:", err)
	// }

	// if err := readMany(usersCollection); err != nil {
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

	// if err := DeleteMany(usersCollection); err != nil {
	// 	fmt.Println("failed to deleteOne:", err)
	// }

	count, err := Count(usersCollection)
	if err != nil {
		fmt.Println("failed to count", err)
	}
	fmt.Println("count:", count)

}
