package main

import (
	"os"
	"log"
    "bufio"
    "strings"
    "fmt"
)

func main()  {
	f, err := os.Open("mylog.log")
	if err != nil {
		log.Fatal(err)		
	}
	defer f.Close()

	r := bufio.NewReader(f)
	
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, "INFO") {
            fmt.Println(s)
        }
	}
}
