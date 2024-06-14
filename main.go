package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"time"
)

const (
	//URL = "https://www.google.com"

	SMTP_SERVER = "smtp.gmail.com"
	SMTP_PORT   = 587
	FromEmail   = "go-url-monitor@gmail.com" //update email
	Password    = "" //update password
	ToEmail     = "" //update email

)

func checkURL(URL string, ch chan<- bool) {
	response, err := http.Get(URL)
	if err != nil {
		fmt.Printf("Error checking %v: %v\n ", URL, err)
		ch <- false
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("%v Status code: %d\n", URL, response.StatusCode)
		ch <- false
		return

	}

	fmt.Printf("%v Status : OK\n", URL)
	ch <- true
}

func sendEmail(subject, body string) error {
	auth := smtp.PlainAuth("", FromEmail, Password, SMTP_SERVER)
	msg := fmt.Sprintf("Subject: %s\r\n\r\n %s", subject, body)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", SMTP_SERVER, SMTP_PORT), auth, FromEmail, []string{ToEmail}, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Welcome to URL Monitoring System")

	//opening the file containing the list of URLs
	file, err := os.Open("urls.txt")
	CheckErr(err)
	defer file.Close()

	//reading the file
	urls := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		urls = append(urls, scanner.Text())

	}

	//Checking for errors in reading the file

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file: ", err)
		return
	}

	//Assigning the intervals after each cycle of URL checking
	interval := 10 * time.Second

	for {
		ch := make(chan bool, len(urls))
		for _, URL := range urls {
			go checkURL(URL, ch)
		}

		allOK := true
		for i := 0; i < len(urls); i++ {
			if !<-ch {
				allOK = false
			}
		}

		if !allOK {
			//send email notif when at least one URL is down
			subject := "URL Monitoring System Alert"
			body := "One or more URLs are down"
			err := sendEmail(subject, body)
			if err != nil {
				fmt.Printf("Error sending email notification: %v\n", err)
			}
		}

		time.Sleep(interval)
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
