package controller

import "net/http"

func (c *BaseController) HelloWorld(w http.ResponseWriter, r *http.Request) {

	c.JsonReponse(w, http.StatusOK, c.CreateMessage("Hello, World!"))
}
