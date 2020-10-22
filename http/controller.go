package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"html/template"
	"os"
	"time"
	"github.com/iamqiutian/uploadFile/g"
	"github.com/iamqiutian/uploadFile/ldap"
	"github.com/iamqiutian/uploadFile/utils"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

func HomePage(res http.ResponseWriter, req *http.Request) {
	if g.Config.LDAP.UseLDAP {
		t, err := template.New("webpage").Parse(g.LoginTpl)
		if err != nil {
			log.Fatal(err)
		}

		err = t.Execute(res, g.Config)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		uid,_ := uuid.NewV4()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uid":      uid,
			"username": "anonymous",
			"exp":      time.Now().Add(time.Minute * 15).Unix(),
		})
		authorization, err := token.SignedString(g.MySigningKey)
		if err != nil {
			res.Write([]byte(err.Error()))
		}
		http.SetCookie(res, &http.Cookie{
			Name:    "token",
			Value:   authorization,
			Expires: time.Now().Add(time.Minute * 15)})

		http.Redirect(res, req, "/choose", 302)
	}
}

func CheckAuth(auth string) bool {
	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		return g.MySigningKey, nil
	})

	if err != nil {
		return false
	}
	if !token.Valid {
		return false
	}
	return true
}

func Login(res http.ResponseWriter, req *http.Request) {

	formUsername := req.FormValue("username")
	formPassword := req.FormValue("password")

	log.Printf("Logging in with user: %s\n", formUsername)

	authenticated := ldap.LDAPAuthUser(formUsername, formPassword)

	if authenticated {
		uid,_ := uuid.NewV4()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uid":      uid,
			"username": formUsername,
			"exp":      time.Now().Add(time.Minute * 15).Unix(),
		})
		authorization, err := token.SignedString(g.MySigningKey)
		if err != nil {
			res.Write([]byte(err.Error()))
		}
		http.SetCookie(res, &http.Cookie{
			Name:    "token",
			Value:   authorization,
			Expires: time.Now().Add(time.Minute * 15)})

		http.Redirect(res, req, "/choose", 302)
	} else {
		http.Redirect(res, req, "/", 302)
	}

}

func ChooseFile(res http.ResponseWriter, req *http.Request) {
	auth, err := req.Cookie("token")
	if err == nil {
		ok := utils.CheckAuth(auth.Value)
		if !ok {
			http.Redirect(res, req, "/", 302)
			return
		}
	}

	t, err := template.New("webpage").Parse(g.UploadTpl)
	if err != nil {
		log.Print(err)
		return
	}

	err = t.Execute(res, g.Config)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
}

func UploadFile(res http.ResponseWriter, req *http.Request) {
	auth, err := req.Cookie("token")
	if err == nil {
		ok := utils.CheckAuth(auth.Value)
		if !ok {
			http.Redirect(res, req, "/", 302)
			return
		}
	}

	uploadedFile, handler, err := req.FormFile("fileupload")
	if err != nil {
		log.Println(err)
		return
	}
	defer uploadedFile.Close()

	rdm := utils.CreateRandom()
	filePath := fmt.Sprintf("%s/%s/%s", g.Config.Upload.Path, rdm, handler.Filename)

	log.Printf("Saving file %s", filePath)

	if err = os.Mkdir(fmt.Sprintf("%s/%s", g.Config.Upload.Path, rdm), os.ModePerm); err != nil {
		log.Println(err)
		return
	}
	saveFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer saveFile.Close()
	io.Copy(saveFile, uploadedFile)

	t, err := template.New("webpage").Parse(g.DoneTpl)
	if err != nil {
		log.Print(err)
		return
	}
	gofilepath := g.Gofilepath {
		Path: rdm,
	}
	err = t.Execute(res, gofilepath)
	if err != nil {
		log.Print("execute: ", err)
		return
	}
}
