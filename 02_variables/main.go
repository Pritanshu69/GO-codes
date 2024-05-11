package main

import "fmt"

func main(){
	var username string  = "Pritanshu";
	fmt.Println(username);
	fmt.Printf("Variable is of type: %T \n",username);

	var isLoggedin bool  = true;
	fmt.Println(isLoggedin);
	fmt.Printf("Variable is of type: %T \n",isLoggedin);
}