package main

import (
  "time"
	"database/sql"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	_ "github.com/mattn/go-sqlite3"
)

type Reservation struct {
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
}

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	db, err := sql.Open("sqlite3", "./wedding-lulab-net.db")
	if err != nil {
		panic(err)
	}

	app.Post("/api/reservations", func(ctx iris.Context) {
		var reservation Reservation
		ctx.ReadJSON(&reservation)

		stmt, _ := db.Prepare("INSERT INTO reservations(name, mobile, created) values(?,?,?)")
		_, err := stmt.Exec(reservation.Name, reservation.Mobile, time.Now().Format("2006-01-02 15:04:05"))
		if err != nil {
			panic(err)
    }
    ctx.Writef("Reservation: %s", reservation.Name)

    ctx.Application().Logger().Infof("Reservation %s, %s", reservation.Name, reservation.Mobile)
	})

	// Start the server using a network address.
	app.Run(iris.Addr(":19090"), iris.WithoutServerError(iris.ErrServerClosed))
}
