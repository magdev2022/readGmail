package main

import (
	"fmt"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func main() {
	fmt.Println("Connecting to server...")
	// Connect to server
	c, err := client.DialTLS("imap.gmail.com:993", nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected")
	// Don't forget to logout
	defer c.Logout()
	// Login
	username := "jin1991919@gmail.com"
	password := "mqspwumvjpybzehy"
	if err := c.Login(username, password); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Logged in successfully")
	// List mailboxes
	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- c.List("", "*", mailboxes)
	}()
	fmt.Println("Mailboxes:")
	for m := range mailboxes {
		fmt.Println("* " + m.Name)
	}
	if err := <-done; err != nil {
		fmt.Println(err)
	}
	// Select INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Flags for INBOX:", mbox.Flags)
	// Get the last 4 messages
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 3 {
		// We're using unsigned integers here, only subtract if the result is > 0
		from = mbox.Messages - 3
	}
	seqset := new(imap.SeqSet)
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, 10)
	done = make(chan error, 1)
	go func() {
		done <- c.Fetch(seqset, []imap.FetchItem{imap.FetchEnvelope}, messages)
	}()

	fmt.Println("Last 4 messages:")
	for msg := range messages {
		fmt.Println("From: " + msg.Envelope.From[0].Address())
	}

	if err := <-done; err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done!")
}
