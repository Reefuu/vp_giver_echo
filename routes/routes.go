package routes

import (
	"net/http"
	"vp_giver_echo/controllers"
	// "vp_giver_echo/middleware"

	"github.com/labstack/echo/v4"
)

// func getUser(c echo.Context) error {
// 	// User ID from path `users/:id`
// 	name := c.Param("name")
// 	return c.String(http.StatusOK, "Hello, "+name)
// }

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Guys`!")
	})

	// e.GET("/user", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, this is user page!!")
	// })

	// e.GET("user/:name", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, this is user page!")
	// })

	// e.GET("/user/:name", getUser)

	e.GET("/pelajaran", controllers.FetchAllPelajaran)

	e.GET("/pelajaran/kelas/:kelas", controllers.FetchPelajaranKls)

	e.POST("/pelajaran", controllers.StorePelajaran)

	// e.PATCH("/mahasiswa", controllers.UpdateMahasiswa)

	// e.DELETE("/mahasiswa", controllers.DeleteMahasiswa)

	// e.GET("generate-hash/:password", controllers.GenerateHashPassword)

	// e.POST("/login", controllers.CheckLogin) //tambah

	return e

}
