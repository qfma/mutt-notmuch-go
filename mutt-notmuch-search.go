package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

// Get the search query from standard input
func get_query() string {
	fmt.Print("Query: ")
	in := bufio.NewReader(os.Stdin)
	query, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return query
}

// Execute the notmuch command using the search query
// notmuch is used with the --output=files option
// that returns a list of files (hits) that match the query.
// The return the hits into a slice of strings
func run_notmuch(query string) []string {
	out, err := exec.Command("notmuch", "search", "--output=files", query).Output()
	if err != nil {
		log.Fatal(err)
	}
	// Convert the []bytes into a string, trim and split
	hits := strings.Split(strings.TrimSpace(string(out)), "\n")
	return hits
}

// Make a folder that is compatible to the Maildir specification
func make_maildir(path string) {
	os.MkdirAll(path, 0700)
	os.Mkdir(filepath.Join(path, "cur"), 0700)
	os.Mkdir(filepath.Join(path, "new"), 0700)
	os.Mkdir(filepath.Join(path, "tmp"), 0700)
}

// Create symlinks in the cur folder of the maildir that link
// to the notmuch hits
func link_results(path string, hits []string) {
	for _, hit := range hits {
		base := filepath.Base(hit)
		link := filepath.Join(path, "cur", base)
		os.Symlink(hit, link)
	}
}

// Empty the cur folder and recreate it afterwards
// Can I just empty all files without deleting the folder?
func empty_maildir(path string) {
	err := os.RemoveAll(filepath.Join(path, "cur"))
	if err != nil {
		log.Fatal(err)
	}
	os.Mkdir(filepath.Join(path, "cur"), 0700)
}

func main() {
	// Get search query string
	query := get_query()
	// Run the notmuch command and return the hits
	hits := run_notmuch(query)

	// Get the home directory of the user
	usr, _ := user.Current()
	home := usr.HomeDir
	// Make the temporary results maildir folder
	mutt_results := filepath.Join(home, ".cache/mutt_results")

	// Test if the output folder exists and empty previous results
	if _, err := os.Stat(mutt_results); err != nil {
		make_maildir(mutt_results)
	} else {
		empty_maildir(mutt_results)
	}

	// Finally create the links to the notmuch hits
	link_results(mutt_results, hits)
}
