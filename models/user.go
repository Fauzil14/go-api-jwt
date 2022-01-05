package models

import (
	"go-api-jwt/utils/token"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Username  string    `json:"username" gorm:"not nul;unique"`
	Email     string    `json:"email" gorm:"not null;unique"`
	Password  string    `json:"password" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// verify password dari inputan dengan hash password yang tersimpan di DB
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string, db *gorm.DB) (string, error) {
	var err error

	usr := User{}

	// get user data if username exist in DB
	// data di ambil menggunakan satu acuan (username) sehingga harus menggunakan fungsi db.Model(struct)
	// tabel yang di cari bedasarkan parameter db.Model(struct{})
	err = db.Model(User{}).Where("username = ?", username).Take(&usr).Error

	if err != nil {
		return "", err
	}

	// verify input password with the from DB
	err = VerifyPassword(password, usr.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	// generate toke if user and password is matched
	token, err := token.GenerateToken(usr.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

// save user when register
// method in user instance
func (usr *User) SaveUser(db *gorm.DB) (*User, error) {
	// hash password from user input
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)

	if errPassword != nil {
		return &User{}, errPassword
	}

	usr.Password = string(hashedPassword)
	usr.Username = html.EscapeString(strings.TrimSpace(usr.Username))

	// create user into DB
	var err error = db.Create(&usr).Error

	if err != nil {
		return &User{}, err
	}

	return usr, nil
}
