package security

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/hopperteam/hopper-account/config"
	"github.com/hopperteam/hopper-account/model"
	"io/ioutil"
)

type SessionUserClaims struct {
	User model.SessionUser `json:"user"`
	jwt.StandardClaims
}

func CreateSession(usr *model.SessionUser, expire int64) (string, error) {
	return jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), &SessionUserClaims{
		User:    *usr,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire,
			Issuer:    "HopperAccount",
		},
	}).SignedString(privateKey)

}

func DecodeSession(sess string) (*model.SessionUser, error) {
	sUsr := &SessionUserClaims{}
	_, err := jwt.ParseWithClaims(sess, sUsr, func(token *jwt.Token) (interface{}, error) {
		return privateKey.Public(), nil
	})

	if err != nil {
		return nil, err
	}

	return &sUsr.User, nil
}

var privateKey *rsa.PrivateKey

func LoadKeys() error {
	dat, err := ioutil.ReadFile(config.Config.RsaPrivateKeyPath)

	if err != nil {
		return err
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(dat)

	return err
}
