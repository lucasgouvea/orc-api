package users

import (
	Database "orc-api/internal/database"
	Errors "orc-api/internal/errors"
	Shared "orc-api/internal/shared"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

const JWT_KEY string = "SECRET"

const TTL = 60

func listUsers(params Shared.Params) ([]UserSchema, error) {
	var schemas []UserSchema

	users := []User{}
	db := Database.GetDB()
	err := db.Limit(params.Limit).Offset(params.Offset).Select("id", "created_at", "name").Find(&users).Error

	for _, u := range users {
		schemas = append(schemas, u.Schema())
	}
	return schemas, err
}

func createUser(user *User) error {
	user.hashPassword()
	db := Database.GetDB()
	err := db.Clauses(clause.Returning{}).Create(&user).Error
	return err
}

func updateUser(user *User) error {
	if len(user.Password) > 0 {
		user.hashPassword()

	}
	db := Database.GetDB()
	res := db.Clauses(clause.Returning{}).Where("id = ?", user.ID).Updates(user)
	if res.RowsAffected == 0 {
		return Errors.ResourceNotFoundErr
	}
	return res.Error
}

func login(schema PostLoginSchema) (*AuthSchema, error) {
	var err error
	var authSchema *AuthSchema

	user := schema.parse()
	pass := user.Password

	db := Database.GetDB()
	res := db.Where("name = ?", user.Name).First(&user)
	if res.RowsAffected == 0 {
		return nil, InvalidUserNameErr
	}
	if res.Error != nil {
		return nil, res.Error
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return nil, InvalidUserPassErr
	}

	authSchema, err = getAuthSchema(user.Name, user.Password)

	if err = updateUser(&User{Model: Shared.Model{ID: user.ID}, Token: authSchema.Token, TokenExpires: authSchema.ExpiresTime}); err != nil {
		return nil, err
	}

	return authSchema, err
}

func getAuthSchema(username string, pass string) (*AuthSchema, error) {
	expirationTime := time.Now().Add(TTL * time.Minute)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return nil, err
	}

	return &AuthSchema{
		Name:        username,
		Token:       tokenString,
		Expires:     claims.ExpiresAt.String(),
		ExpiresTime: claims.ExpiresAt.Time,
	}, nil
}
