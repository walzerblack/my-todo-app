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


	func main(){

		fmt.Println("Welcome to the **|ToDo App|** .")




	}