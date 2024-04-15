package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"sync"
)

type Course struct {
	Name string `json:"name"`
}

var (
	courses = []*Course{
		{"Laajalahti"},
		{"Köykäri"},
		{"Kippassuo"},
		{"Laajvuori"},
		{"Beast"},
	}
	seq  = 1
	lock = sync.Mutex{} // what this is:D
)

func getAllCourses(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, courses)
}
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":1323"))
}
