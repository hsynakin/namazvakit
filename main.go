package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hsynakin/namazvakit/apiroot"
)

func main() {

	app := gin.Default()

	api := app.Group("/api")

	apiroot.Apiroot(api)

	app.Run(":1432")

	/*fmt.Println(res)
	fmt.Println(string(body))*/

}
