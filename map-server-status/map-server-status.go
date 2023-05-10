//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import "fmt"

const (
	Online = iota
	Offline
	Maintenance
	Retired
)

func printServerStatus(servers map[string]int) {
	fmt.Println("\nThere are", len(servers), " servers")

	stats := make(map[int]int)

	for _, status := range servers {
		switch status {
		case Online:
			stats[Online]++
		case Offline:
			stats[Offline]++
		case Maintenance:
			stats[Maintenance]++
		case Retired:
			stats[Retired]++
		default:
			panic("Unhandled Request")
		}
	}

	fmt.Println("# of Online Servers", stats[Online])
	fmt.Println("# of Offline Servers", stats[Offline])
	fmt.Println("# of Maintenance Servers", stats[Maintenance])
	fmt.Println("# of Retired Servers", stats[Retired])

}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatus := make(map[string]int)

	for _, server := range servers {
		serverStatus[server] = Online
	}

	printServerStatus(serverStatus)
	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline
	printServerStatus(serverStatus)

	for server, status := range serverStatus {
		if status == Online {
			serverStatus[server] = Maintenance
		}
	}
	printServerStatus(serverStatus)
}
