package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func CallGemini(prompt string) {
	godotenv.Load(".env")

	client := &http.Client{
		Timeout: 50*time.Second,
	}

	geminiURL := "https://generativelanguage.googleapis.com/v1beta/interactions"

	reqBody := []byte(fmt.Sprintf(`{"model": "gemini-3.5-flash","input": "%s"}`,prompt))

	req,err := http.NewRequest(http.MethodPost,geminiURL,bytes.NewBuffer(reqBody))
	if err!=nil{
		log.Println("Error Sending Gemini Request!")
		return
	}

	req.Header.Add("x-goog-api-key",os.Getenv("GEMINI"))
	req.Header.Add("Content-Type","application/json")

	resp,err := client.Do(req)
	if err!=nil{
		log.Println("Error Sending HTTP Request!",err)
		return
	}
	defer resp.Body.Close()

	resBody, _ := io.ReadAll(resp.Body)

	var resps any
	err = json.Unmarshal(resBody, &resps)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resBody))

}