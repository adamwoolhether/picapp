package controllers

import (
	"fmt"
	"log"
	"net/http"
	"picapp/models"
	"picapp/rand"
	"picapp/views"
)

// NewUsers creates a new Users controller. To be used during initial setup.
// If templates are incorrectly parsed, a panic will occur.
func NewUsers(us models.UserService) *Users {
	return &Users{
		NewView:   views.NewView("bootstrap", "users/new"),
		LoginView: views.NewView("bootstrap", "users/login"),
		us:        us,
	}
}

type Users struct {
	NewView   *views.View
	LoginView *views.View
	us        models.UserService
}

// New renders the form allowing users to create a new account
// GET /signup
func (g *Users) New(w http.ResponseWriter, r *http.Request) {
	g.NewView.Render(w, nil)
}

type SignupForm struct {
	Name     string `scheme:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create process the signup form after user submission
// POST /signup
func (g *Users) Create(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		log.Println(err)
		vd.SetAlert(err)
		g.NewView.Render(w, vd)
		return
	}
	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := g.us.Create(&user); err != nil {
		vd.SetAlert(err)
		g.NewView.Render(w, vd)
		return
	}
	if err := g.signIn(w, &user); err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Login verifies the provided email-addy & password, logging in the user if correct.
// POST /login
func (g *Users) Login(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	form := LoginForm{}
	if err := parseForm(r, &form); err != nil {
		log.Println(err)
		g.NewView.Render(w, vd)
		return
	}

	user, err := g.us.Authenticate(form.Email, form.Password)
	if err != nil {
		switch err {
		case models.ErrNotFound:
			vd.AlertError("Email address not found")
		default:
			vd.SetAlert(err)
		}
		g.LoginView.Render(w, vd)
		return
	}

	err = g.signIn(w, user)
	if err != nil {
		vd.SetAlert(err)
		g.LoginView.Render(w, vd)
		return
	}
	http.Redirect(w, r, "/cookietest", http.StatusFound)
}

// signIn signs in a given user after account creation and sets cookies
func (g *Users) signIn(w http.ResponseWriter, user *models.User) error {
	if user.Remember == "" {
		token, err := rand.RememberToken()
		if err != nil {
			return err
		}
		user.Remember = token
		err = g.us.Update(user)
		if err != nil {
			return err
		}
	}
	cookie := http.Cookie{
		Name:     "remember_token",
		Value:    user.Remember,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie) // Must write cookie to http.ResponseWriter before writing with Fprint
	return nil
}

// CookieTest displays the cookies set on the current user
func (g *Users) CookieTest(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("remember_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := g.us.ByRemember(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, user)
}
