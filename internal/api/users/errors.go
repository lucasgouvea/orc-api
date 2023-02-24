package users

import (
	"errors"
	"strconv"
)

var InvalidUserNameErr = errors.New("Invalid user name.")

func InvalidUserPassErr(attempts int64) error {
	return errors.New("Invalid user password, attempts: " + strconv.FormatInt(attempts, 10))
}

var BlockedUserErr = errors.New("User has been blocked.")

var InvalidAuthHeaderErr = errors.New("Could not parse Authorization header.")
var InvalidJWTUserErr = errors.New("JWT user is not matching with the one issued.")
var InvalidJWTTokenErr = errors.New("JWT token is not matching with the one issued.")
