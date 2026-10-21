package main

import "fmt"

type Student struct{
	stu_id int
	name string
	Mark
}

type Mark struct{
	eng int
	math int
	sci int
}

func (m Mark) comp_avg() int{
	avg := (m.eng + m.math + m.sci)/3
	return avg
}

func main(){
	//Get the input from the user
	stu_list := []Student{}
	temp := Student{}
	var choice bool

	for true{
		fmt.Print("Enter the id: ")
		fmt.Scanln(&(temp.stu_id))
		fmt.Print("Enter the name")
		fmt.Scanln(&(temp.name))
		fmt.Print("Enter the eng mark")
		fmt.Scanln(&(temp.eng))
		fmt.Print("Enter the math mark")
		fmt.Scanln(&(temp.math))
		fmt.Print("Enter the sci mark")
		fmt.Scanln(&(temp.sci))

		stu_list = append(stu_list, temp)

		fmt.Println("Do you want to continue: ")
		fmt.Scanln(&choice)

		if choice == false{
			break
		}
	}
	fmt.Printf("%v \n", stu_list)
	// loop through the stu list to calc and display student info
	for _, val:= range stu_list{
		//display details
		fmt.Printf("Name: %v ", val.name)
		fmt.Printf("Avg: %v \n", val.comp_avg())
	}

}