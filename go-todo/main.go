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

func main() {

	fmt.Println("========= TODO APP =========")

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

	var todos []Todo

	choose := "y"
	reader := bufio.NewReader(os.Stdin)
	todoId := 1

	for {
		if choose == "y" || choose == "Y" {
			fmt.Printf("Enter the todo title: ")
			todoTitle, _ := reader.ReadString('\n')
			todoTitle = strings.TrimSpace(todoTitle)

			todo := Todo{
				Id:        todoId,
				Title:     todoTitle,
				Completed: false,
			}

			todos = append(todos, todo)
			todoId++
			fmt.Println("Do you want to add new todo (y/n)")
			// fmt.Scan(&choose)
			choose, _ = reader.ReadString('\n')
			choose = strings.TrimSpace(choose)
		} else if choose == "n" || choose == "N" {
			break
		} else {
			fmt.Println("Choose a valid choice (y/n)")
			// fmt.Scan(&choose)
			choose, _ = reader.ReadString('\n')
			choose = strings.TrimSpace(choose)
		}
	}

	fmt.Println(todos)
}
