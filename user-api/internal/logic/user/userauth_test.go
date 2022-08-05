package user

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)



func TestGeneratorToken(t *testing.T) {
	jwt := new(TokenJWT)
	jwt.ttl = time.Minute * 60 * 480
	jwt.key = "abcdefg123"
}

type TokenJWT struct {
	signMethod jwt.SigningMethod
	key        interface{}
	ttl        time.Duration
}

func generaToken(jwt *TokenJWT) {
	if jwt == nil {
		return
	}
	//sign, err := jwt.Sign(1)

}

type Claims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

//func (jwt TokenJWT) Sign(userId int) (string, error) {
//	var claims = Claims{
//		userId,
//		jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(jwt.ttl).Unix(),
//			Issuer:    "fogcloud",
//		},
//	}
//	return jwt.NewWithClaims(t.signMethod, claims).SignedString(t.key)
//}
