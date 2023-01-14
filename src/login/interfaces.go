package login

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/shubhamstark/dependency-injection.git/src/user"
)

// Validator will validate id and credentials. Id can be email, phone, pan etc.
type IValidator interface {
	Validate(id, credentials string) (isValid bool, err error) // stateless
}

type IPayloadGetter interface {
	// Get will return data which will be sent to user after successful login
	Get() Payload // stateless
}

type IFailedAttempts interface {
	// Get will fetch failed attempts. It will fetch user identity from initializer. This is done to force the instance to be used with only one user.
	Get() (int, error)

	// Get will set failed attempts.
	Set(attempts int) error
}

type ITokenHandler interface {
	IssueTokenPair(claims jwt.Claims) (idToken string, refreshToken string, err error)
	GenerateJWT(claims jwt.Claims) (idToken string, err error)
	ExtendExpiry(idToken string, refreshToken string) (extendedJWT string, err error)

	Invalidate(refreshToken string) error
	InvalidateAll(refreshToken string) error
}

type IUserGetter interface { // stateful
	Get() (user.User, error)
}
