package app





import (
	"fmt"
	"net/http"
	"time"

	"github.com/andrewarrow/feedback/router"
)

func handleWelcomeCreate(c *router.Context) {
	//c.ReadFormValuesIntoParams("")

	returnPath := "/welcomes"

	now := time.Now().Unix()
	c.Params = map[string]any{}
	c.Params["user_id"] = c.User["id"]
	c.Params["name"] = fmt.Sprintf("Untitled %d", now)
	c.Params["street1"] = "123 Main St."
	c.Params["city"] = "Los Angeles"
	c.Params["state"] = "CA"
	c.Params["zip"] = "90066"
	c.Params["country"] = "USA"
	message := c.ValidateCreate("welcome")
	if message != "" {
		router.SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath, 302)
		return
	}
	message = c.Insert("welcome")
	if message != "" {
		router.SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath, 302)
		return
	}
	http.Redirect(c.Writer, c.Request, returnPath, 302)
}