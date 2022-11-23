package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bl0ckp1n9/bl0ckp1n9Coin/blockchain"
	"github.com/bl0ckp1n9/bl0ckp1n9Coin/utils"
)


const port string = ":4000"

type URL string 

type URLDescription struct {
	URL URL `json:"url"`
	Method string `json:"method"`
	Description string `json:"description"`
	Payload string `json:"payload,omitempty"`
}

type AddBlockBody struct {
	Message string
}

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s",port,u)
	return []byte(url), nil
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription {
		{
			URL: URL("/") ,
			Method : "GET",
			Description: "See Documentation",
		},
		{
			URL: URL("/blocks") ,
			Method : "POST",
			Description: "Add A Block",
			Payload: "data:string",
		},
		{
			URL: URL("/blocks/{id}") ,
			Method : "POST",
			Description: "See A Block",
		},
	}
	rw.Header().Add("Content-Type","application/json")
	json.NewEncoder(rw).Encode(data)

}


func blocks(rw http.ResponseWriter,r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type","application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody AddBlockBody;
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated) // StatusCreated Code = 201
	}
}

func main() {

	http.HandleFunc("/",documentation)
	http.HandleFunc("/blocks",blocks)
	fmt.Printf("Listening http://localhost%s\n",port)
	log.Fatal(http.ListenAndServe(port,nil))
} 

