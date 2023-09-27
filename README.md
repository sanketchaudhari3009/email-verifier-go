# Email Verifier Go

**Email Verifier Go** is a simple command-line tool to check domain information, including MX records, SPF records, and DMARC records.

## Features

- Verify domain information:
  i] Check MX (Mail Exchange) records.
  ii] Check SPF (Sender Policy Framework) records.
  iii] Check DMARC (Domain-based Message Authentication, Reporting, and Conformance) records.

## Usage

To verify a domain, simply enter the domain name when prompted:

```shell
$ go run main.go
Enter a domain (or 'exit' to quit): example.com
