package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendMail(email, address, productName string) {
	url := "http://localhost:8002/mailer"

	mailBody := map[string]string{
		"buyer_email":   email,
		"buyer_address": address,
		"product_name":  productName,
	}

	mailBodyMarshal, _ := json.Marshal(mailBody)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(mailBodyMarshal))
	if err != nil {
		panic(err)
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	fmt.Println("response status", res.Status)
	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body), "response body")

}
