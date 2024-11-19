package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"slices"
	"strings"
)

type User struct {
	Name  string
	Email string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the names you want to give secret santas. Make sure theyre email is included in the format Name:Email. Enter 'done' when done.")

	// While loop
	var names []string
	var emails = make(map[string]string)
	for {
		text, _ := reader.ReadString('\n')

		if text == "done\n" {
			fmt.Println("Completed!")
			break
		}

		// Remove newline character
		text = text[:len(text)-1]
		// Split text by ";"
		splitText := strings.Split(text, ":")
		name := splitText[0]
		names = append(names, name)
		emails[name] = splitText[1]
	}

	// Print the names
	pairs := giveSecretSantas(names)

	sendSecretsEmails(pairs, emails)
}

var FROM = os.Getenv("FROM")
var PORT = os.Getenv("PORT")
var HOST = os.Getenv("HOST")
var AUTH = smtp.PlainAuth("", FROM, os.Getenv("MAIL_PASSWORD"), HOST)

func sendSecretsEmails(pairs []Pair, emails map[string]string) {
	for _, pair := range pairs {
		// Send email to pair.From
		email := emails[pair.From]

		recipients := []string{email}

		msg := []byte("From: " + FROM + "\r\n" +
			"To: " + email + "\r\n" +
			"Subject: Test mail\r\n\r\n" +
			"Hello " + pair.From + ",\r\n" +
			"You have been assigned to give a book to " + pair.To + ".\r\n" +
			"Good luck!\r\n")

		err := smtp.SendMail(HOST+":"+PORT, AUTH, FROM, recipients, msg)

		if err != nil {
			fmt.Println("Error sending email to " + pair.From)
		}

		// sENT
		fmt.Println("Sent email to " + pair.From)
	}
}

type Pair struct {
	From string
	To   string
}

// Return an array of String tuples
func giveSecretSantas(names []string) []Pair {
	var pairs []Pair
	// List of names from array names
	var namesList = slices.Clone(names)

	for _, name := range names {
		// Get a random name from namesList
		// Remove that name from namesList
		// Add the pair to the name array

		index := rand.Intn(len(namesList))

		santa := namesList[index]

		pairs = append(pairs, Pair{name, santa})

		namesList = append(namesList[:index], namesList[index+1:]...)
	}

	return pairs
}
