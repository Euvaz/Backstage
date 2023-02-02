package main
 
import (
    "database/sql"
    "encoding/base64"
    "fmt"
    _ "github.com/jackc/pgx/v5/stdlib"
    "github.com/spf13/viper"
    "log"
    "math/rand"
    "os"
)
 
const (
    alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// Function to initialize database tables
func initTables(db *sql.DB) {
    log.Printf("Initializing Tables...")
    initTableWorkers(db)
    initTableTokens(db)
    defer log.Printf("Tables successfully initialized")
}

// Function to initialize the "workers" table
func initTableWorkers(db *sql.DB) {
    log.Printf("Creating \"workers\" table if not already present...")
    var execStr string = `CREATE TABLE IF NOT EXISTS workers (
                            id SERIAL PRIMARY KEY,
                            address INET,
                            port INTEGER,
                            name TEXT,
                            UNIQUE (address, port)
                          )`
    db.Exec(execStr)
    log.Printf("Success")
}

// Function to initialize the "tokens" table
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

// Function to generate a random alphanumeric string of set length
func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = alphanum[rand.Intn(len(alphanum))]
    }
    return string(b)
}

// Function to generate an enrollment token
func genEnrollmentToken(db *sql.DB, host string, port int) {
    var key string = RandStringBytes(50)
    var enrollmentToken string = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("{addr:\"%s:%v\",key:\"%s\"}", host, port, key)))
    var execStr string = fmt.Sprintf(`INSERT INTO tokens (id, name, created)
                                      VALUES (DEFAULT, '%s', CURRENT_TIMESTAMP)`, key)
    db.Exec(execStr)
    log.Printf("Generated Token: \"%s\"", enrollmentToken)
}

// Function to enroll a worker in the cluster
func enrollWorker(db *sql.DB) {
    var workerAddress string = "10.13.0.25"
    var workerPort int = 3802
    var workerName string = "worker-1"
    var execStr string = fmt.Sprintf(`INSERT INTO workers (id, address, port, name)
                                      VALUES (DEFAULT, '%s', %v, '%s')`,
                                      workerAddress, workerPort, workerName)
    db.Exec(execStr)
    log.Printf("worker \"%s\" Enrolled", workerName)
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
    defer log.Printf("Database connection closed")
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatalln(err)
        os.Exit(1)
    }
    log.Printf("Connection established")

    initTables(db)

    //genEnrollmentToken(db, vi.GetString("host"), vi.GetInt("port"))
}
