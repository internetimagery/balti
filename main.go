package main

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

// Shared info across all pages.
const TITLE = "Balti Menu Order Form v1.0"


// A single order
type Order struct {
  Name string `form:"name" binding:"required"`
  Meal string `form:"meal" binding:"required"`
  Spice string `form:"spice" binding:"required"`
  Side string `form:"side" binding:"required"`
  Notes string `form:"notes"`
}



func main()  {


  // Store orders
  orders := make(map[string]Order)

  // Set dev mode
  // gin.SetMode(gin.ReleaseMode)
  gin.SetMode(gin.DebugMode)

  // Set up our router
  router := gin.Default()

  // Load our template files and initialize some global variables
  router.LoadHTMLGlob("templates/*.html")

  // Set up our routers
  router.GET("/", func(c *gin.Context) {
    // Menu
    c.HTML(http.StatusOK, "index.html", gin.H{
      "title": TITLE,
    })
  })
  router.POST("/", func(c *gin.Context) {
    id := c.ClientIP() // Collect identifying information for the person
    var order Order
    err := c.Bind(&order)
    if err == nil {
      orders[id] = order
      fmt.Println("Order placed by:", id, order)
    } else {
      fmt.Println("ERROR:", err)
    }

    // Place order
    c.HTML(http.StatusOK, "order.html", gin.H{
      "title": TITLE,
      "order": orders[id],
    })
  })
  router.GET("/view", func(c *gin.Context) {
    // View orders
    c.HTML(http.StatusOK, "view.html", gin.H{
      "title": TITLE,
      "orders": orders,
    })
  })
  // Load static files.
  router.Static("/static", "./static")



  // Start server on port "balt" :8411
  port := ":8411"
  fmt.Println("Collecting orders on port", port)
  router.Run(port)

}
