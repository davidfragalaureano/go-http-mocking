package main

import (
	"fmt"
	"io/ioutil"

	"github.com/http-mocking/restclient"
)

func MakeGetRequest(url string) (string, error) {

	// Making HTTP post request
	res, err := restclient.Get(url, nil, nil)

	if err != nil {
		return "", fmt.Errorf("Unable to request via GET: %v", err)
	}

	defer res.Body.Close()

	pokemons, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", fmt.Errorf("Error unmarshalling body %v", err)
	}

	fmt.Printf("%+v\n", string(pokemons))

	return string(pokemons), nil
}

func main() {
	MakeGetRequest("https://pokeapi.co/api/v2/pokemon/ditto")
}
