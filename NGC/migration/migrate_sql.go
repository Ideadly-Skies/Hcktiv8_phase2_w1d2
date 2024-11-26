package migration

/*
	Usage: go run . [file.sql] from the parent directory
	to populate the db using the DDL & DML Query
*/
import (
	"database/sql"
	"fmt"
	"io/ioutil"

	"os"
	"w2/d2/db"
	"strings"
)

func MigrateData() {
	// use recover to handle any potential panics
	defer HandlePanic()

	// Get the SQL filename from command line argument
	args := os.Args[1:]

	if len(args) < 1 {
		panic("No SQL file provided!")
	}
	
	// read the filename from the first argument
	filename := args[0]

	// read SQL commands from the file w2/d2/NGC/migration/query.sql
	
	sqlCommands, err := ReadSQLCommands(filename)

	if err != nil {
		panic(err)
	}

	// Connect to the database
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)		
	}

	defer db.Close()

	// execute SQL commands
	err = ExecuteSQLCommands(db, sqlCommands)
	if err != nil {
		panic(err)
	}

	fmt.Println("All Tables Created Successfully!\n")
}

// func to handle panic using recover
func HandlePanic(){
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic", r)
	}
}

// function to read SQL commands from a file
func ReadSQLCommands(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// function to execute SQL commands on the db
func ExecuteSQLCommands(db *sql.DB, commands string) error {
	statements := strings.Split(commands, ";") // split SQL commands by semicolon
	for _, stmt := range statements {
		stmt = strings.TrimSpace(stmt)		   // Trim whitespace
		if stmt != "" {
			_, err := db.Exec(stmt)
			if err != nil {
				return fmt.Errorf("failed to execute statement %q: %w", stmt, err)
			}
		}
	}
	return nil
}
