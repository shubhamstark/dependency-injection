package login

import "github.com/golang-jwt/jwt/v4"

// User will provide some id
// MPIN or password

// we will return some initial login payload and tokens

type Claims struct {
	UserID string
	jwt.Claims
}

type Payload struct {
}

type Response struct {
	Payload
	IDToken      string
	RefreshToken string
}

type Login struct {
	FailedAttempts IFailedAttempts
	Validator      IValidator
	PayloadGetter  IPayloadGetter
	TokenHandler   ITokenHandler
	UserGetter     IUserGetter
}

// id could be email, phone, pan
func (l Login) Do(id string, credentals string) (Response, error) {

	failedAttempts, err := l.FailedAttempts.Get()
	if err != nil {
		// TODO handle
		return Response{}, err
	}

	if failedAttempts > 5 {
		// TODO handle
		return Response{}, err
	}

	credsCorrect, err := l.Validator.Validate(id, credentals)
	if err != nil {
		// TODO handle
		return Response{}, err
	}

	if !credsCorrect {
		// TODO handle
		l.FailedAttempts.Set(failedAttempts + 1)
		return Response{}, err
	}
	payload := l.PayloadGetter.Get()

	user, err := l.UserGetter.Get()
	if err != nil {
		// TODO handle
		return Response{}, err
	}

	claims := Claims{
		UserID: user.UserID,
	}

	idToken, refresh, err := l.TokenHandler.IssueTokenPair(claims)
	if err != nil {
		// TODO handle
		return Response{}, err
	}

	return Response{
		Payload:      payload,
		IDToken:      idToken,
		RefreshToken: refresh,
	}, nil
}
