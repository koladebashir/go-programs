package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"log"
)

func GetBatteryPercentage() string {
	// pmset is a terminal command for power management
	cmd := exec.Command("pmset", "-g", "batt")

	// Ouput() returns the output of the command as a slice of bytes
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Command exited with error:", err)
	}
	o := string(output)


	s := strings.Split(o, ";")
	bp := s[0]

	str := strings.Split(bp, ")")[1]
	return strings.Split(strings.TrimSpace(str), "%")[0]
	
}

func main() {

	if len(os.Args) != 2 {
		log.Fatal(errors.New("insufficent number of arguments provided"))
	}

	// Takes the users specified battery level from the argument provided when the script runs
	userBatteryLevel := os.Args[1]

	for {
		batteryPercent := GetBatteryPercentage()
		if batteryPercent == userBatteryLevel {

			// osascript is used to execute Javascript and Applescript commands
			s := fmt.Sprintf(`display notification "Battery Level reached %v" with title "Battery Notification"`, batteryPercent)
			cmd2 := exec.Command("osascript", "-e", s)
			cmd2.Output()
			break
		}
		time.Sleep(2 * time.Second)

	}

	fmt.Println("Exiting program...")
}