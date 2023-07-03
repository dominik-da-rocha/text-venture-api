package v1

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
)

func GuestBook(router *mux.Router, database string) {
   db, err := initializeDB(database)   
   if (err != nil) {
      log.Fatal(err.Error())
   }
   
   count, err := getRowCount(db)
   if (err != nil) {
      log.Fatal(err.Error())
   }
   
   if (count <=0) {
      err= fillDatabaseWithEntries(db)
      if (err != nil) {
         log.Fatal(err.Error())
      }
   }

   entries, err := getGuestBookEntries(db);
   if err != nil {
      log.Fatal(err.Error())
   }

	url := "/api/v1/guest-book"
	router.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "application/json")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

      json, err := json.Marshal(entries);
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		bytes, err := w.Write(json)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Printf("handled %d bytes %d %s %s", bytes, http.StatusOK, http.StatusText(http.StatusOK), url)
	}).Methods(http.MethodGet)

	router.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {

         var entry guestBookEntry
			err := json.NewDecoder(r.Body).Decode(&entry)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

         entries = append([]guestBookEntry{entry}, entries...)
         saveGuestBookEntry(db, entry)        

			w.WriteHeader(http.StatusCreated)
			log.Printf("handled %d %s %s", http.StatusCreated, http.StatusText(http.StatusCreated), url)
	}).Methods(http.MethodPost)   
}


type guestBookEntry struct {
   ID      int`json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func initializeDB(database string) (*sql.DB, error) {
   err := os.MkdirAll(database, os.ModePerm)
	if err != nil {
		return nil, err
	}

   databaseFullPath := filepath.Join(database, "guestbook.db")

	db, err := sql.Open("sqlite3", databaseFullPath) // Open the SQLite database
	if err != nil {
		return nil, err
	}

	// Create the guestbook_entries table if it doesn't exist
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS guestbook_entries (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			comment TEXT
		);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func saveGuestBookEntry(db *sql.DB, entry guestBookEntry) error {
	insertSQL := "INSERT INTO guestbook_entries (name, comment) VALUES (?, ?)"
	_, err := db.Exec(insertSQL, entry.Name, entry.Comment)
	return err
}

func getGuestBookEntries(db *sql.DB) ([]guestBookEntry, error) {
	query := "SELECT id, name, comment FROM guestbook_entries ORDER BY id DESC"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	entries := []guestBookEntry{}
	for rows.Next() {
		var entry guestBookEntry
		err := rows.Scan(&entry.ID, &entry.Name, &entry.Comment)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func getRowCount(db *sql.DB) (int, error) {
	query := "SELECT COUNT(*) FROM guestbook_entries"
	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func fillDatabaseWithEntries(db *sql.DB) error {
   var entries = []guestBookEntry{
      {         
         Comment: "Mind-bending and captivating! A journey through dimensions that will leave you questioning reality.",
         Name:    "The Stellar Paradox",
      },
      {         
         Comment: "A visionary masterpiece! This book pushes the boundaries of imagination and explores the depths of human existence.",
         Name:    "Nebula's End",
      },
      {         
         Comment: "Prepare to be transported to a dystopian future where technology reigns supreme. A thrilling and thought-provoking adventure!",
         Name:    "Quantum Nexus",
      },
	}

	// Prepare the SQL statement for insertion
	insertSQL := "INSERT INTO guestbook_entries (name, comment) VALUES (?, ?)"
	stmt, err := db.Prepare(insertSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insert each entry into the database
	for _, entry := range entries {
		_, err := stmt.Exec(entry.Name, entry.Comment)
		if err != nil {
			return err
		}
	}

	return nil
}


