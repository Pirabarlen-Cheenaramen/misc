//This is NOT complete
//<pirabarlen cheenaramen> selven@hackers.mu
//used along ipset.go for testing.
package main
import (

	"fmt"
	"github.com/pirabarlen/ipset"
)

func main() {

	var err error
	var output string
	var typenet = []string {"net","net"}
	var createit =map[string] string{
		"timeout"  : "30",
		"counters" : "",
	}
	output, err = ipset.Create("myset", "hash", typenet, createit)
	fmt.Println("Done", output ,err)

}
