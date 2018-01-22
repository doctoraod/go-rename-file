package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	files, err := ioutil.ReadDir("../../Downloads/Music")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		value := f.Name()
		substring := value[3:len(value)] // Remove 3 string first
		fmt.Println(f.Name())
		fmt.Println("     > " + substring)
		os.Rename(("../../Downloads/Music/" + f.Name()), ("../../Downloads/Music/" + substring))
	}
}
