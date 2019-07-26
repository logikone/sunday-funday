package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/logikone/sunday-funday/store"
)

type Server struct {
	*gin.Engine
}

type User struct {
	Name     string `db:"user_name"`
	ID       string `db:"user_id"`
	Password string `db:"user_password"`
}

func (s Server) Run() error {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	engine.GET("/", rootHandler)

	db := store.Must(store.Connect("sqlite3", "database.sqlite"))

	var users []User

	if err := db.Select(&users, "SELECT * FROM users"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%#v\n", users)

	return engine.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
