package main

import ("github.com/jaenmo/Hi/greeting"
		"github.com/jaenmo/Hi/greeting/Spanish"
		"github.com/sirupsen/logrus")

func main(){
	greeting.Hello()
	greeting.Hi()
	spanish.Hola()
	spanish.Adios()
	logrus.Println("Oh mininio")
}