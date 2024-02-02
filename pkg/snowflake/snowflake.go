package snowflake

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

var snowf *snowflake.Node

// func InitSnowflake(startTime string, machineId int64) (err error) {
func init() {
	// Create a new Node with a Node number of 1
	snowf, _ = snowflake.NewNode(10)
}

// Generate a snowflake ID.
func GetId() (id int64) {
	id = snowf.Generate().Int64()
	fmt.Println("id:", id)
	return
}

func getId() (id snowflake.ID) {
	id = snowf.Generate()
	return
}

func main() {
	id := getId()
	// Print out the ID in a few different ways.
	fmt.Printf("Int64  ID: %d\n", id)
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())

	// Print out the ID's timestamp
	fmt.Printf("ID Time  : %d\n", id.Time())

	// Print out the ID's node number
	fmt.Printf("ID Node  : %d\n", id.Node())

	// Print out the ID's sequence number
	fmt.Printf("ID Step  : %d\n", id.Step())

	// Generate and print, all in one.
	fmt.Printf("ID       : %d\n", snowf.Generate().Int64())
}
