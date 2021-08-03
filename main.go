package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		authRequired()
	}

	authToken := os.Args[1]

	_, authenticated := checkToken(authToken)

	if authenticated {
		shell := exec.Command("tmux", "new", "-A", "-s",  "HR")
		shell.Stdin = os.Stdin
		shell.Stdout = os.Stdout
		shell.Stderr = os.Stderr
		fmt.Println("Auth Successfull")
		fmt.Println("Spawning a bash process")
		err := shell.Run()

		if err != nil {
			authRequired()
		}

		return
	} else {
		authRequired()
	}
}

func checkToken(token string) (string, bool) {
	data, err := base64.StdEncoding.DecodeString(token)
	envToken := []byte(os.Getenv("SUDORANK_PASS"))

	if (err == nil) && (string(data) == string(envToken)) {
		fmt.Println("Authenticated")
		return "", true
	}

	return string(data), false
}

func authRequired() {
	fmt.Println("Authentication Required")
	time.Sleep(600 * time.Second)
	os.Exit(0)
}
