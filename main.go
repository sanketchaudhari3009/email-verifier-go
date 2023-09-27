package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Print("Enter a domain (or 'exit' to quit): ")
		scanner := bufio.NewScanner(os.Stdin)

		// Read the user's input
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		domain, mx, spf, spfRecord, dmarc, dmarcRecord := checkDomain(input)
		printDomainInfo(domain, mx, spf, spfRecord, dmarc, dmarcRecord)
	}
}

func checkDomain(domain string) (string, string, string, string, string, string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecord, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecord) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	return domain, yesNo(hasMX), yesNo(hasSPF), spfRecord, yesNo(hasDMARC), dmarcRecord
}

func yesNo(flag bool) string {
	if flag {
		return "Yes"
	}
	return "No"
}

func printDomainInfo(domain, mx, spf, spfRecord, dmarc, dmarcRecord string) {

	fmt.Println(strings.Repeat("___", len(dmarcRecord)/2))

	fmt.Println("\nDomain:", domain)
	fmt.Println("Has MX Record:", mx)

	// Print separator line
	fmt.Println(strings.Repeat("___", len(dmarcRecord)/2))

	fmt.Println("\nHas SPF Record:", spf)
	fmt.Println("SPF Record:", spfRecord)

	// Print separator line
	fmt.Println(strings.Repeat("___", len(dmarcRecord)/2))

	fmt.Println("\nHas DMARC Record:", dmarc)
	fmt.Println("DMARC Record:", dmarcRecord)

	fmt.Println(strings.Repeat("___", len(dmarcRecord)/2))
}
