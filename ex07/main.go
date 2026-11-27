package main

import "fmt"

func main() {
	var modules = map[int]string{
		104: "Datenmodell implementieren",
		117: "Netzwerk mit Horst",
		346: "Cloudlösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 346)
	modules[254] = "Nicolette Modul"
	modules[117] = "Bienen-Modul"
	fmt.Println(modules)
}
