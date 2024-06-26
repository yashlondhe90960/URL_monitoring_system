# URL_monitoring_system

## Overview
This is a URL monitoring system written in Go. It reads a list of URLs from a file, checks the status of each URL at regular intervals, and sends an email notification if any of the URLs are down.

## Prerequisites
- Go (version 1.14 or later)
- A text file named `urls.txt` containing the URLs to be monitored, one per line
- SMTP server details and email credentials for sending email notifications

## Installation
1. Clone this repository or download the Go file.
2. Install Go if you haven't already. You can download it from the official Go website.
3. Create a `urls.txt` file in the same directory as your Go program and add the URLs you want to monitor, one per line.

## Usage
1. Update the `FromEmail`, `Password`, and `ToEmail` constants in the code with your email credentials and the recipient's email address.
2. Open a terminal, navigate to the directory containing your Go program, and run the command `go run main.go`.

## Features
- Reads URLs from a file
- Checks the status of each URL at regular intervals
- Sends an email notification if any of the URLs are down

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.
