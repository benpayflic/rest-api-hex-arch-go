package database

import (
	"database/sql"
	"log"

	"github.com/benpayflic/rest-api-hex-arch-go/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

type Adapter struct {
	conn   *sql.DB
	config *config.Config
}

func NewAdapter(config *config.Config) (*Adapter, error) {
	conn := &sql.DB{}
	return &Adapter{conn: conn, config: config}, nil
}

func (dba *Adapter) Connect() {

	log.Printf("Connecting to database at: %v ... \n", dba.config.DbPath)
	connection, err := sql.Open("sqlite3", dba.config.DbPath)
	if err != nil {
		log.Fatalln("Failed to connect to Database: ", err)
	}
	(*dba).conn = connection
	log.Printf("Connected to database at: %v !\n", dba.config.DbPath)

	// Create customer table
	statement, err :=
		dba.conn.Prepare("CREATE TABLE IF NOT EXISTS Cats (ID INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255) NOT NULL UNIQUE, breed VARCHAR(255) NOT NULL, dob VARCHAR(255) NOT NULL, humanYears INTEGER NOT NULL, catFact VARCHAR(255) NOT NULL, CreatedAt DATETIME NOT NULL, UpdatedAt DATETIME NULL)")
	if err != nil {
		log.Fatalln("Failed to connect to Database: ", err)
	}
	statement.Exec()
}
