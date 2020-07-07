package powerdns

import (
	"fmt"
	"log"
)

// Error structure with customised error message
type Error struct {
	Status  string
	Message string `json:"error"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%v %v", e.Status, e.Message)
}

// usage prints both create and delete usage
func usage() {
	log.Print(createUsage)
	log.Print(deleteUsgage)
}

// logerr error check and logging
func logerr(err error, x ...int) {
	if err != nil {
		log.Fatalln(err)
	}
}
