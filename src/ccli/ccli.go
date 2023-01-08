package ccli

import (
	"cbotcli/src/lib/httpcli"
	"os"
)

type Ccli struct {
	http       *httpcli.Httpcli
	configName string
	Host       string
	Port       int32
}

func NewCcli(host string, port int32) *Ccli {
	c := new(Ccli)
	cnames := os.Args[1:]
	cname := cnames[0]

	c.http = httpcli.NewHttpcli()
	c.configName = cname
	c.Host = host
	c.Port = port
	return c
}

func (c *Ccli) Cname() string {
	return c.configName
}

func (c *Ccli) Http() *httpcli.Httpcli {
	return c.http
}

// func (c *Ccli) Do() {
// 	h := httpcli.NewHttpcli()
// 	h.Get()
// }
