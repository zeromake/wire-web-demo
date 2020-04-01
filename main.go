package main

import (
	"fmt"
)

func main() {
	server, err := NewServer()
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}
}
