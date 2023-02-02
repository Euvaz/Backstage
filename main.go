package main
 
import (
    "database/sql"
    "log"
    "os"
    "fmt"
    _ "github.com/jackc/pgx/v5/stdlib"
)
 
const (
    dbHost     = "localhost"
    dbPort     = 5432
    dbUser     = "user"
    dbPass     = "pass"
    dbName     = "db"
)

func initTables(db *sql.DB) {
    log.Printf("Initializing Tables...")
    initTableNodes(db)
    initTableTokens(db)
    defer log.Printf("Tables successfully initialized")
}

func initTableNodes(db *sql.DB) {
    log.Printf("Creating \"nodes\" table if not already present...")
    var execStr string = `CREATE TABLE IF NOT EXISTS nodes (
                            id SERIAL PRIMARY KEY,
                            address INET,
                            port INTEGER,
                            name TEXT,
                            UNIQUE (address, port)
                          )`
    db.Exec(execStr)
    log.Printf("Success")
}

func initTableTokens(db *sql.DB) {
    log.Printf("Creating \"tokens\" table if not alrady present...")
    var execStr string = `CREATE TABLE IF NOT EXISTS tokens (
                            id SERIAL PRIMARY KEY,
                            created TIMESTAMP
                          )`
    db.Exec(execStr)
    log.Printf("Success")
}

func enrollNode(db *sql.DB) {
    var nodeAddress string = "10.13.0.25"
    var nodePort int = 3802
    var nodeName string = "node-1"
    var execStr string = fmt.Sprintf(`INSERT INTO nodes (id, address, port, name)
                                      VALUES (DEFAULT, '%s', %v, %s)`,
                                      nodeAddress, nodePort, nodeName)
    db.Exec(execStr)
    log.Printf("Node %s enrolled", nodeName)
}

func main() {
    log.SetFlags(log.Lshortfile)
    log.SetPrefix("Backstage: ")

    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
                             dbHost, dbPort, dbUser, dbPass, dbName)

    log.Printf("Connecting to database...")
    db, err := sql.Open("pgx", psqlconn)
    if err != nil {
        log.Fatalln(err)
        os.Exit(1)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatalln(err)
        os.Exit(1)
    }
    log.Printf("Connection established")

    initTables(db)
}
