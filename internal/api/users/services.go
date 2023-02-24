package users

import (
	"net/http"
	Database "orc-api/internal/database"
	Errors "orc-api/internal/errors"
	Shared "orc-api/internal/shared"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

const TTL = 60

/* USER */

func listUsers(params Shared.Params) ([]UserSchema, error) {
	schemas := make([]UserSchema, 0)
	users := []User{}
	db := Database.GetDB()
	err := db.Limit(params.Limit).Offset(params.Offset).Select("id", "created_at", "name").Find(&users).Error

	for _, u := range users {
		schemas = append(schemas, u.Schema())
	}
	return schemas, err
}

func findByUserName(name string) (*User, error) {
	var user *User

	db := Database.GetDB()
	res := db.Model(&User{}).Where("name = ?", name).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
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

/* LOGIN */

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
	jwt_key := Shared.GetEnvVars().JWT_KEY
	tokenString, err := token.SignedString([]byte(jwt_key))
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

func parseJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		jwt_key := Shared.GetEnvVars().JWT_KEY
		return []byte(jwt_key), nil
	})
}

var ValidateJWTHandler = func(c *gin.Context) {
	var user *User
	var err error
	var matchedRegExp bool

	if c.Request.URL.Path == "/v1/login" {
		return
	}

	auth := c.GetHeader("Authorization")
	matchedRegExp, err = regexp.MatchString("^Bearer.+$", auth)
	if matchedRegExp == false || err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": InvalidAuthHeaderErr.Error()})
		return
	}
	tokenString := auth[len("Bearer "):]
	token, err := parseJWT(tokenString)
	if token.Valid == false {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	name := token.Claims.(jwt.MapClaims)["username"].(string)
	if user, err = findByUserName(name); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": InvalidJWTUserErr.Error()})
		return
	}
	if tokenString != user.Token {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": InvalidJWTTokenErr.Error()})
		return
	}

}
