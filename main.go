package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

// ANSI color codes
const (
	green = "\033[32m"
	red   = "\033[31m"
	reset = "\033[0m"
)

// Function to send OPTIONS request and check for PUT method
func checkPUTMethod(domain string) bool {
	client := &http.Client{
		Timeout: time.Second * 10, // 10 seconds timeout for each request
	}

	req, err := http.NewRequest("OPTIONS", domain, nil)
	if err != nil {
		fmt.Printf("Error creating request for domain %s: %v\n", domain, err)
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request to domain %s: %v\n", domain, err)
		return false
	}
	defer resp.Body.Close()

	allowHeader := resp.Header.Get("Allow")
	return strings.Contains(allowHeader, "PUT")
}

// Worker function to process domains
func worker(domains <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for domain := range domains {
		if !strings.HasPrefix(domain, "http://") && !strings.HasPrefix(domain, "https://") {
			domain = "https://" + domain
		}

		if checkPUTMethod(domain) {
			fmt.Printf("%sVulnerable: %s%s\n", red, domain, reset)
			results <- domain
		} else {
			fmt.Printf("%sNot Vulnerable: %s%s\n", green, domain, reset)
		}
	}
}

func main() {
	// Define command-line flags
	inputFile := flag.String("l", "", "Input file containing list of domains (required)")
	outputFile := flag.String("o", "output.txt", "Output file to save results")
	concurrency := flag.Int("c", 10, "Number of concurrent workers")

	flag.Parse()

	// Check if the required -l flag is provided
	if *inputFile == "" {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Open the input file
	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Printf("Error opening input file: %v\n", err)
		return
	}
	defer file.Close()

	// Create the output file
	output, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		return
	}
	defer output.Close()

	writer := bufio.NewWriter(output)

	// Channels for domains and results
	domains := make(chan string, *concurrency)
	results := make(chan string, *concurrency)

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go worker(domains, results, &wg)
	}

	// Read domains from the input file
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			domain := scanner.Text()
			domains <- domain
		}
		close(domains)
	}()

	// Wait for workers to finish and close results channel
	go func() {
		wg.Wait()
		close(results)
	}()

	// Write results to output file
	for domain := range results {
		_, err := writer.WriteString(domain + "\n")
		if err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
		}
	}

	// Flush the writer buffer
	writer.Flush()
	fmt.Println("Check completed. Results saved to", *outputFile)
}
