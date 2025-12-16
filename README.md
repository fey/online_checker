# Players Online Checker

A small Go CLI tool that retrieves the number of online players from a remote API and stores the results in a local SQLite database.

## Requirements

- Go 1.23+ (or compatible)
- SQLite3
- A valid CA certificate (`.pem`) if the API uses a self-signed certificate

## Usage

```bash
./checks -h
```

### Example

```bash
./checker -ca certs/ca.pem -db /var/data/players.sqlite3 -url https://example.com/get_online
```
