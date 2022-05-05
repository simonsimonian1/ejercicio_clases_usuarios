package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/simonsimonian1/modulo7_GoWeb/clase3/C3-tm-estructuraDeCapas/EjerciciosPersonales/grupo4_clase7_tm/internal/users"
	"github.com/simonsimonian1/modulo7_GoWeb/clase3/C3-tm-estructuraDeCapas/EjerciciosPersonales/grupo4_clase7_tm/pkg/web"
)

type user struct {
	Nombre          string  `form:"nombre" json:"nombre"`
	Apellido        string  `form:"apellido" json:"apellido"`
	Email           string  `form:"email" json:"email"`
	Edad            int     `form:"edad" json:"edad"`
	Altura          float64 `form:"altura" json:"altura"`
	Activo          *bool   `form:"activo" json:"activo"`
	FechaDeCreacion string  `form:"fecha_de_creacion" json:"fecha_de_creacion"`
}

type User struct {
	service users.Service
}

// ListUsers godoc
// @Summary List users
// @Tags Users
// @Description get users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /users [get]
func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/* token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		} */

		allUsers, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		if len(allUsers) == 0 {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "No hay usuarios cargados aún"))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, allUsers, ""))
	}
}

// StoreUser godoc
// @Summary Store users
// @Tags users
// @Description store users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body user true "User to store"
// @Success 200 {object} web.Response
// @Router /users [post]
func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/* token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "Token inválido"))
			return
		} */
		var userRequest user
		if err := ctx.ShouldBindJSON(&userRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		// Validaciones
		if userRequest.Nombre == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El nombre no puede ser vacío"))
			return
		}
		if userRequest.Apellido == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El apellido no puede ser vacío"))
			return
		}
		if userRequest.Email == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "El email no puede ser vacío"))
			return
		}
		if userRequest.FechaDeCreacion == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "La fecha no puede ser vacía"))
			return
		}
		if userRequest.Altura == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "La altura del usuario no puede ser cero"))
			return
		}
		if userRequest.Edad == 0 {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "La edad del usuario no puede ser cero"))
			return
		}
		if userRequest.Activo == nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Se debe indicar si el usuario está o no activo"))
			return
		}

		// Enviamos los datos del usuario al método store del service
		user, err := c.service.Store(userRequest.Nombre, userRequest.Apellido, userRequest.Email, userRequest.FechaDeCreacion, userRequest.Edad, userRequest.Altura, userRequest.Activo)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, user, ""))
	}
}

// UpdateUser godoc
// @Summary Update users
// @Tags users
// @Description update users
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body requestUser true "User to update"
// @Param id path int true "User ID"
// @Success 200 {object} web.Response
// @Router /users/{id} [put]
func (c *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		var userRequest user
		err = ctx.ShouldBindJSON(&userRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if userRequest.Nombre == "" || userRequest.Apellido == "" || userRequest.Email == "" || userRequest.FechaDeCreacion == "" || userRequest.Edad == 0 || userRequest.Altura == 0.0 || userRequest.Activo == nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Todos los campos son requeridos"))
			return
		}
		userSended, err := c.service.Update(id, userRequest.Nombre, userRequest.Apellido, userRequest.Email, userRequest.FechaDeCreacion, userRequest.Edad, userRequest.Altura, userRequest.Activo)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusAccepted, web.NewResponse(http.StatusAccepted, userSended, ""))
	}
}

// UpdateUserNameAndSurname godoc
// @Summary UpdateNameAndSurname users
// @Tags users
// @Description update username and usersurname
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body requestUser true "User to update"
// @Param id path int true "User ID"
// @Success 200 {object} web.Response
// @Router /users/{id} [put]
func (c *User) UpdateNameAndSurname() gin.HandlerFunc {
	type userToUpdateNameAndSurname struct {
		Nombre   string `form:"nombre" json:"nombre"`
		Apellido string `form:"apellido" json:"apellido"`
	}
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		var userRequest userToUpdateNameAndSurname
		err = ctx.ShouldBindJSON(&userRequest)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		if userRequest.Nombre == "" || userRequest.Apellido == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "Todos los campos son requeridos"))
			return
		}
		userSended, err := c.service.UpdateNameAndSurname(id, userRequest.Nombre, userRequest.Apellido)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusAccepted, web.NewResponse(http.StatusAccepted, userSended, ""))
	}
}

func (c *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusAccepted, web.NewResponse(http.StatusAccepted, nil, ""))
	}
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}
