package sapi

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/guregu/null"
)

func randomTokenID() string {
	var n = 32
	var characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ013456789-._~+/")
	max := big.NewInt(int64(len(characters)))
	b := make([]rune, n)
	for i := range b {
		j, _ := rand.Int(rand.Reader, max)
		b[i] = characters[j.Uint64()]
	}
	return string(b)
}

// Jeton d'accès
type Token struct {
	ID                string      `json:"access_token"`
	TokenType         string      `json:"token_type"`
	UserID            null.String `json:"user_id"`
	AuthorizationCode null.String `json:"autorization_code"`
	GrantType         string      `json:"grant_type"`
	Created           time.Time   `json:"created"`
	Expires           time.Time   `json:"expires"`
	ExpiresIn         int         `json:"expires_in"`
}

type OAuthToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewUserToken(username, password string) (*Token, error) {
	u, err := userQueryNoAuth(username)
	if err == nil {
		if u.HashMatch(password) {
			var id string
			var created, expires time.Time
			rt := randomTokenID()
			for true {
				err = db.QueryRowx("INSERT INTO token(id, contact_id, grant_type, created, expires) VALUES ($1, $2, 'password', now(), now() + interval '1 hour') RETURNING id, created, expires", rt, u.ID).Scan(&id, &created, &expires)
				if err == nil {
					break
				} else {
					time.Sleep(1000 * time.Millisecond)
					logE.Println("inserting token failed...")
				}
			}
			logI.Printf("%s logged in successfully (%s valid until %s)", u.ID, id, expires)
			return &Token{ID: id, UserID: null.StringFrom(u.ID), TokenType: "bearer", Created: created, Expires: expires, ExpiresIn: 3600}, err
		} else {
			return nil, &AuthFailedError{msg: "bad username or password"}
		}
	} else {
		return nil, &AuthFailedError{msg: "bad username or password"}
	}
}

func TokenQuery(id string) (*Token, error) {
	t := Token{ID: id}
	err := db.QueryRowx("SELECT id, contact_id, authorization_code, grant_type, created, expires FROM token WHERE id = $1 AND expires > NOW()", t.ID).StructScan(&t)
	return &t, err
}

func (t *Token) Delete(uc *UserContext) error {
	r := db.QueryRowx("DELETE FROM token WHERE id=$1", t.ID)
	return r.Err()
}

func (t *Token) AsOAuth() *OAuthToken {
	oat := new(OAuthToken)
	oat.AccessToken = t.ID
	oat.TokenType = t.TokenType
	oat.ExpiresIn = t.ExpiresIn
	logD.Println(oat)
	return oat
}
