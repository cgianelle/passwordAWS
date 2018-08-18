package main

import (
	"context"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cgianelle/password"
)

type MyEvent struct {
	Length int    `json:"length"`
	Type   string `json:"type"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var pwd string

	alphaPassword := password.Password{"aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ", r1}
	numericPassword := password.Password{"0123456789", r1}
	specialPassword := password.Password{"!@#$%^&*()-_+=,.?/:;'<>?/", r1}

	createPassword := password.PasswordBuilder(event.Length, r1)

	switch event.Type {
	case "alpha":
		pwd = createPassword(alphaPassword)
	case "alphaNum":
		pwd = createPassword(alphaPassword, numericPassword)
	case "special":
		pwd = createPassword(alphaPassword, numericPassword, specialPassword)
	}

	return pwd, nil
}

func main() {
	lambda.Start(HandleRequest)
}
