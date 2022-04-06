package controllers

import (
	"context"
	"encoding/json"
	"fmt"

	"golang-backend/database"
	"golang-backend/model"

	"log"
	"net/http"

	// "os"
	// "strings"
	"time"

	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/mongo/options"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var SECRET = "ECOMMERCE"
var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var itemCollection *mongo.Collection = database.OpenCollection(database.Client, "item")
var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.Password = HashPassword(user.Password)
	err := userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&user)
	if err == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("User already exists")
		return
	}
	insertResult, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Fatal(err)
		return
	}
	fmt.Println(r.Body)
	fmt.Println("Inserted a Single User ", insertResult.InsertedID)
	var data = map[string]string{"success": "true", "user": user.Name}
	json.NewEncoder(w).Encode(data)
}

func GenerateJWT(id string) (string, error) {
	var mySigningKey = []byte(SECRET)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 7).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatalf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	var dbUser model.User

	json.NewDecoder(r.Body).Decode(&user)
	err := userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		log.Fatal(err)
	}
	if dbUser.Email == "" {
		json.NewEncoder(w).Encode("User Does Not Exist!")
	} else {
		userPassword := []byte(user.Password)
		dbUserPassword := []byte(dbUser.Password)

		passwordErr := bcrypt.CompareHashAndPassword(dbUserPassword, userPassword)
		if passwordErr != nil {
			fmt.Println("Incorrect Password!")
			json.NewEncoder(w).Encode("Incorrect Password!")
			return
		}
		token, err := GenerateJWT(dbUser.ID.Hex())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(token)
		var data = map[string]string{"token": token, "success": "true", "user": dbUser.Name}
		fmt.Println(data)

		json.NewEncoder(w).Encode(data)
	}
}

func SampleProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var items []model.Item
	fmt.Print(items)
	cur, err := itemCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(cur)
	for cur.Next(context.Background()) {
		var item model.Item
		err := cur.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	json.NewEncoder(w).Encode(items)
}

func EachProductGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	var eachItem model.Item

	params := mux.Vars(r)
	name := params["name"]
	fmt.Println(name)
	err := itemCollection.FindOne(context.Background(), bson.M{"name": name}).Decode(&eachItem)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(eachItem)
}

func CartAdd(w http.ResponseWriter, r *http.Request) {

	var order model.Order
	json.NewDecoder(r.Body).Decode(&order)
	fmt.Print(order)
	order.Total = order.Price * order.Quantity
	fmt.Print((order.Total))
	insertResult, err := orderCollection.InsertOne(context.Background(), order)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(insertResult)
}

func FullCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var orders []model.Order
	cur, err := orderCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(cur)
	for cur.Next(context.Background()) {
		var order model.Order
		err := cur.Decode(&order)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	json.NewEncoder(w).Encode(orders)
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	fmt.Println(id)
	result, err := orderCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
}
