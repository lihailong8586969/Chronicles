package user

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/tghastings/blog/config/db"
	"golang.org/x/crypto/bcrypt"
)

var mySigningKey = []byte("pleasedonthackmebro")

// Message describes the JSON object message
type Message struct {
	Type    string
	Message string
}

// APIToken is used to generate an API JSON message
type APIToken struct {
	APIToken string
}

// JSONTOKEN is used to make the JWT
type JSONTOKEN struct {
	Token string
}

// User describes the user database
type User struct {
	gorm.Model
	Username string
	Password string
	APIToken string
	Email    string
}

// Cookie describes the cookies that we use to store session information
type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string

	// MaxAge=0 means no 'Max-Age' attribute specified.
	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
	// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HTTPOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}

// Index displays all users
func Index(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.DB.Find(&users)
	js, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Create creates a new user
func Create(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Fatal(err)
	}
	// Make token
	token := tokenGenerator()
	JSONtoken := APIToken{token}
	js, err := json.Marshal(JSONtoken)
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	fmt.Println(newUser.Password)
	newUser.Password = hashAndSalt([]byte(newUser.Password))
	fmt.Println(newUser.Password)
	err = json.Unmarshal(js, &newUser)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if err != nil {
		// handle error
		log.Fatal(err)
	}
	// Create
	db.DB.Create(&newUser)

	//json resp
	msg := Message{"success", "User added."}
	js, err = json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Show shows all users
func Show(w http.ResponseWriter, r *http.Request) {
	var user User
	userID := strings.TrimPrefix(r.URL.Path, "/admin/user/")
	if userID == "" {
		log.Println("ID Param 'key' is missing")
		return
	}
	db.DB.Find(&user, userID)
	js, err := json.Marshal(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Update updates one user record by id
func Update(w http.ResponseWriter, r *http.Request) {
	var user User
	userID := strings.TrimPrefix(r.URL.Path, "/admin/user/update/")
	if userID == "" {
		log.Println("ID Param 'key' is missing")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// fmt.Println(user)
	db.DB.Model(&user).Where("ID = ?", userID).Updates(user)
	//json resp
	msg := Message{"success", user.Username + "'s account updated."}
	js, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Delete removes a user by id
func Delete(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// fmt.Println(user.ID)
	// unmarshal content to ApplicantJSON
	db.DB.Find(&user, user.ID)
	db.DB.Delete(&user)
}

// FirstUser is called from main.go and is used to create an admin user on first run.
func FirstUser() {
	var userCount int
	var users []User
	db.DB.Find(&users).Count(&userCount)
	if userCount == 0 {
		db.DB.Create(&User{Username: "root", Password: "$2a$04$7ZZOLkODB70E5UL9UqvGzuPnqfaCZjKVUd7UhYP4jRywU/gOzHomS"})
		fmt.Println("Created new user `root` password is `12345`")
	}
}

// Auth checks a users creds
func Auth(w http.ResponseWriter, r *http.Request) {
	var user User
	fmt.Printf("%+v", r)
	// var count int
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	// fmt.Println(user.Username)
	plainPassword := user.Password
	// fmt.Println(plainPassword)
	// unmarshal content to ApplicantJSON
	db.DB.Where("username = ?", user.Username).Find(&user)
	// fmt.Println(user.Password)

	// do the passwords match?
	if comparePasswords(user.Password, []byte(plainPassword)) {
		//json resp
		newToken := GenerateJWT(user.Username)
		// JWT := JSONTOKEN{newToken}
		// js, err := json.Marshal(JWT)
		// if err != nil {
		//   log.Println(err)
		// }
		// w.Header().Set("Content-Type", "application/json")
		// w.Write(js)

		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "username", Value: user.Username, Expires: expiration}
		http.SetCookie(w, &cookie)
		cookie = http.Cookie{Name: "Token", Value: newToken, Expires: expiration}
		http.SetCookie(w, &cookie)

	} else {
		fmt.Fprintf(w, "Bad username and/or password")
	}
}

// hashAndSalt allows us to hash a password and salt it
func hashAndSalt(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// comparePasswords checks the passwords
func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// GenerateJWT generates the JSON web token
func GenerateJWT(username string) string {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = username

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "Error, unable to make JWT in user.go"
	}
	//json resp
	return tokenString
}

// tokenGenerator generates the API token
func tokenGenerator() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
