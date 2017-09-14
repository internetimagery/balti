package main

import (
  // "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

func main()  {
  // Set up our router
  router := gin.Default()

  // Load our template files and initialize some global variables
  router.LoadHTMLGlob("templates/*.html")
  title := "Balti Menu Order"

  // Set up our routers
  router.GET("/", func(c *gin.Context) {
    // Menu
    c.HTML(http.StatusOK, "index.html", gin.H{
      "title": title,
    })
  })
  router.POST("/", func(c *gin.Context) {
    // Place order
    c.HTML(http.StatusOK, "order.html", gin.H{
      "title": title,
    })
  })
  router.GET("/view", func(c *gin.Context) {
    // View orders
    c.HTML(http.StatusOK, "view.html", gin.H{
      "title": title,
    })
  })



  // Start server on port 6070
  port := ":8080"
  router.Run(port)

}
