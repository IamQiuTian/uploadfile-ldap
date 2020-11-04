package ldap

import (
	"crypto/tls"
	"fmt"
	"log"
	"github.com/go-ldap/ldap"
	"github.com/iamqiutian/uploadFile/g"
)

func LDAPAuthUser(username string, password string) bool {
	bindusername := g.Config.LDAP.BindDN
	bindpassword := g.Config.LDAP.BindPassword

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", g.Config.LDAP.Host, 389))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
	if err != nil {
		log.Fatal(err)
	}

	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Fatal(err)
	}

	searchRequest := ldap.NewSearchRequest(
		g.Config.LDAP.Base,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organizationalPerson)(uid=%s))", username),
		[]string{"dn"},
		nil,
	)
	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Println(err)
		return false
	}

	if len(sr.Entries) != 1 {
		log.Println("User does not exist or too many entries returned")
		return false
	}

	userdn := sr.Entries[0].DN

	err = l.Bind(userdn, password)
	if err != nil {
		log.Println(err)
		return false
	}

	err = l.Bind(bindusername, bindpassword)
	if err != nil {
		log.Println(err)
	}
	return true
}
