package go_token_test

import (
	"encoding/json"
	"github.com/orange-juzipi/go-token"
	"testing"
	"time"
)

const secret = "go-token"

func TestJwtSign(t *testing.T) {
	jwt := go_token.New(secret)
	tokenString, err := jwt.JwtSign(1, time.Hour*24)
	if err != nil {
		t.Errorf("Error signing:%v", err)
		return
	}
	t.Log(tokenString)

	parse, err := jwt.JwtParse(tokenString)
	if err != nil {
		t.Errorf("Error parsing:%v", err)
		return
	}
	body, _ := json.MarshalIndent(parse, "", "\t")
	t.Log(string(body))
}
