package users

import (
	//"fmt"

	"fmt"

	"github.com/simonsimonian1/modulo7_GoWeb/clase3/C3-tm-estructuraDeCapas/EjerciciosPersonales/grupo4_clase7_tm/pkg/store"
)

type User struct {
	Id              int     `form:"id" json:"id"`
	Nombre          string  `form:"nombre" json:"nombre"`
	Apellido        string  `form:"apellido" json:"apellido"`
	Email           string  `form:"email" json:"email"`
	Edad            int     `form:"edad" json:"edad"`
	Altura          float64 `form:"altura" json:"altura"`
	Activo          *bool   `form:"activo" json:"activo"`
	FechaDeCreacion string  `form:"fecha_de_creacion" json:"fecha_de_creacion"`
}

var users []User

type Repository interface {
	GetAll() ([]User, error)
	Store(id int, nombre, apellido, email, fechaDeCreacion string, edad int, altura float64, activo *bool) (User, error)
	LastID() (int, error)
	Update(id int, nombre string, apellido string, email string, fechaDeCreacion string, edad int, altura float64, activo *bool) (User, error)
	UpdateNameAndSurname(id int, nombre string, apellido string) (User, error)
	Delete(id int) error
}

type repository struct {
	dataBase store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		dataBase: db,
	}
}

/* type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
} */

func (r *repository) GetAll() ([]User, error) {
	err := r.dataBase.Read(&users)
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

func (r *repository) LastID() (int, error) {
	var users []User
	if err := r.dataBase.Read(&users); err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, nil
	}
	return users[len(users)-1].Id, nil
}

func (r *repository) Store(id int, nombre, apellido, email, fechaDeCreacion string, edad int, altura float64, activo *bool) (User, error) {
	var users []User

	// Intento leer primero
	err := r.dataBase.Read(&users)
	if err != nil {
		return User{}, err
	} // coloco en la lista de usuario lo que leo del json
	u := User{id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion} // instancio un usuario
	users = append(users, u)                                                      // appendo a la lista de usuarios el último usuario
	if err := r.dataBase.Write(users); err != nil {                               // intento escribir en el archivo json el nuevo usuario
		return User{}, err
	}

	return u, nil
}

func (r *repository) Update(id int, nombre, apellido, email, fechaDeCreacion string, edad int, altura float64, activo *bool) (User, error) {
	u := User{id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion}
	updated := false
	err := r.dataBase.Read(&users)
	if err != nil {
		return User{}, err
	}
	tamañoUsers := len(users)
	fmt.Print(tamañoUsers)
	for i := range users {
		if users[i].Id == u.Id {
			users[i] = u
			if err := r.dataBase.Write(users); err != nil { // intento escribir en el archivo json el nuevo usuario
				return User{}, err
			}
			updated = true
		}
	}
	if !updated {
		return User{}, fmt.Errorf("Usuario número %d no encontrado", id)
	}
	return u, nil
}

func (r *repository) UpdateNameAndSurname(id int, nombre, apellido string) (User, error) {
	var user User
	updated := false
	err := r.dataBase.Read(&users)
	if err != nil {
		return User{}, err
	}

	for i := range users {
		//fmt.Printf("User i --> %+v  \n Id --> %d \n", users[i], i)
		if users[i].Id == id {
			users[i].Nombre = nombre
			users[i].Apellido = apellido
			if err := r.dataBase.Write(users); err != nil { // intento escribir en el archivo json el nuevo usuario
				return User{}, err
			}
			user = users[i]
			updated = true
		}
	}
	if !updated {
		return User{}, fmt.Errorf("Usuario número %d no encontrado", id)
	}
	return user, nil
}

func (r *repository) Delete(id int) error {
	var index int
	localDeleted := false
	err := r.dataBase.Read(&users)
	if err != nil {
		return err
	}
	for i := range users {
		if users[i].Id == id {
			index = i
			localDeleted = true
			break
		}
	}
	if !localDeleted {
		return fmt.Errorf("Usuario número %d no encontrado", id)
	}
	users = append(users[:index], users[index+1:]...)
	if err = r.dataBase.Write(&users); err != nil { // intento escribir en el archivo json la nueva lista de usuarios pero con el usuario eliminado
		return err
	}
	return nil
}
