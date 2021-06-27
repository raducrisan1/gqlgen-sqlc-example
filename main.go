package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raducrisan1/gqlgen-sqlc-example/gqlgen"
	"github.com/raducrisan1/gqlgen-sqlc-example/pg"
)

func main() {
	db, err := pg.Open("dbname=gqlgen_sqlc_example_db sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// initialize the repository
	repo := pg.NewRepository(db)

	r := gin.Default()
	r.POST("/query", gin.WrapH(gqlgen.NewHandler(repo)))
	r.GET("/", gin.WrapH(gqlgen.NewPlaygroundHandler("/query")))
	r.Run()
}
