package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var methods = []string{"GET", "POST", "PUT", "DELETE"}
var urls = []string{
	"/index.html",
	"/about",
	"/contact",
	"/info",
	"/login",
	"/blog",
}
var statuses = []string{"200", "200", "200", "200", "200", "200", "200", "200", "404", "500", "301"}
var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:130.0) Gecko/20100101 Firefox/130.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36 Edg/129.0.0.0",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:131.0) Gecko/20100101 Firefox/131.0",
}
var referrers = []string{"https://example.com", "https://foobar.org", "-"}

func randomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(255), rand.Intn(255), rand.Intn(255), rand.Intn(255))
}

func logTimestamp() string {
	return time.Now().Format("02/Jan/2006:15:04:05 -0700")
}

func generateLogLine(logFormat string) string {
	ip := randomIP()
	timestamp := logTimestamp()
	method := methods[rand.Intn(len(methods))]
	url := urls[rand.Intn(len(urls))]
	status := statuses[rand.Intn(len(statuses))]
	size := fmt.Sprintf("%d", rand.Intn(5000)+500)
	userAgent := userAgents[rand.Intn(len(userAgents))]
	referrer := referrers[rand.Intn(len(referrers))]

	if logFormat == "nginx" {
		return fmt.Sprintf(`%s - - [%s] "%s %s HTTP/1.1" %s %s "%s" "%s"`, ip, timestamp, method, url, status, size, referrer, userAgent)
	} else if logFormat == "apache" {
		return fmt.Sprintf(`%s - - [%s] "%s %s HTTP/1.1" %s %s "%s" "%s"`, ip, timestamp, method, url, status, size, referrer, userAgent)
	}
	return ""
}

func writeLog(logLine string, logFilePath string) {
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(logLine + "\n"); err != nil {
		fmt.Printf("Error writing to log file: %v\n", err)
	}
}

func generateLogs(logFormat string, logFilePath string, interval time.Duration) {
	for {
		logLine := generateLogLine(logFormat)
		writeLog(logLine, logFilePath)
		time.Sleep(interval)
	}
}

func main() {
	logFormat := flag.String("format", "nginx", "Log format: 'nginx' or 'apache'")
	logFilePath := flag.String("path", "access.log", "Path to the log file")
	interval := flag.Duration("interval", 2*time.Second, "Time interval between log entries (e.g., 2s, 500ms)")

	flag.Parse()

	if !strings.EqualFold(*logFormat, "nginx") && !strings.EqualFold(*logFormat, "apache") {
		fmt.Println("Invalid log format. Please choose either 'nginx' or 'apache'.")
		return
	}

	generateLogs(*logFormat, *logFilePath, *interval)
}
