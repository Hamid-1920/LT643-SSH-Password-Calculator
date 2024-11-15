package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Print("LT643 SSH password generator by Hamid1920\n-------------------------------------------------------\n")

	var macAddress string
	if len(os.Args) > 1 {
		macAddress = os.Args[1]
	} else {
		fmt.Print("Enter MAC address: ")
		reader := bufio.NewReader(os.Stdin)
		macAddress, _ = reader.ReadString('\n')
	}
	macAddress = strings.ReplaceAll(strings.TrimSpace(macAddress), ":", "")
	if isValidMACAddress(macAddress) {
		password := strings.ToUpper(macAddress)
		deviceTypes := []string{"MN6200D", "HA6400", "TF i60 B1"}
		iterationCount := 1365
		derivedKeyLength := 6

		for _, deviceType := range deviceTypes {
			salt := "Argt3l3comErwin" + deviceType + password
			derivedKey := pbkdf2.Key([]byte(password), []byte(salt), iterationCount, derivedKeyLength, sha256.New)
			fmt.Printf("Generated Password for %s: %x\n", deviceType, derivedKey)
		}
	} else {
		fmt.Println("Invalid MAC address")
	}

}
func isValidMACAddress(macAddress string) bool {

	macPattern := `^[0-9A-Fa-f]{12}$`
	re := regexp.MustCompile(macPattern)
	return re.MatchString(macAddress)
}
