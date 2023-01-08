package etoe

import (
	"cbotcli/src/ccli"
	u "cbotcli/src/util"
	"fmt"
)

/***  execute etoe test with cmd: rspec  ***/

type Etoe struct{}

func NewEtoe() *Etoe {
	obj := new(Etoe)
	return obj
}

func (e *Etoe) Execute(c *ccli.Ccli) {
	/***  ete-test: basic case  ***/
	{
		// GET /m1/script/vim
		path1 := fmt.Sprintf("/m1/script/%s", c.Cname())
		url1 := fmt.Sprintf("%s:%d%s", c.Host, c.Port, path1)
		resp1 := c.Http().Get(url1)
		defer resp1.Body.Close()
		u.Logg(u.ReadResp(resp1))

		// POST /m1/signup {"email":"yale@gmail.com", "password":"12345"}
		path2 := "/m1/signup"
		payload := `{"email":"yale@gmail.com", "password":"12345"}`
		url2 := fmt.Sprintf("%s:%d%s", c.Host, c.Port, path2)
		resp2 := c.Http().Post(url2, payload)
		defer resp2.Body.Close()
		u.Logg(u.ReadResp(resp2))
		u.Logg(fmt.Sprintf("%q", u.ReadResp(resp2)))
	}
}
