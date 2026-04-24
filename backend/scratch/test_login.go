package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func testLogin(email, password string) {
	fmt.Printf("Testing login for %s... ", email)
	loginData := map[string]string{
		"email":    email,
		"password": password,
	}
	jsonData, _ := json.Marshal(loginData)

	resp, err := http.Post("http://localhost:8080/api/v1/auth/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Status: %s | Body: %s\n", resp.Status, string(body))
}

func main() {
	testLogin("admin@kreatif.id", "password")
	testLogin("manager@kreatif.id", "password")
	testLogin("controller@kreatif.id", "password")
	testLogin("user@kreatif.id", "password")
}
