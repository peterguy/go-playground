package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/keegancsmith/sqlf"
	_ "github.com/lib/pq"
)

func main() {
	// Database connection settings
	dsn := "host=localhost port=5433 user=myuser password=mypassword dbname=mydatabase sslmode=disable"

	// Connect to PostgreSQL
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Enable SQL logging for this session
	_, err = db.Exec("SET log_statement = 'all';")
	if err != nil {
		log.Fatalf("Failed to enable logging: %v", err)
	}

	// Define test cases
	testCases := []string{
		"name'; DROP TABLE users; --", // SQL injection attempt
		"O'Brien",                     // Single quote
		`C:\Users\John`,               // Backslash
		`C:\_sers\%John`,              // Backslash escaping a special character
		`C:\\_sers\\%John`,            // Double backslash - escape the escape character for a literal
		`C:\\Users\\John`,             // Double backslash - escape the escape character for a literal
		"100% Real",                   // Percent sign
		"test_case",                   // Underscore
		"北京",                          // Unicode
	}

	// Run the test cases
	for _, input := range testCases {
		// Use sqlf to construct the query safely
		query := sqlf.Sprintf("SELECT id, name FROM users WHERE name ILIKE %s", "%"+input+"%")

		// Print the raw query and arguments
		fmt.Printf("\nInput: %s\n", input)
		fmt.Println("Raw Query:", query.Query(sqlf.PostgresBindVar))
		fmt.Println("Arguments:", query.Args())

		// Execute the query
		rows, err := db.Query(query.Query(sqlf.PostgresBindVar), query.Args()...)
		if err != nil {
			log.Printf("Query failed for input '%s': %v\n", input, err)
			continue
		}
		defer rows.Close()

		fmt.Println("Results:")

		// Print query results
		found := false
		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				log.Fatalf("Failed to scan row: %v", err)
			}
			fmt.Printf("  ID: %d | Name: %s\n", id, name)
			found = true
		}

		if !found {
			fmt.Println("  (No matches)")
		}

		if err := rows.Err(); err != nil {
			log.Fatalf("Row iteration error: %v", err)
		}
	}

	fmt.Println("\nAll tests completed!")
}
