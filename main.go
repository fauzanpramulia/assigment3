package main

//Based Template Repo @Nurfaizal
//https://github.com/mochamadnurfaizal/golang-third-assignment
//Modified By Fauzan
import (
	"assigment3/controllers"
	"assigment3/database"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

const PORT = ":8088"

func main() {
	database.Connect()

	ticker := time.NewTicker(15 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				controllers.UpdateData()
			}
		}
	}()
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Aplikasi Buat Baca Status Udara dan Status Air"
		return ctx.String(http.StatusOK, data)
	})
	r.POST("/update/:id", controllers.UpdateEnvirontmentByGorm)

	r.Logger.Fatal(r.Start(PORT))
}
