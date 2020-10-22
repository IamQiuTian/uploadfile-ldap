package main

import (
	"log"
	"flag"
	"github.com/iamqiutian/uploadFile/g"
	"github.com/iamqiutian/uploadFile/http"
	"github.com/BurntSushi/toml"
)

func init() {
	ConfigFile := flag.String("conf", "", "Config file for this listener and ldap configs")
	flag.Parse()

	if _, err := toml.DecodeFile(*ConfigFile, &g.Config); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.Start()
}