package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	lnAddr := "0.0.0.0:" + port

	ln, err := net.Listen("tcp", lnAddr)
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Printf("Server listening at %s \n", lnAddr)
	for {
		sCon, err := ln.Accept()
		if err != nil {

		}
		target := os.Getenv("ADDR")
		if target == "" {
			addr, err := ioutil.ReadFile("/tmp/addr")
			if err != nil {
				target = "www.google.com:80"
			} else {
				target = strings.Trim(string(addr), "\n")
			}

		}
		fmt.Println(sCon.RemoteAddr().String() + "->" + target)
		dCon, err := net.Dial("tcp", target)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// ----------forward-----------
		go func() {
			io.Copy(sCon, dCon)
			sCon.Close()
		}()

		go func() {
			io.Copy(dCon, sCon)
			dCon.Close()
		}()
	}
}
