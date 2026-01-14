package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"net/http"

	"crypto/tls"
	"crypto/x509"

	_ "modernc.org/sqlite"
)

type Check struct {
	Online int `json:"online"`
}

func main() {
	url := flag.String("url", "https://example.com", "url to api")
	caPath := flag.String("ca", "ca.pem", "path to CA certificate")
	dbPath := flag.String("db", "data.sqlite3", "path to sqlite3 storage")

	flag.Parse()

	db, err := initDb(*dbPath)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	resp, err := request(*url, *caPath)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var check Check
	err = json.NewDecoder(resp.Body).Decode(&check)

	if err != nil {
		log.Fatal(err)
	}

	err = safeCheck(db, check)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done!")
}

func initDb(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", fmt.Sprintf("%s?_journal_mode=WAL&_busy_timeout=5000", dbPath))

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS checks (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		current_online INTEGER NOT NULL,
		created_at DATETIME NOT NULL
	);`)

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func request(url, caPath string) (*http.Response, error) {
	caCert, err := os.ReadFile(caPath)

	if err != nil {
		return nil, err
	}

	caPool := x509.NewCertPool()

	if !caPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to append CA certificate")
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caPool,
			},
		},
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	return resp, nil
}

func safeCheck(db *sql.DB, record Check) error {
	_, err := db.Exec("INSERT INTO checks(current_online, created_at) VALUES(?, ?)", record.Online, time.Now().UTC())

	return err
}
