package main
 
import (
    "github.com/Euvaz/Backstage-Hive/cmd"
)
 
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
    cmd.Execute()

//    genEnrollmentToken(db, vi.GetString("host"), vi.GetInt("port"))
}
