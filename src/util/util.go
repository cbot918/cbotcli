package util

import (
	"fmt"
	"log"
)

func Checke(err error, message string) {
	if err != nil {
		log.Fatal(err)
		fmt.Println(message)
	}
}

func Logg(message interface{}) {
	fmt.Println(message)
}
