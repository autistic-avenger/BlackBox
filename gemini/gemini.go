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
	"github.com/charmbracelet/glamour"
)

type geminiResponse struct {
	Steps   []geminiStep `json:"steps"`
}

type geminiStep struct {
	Content []geminiStepContent `json:"content,omitempty"`
}

type geminiStepContent struct {
	Text string `json:"text"`
}



func CallGemini(prompt string) {
	godotenv.Load(".env")

	client := &http.Client{
		Timeout: 50*time.Second,
	}

	geminiURL := "https://generativelanguage.googleapis.com/v1beta/interactions"
	fmt.Println("Getting Reply from gemini...")
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

	var formattedOutput geminiResponse
	err = json.Unmarshal(resBody, &formattedOutput)
	if err != nil {
		log.Fatal(err)
	}
	
	mdOutput := formattedOutput.Steps[1].Content[0].Text
	Finaloutput,_ := glamour.Render(mdOutput,"dark")
	fmt.Println(Finaloutput)

}