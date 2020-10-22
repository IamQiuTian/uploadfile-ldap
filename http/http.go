package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/iamqiutian/uploadFile/g"
)


func Start() {
	listenPort := fmt.Sprintf(":%d", g.Config.Listen.Port)

	if g.Config.Listen.SSL == true {
		log.Printf("Listening on port %d without SSL", g.Config.Listen.Port)
		err := http.ListenAndServeTLS(":"+listenPort, g.Config.Listen.Cert, g.Config.Listen.Key, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	} else {
		log.Printf("Listening on port %d without SSL", g.Config.Listen.Port)
		err := http.ListenAndServe(listenPort, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	}
}
