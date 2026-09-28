// 1. URI string for connection
// 2. Create mongodb clients
// 3. Connect to mongodb
// 4. disconnect just befor main function finished
// 5. ping the database

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	// "go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Product struct {
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Price       float64   `json:"price" bson:"price"`
	Category    string    `json:"category" bson:"category"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func connectDB() {

	uri := os.Getenv("MONGODB_URI")
	// uri := "mongodb+srv://xxtracymiller_db_user:EjmO1wAE3ueMGt0A@cluster0.y4mb6h9.mongodb.net"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	// 2. Set client options and apply the URI
	clientOpts := options.Client().ApplyURI(uri)

	// 3. Connect to MongoDB
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		log.Fatalf("Failed to create Mongo client: %v", err)
	}

	// 4. Defer disconnect to ensure resources are cleaned up when main exits
	defer func() {
		// Use a short, dedicated timeout context for disconnecting
		disconnectCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := client.Disconnect(disconnectCtx)
		if err != nil {
			log.Fatalf("Error during MongoDB disconnect: %v", err)
		}
	}()

	// 5. Ping the database to verify the connection is alive
	pingCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(pingCtx, nil)
	if err != nil {
		log.Fatalf("Could not ping MongoDB: %v", err)
	}

	fmt.Println("Database connected successfully.")

	// database := client.Database("test_database")
	// collection := database.Collection("test_collection")

	// fmt.Println("Collection: ", collection)
}

func insertData(p Product) bool {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. ")
	}
	clientOpts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(clientOpts)
	if err != nil {
		log.Fatalf("Failed to create Mongo client: %v", err)
	}

	coll := client.Database("go_dbconnect").Collection("products")

	// Find one operation
	filter := bson.M{"name": p.Name}
	var findResult bson.M
	findErr := coll.FindOne(context.TODO(), filter).Decode(&findResult)
	res, _ := bson.MarshalExtJSON(findResult, false, false) // bson to json conversion
	fmt.Println("findResult", string(res))                  // findResult {"category":"men's clothing","_id":{"$oid":"6aba18bd2a481b9eec35bd1a"},"name":"Mens Casual Premium Slim Fit T-Shirts","description":"Slim-fitting style, contrast raglan long sleeve, three-button henley placket, light weight & soft fabric for breathable and comfortable wearing. And Solid stitched shirts with round neck made for durability and a great fit for casual fashion wear and diehard baseball fans. The Henley style round neckline includes a three-button placket.","price":100.0}
	fmt.Println("findResult", findResult)                   // findResult {"_id":{"$oid":"6aba18bd2a481b9eec35bd1a"},"name":"Mens Casual Premium Slim Fit T-Shirts","description":"Slim-fitting style, contrast raglan long sleeve, three-button henley placket, light weight & soft fabric for breathable and comfortable wearing. And Solid stitched shirts with round neck made for durability and a great fit for casual fashion wear and diehard baseball fans. The Henley style round neckline includes a three-button placket.","price":{"$numberDouble":"100.0"},"category":"men's clothing"}
	if findErr == nil {
		return true
	}

	// Insert one operation
	result, err := coll.InsertOne(context.TODO(), p)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return false
}

// Handler function
func addProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product Product

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		response := struct {
			Response string `json:"response"`
		}{Response: "Please send some valid data"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if strings.TrimSpace(product.Name) == "" {
		response := struct {
			Response string `json:"response"`
		}{Response: "Name must be required"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	} else if strings.TrimSpace(product.Description) == "" {
		response := struct {
			Response string `json:"response"`
		}{Response: "Description must be required"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	} else if product.Price == 0 {
		response := struct {
			Response string `json:"response"`
		}{Response: "Price must be required"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	} else if strings.TrimSpace(product.Category) == "" {
		response := struct {
			Response string `json:"response"`
		}{Response: "Category must be required"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	found := insertData(product)
	if found {
		response := struct {
			Response string `json:"response"`
		}{Response: "Same name already exist"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := struct {
		Response Product `json:"response"`
	}{Response: product}

	json.NewEncoder(w).Encode(response)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
		return
	}

	clientOpt := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(clientOpt)
	if err != nil {
		log.Fatalf("Failed to create Mongo client: %v", err)
	}

	coll := client.Database("go_dbconnect").Collection("products")

	cursor, cursorErr := coll.Find(context.TODO(), bson.D{{}})
	if cursorErr != nil {
		log.Fatal(cursorErr)
		return
	}
	defer cursor.Close(context.TODO())

	var products []Product
	productsErr := cursor.All(context.TODO(), &products)
	if productsErr != nil {
		log.Fatal(productsErr)
	}

	if len(products) == 0 {
		response := struct {
			Response string `json:"response" bson:"response"`
		}{Response: "No data found"}
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := struct {
		Message string    `json:"message" bson:"message"`
		Data    []Product `json:"data" bson:"data"`
	}{Message: "Data fetched successfully", Data: products}

	json.NewEncoder(w).Encode(response)

	// fmt.Println(products)
}

func getOneProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
		return
	}

	clientOpts := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(clientOpts)
	if err != nil {
		log.Fatalf("Failed to create Mongo client: %v", err)
		return
	}

	// coll := client.Database("go_dbconnect").Collection("products")
	coll := client.Database("go_dbconnect").Collection("products")

	idStr := r.PathValue("id")
	idStr = strings.TrimSpace(idStr)

	// 2. Convert the string to a MongoDB ObjectID
	objID, objIDErr := bson.ObjectIDFromHex(idStr)
	if objIDErr != nil {
		// Handle invalid ObjectID format (e.g., if it's not 24 hex characters)
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	filter := bson.M{"_id": objID}
	var data Product
	findOneProductErr := coll.FindOne(context.TODO(), filter).Decode(&data)
	if findOneProductErr != nil {
		response := struct {
			Message string `json:"message" bson:"message"`
			Data    string `json:"data" bson:"data"`
		}{Message: "No product found with given id.", Data: ""}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)

		return
	}

	// json.NewDecoder(res).Decode(&data)

	response := struct {
		Message string  `json:"message" bson:"message"`
		Data    Product `json:"data" bson:"data"`
	}{Message: "Data fetched successfully.", Data: data}

	json.NewEncoder(w).Encode(response)
}

func removeOneProduct(w http.ResponseWriter, r *http.Request){}

func main() {
	// Load .env file first
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file. |", envErr)
	}

	http.HandleFunc("POST /add-products", addProducts)
	http.HandleFunc("GET /products", getProducts)
	http.HandleFunc("GET /product/{id}", getOneProduct)
	http.HandleFunc("DELETE /product/{id}", removeOneProduct)

	fmt.Println("Server is listiing on port 3000")
	connectDB()

	// p := Product{
	// 	Name:        "Fjallraven - Foldsack No. 1 Backpack, Fits 15 Laptops",
	// 	Description: "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
	// 	Price:       109.95,
	// 	Category:    "men's clothing",
	// }

	// insertData(p)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("Error while listining.")
		return
	}
}
