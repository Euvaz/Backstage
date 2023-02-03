package main
 
import (
    "database/sql"
    "encoding/base64"
    "fmt"
    _ "github.com/jackc/pgx/v5/stdlib"
    "github.com/spf13/viper"
    "log"
    "crypto/rand"
    _ "os"
    "github.com/Euvaz/Backstage-Hive/db"
)
 
// Function to initialize database tables
func initTables(db *sql.DB) {
    log.Printf("Initializing Tables...")
    initTableSwarm(db)
    initTableTokens(db)
    defer log.Printf("Tables successfully initialized")
}

// Function to initialize the "swarm" table
func initTableSwarm(db *sql.DB) {
    log.Printf("Creating \"swarm\" table if not already present...")
    var execStr string = `CREATE TABLE IF NOT EXISTS swarm (
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
    randomBytes := make([]byte, 64)
    _, err := rand.Read(randomBytes)
    if err != nil {
        log.Println(err)
    }
    return base64.StdEncoding.EncodeToString(randomBytes)[:n]
}

// Function to generate an enrollment token
func genEnrollmentToken(db *sql.DB, host string, port int) {
    var key string = RandStringBytes(50)
    var enrollmentToken string = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("{addr:\"%s:%v\",key:\"%s\"}", host, port, key)))
    var execStr string = fmt.Sprintf(`INSERT INTO tokens (id, name, created)
                                      VALUES (DEFAULT, '%s', CURRENT_TIMESTAMP)`, key)
    db.Exec(execStr)
    log.Printf("Generated Token: \"%s\"", enrollmentToken)
    log.Printf("Generated Key: \"%s\"", key)
}

// Function to enroll a drone in the swarm
func enrollDrone(db *sql.DB) {
    var droneAddress string = "10.13.0.25"
    var dronePort int = 3802
    var droneName string = "drone-1"
    var execStr string = fmt.Sprintf(`INSERT INTO swarm (id, address, port, name)
                                      VALUES (DEFAULT, '%s', %v, '%s')`,
                                      droneAddress, dronePort, droneName)
    db.Exec(execStr)
    log.Printf("drone \"%s\" Enrolled", droneName)
}

func main() {
    log.SetFlags(log.Lshortfile)
    log.SetPrefix("Backstage: ")

    vi := viper.New()
    vi.SetConfigFile("config.yaml")
    vi.ReadInConfig()

    db := pg.Connect(vi.GetString("dbHost"), vi.GetInt("dbPort"), vi.GetString("dbUser"),
                     vi.GetString("dbPass"), vi.GetString("dbName"))
    defer pg.Disconnect(db)

    //initTables(db)

    //genEnrollmentToken(db, vi.GetString("host"), vi.GetInt("port"))
}
