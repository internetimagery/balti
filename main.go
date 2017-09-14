package main

import (
  "fmt"
  "net/http"
  "github.com/gin-gonic/gin"
)

// Shared info across all pages.
const TITLE = "Balti Menu Order Form"


// A single order
type Order struct {
  Meal string `form:"meal" binding:"required"`
  Spice string `form:"spice" binding:"required"`
  Naan string `form:"naan" binding:"required"`
  Notes string `form:"notes"`
}



func main()  {

  orders := make(map[string]Order)

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
      fmt.Println("Your order:\n")
      fmt.Println("meal:", order.Meal)
      fmt.Println("spice:", order.Spice)
      fmt.Println("naan:", order.Naan)
      fmt.Println("notes:", order.Notes)
      fmt.Println(orders)
    } else {
      fmt.Println(err)
    }

    fmt.Println("Your IP is:", c.ClientIP())

    // Place order
    c.HTML(http.StatusOK, "order.html", gin.H{
      "title": TITLE,
    })
  })
  router.GET("/view", func(c *gin.Context) {
    // View orders
    c.HTML(http.StatusOK, "view.html", gin.H{
      "title": TITLE,
    })
  })



  // Start server on port 6070
  port := ":8080"
  router.Run(port)

}
