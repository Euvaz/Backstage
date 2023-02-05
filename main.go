package main

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
)

func main() {
	viper := viper.New()
	viper.SetConfigFile("config.toml")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(err.Error())
	}

	viper.SetDefault("host", "localhost")
	viper.SetDefault("port", 6789)
	viper.SetDefault("dbHost", "localhost")
	viper.SetDefault("dbPort", 5432)
	viper.SetDefault("dbUser", "backstage")
	viper.SetDefault("dbPass", "backstage")
	viper.SetDefault("dbName", "backstage")

	db := getDB(viper.GetString("dbHost"), viper.GetInt("dbPort"), viper.GetString("dbUser"), viper.GetString("dbPass"), viper.GetString("dbName"))

    // Add root command
	cmd := &cobra.Command {
		Use:   "Backstage-Hive",
		Short: "Short Desc",
		Long:  `Long Desc`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
		},
		Run: func(cmd *cobra.Command, args []string) {
			logger.Info("Starting server")
			initDB(db)
			defer closeDB(db)
		},
	}

    // Add command
    createCmd := &cobra.Command {
        Use:   "create",
        Short: "Short Desc",
        Long:  `Long Desc`,
    }

    // Add subcommand
    createTokenCmd := &cobra.Command {
        Use:   "token",
        Short: "Short Desc",
        Long:  `Long Desc`,
        Run: func(cmd *cobra.Command, args []string) {
            genEnrollmentToken(db, viper.GetString("host"), viper.GetInt("port"))
        },
    }

    // Add command
    getCmd := &cobra.Command {
        Use:   "get",
        Short: "Short Desc",
        Long:  `Long Desc`,
    }

    // Add commands
    cmd.AddCommand(createCmd)
    cmd.AddCommand(getCmd)

    // Add subcommands
    createCmd.AddCommand(createTokenCmd)


	err = cmd.Execute()
	if err != nil {
		logger.Fatal(err.Error())
	}
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

func initDB(db *sql.DB) {
	var err error

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

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS permissions (
                        id SERIAL PRIMARY KEY,
                        name TEXT
                      )`)
	if err != nil {
		logger.Fatal(err.Error())
	}

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

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tokens (
                        id SERIAL PRIMARY KEY,
                        key TEXT,
                        created TIMESTAMP,
                        UNIQUE (key)
                      )`)
	if err != nil {
		logger.Fatal(err.Error())
	}

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

	logger.Info("Tables successfully initialized")
}

func getDB(host string, port int, user string, pass string, name string) *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name)

	logger.Debug("Connecting to database...")
	database, err := sql.Open("pgx", psqlconn)
	if err != nil {
		logger.Fatal(err.Error())
	}

	err = database.Ping()
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Debug("Connection established")

	return database
}

func closeDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("Database connection closed")
}

// Function to generate an enrollment token
func genEnrollmentToken(db *sql.DB, host string, port int) {
    logger.Debug("Creating token...")
    var key string = RandStringBytes(50)
    var enrollmentToken string = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf(`{"addr":"%s:%v","key":"%s"}`, host, port, key)))
   var execStr string = fmt.Sprintf(`INSERT INTO tokens (id, key, created)
                                      VALUES (DEFAULT, '%s', CURRENT_TIMESTAMP)`, key)
    db.Exec(execStr)
    fmt.Println("Generated Token:", enrollmentToken)
    logger.Debug("Created token")
}

// Function to enroll a Drone into the Hive inventory
//func enrollDrone(db *sql.DB) {
//    var droneAddress string = "10.13.0.25"
//    var dronePort int = 3802
//    var droneName string = "drone-1"
//    var execStr string = fmt.Sprintf(`INSERT INTO drones (id, address, port, name)
//                                      VALUES (DEFAULT, '%s', %v, '%s')`,
//                                      droneAddress, dronePort, droneName)
//    db.Exec(execStr)
//    logger.Infof("drone \"%s\" Enrolled", droneName)
//}
