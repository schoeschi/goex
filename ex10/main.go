package main

import "fmt"

func main() {
	var users = []string{"alice", "bob", "charlene", "dan", "erica", "freddy"}
	var servers = []string{"web", "db", "proxy", "backup"}
	var needed = make([]string, 0, len(users)*len(servers))
	for _, u := range users {
		for _, s := range servers {
			hostname := fmt.Sprintf("%s_%s", u, s)
			needed = append(needed, hostname)
			fmt.Println(len(needed), cap(needed))
		}
	}
	fmt.Printf("needed %d servers: %v", len(needed), needed)
}
