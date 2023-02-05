package cmd

import (
    "crypto/rand"
    "database/sql"
    "encoding/base64"
    "fmt"
    "github.com/Euvaz/Backstage-Hive/logger"
    _ "github.com/gin-gonic/gin"
    _ "github.com/jackc/pgx/v5/stdlib"
	"github.com/spf13/cobra"
    "github.com/spf13/viper"
	"os"
)

// Function to initialize Viper
func initViper() *viper.Viper {
    viper := viper.New()
    viper.SetConfigFile("config.toml")
    err := viper.ReadInConfig()
    if err != nil {
        logger.Fatal(err.Error())
        os.Exit(1)
    }

    viper.SetDefault("host", "localhost")
    viper.SetDefault("port", 6789)
    viper.SetDefault("dbHost", "localhost")
    viper.SetDefault("dbPort", 5432)
    viper.SetDefault("dbUser", "backstage")
    viper.SetDefault("dbPass", "backstage")
    viper.SetDefault("dbName", "backstage")

    return viper
}

// Function to initialize database
func initDB(db *sql.DB) {
    var err error

    logger.Info("Initializing Tables...")
    defer logger.Info("Tables successfully initialized")
    
    // Create "drones" table
    logger.Info("Creating \"drones\" table if not already present...")
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS drones (
                        id SERIAL PRIMARY KEY,
                        address INET,
                        port INTEGER,
                        name TEXT,
                        UNIQUE (address, port),
                        UNIQUE (name)
                      )`)
    if err != nil {
        logger.Fatal(err.Error())
    }
    logger.Info("Success")

    // Create "permissions" table
    logger.Info("Creating \"permissions\" table if not already present...")
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS permissions (
                        id SERIAL PRIMARY KEY,
                        name TEXT
                      )`)
    if err != nil {
        logger.Fatal(err.Error())
    }
    logger.Info("Success")

    // Create "groups" table
    logger.Info("Creating \"groups\" table if not already present...")
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS groups (
                        id SERIAL PRIMARY KEY,
                        name TEXT,
                        permissions_id SERIAL
                        REFERENCES permissions (id),
                        UNIQUE (name)
                      )`)
    if err != nil {
        logger.Fatal(err.Error())
    }
    logger.Info("Success")

    // Create "swarms" table
    logger.Info("Creating \"swarms\" table if not already present...")
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS swarms (
                        id SERIAL PRIMARY KEY,
                        name TEXT,
                        drones_id SERIAL
                        REFERENCES drones (id),
                        UNIQUE (name)
                      )`)
    if err != nil {
        logger.Fatal(err.Error())
    }
    logger.Info("Success")

    // Create "tokens" table
    logger.Info("Creating \"tokens\" table if not alrady present...")
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS tokens (
                        id SERIAL PRIMARY KEY,
                        key TEXT,
                        created TIMESTAMP,
                        UNIQUE (key)
                      )`)
    if err != nil {
        logger.Fatal(err.Error())
    }
    logger.Info("Success")

    // Create "users" table
    logger.Info("Creating \"users\" table if not already present...")
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
                        id SERIAL PRIMARY KEY,
                        name TEXT,
                        groups_id SERIAL
                        REFERENCES groups (id),
                        pass TEXT,
                        created TIMESTAMP,
                        UNIQUE (name)
                      )`)
    if err != nil {
        logger.Fatal(err.Error())
    }
    logger.Info("Success")
}

// Function to generate a random alphanumeric string of set length
func RandStringBytes(n int) string {
    randomBytes := make([]byte, 64)
    _, err := rand.Read(randomBytes)
    if err != nil {
        logger.Fatal(err.Error())
    }
    return base64.StdEncoding.EncodeToString(randomBytes)[:n]
}

func dbConnect(host string, port int, user string, pass string, name string) *sql.DB {

    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
                             host, port, user, pass, name)

    // Connect to database
    logger.Info("Connecting to database...")
    database, err := sql.Open("pgx", psqlconn)
    if err != nil {
        logger.Fatal(err.Error())
        os.Exit(1)
    }

    // Verify database connection
    err = database.Ping()
    if err != nil {
        logger.Fatal(err.Error())
        os.Exit(1)
    }
    logger.Info("Connection established")

    return database
}

func dbClose(database *sql.DB) {
    db.Close()
    logger.Info("Database connection closed")
}

var (
    vi = initViper()
    db = dbConnect(vi.GetString("dbHost"), vi.GetInt("dbPort"), vi.GetString("dbUser"),
                   vi.GetString("dbPass"), vi.GetString("dbName"))

    HiveCmd = &cobra.Command {
	    Use:   "Backstage-Hive",
	    Short: "Short Desc",
	    Long:  `Long
                Desc`,

        PersistentPreRun: func(cmd *cobra.Command, args []string) {
        },

        Run: func(cmd *cobra.Command, args []string) {
            // Initialize database
            initDB(db)
            defer dbClose(db)
        },
    }
)

func Execute() {
	err := HiveCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	HiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
