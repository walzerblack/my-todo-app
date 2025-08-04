package main

import (
	"fmt"
	"time"
)



	type States struct{
			Todo bool
			Inprogress bool
			Done bool
	}

	type Properties struct{
		Id uint32
		Description string
		Status *States
		CreatedAt time.Time
		UpdatedAt *time.Time
	}

	type Todos []Properties

	func Getstatus(s *States) string {
		if s.Todo{
			return "To Do"
		}else if s.Inprogress{
			return "InProgress"
		}else if s.Done {
			return "Done"
		}
		return "Unkown Please pick a Valid one"
	}


	func main(){
		fmt.Println("                          ")
		fmt.Println("                    ----------------------------------------------")
		fmt.Println("                         Welcome to your first **|ToDo App|** .")
		fmt.Println("                           *_* we got so far didn't We ")
		fmt.Println("                             **|Current Tasks Todo|**")
		fmt.Println("                    ----------------------------------------------")
		fmt.Printf("                           \n%-5s %-25s %-15s %-20s\n", "ID", "Description", "Status", "Created At")
		fmt.Println("-------------------------------------------------------------------------------")
		




	}