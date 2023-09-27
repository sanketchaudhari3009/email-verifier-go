# Email Verifier Go

![GitHub](https://img.shields.io/github/license/sanketchaudhari3009/email-verifier-go)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sanketchaudhari3009/email-verifier-go)

**Email Verifier Go** is a simple command-line tool to check domain information, including MX records, SPF records, and DMARC records.

## Features

- Verify domain information:
  - Check MX (Mail Exchange) records.
  - Check SPF (Sender Policy Framework) records.
  - Check DMARC (Domain-based Message Authentication, Reporting, and Conformance) records.

## Usage

To verify a domain, simply enter the domain name when prompted:

```shell
$ go run main.go
Enter a domain (or 'exit' to quit): example.com
