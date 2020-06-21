package helpers

import "fmt"

func GetHa1(username, domain, password string) string {
	return fmt.Sprintf("%s:%s:%s", username, domain, password)
}

func GetHa1b(username, domain, password string) string {
	return fmt.Sprintf("%s@%s:%s:%s", username, domain, domain, password)
}

func GetSip(username, domain string) string {
	return fmt.Sprintf("sip:%s@%s", username, domain)
}
