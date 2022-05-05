package main

import (
	"errors"
	"fmt"
	"os"

	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/simonsimonian1/modulo7_GoWeb/clase3/C3-tm-estructuraDeCapas/EjerciciosPersonales/grupo4_clase7_tm/cmd/server/handler"
	"github.com/simonsimonian1/modulo7_GoWeb/clase3/C3-tm-estructuraDeCapas/EjerciciosPersonales/grupo4_clase7_tm/docs"
	"github.com/simonsimonian1/modulo7_GoWeb/clase3/C3-tm-estructuraDeCapas/EjerciciosPersonales/grupo4_clase7_tm/internal/users"
	"github.com/simonsimonian1/modulo7_GoWeb/clase3/C3-tm-estructuraDeCapas/EjerciciosPersonales/grupo4_clase7_tm/pkg/auth"
	"github.com/simonsimonian1/modulo7_GoWeb/clase3/C3-tm-estructuraDeCapas/EjerciciosPersonales/grupo4_clase7_tm/pkg/store"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/go-sql-driver/mysql"
)

var (
	runErr    = errors.New("Che te esta rompiendo cuando lo corres")
	StorageDB *sql.DB
)

// @title API usuarios
// @version 1.0
// @description Esta API maneja la informacion de los usuarios.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	_ = godotenv.Load() // carga del atchivo .env
	// Conexion con la bdd
	/* dataSource := "root@tcp(localhost:3306)/ejercicioREST_usuarios"
	dataBase, _ := sql.Open("mysql", dataSource) */
	database := store.New(store.FileType, "")
	repository := users.NewRepository(database)
	service := users.NewService(repository)
	userController := handler.NewUser(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := r.Group("/users")
	users.Use(auth.TokenAuthMiddleware()) // Autenticamos previo a cualquier request
	users.POST("/", userController.Store())
	users.GET("/", userController.GetAll())
	users.PUT("/", userController.Update())
	users.PATCH("/", userController.UpdateNameAndSurname())
	users.DELETE("/", userController.Delete())
	err := r.Run()
	if err != nil {
		fmt.Print(runErr)
	}
}

/* func init() {

	// Open inicia un pool de conexiones. Sólo abrir una vez
	var err error
	StorageDB, err =
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database Configured")
} */
