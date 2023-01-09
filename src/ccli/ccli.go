package ccli

import (
	"cbotcli/src/lib/httpcli"
	"os"
)

type Ccli struct {
	Http  *httpcli.Httpcli
	Cname string
	Host  string
	Port  int32
}

func NewCcli(host string, port int32) *Ccli {
	c := new(Ccli)
	cnames := os.Args[1:]
	cname := cnames[0]

	c.Http = httpcli.NewHttpcli()
	c.Cname = cname
	c.Host = host
	c.Port = port
	return c
}

// func (c *Ccli) Cname() string {
// 	return c.configName
// }

// func (c *Ccli) Http() *httpcli.Httpcli {
// 	return c.http
// }
