package main

import (
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
