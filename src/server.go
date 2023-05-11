package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

type App struct {
	DB *sql.DB
}

// items will return all the items from the database.
func (app *App) items(c *gin.Context) {
	items, err := GetItems(app.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, items)
	}
}

// add will add the specified item to the database.
func (app *App) add(c *gin.Context) {
	var item Todo

	// Receive the item
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add to database
	if err := AddItem(app.DB, &item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.Status(http.StatusOK)
	}
}

func (app *App) delete(c *gin.Context) {
	// Get the id and convert it to an int
	id, _ := strconv.Atoi(c.Param("id"))
	if err := DeleteItem(app.DB, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.Status(http.StatusOK)
	}
}

func (app *App) update(c *gin.Context) {
	// Get the id and convert it to an int
	id, _ := strconv.Atoi(c.Param("id"))

	// Receive the item
	var item Todo
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the item
	if err := UpdateItem(app.DB, id, &item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.Status(http.StatusOK)
	}
}

func (app *App) Initialize() error {
	// Set up db
	var err error
	app.DB, err = connect()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// Set up logger
	log.New(os.Stdout, "gin: ", log.LstdFlags|log.Lshortfile)

	// Initialize app
	app := App{}
	if err := app.Initialize(); err != nil {
		log.Fatal(err)
	}

	// Creates a gin router with default middleware
	r := gin.Default()

	r.GET("/getItems", app.items)
	r.POST("/addItem", app.add)
	r.DELETE("/deleteItem/:id", app.delete)
	r.PATCH("/updateItem/:id", app.update)

	// listen and serve
	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
