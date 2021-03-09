package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
	"lenslocked.com/views"
)

func NewUsersController() *User {
	return &User{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

func (u *User) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	decoder := schema.NewDecoder()
	form := SignupForm{}
	if err := decoder.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}

type User struct {
	NewView *views.View
}

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}
