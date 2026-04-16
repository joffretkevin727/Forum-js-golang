package controller

import "fmt"

func Message() {
	Name := [3]string{"controller", "model", "router"}
	fmt.Println("voici " + Name[0] + ".go !")
}
