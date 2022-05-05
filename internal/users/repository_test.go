/* package users

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubDB struct {
	beforeUpdate bool
}

func (stubDB *stubDB) Read(data interface{}) error {
	stubDB.beforeUpdate = true
	var users []User
	user1 := User{
		Id:              1,
		Nombre:          "Juan",
		Apellido:        "Tech",
		Email:           "",
		Edad:            10,
		Altura:          180.0,
		Activo:          nil,
		FechaDeCreacion: "",
	}
	user2 := User{
		Id:              2,
		Nombre:          "Pepe",
		Apellido:        "Tech",
		Email:           "",
		Edad:            10,
		Altura:          180.0,
		Activo:          nil,
		FechaDeCreacion: "",
	}
	users = append(users, user1, user2)

	dataEncoded, _ := json.Marshal(users)
	err := json.Unmarshal(dataEncoded, &data)
	if err != nil {
		return err
	}
	return nil
}

func (stubDB stubDB) Write(data interface{}) error {

	return nil
}

func TestGetAll(t *testing.T) {
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
	storeStub := stubDB{}
	repository := NewRepository(&storeStub)

	// cuando
	resultado, _ := repository.GetAll()

	// entonces
	assert.Equal(t, users, resultado)
}

func TestStoreRepository(t *testing.T) {
	// Dado
	userToSaveForTesting := User{
		Id:              1,
		Nombre:          "Usuario",
		Apellido:        "ToSave",
		Email:           "",
		Edad:            10,
		Altura:          200.0,
		Activo:          nil,
		FechaDeCreacion: "",
	}
	storeStub := stubDB{}
	repository := NewRepository(&storeStub)

	// cuando
	resultado, _ := repository.Store(userToSaveForTesting.Id, userToSaveForTesting.Nombre, userToSaveForTesting.Apellido, userToSaveForTesting.Email, userToSaveForTesting.FechaDeCreacion, userToSaveForTesting.Edad, userToSaveForTesting.Altura, userToSaveForTesting.Activo)

	// entonces
	assert.Equal(t, userToSaveForTesting, resultado)
}

func TestUpdateNameAndSurname(t *testing.T) {
	userUpdated := User{
		Id:              1,
		Nombre:          "Nuevo nombre",
		Apellido:        "Nuevo apellido",
		Email:           "",
		Edad:            10,
		Altura:          180.0,
		Activo:          nil,
		FechaDeCreacion: "",
	}
	storeStub := stubDB{beforeUpdate: false}
	repository := NewRepository(&storeStub)

	// cuando
	resultado, _ := repository.UpdateNameAndSurname(1, "Nuevo nombre", "Nuevo apellido")

	// entonces
	assert.Equal(t, userUpdated, resultado)
	assert.True(t, storeStub.beforeUpdate)
}

func TestLastId(t *testing.T) {
	// dado
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
	storeStub := stubDB{}
	repository := NewRepository(&storeStub)
	expected := len(users)

	// cuando
	resultado, _ := repository.LastID()

	// entonces
	assert.Equal(t, resultado, expected)
} */
