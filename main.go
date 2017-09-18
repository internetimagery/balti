package main

import (
  "os"
  "fmt"
  "strings"
  "io/ioutil"
  "net/http"
  "encoding/json"
  "path/filepath"
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

type Meal struct {
  Name string
  Desc string
}

type Category struct {
  Name string
  Meals []Meal
}

type Menu map[string]Category

type Food struct {
  Mains Menu
  Spice []string
  Side []string
}

func main()  {
  // Where are we?
  app_path, _ := os.Executable()
  root := filepath.Dir(app_path)

  // Load menu information
  data, err := ioutil.ReadFile(filepath.Join(root, "menu.json"))
  if err != nil {
    panic(err)
  }
  var menu Food
  err = json.Unmarshal(data, &menu)
  if err != nil {
    panic(err)
  }

  // Custom quote
  quote := strings.Join(os.Args[1:], " ")

  // Store orders
  orders := make(map[string]Order)

  // Set dev mode
  // gin.SetMode(gin.ReleaseMode)
  gin.SetMode(gin.DebugMode)

  // Set up our router
  router := gin.Default()

  // Load our template files and initialize some global variables

  router.LoadHTMLGlob(filepath.Join(root, "templates", "*.html"))

  // Set up our routers
  router.GET("/", func(c *gin.Context) {
    // Menu
    c.HTML(http.StatusOK, "index.html", gin.H{
      "title": TITLE,
      "quote": quote,
      "menu": menu,
    })
  })
  router.POST("/", func(c *gin.Context) {
    id := c.ClientIP() // Collect identifying information for the person
    var order Order
    err := c.Bind(&order)
    if err == nil {
      orders[id] = order
      fmt.Println("Order placed by:", id, order.Name, order)
      // Place order
      c.HTML(http.StatusOK, "order.html", gin.H{
        "title": TITLE,
        "order": orders[id],
        "quote": "Thank you for your order.",
      })
    } else {
      // Bad order. Ask to repeat.
      c.HTML(http.StatusOK, "index.html", gin.H{
        "title": TITLE,
        "quote": "There was an issue with your order. Please try again.",
      })
    }

  })
  router.GET("/view", func(c *gin.Context) {
    // View orders
    c.HTML(http.StatusOK, "view.html", gin.H{
      "title": TITLE,
      "orders": orders,
      "quote": "Below are all recent orders, for your booking pleasure.",
    })
  })
  // Load static files.
  router.Static("/static", filepath.Join(root, "static"))



  // Start server on port "balt" :8411
  port := ":8411"
  fmt.Println("Collecting orders on port", port)
  router.Run(port)
}
