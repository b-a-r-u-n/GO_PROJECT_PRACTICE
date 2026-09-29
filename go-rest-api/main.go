// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// type Article struct {
// 	Id      string `json:"id"`
// 	Title   string `json:"title"`
// 	Desc    string `json:"desc"`
// 	Content string `json:"content"`
// }

// var Articles []Article

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the homepage!!")
// 	fmt.Println("Endpoint Hit.. | Home page")
// }
// func returnAllArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit.. | Articles page")
// 	json.NewEncoder(w).Encode(Articles)

// }

// func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit.. | Article page")
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	// fmt.Fprintf(w, "Id:"+id)

// 	for _, article := range Articles {
// 		if article.Id == id {
// 			json.NewEncoder(w).Encode(article)
// 			return
// 		}
// 	}

// 	response := struct {
// 		Error string `json:"error"`
// 	}{Error: "No data found for given id."}
// 	json.NewEncoder(w).Encode(response)
// }

// func createNewArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit.. | Create article page")

// 	reqBody, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Error reading request body", http.StatusBadRequest)
// 		return
// 	}
// 	// fmt.Fprintf(w, "%+v", string(reqBody))

// 	var article Article
// 	err = json.Unmarshal(reqBody, &article)
// 	if err != nil {
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}

// 	Articles = append(Articles, article)

// 	json.NewEncoder(w).Encode(article)
// }

// func deleteArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit.. | Delete article page")

// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	for index, article := range Articles {
// 		if article.Id == id {
// 			Articles = append(Articles[:index], Articles[index+1:]...)
// 			response := struct {
// 				Message string `json:"message"`
// 			}{Message: "Article deleted successfully."}
// 			json.NewEncoder(w).Encode(response)
// 			return
// 		}
// 	}

// 	response := struct {
// 		Error string `json:"error"`
// 	}{Error: "Article not found."}
// 	json.NewEncoder(w).Encode(response)
// }

// func updateArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit.. | Update article page")

// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	reqBody, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Error reading request body", http.StatusBadRequest)
// 		return
// 	}

// 	var bodyData Article

// 	err = json.Unmarshal(reqBody, &bodyData)
// 	if err != nil {
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}

// 	for index, article := range Articles {
// 		if article.Id == id {
// 			Articles[index].Title = bodyData.Title
// 			Articles[index].Content = bodyData.Content
// 			Articles[index].Desc = bodyData.Desc

// 			json.NewEncoder(w).Encode(article)
// 			return
// 		}
// 	}

// 	http.Error(w, "Invalid id", http.StatusBadRequest)
// }

// // go router
// // func handleRequests() {
// // 	http.HandleFunc("/", homePage)
// // 	http.HandleFunc("/articles", returnAllArticles)

// // 	fmt.Println("Server is running on port 3000")
// // 	log.Fatal(http.ListenAndServe(":3000", nil))
// // }

// // github.com/gorilla/mux router
// func handleRequests() {
// 	myRouter := mux.NewRouter().StrictSlash(true)

// 	myRouter.HandleFunc("/", homePage)
// 	myRouter.HandleFunc("/articles", returnAllArticles)
// 	myRouter.HandleFunc("/article/{id}", returnSingleArticle).Methods("GET")
// 	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
// 	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
// 	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")

// 	fmt.Println("Server is running on port 3000")
// 	log.Fatal(http.ListenAndServe(":3000", myRouter))
// }

// func main() {

// 	Articles = []Article{
// 		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
// 		Article{Id: "2", Title: "Hyy", Desc: "Article Description", Content: "Article Content"},
// 		Article{Id: "3", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
// 		Article{Id: "4", Title: "Hyy", Desc: "Article Description", Content: "Article Content"},
// 		Article{Id: "5", Title: "Hyy", Desc: "Article Description", Content: "Article Content"},
// 	}

// 	handleRequests()
// }

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("<h1>Home page</h1>")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Articles)
}

func getOneArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, article := range Articles {
		if article.Id == id {
			json.NewEncoder(w).Encode(article)
			return
		}
	}

	http.Error(w, "Please provide a valid id!!!", http.StatusBadRequest)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error while reading.", http.StatusBadRequest)
		return
	}

	var article Article
	err = json.Unmarshal(reqBody, &article)
	if err != nil {
		http.Error(w, "Invalid JSON/Invalid body", http.StatusBadRequest)
		return
	}

	for _, art := range Articles {
		if art.Title == article.Title {
			http.Error(w, "Title already exist", http.StatusBadRequest)
			return
		}
	}

	article.Id = strconv.Itoa(len(Articles) + 1)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error while reading", http.StatusBadRequest)
		return
	}

	var bodyData Article
	err = json.Unmarshal(reqBody, &bodyData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for index, article := range Articles {
		if article.Id == id {
			if bodyData.Content == "" {
				bodyData.Content = article.Content
			}
			if bodyData.Desc == "" {
				bodyData.Desc = article.Desc
			}
			if bodyData.Title == "" {
				bodyData.Title = article.Title
			}
			Articles[index].Content = bodyData.Content
			Articles[index].Desc = bodyData.Desc
			Articles[index].Title = bodyData.Title

			json.NewEncoder(w).Encode(bodyData)
			return
		}
	}

	http.Error(w, "Invalid id", http.StatusBadRequest)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			json.NewEncoder(w).Encode("Article deleted successfully")
			return
		}
	}

	http.Error(w, "Invalid id", http.StatusBadRequest)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/articles", getAllArticles).Methods("GET")
	myRouter.HandleFunc("/article/{id}", getOneArticle).Methods("GET")
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")

	fmt.Println("Server is running on port 3000")
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {

	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hyy", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "3", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "4", Title: "Hyy", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "5", Title: "Hyy", Desc: "Article Description", Content: "Article Content"},
	}

	handleRequests()
}
