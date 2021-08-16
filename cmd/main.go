package main

import (
	"errors"
	//"fmt"
	"os"
	"strings"

	"github.com/1Password/connect-sdk-go/connect"
)

// FIXME modify the default string to match the input
type Config struct {
	Password string `opitem:"DemoPassword" opfield:".password"`
}

var client connect.Client

func RetrieveSecret(variableName string, client connect.Client) (string, error) {
	// Parse input to identify the secret item and field
	var secretKey = strings.Split(variableName, "/")
	if len(secretKey) != 3 {
		return "", errors.New("Secret path incorrect, expect 3 elements, got: " + variableName)
	}
	provider := secretKey[0]
	itemUUID := secretKey[1]
	itemField := secretKey[2]
	println("Receiving: ", provider, itemUUID, itemField)

	if provider != "1password" {
		return "", errors.New("Unsupported provider, expected '1password', got: " + provider)
	}

	// key := fmt.Sprintf(`opuuid: "%s" opfield:".%s"`, itemUUID, itemField)
	// NOT WORKING
	//type Config struct {
	//	Password string;
	//n}
	// NOT WORKING: item := Config{key}
	var item Config
	connect.Load(client,&item)

	return item.Password, nil
}

func main() {
	// NOTE: Requires: OP_CONNECT_TOKEN, OP_CONNECT_HOST, OP_VAULT
	client, err := connect.NewClientFromEnvironment()
	if err != nil {
		panic(err)
	}

	secret, err := RetrieveSecret(os.Args[1], client)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write([]byte(secret))
}
