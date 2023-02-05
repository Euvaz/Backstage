package main
 
import (
    "github.com/Euvaz/Backstage-Hive/cmd"
    "log"
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
//func enrollDrone(db *sql.DB) {
//    var droneAddress string = "10.13.0.25"
//    var dronePort int = 3802
//    var droneName string = "drone-1"
//    var execStr string = fmt.Sprintf(`INSERT INTO drones (id, address, port, name)
//                                      VALUES (DEFAULT, '%s', %v, '%s')`,
//                                      droneAddress, dronePort, droneName)
//    db.Exec(execStr)
//    log.Printf("drone \"%s\" Enrolled", droneName)
//}

func main() {
    log.SetFlags(log.Lshortfile)
    log.SetPrefix("Backstage-Hive: ")

    HiveCmd.Execute()

//    genEnrollmentToken(db, vi.GetString("host"), vi.GetInt("port"))
}
