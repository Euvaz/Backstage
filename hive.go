package main
 
import (
    "database/sql"
    "fmt"
    "github.com/Euvaz/Backstage-Hive/cmd"
    _ "github.com/gin-gonic/gin"
    _ "github.com/jackc/pgx/v5/stdlib"
    "github.com/spf13/viper"
    "log"
    _ "net/http"
    "os"
)
 
// Function to generate an enrollment token
//func genEnrollmentToken(db *sql.DB, host string, port int) {
//    var key string = RandStringBytes(50)
//    var enrollmentToken string = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"addr":"%s:%v","key":"%s"}`, host, port, key)))
//    var execStr string = fmt.Sprintf(`INSERT INTO tokens (id, key, created)
//                                      VALUES (DEFAULT, '%s', CURRENT_TIMESTAMP)`, key)
//    db.Exec(execStr)
//    log.Printf("Generated Token: \"%s\"", enrollmentToken)
//    log.Printf("Generated Key: \"%s\"", key)
//}

// Function to enroll a Drone into the Hive inventory
func enrollDrone(db *sql.DB) {
    var droneAddress string = "10.13.0.25"
    var dronePort int = 3802
    var droneName string = "drone-1"
    var execStr string = fmt.Sprintf(`INSERT INTO drones (id, address, port, name)
                                      VALUES (DEFAULT, '%s', %v, '%s')`,
                                      droneAddress, dronePort, droneName)
    db.Exec(execStr)
    log.Printf("drone \"%s\" Enrolled", droneName)
}

func main() {
    log.SetFlags(log.Lshortfile)
    log.SetPrefix("Backstage-Hive: ")

    vi := viper.New()
    vi.SetConfigFile("config.yaml")
    err := vi.ReadInConfig()
    if err != nil {
        log.Fatalln(err)
        os.Exit(1)
    }

    vi.SetDefault("host", "localhost")
    vi.SetDefault("port", 6789)
    vi.SetDefault("dbHost", "localhost")
    vi.SetDefault("dbPort", 5432)
    vi.SetDefault("dbUser", "backstage")
    vi.SetDefault("dbPass", "backstage")
    vi.SetDefault("dbName", "backstage")

//    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
//                             vi.GetString("dbHost"), vi.GetInt("dbPort"), vi.GetString("dbUser"),
//                             vi.GetString("dbPass"), vi.GetString("dbName"))
//
//    // Connect to database
//    log.Printf("Connecting to database...")
//    db, err := sql.Open("pgx", psqlconn)
//    if err != nil {
//        log.Fatalln(err)
//        os.Exit(1)
//    }
//    defer log.Printf("Database connection closed")
//    defer db.Close()
//
//    // Verify database connection
//    err = db.Ping()
//    if err != nil {
//        log.Fatalln(err)
//        os.Exit(1)
//    }
//    log.Printf("Connection established")
//
//    // Initialize database
//    initDB(db)

    HiveCmd.Execute()

    //genEnrollmentToken(db, vi.GetString("host"), vi.GetInt("port"))
}
