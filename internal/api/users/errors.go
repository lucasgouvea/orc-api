package users

import "errors"

var InvalidUserNameErr = errors.New("Invalid user name.")
var InvalidUserPassErr = errors.New("Invalid user password.")

var InvalidAuthHeaderErr = errors.New("Could not parse Authorization header.")
var InvalidJWTUserErr = errors.New("JWT user is not matching with the one issued.")
var InvalidJWTTokenErr = errors.New("JWT token is not matching with the one issued.")
