package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	// getting a response and an optional error
	// using the http package to get the robots.txt file
	res, err := http.Get("http://www.google.com/robots.txt")

	if err != nil { // check to see if the error is nil
		// and if not proceed
		log.Fatal(err)
	}

	// close the body of the response to let the web know that
	// we are done working with it

	defer res.Body.Close()

	// get the response and use the ReadAll from ioutil pacakger
	robots, err := ioutil.ReadAll(res.Body)

	// checking to see fif the read operation fail
	// checking for the error
	if err != nil {
		log.Fatal(err)
	}

	// print the value of the file
	fmt.Printf("%s", robots)
}
