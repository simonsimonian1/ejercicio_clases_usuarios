/* package users

import (
	"encoding/json"
	"testing"

	"github.com/simonsimonian1/modulo7_GoWeb/clase3/C3-tm-estructuraDeCapas/EjerciciosPersonales/grupo4_clase7_tm/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceGetAll(t *testing.T) {
	users := []User{
		{
			Id:              1,
			Nombre:          "Juan",
			Apellido:        "Tech",
			Email:           "",
			Edad:            10,
			Altura:          180.0,
			Activo:          nil,
			FechaDeCreacion: "",
		}, {
			Id:              2,
			Nombre:          "Pepe",
			Apellido:        "Tech",
			Email:           "",
			Edad:            10,
			Altura:          180.0,
			Activo:          nil,
			FechaDeCreacion: "",
		},
	}
	dataJson, _ := json.Marshal(users)
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, users, result)
	assert.Nil(t, err)
}

func TestStore(t *testing.T) {
	user := User{
		Nombre:          "Juan",
		Apellido:        "Tech",
		Email:           "",
		Edad:            10,
		Altura:          180.0,
		Activo:          nil,
		FechaDeCreacion: "",
	}

	encodedData, _ := json.Marshal([]User{})
	dbMock := store.Mock{
		Data: encodedData,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)
	result, _ := myService.Store(user.Nombre, user.Apellido, user.Email, user.FechaDeCreacion, user.Edad, user.Altura, user.Activo)

	assert.Equal(t, user.Nombre, result.Nombre)
	assert.Equal(t, user.Apellido, result.Apellido)
	assert.Equal(t, user.Email, result.Email)
	assert.Equal(t, user.FechaDeCreacion, result.FechaDeCreacion)
	assert.Equal(t, user.Edad, result.Edad)
	assert.Equal(t, user.Altura, result.Altura)
	assert.Equal(t, user.Activo, result.Activo)
	assert.Equal(t, 1, result.Id)
}

func TestServiceUpdate(t *testing.T) {
	// Dado
	user := User{
		Id:              1,
		Nombre:          "Juan",
		Apellido:        "Tech",
		Email:           "",
		Edad:            10,
		Altura:          180.0,
		Activo:          nil,
		FechaDeCreacion: "",
	}

	userToUpdate := User{
		Id:              1,
		Nombre:          "Juan Nuevo",
		Apellido:        "Tech Nuevo",
		Email:           "Nuevo",
		Edad:            10,
		Altura:          180.0,
		Activo:          nil,
		FechaDeCreacion: "",
	}

	dataJson, _ := json.Marshal([]User{user})
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	// Cuando
	result, _ := myService.Update(userToUpdate.Id, userToUpdate.Nombre, userToUpdate.Apellido, userToUpdate.Email, userToUpdate.FechaDeCreacion, userToUpdate.Edad, userToUpdate.Altura, userToUpdate.Activo)

	// Entonces
	assert.Equal(t, userToUpdate, result)
}

func TestServiceDelete(t *testing.T) {
	// Dado
	users := []User{
		{
			Id:              1,
			Nombre:          "Juan",
			Apellido:        "Tech",
			Email:           "",
			Edad:            10,
			Altura:          180.0,
			Activo:          nil,
			FechaDeCreacion: "",
		}, {
			Id:              2,
			Nombre:          "Pepe",
			Apellido:        "Tech",
			Email:           "",
			Edad:            10,
			Altura:          180.0,
			Activo:          nil,
			FechaDeCreacion: "",
		},
	}
	dataJson, _ := json.Marshal(users)
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)
	userExpected := []User{
		{
			Id:              1,
			Nombre:          "Juan",
			Apellido:        "Tech",
			Email:           "",
			Edad:            10,
			Altura:          180.0,
			Activo:          nil,
			FechaDeCreacion: "",
		},
	}
	var result []User

	// Cuando
	err := myService.Delete(len(users))
	assert.Nil(t, err)

	// Entonces
	if err != nil {
		err = json.Unmarshal(dbMock.Data, &result)
		if err == nil {
			assert.Equal(t, userExpected, result)
		}
	}

}

func TestServiceDeleteConIdInvalido(t *testing.T) {
	// Dado
	users := []User{
		{
			Id:              1,
			Nombre:          "Juan",
			Apellido:        "Tech",
			Email:           "",
			Edad:            10,
			Altura:          180.0,
			Activo:          nil,
			FechaDeCreacion: "",
		}, {
			Id:              2,
			Nombre:          "Pepe",
			Apellido:        "Tech",
			Email:           "",
			Edad:            10,
			Altura:          180.0,
			Activo:          nil,
			FechaDeCreacion: "",
		},
	}
	dataJson, _ := json.Marshal(users)
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)
	var result []User

	// Cuando
	err := myService.Delete(len(users) + 1)

	// Entonces
	if err == nil {
		err = json.Unmarshal(dbMock.Data, &result)
		if err == nil {
			assert.NotNil(t, err)
			assert.Equal(t, users, result)
		}
	}
}

func TestServiceUpdateNameAndSurname(t *testing.T) {
	// Dado
	user := User{
		Id:              1,
		Nombre:          "Juan",
		Apellido:        "Tech",
		Email:           "",
		Edad:            10,
		Altura:          180.0,
		Activo:          nil,
		FechaDeCreacion: "",
	}

	userToUpdate := User{
		Id:       1,
		Nombre:   "Juan Nuevo",
		Apellido: "Tech Nuevo",
	}

	dataJson, _ := json.Marshal([]User{user})
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)

	// Cuando
	result, _ := myService.UpdateNameAndSurname(userToUpdate.Id, userToUpdate.Nombre, userToUpdate.Apellido)

	// Entonces
	assert.Equal(t, userToUpdate.Nombre, result.Nombre)
	assert.Equal(t, userToUpdate.Apellido, result.Apellido)
}
*/