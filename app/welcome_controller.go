package app

import (
	"github.com/andrewarrow/feedback/router"
)

func HandleWelcome(c *router.Context, second, third string) {
	if second == "" && third == "" && c.Method == "GET" {
		handleWelcomeIndex(c)
		return
	}
	c.NotFound = true
}

func handleWelcomeIndex(c *router.Context) {
	topVars := map[string]any{}
	send := map[string]any{}
	send["top"] = c.Template("welcomes_list_top.html", topVars)
	c.SendContentInLayout("generic_top_bottom.html", send, 200)
}
