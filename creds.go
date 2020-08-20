package utils

import (
	"fmt"
	"log"

	"github.com/havoc-io/go-keytar"
)

// Constants
const (
	Service = "Now CLI"
)

func setCredentials(Account string, Password string) {
	//fmt.Println("In setCredentials fn")
	//fmt.Println("Service: ", Service)
	//fmt.Println("Account: ", Account)
	//fmt.Println("Password: ", Password)

	// Delete existing entry
	deletePassword(Service, Account)

	// Add new entry
	addPassword(Service, Account, Password)
}

func getCredentials() string {
	// TODO: Extract username, password, host, method

	return getPassword(Service, "")
}

// Convenience methods for keytar
func initKeychain() {
	var err error
	keychain, err = keytar.GetKeychain()
	if err != nil {
		log.Fatal("Unable to create keychain")
	}
}

func addPassword(Service string, Account string, Password string) {
	// err := keychain.AddPassword("example.org", "George", "deCr37")
	err := keychain.AddPassword(Service, Account, Password)
	if err != nil {
		fmt.Println("Password addition failed: ", err)
	} else {
		fmt.Println("Creds saved in keychain")
	}
}

func deletePassword(Service string, Account string) {
	err := keychain.DeletePassword(Service, Account)
	if err != nil {
		fmt.Println("Password deletion failed - May be new entry: ", err)
	}
}

func getPassword(Service string, Account string) string {
	password, err := keychain.GetPassword(Service, Account)
	if err != nil {
		fmt.Println("Password retrieval failed: ", err)
	} else {
		//fmt.Println("password retrieved : ", password)
	}
	return password
}
