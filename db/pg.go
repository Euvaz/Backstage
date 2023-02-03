package pg

import (
    "database/sql"
    "fmt"
    "log"
    "os"
)

func Connect(host string, port int, user string, pass string, name string) *sql.DB {
    conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
                        host, port, user, pass, name)
    log.Printf("Connecting to database...")
    state, err := sql.Open("pgx", conn)
    if err != nil {
        log.Fatalln(err)
        os.Exit(1)
    }

    err = state.Ping()
    if err != nil {
        log.Fatalln(err)
        os.Exit(1)
    }
    log.Printf("Connection established")
    return state
}

func Disconnect(state *sql.DB) {
    log.Printf("Disconnecting from database...")
    state.Close()
    log.Printf("Disconnected from database")
}
