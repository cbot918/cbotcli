package main

import (
	"cbotcli/src/ccli"
	e "cbotcli/test/etoe"
)

const (
	HOST = "http://127.0.0.1"
	PORT = 8181
)

func main() {

	c := ccli.NewCcli(HOST, PORT)

	e.NewEtoe().Execute(c)

}
