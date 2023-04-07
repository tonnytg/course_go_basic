package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"time"
	"encoding/json"
)

type Family []struct {
	CreatedAt  time.Time `json:"createdAt"`
	Name       string    `json:"name"`
	Job        string    `json:"job"`
	FamilyName string    `json:"familyName"`
	ID         string    `json:"id"`
}

func main() {

	url := "https://610aa52552d56400176afebe.mockapi.io/api/v1/friendlist"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f := Family{}
	err = json.Unmarshal(body, &f);

	for i, v := range f {
		fmt.Println("Elemento:", i)
		fmt.Println(v.Job)
		fmt.Println(v.Name)
	}
}