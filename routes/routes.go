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

	e.GET("/buku", controllers.FetchAllBuku)

	e.GET("/buku/pelajaran/:pelajaran", controllers.FetchBukuPljrn)

	e.POST("/buku", controllers.StoreBuku)

	e.GET("/bab", controllers.FetchAllBab)

	e.GET("/bab/buku/:buku", controllers.FetchBabBuku)

	e.POST("/bab", controllers.StoreBab)

	e.GET("/subbab", controllers.FetchAllSubbab)

	e.GET("/subbab/bab/:bab", controllers.FetchSubbabBab)

	e.POST("/subbab", controllers.StoreSubbab)

	// e.PATCH("/mahasiswa", controllers.UpdateMahasiswa)

	// e.DELETE("/mahasiswa", controllers.DeleteMahasiswa)

	e.GET("/users", controllers.FetchAllUsers)

	e.POST("/login", controllers.CheckLogin)

	e.POST("/register", controllers.RegisterNewAcc)

	e.PATCH("/koin", controllers.AddCoin)

	e.GET("/quiz", controllers.FetchAllQuiz)

	e.GET("/quiz/subbab/:subbab", controllers.FetchQuizSubbab)

	e.POST("/quiz", controllers.StoreQuiz)
	

	return e

}
