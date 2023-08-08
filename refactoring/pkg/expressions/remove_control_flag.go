package expressions

import "fmt"

func sendAlert(name string) {
	fmt.Println("alert name: ", name)
}

// --- origin
func checkSecurityOri(people []string) {
	found := false
	for _, p := range people {
		if !found {
			if p == "Don" {
				sendAlert(p)
				found = true
			}
			if p == "John" {
				sendAlert(p)
				found = true
			}
		}
	}
}

// --- modify
func checkSecurity(people []string) {
	for _, p := range people {
		if !isSecurityName(p) {
			sendAlert(p)
			break
		}
	}
}

func isSecurityName(name string) bool {
	return name != "John" && name != "Don"
}

func TestRemoveControlFlag() {
	people := []string{"ayx", "cc"}
	checkSecurityOri(people)
	checkSecurity(people)

	people1 := []string{"xx", "John"}
	checkSecurityOri(people1)
	checkSecurity(people1)
}
