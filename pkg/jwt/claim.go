package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/bagasunix/bank-ina/server/domains/entities"
)

type Claims struct {
	User *entities.User `json:"user,omitempty"`
	jwt.StandardClaims
}

type ClaimsBuilder struct {
	claims *Claims
}

func (c *ClaimsBuilder) User(user *entities.User) *ClaimsBuilder {
	c.claims.User = user
	return c
}

func (c *ClaimsBuilder) ExpiresAt(expiresAt time.Time) *ClaimsBuilder {
	c.claims.ExpiresAt = expiresAt.Unix()
	return c
}

func (c *ClaimsBuilder) Build() *Claims {
	return c.claims
}

func NewClaimsBuilder() *ClaimsBuilder {
	a := new(ClaimsBuilder)
	a.claims = new(Claims)
	a.claims.StandardClaims = jwt.StandardClaims{}
	return a
}
