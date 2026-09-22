package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Id        int
	Title     string
	Completed bool
}

func addTodo(todoId int, todoTitle string, todos *[]Todo){
	todo := Todo{
		Id: todoId,
		Title: todoTitle,
		Completed: false,
	}

	*todos = append(*todos, todo)
}

func main() {

	fmt.Println("========= TODO APP =========")

	//                              METHOD - 1
	// todoOne := Todo{
	// 	Id:        1,
	// 	Title:     "Todo one title",
	// 	Completed: false,
	// }

	// todoTwo := Todo{
	// 	Id:        2,
	// 	Title:     "Todo two title",
	// 	Completed: false,
	// }

	// todoThree := Todo{
	// 	Id:        3,
	// 	Title:     "Todo three title",
	// 	Completed: true,
	// }

	// todos := []Todo{todoOne, todoTwo, todoThree}

	// for _, todo := range todos {
	// 	fmt.Printf("Todo Id: %d, Todo Title: %s, Todo Completed: %t\n",todo.Id, todo.Title, todo.Completed)
	// }

	//                              METHOD - 2
	// var todos []Todo

	// choose := "y"
	// reader := bufio.NewReader(os.Stdin)
	// todoId := 1

	// for {
	// 	if choose == "y" || choose == "Y" {
	// 		fmt.Printf("Enter the todo title: ")
	// 		todoTitle, _ := reader.ReadString('\n')
	// 		todoTitle = strings.TrimSpace(todoTitle)

	// 		todo := Todo{
	// 			Id:        todoId,
	// 			Title:     todoTitle,
	// 			Completed: false,
	// 		}

	// 		todos = append(todos, todo)
	// 		todoId++
	// 		fmt.Println("Do you want to add new todo (y/n)")
	// 		// fmt.Scan(&choose)
	// 		choose, _ = reader.ReadString('\n')
	// 		choose = strings.TrimSpace(choose)
	// 	} else if choose == "n" || choose == "N" {
	// 		break
	// 	} else {
	// 		fmt.Println("Choose a valid choice (y/n)")
	// 		// fmt.Scan(&choose)
	// 		choose, _ = reader.ReadString('\n')
	// 		choose = strings.TrimSpace(choose)
	// 	}
	// }

	// fmt.Println(todos)

	// var todos []Todo
	// todoId := 1
	// choice := "y"
	// reader := bufio.NewReader(os.Stdin)

	// for {
	// 	if choice == "y" || choice == "Y"{
	// 		fmt.Printf("Enter the todo title: ")
	// 		todoTitle, readerErr := reader.ReadString('\n')
	// 		if readerErr != nil {
	// 			fmt.Println("Error while reading |", readerErr)
	// 			return
	// 		}
	// 		todoTitle = strings.TrimSpace(todoTitle)
	// 		addTodo(todoId, todoTitle, &todos)

	// 		fmt.Printf("Do you want to add another todo (y/n): ")
	// 		var choiceErr error
	// 		choice, choiceErr = reader.ReadString('\n')
	// 		if choiceErr != nil{
	// 			fmt.Println("Error while reading |", choiceErr)
	// 			break
	// 		}
	// 		choice = strings.TrimSpace(choice)
	// 		todoId++
	// 	} else if choice == "n" || choice == "N"{
	// 		break
	// 	} else {
	// 		fmt.Printf("Invalid choice choose again (y/n): ")
	// 		var choiceErr error
	// 		choice, choiceErr = reader.ReadString('\n')
	// 		if choiceErr != nil{
	// 			fmt.Println("Error while reading |", choiceErr)
	// 			break
	// 		}
	// 		choice = strings.TrimSpace(choice)
	// 	}
	// }
	// fmt.Println("Todos: ", todos)


	var todos []Todo
	todoId := 1
	choice := "1"
	reader := bufio.NewReader(os.Stdin)

	for {
		if choice == "1"{
			fmt.Printf("Enter the todo title: ")
			todoTitle, readerErr := reader.ReadString('\n')
			if readerErr != nil {
				fmt.Println("Error while reading |", readerErr)
				return
			}
			todoTitle = strings.TrimSpace(todoTitle)
			addTodo(todoId, todoTitle, &todos)

			fmt.Println("1. Add todo")
			fmt.Println("2. Remove todo")
			fmt.Println("3. Update todo")
			fmt.Println("4. Exit")
			fmt.Printf("Enter your choice: ")
			var choiceErr error
			choice, choiceErr = reader.ReadString('\n')
			if choiceErr != nil{
				fmt.Println("Error while reading |", choiceErr)
				break
			}
			choice = strings.TrimSpace(choice)
			todoId++
		} else if choice == "4" {
			fmt.Println("Goodbye..")
			break
		} else {
			fmt.Println("Invalid choice choose again")
			fmt.Println("1. Add todo")
			fmt.Println("2. Remove todo")
			fmt.Println("3. Update todo")
			fmt.Println("4. Exit")
			fmt.Printf("Enter your choice: ")
			var choiceErr error
			choice, choiceErr = reader.ReadString('\n')
			if choiceErr != nil{
				fmt.Println("Error while reading |", choiceErr)
				break
			}
			choice = strings.TrimSpace(choice)
		}
	}
	fmt.Println("Todos: ", todos)
}
