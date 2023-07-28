package app





import (
	"net/http"

	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func handleWelcomeShowPost(c *router.Context, guid string) {
	cols, editable := router.GetEditableCols(c, "welcome")
	list := []string{}
	for _, item := range cols {
		if router.IsEditable(item, editable) == false {
			continue
		}
		list = append(list, item)
	}
	list = append(list, "submit")
	c.ReadFormValuesIntoParams(list...)
	submit := c.Params["submit"].(string)
	if submit != "save" {
		//handleFooCreate(c, guid)
		return
	}

	c.ValidateUpdate("welcome")
	message := c.ValidateUpdate("welcome")
	returnPath := "/welcomes"
	if message != "" {
		router.SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath+"/"+guid, 302)
		return
	}
	message = c.Update("welcome", "where guid=", guid)
	if message != "" {
		router.SetFlash(c, message)
		http.Redirect(c.Writer, c.Request, returnPath+"/"+guid, 302)
		return
	}
	http.Redirect(c.Writer, c.Request, returnPath, 302)
}

func handleWelcomeShow(c *router.Context, guid string) {
	item := c.One("welcome", "where guid=$1", guid)
	regexMap := map[string]string{}
	cols, editable := router.GetEditableCols(c, "welcome")
	//cols = append(cols, "save")
	//editable["save"] = "save"

	colAttributes := map[int]string{}
	colAttributes[1] = "w-3/4"

	m := map[string]any{}
	headers := []string{"field", "value"}

	params := map[string]any{}
	params["item"] = item
	params["editable"] = editable
	params["regex_map"] = regexMap
	m["headers"] = headers
	m["cells"] = c.MakeCells(util.ToAny(cols), headers, params, "_welcome_show")
	m["col_attributes"] = colAttributes
	m["save"] = true
	topVars := map[string]any{}
	topVars["name"] = item["name"]
	topVars["guid"] = guid
	send := map[string]any{}
	send["bottom"] = c.Template("table_show.html", m)
	send["top"] = c.Template("welcomes_top.html", topVars)

	c.SendContentInLayout("generic_top_bottom.html", send, 200)
}