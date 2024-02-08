package main

import (
    "net/http"
	"github.com/gin-gonic/gin"
	"log"
)

// getAlbums responds with the list of all albums as JSON.
func passwordcheck(c *gin.Context) {

	//Assume this call hit the database
	//and cross checking uname and pwd from okta pareter

	var jsonResponse = "{\"commands\": [{\"type\": \"com.okta.action.update\", \"value\": {\"credential\": \"VERIFIED\"}}]}";
    c.IndentedJSON(http.StatusOK, jsonResponse)
}

func main() {
    router := gin.Default()
    router.GET("/", passwordcheck)
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", router)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
