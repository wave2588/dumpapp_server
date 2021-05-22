package util

import (
	"errors"
	"time"

	"dumpapp_server/pkg/common/constant"
	"github.com/dgrijalva/jwt-go"
)

var hmacSampleSecret = []byte("dump_project_technology")

type DumpCustomClaim struct {
	MemberID int64
	jwt.StandardClaims
}

var GuestCrabCustomClaim = DumpCustomClaim{
	StandardClaims: jwt.StandardClaims{
		Audience: constant.MemberTypeGuest,
		Issuer:   "dump_project",
		Subject:  "dump_guest_ticket",
	},
}

var RegisterCrabCustomClaim = DumpCustomClaim{
	StandardClaims: jwt.StandardClaims{
		Audience: constant.MemberTypeRegister,
		Issuer:   "dump_project",
		Subject:  "dump_register_ticket",
	},
}

func GenerateRegisterTicket(memberID int64) (string, error) {
	registerClaim := GenerateRegisterClaim(memberID)
	registerTicket, err := GenerateTicket(registerClaim)
	return registerTicket, err
}

func GenerateRegisterClaim(memberID int64) DumpCustomClaim {
	timeDelta, _ := time.ParseDuration("720h")
	generatedClaim := DumpCustomClaim{
		memberID,
		jwt.StandardClaims{
			Issuer:    RegisterCrabCustomClaim.Issuer,
			Subject:   RegisterCrabCustomClaim.Subject,
			Audience:  RegisterCrabCustomClaim.Audience,
			ExpiresAt: time.Now().Add(timeDelta).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}
	return generatedClaim
}

func GenerateTicket(claim DumpCustomClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(hmacSampleSecret)
	return tokenString, err
}

func ParseTicket(tokenString string) (DumpCustomClaim, error) {
	var outputClaim DumpCustomClaim
	_, err := jwt.ParseWithClaims(tokenString, &outputClaim, VerifyMethod)
	return outputClaim, err
}

func VerifyMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("SigningMethod not correct")
	}
	return hmacSampleSecret, nil
}
