# sdigo
Intrusion Detection System by Golang language

This project is not commercial. The goal of the development was to gain knowledge on working with packages using the Go programming languag.

Found errors, most likely, will not be eliminated. But you can do it for me (if you have nothing to do 😜)

# Run

```bash
go run cmd/main.go --file=resources/pcap/*.pcap --rules=resources/rules/*.json
```

## Rules

All rights from this file will be executed with the prefix "AND" and will be executed together with each other

You can add your own rules to the base list with the following construct:

```json
[
  {
    "rule_value": "HTTP"
  },
  {
    "rule_value": "GET"
  },
  {
    "rule_value": "## your new rule ##"
  }
]
```
