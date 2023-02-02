package main
 
import (
    "database/sql"
    "encoding/base64"
    "log"
    "os"
    "fmt"
    _ "github.com/jackc/pgx/v5/stdlib"
    "github.com/spf13/viper"
    "math/rand"
)
 
const (
    alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
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
                            name TEXT,
                            created TIMESTAMP
                          )`
    db.Exec(execStr)
    log.Printf("Success")
}

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = alphanum[rand.Intn(len(alphanum))]
    }
    return string(b)
}

func genEnrollmentToken(db *sql.DB, host string, port int) {
    var enrollmentToken string = RandStringBytes(50)
    var enrollmentTokenJSON string = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("{addr:\"%s:%v\",token:\"%s\"}", host, port, enrollmentToken)))
    var execStr string = fmt.Sprintf(`INSERT INTO tokens (id, name, created)
                                      VALUES (DEFAULT, '%s', CURRENT_TIMESTAMP)`,
                                      enrollmentToken)
    db.Exec(execStr)
    log.Printf("Token Generated")
}

func enrollNode(db *sql.DB) {
    var nodeAddress string = "10.13.0.25"
    var nodePort int = 3802
    var nodeName string = "node-1"
    var execStr string = fmt.Sprintf(`INSERT INTO nodes (id, address, port, name)
                                      VALUES (DEFAULT, '%s', %v, '%s')`,
                                      nodeAddress, nodePort, nodeName)
    db.Exec(execStr)
    log.Printf("Node \"%s\" enrolled", nodeName)
}

func main() {
    log.SetFlags(log.Lshortfile)
    log.SetPrefix("Backstage: ")

    vi := viper.New()
    vi.SetConfigFile("config.yaml")
    vi.ReadInConfig()

    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
                             vi.GetString("dbHost"), vi.GetInt("dbPort"), vi.GetString("dbUser"),
                             vi.GetString("dbPass"), vi.GetString("dbName"))

    log.Printf("Connecting to database...")
    db, err := sql.Open("pgx", psqlconn)
    if err != nil {
        log.Fatalln(err)
        os.Exit(1)
    }
    defer db.Close()
    defer log.Printf("Database connection closed")

    err = db.Ping()
    if err != nil {
        log.Fatalln(err)
        os.Exit(1)
    }
    log.Printf("Connection established")

    initTables(db)

    //genEnrollmentToken(db, vi.GetString("host"), vi.GetInt("port"))
}
