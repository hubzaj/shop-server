package main

import "fmt"

func main() {
	server := StartShop()
	address := server.Addr
	fmt.Println(address)
}
