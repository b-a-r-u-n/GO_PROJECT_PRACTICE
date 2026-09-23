package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Todo struct {
	Id        int
	Title     string
	Completed bool
}

func addTodo(todoId int, todoTitle string, todos *[]Todo) {
	todo := Todo{
		Id:        todoId,
		Title:     todoTitle,
		Completed: false,
	}

	*todos = append(*todos, todo)
}

func removeTodo(todoId int, todos *[]Todo) error {

	for _, todo := range *todos {
		if todo.Id == todoId {
			*todos = slices.DeleteFunc(*todos, func(todo Todo) bool {
				return todo.Id == todoId
			})
			return nil
		}
	}
	return fmt.Errorf("Todo not found")
}

func updateTodo(todoId int, todos *[]Todo) error {
	for index, todo := range *todos {
		if todo.Id == todoId {
			(*todos)[index].Completed = true
			return nil
		}
	}
	return fmt.Errorf("Todo not found")
}

func showMenu() {
	fmt.Printf("\n\n")
	fmt.Println("1. Add todo")
	fmt.Println("2. Remove todo")
	fmt.Println("3. Complete todo")
	fmt.Println("4. View all todos")
	fmt.Println("5. Exit")
	fmt.Printf("Enter your choice: ")
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
		if choice == "1" {
			fmt.Printf("\n\n")
			fmt.Printf("Enter the todo title: ")
			todoTitle, readerErr := reader.ReadString('\n')
			if readerErr != nil {
				fmt.Println("Error while reading |", readerErr)
				return
			}
			todoTitle = strings.TrimSpace(todoTitle)
			addTodo(todoId, todoTitle, &todos)

			showMenu()
			var choiceErr error
			choice, choiceErr = reader.ReadString('\n')
			if choiceErr != nil {
				fmt.Println("Error while reading |", choiceErr)
				break
			}
			choice = strings.TrimSpace(choice)
			todoId++
		} else if choice == "2" {
			if len(todos) == 0 {
				fmt.Println("You don't have todo to remove | Please add some.. ")
				choice = "1"
				continue
			}
			fmt.Printf("Enter the id of the todo you want to remove: ")
			// var removeTodoId int
			// fmt.Scan(&removeTodoId)
			removeTodoId, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error while reading |", err)
				break
			}
			removeTodoId = strings.TrimSpace(removeTodoId)
			var removeTodoIdINTErr error
			var removeTodoIdINT int
			removeTodoIdINT, removeTodoIdINTErr = strconv.Atoi(removeTodoId)
			if removeTodoIdINTErr != nil {
				fmt.Println("Error while intiger conversion |", removeTodoIdINTErr)
			}
			todoRemoveError := removeTodo(removeTodoIdINT, &todos)
			if todoRemoveError != nil {
				fmt.Println(todoRemoveError)
				fmt.Println("Wrong id!! Please choose a valid id")
				continue
			}
			fmt.Println("Todo removed successfully!!")

			showMenu()
			var choiceErr error
			choice, choiceErr = reader.ReadString('\n')
			if choiceErr != nil {
				fmt.Println("Error while reading |", choiceErr)
				break
			}
			choice = strings.TrimSpace(choice)
		} else if choice == "3" {
			if len(todos) == 0 {
				fmt.Println("You don't have todo to remove | Please add some.. ")
				choice = "1"
				continue
			}
			fmt.Printf("Enter the id of the todo you want to update: ")
			UpdateTodoId, UpdateTodoIdErr := reader.ReadString('\n')
			if UpdateTodoIdErr != nil {
				fmt.Println("Error while reading |", UpdateTodoIdErr)
			}
			UpdateTodoId = strings.TrimSpace(UpdateTodoId)
			var updateTodoIdINTErr error
			var updateTodoIdINT int
			updateTodoIdINT, updateTodoIdINTErr = strconv.Atoi(UpdateTodoId)
			if updateTodoIdINTErr != nil {
				fmt.Println("Error while intiger conversion |", updateTodoIdINTErr)
				break
			}
			todoUpdateError := updateTodo(updateTodoIdINT, &todos)
			if todoUpdateError != nil {
				fmt.Println(todoUpdateError)
				fmt.Println("Wrong id!! Please choose a valid id")
				continue
			}

			fmt.Println("Todo updated successfully!!")

			showMenu()
			var choiceErr error
			choice, choiceErr = reader.ReadString('\n')
			if choiceErr != nil {
				fmt.Println("Error while reading |", choiceErr)
				break
			}
			choice = strings.TrimSpace(choice)
		} else if choice == "4" {
			if len(todos) == 0 {
				fmt.Println("You don't have todos | Please add some.. ")
				choice = "1"
				continue
			}
			fmt.Println("Todos: ", todos)

			showMenu()
			var choiceErr error
			choice, choiceErr = reader.ReadString('\n')
			if choiceErr != nil {
				fmt.Println("Error while reading |", choiceErr)
				break
			}
			choice = strings.TrimSpace(choice)
		} else if choice == "5" {
			fmt.Println("Goodbye..")
			break
		} else {
			fmt.Println("Invalid choice choose again")

			showMenu()
			var choiceErr error
			choice, choiceErr = reader.ReadString('\n')
			if choiceErr != nil {
				fmt.Println("Error while reading |", choiceErr)
				break
			}
			choice = strings.TrimSpace(choice)
		}
	}
	// fmt.Println("Todos: ", todos)
}
