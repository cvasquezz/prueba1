package main

import (
	"fmt"

	external "github.com/cvasquezz/readEnvGO"
)

var wea = external.Wea{}

type User struct {
	external.User
	Apellido string
}

func main() {
	fmt.Printf("Hola Mundo \n")
	wea.Default = false
	user := User{}
	user.Nombre = "pepe"
	user.Apellido = "perez"
	wea.User = user

	user2 := external.User{}
	user2.Nombre = "Juan"

	wea.ReadEnv(&user2)
	fmt.Print("Fin")

}
