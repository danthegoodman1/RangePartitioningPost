package main

import (
	"fmt"
	"github.com/danthegoodman1/RangePartitioningPost/database"
)

func main() {
	fmt.Println("Running range splitting!")

	// First let's insert 500 records and see the DB

	db := database.NewDB(600) // 600 record limit per range
	for i := 1; i <= 500; i++ {
		v := fmt.Sprintf("%05d", i)
		// Just set itself in
		db.Set(v, v)
	}

	// Debug it
	fmt.Printf("Current DB:\n%s\n", db.DebugRanges())

	// Now, let's insert again to exceed the range
	for i := 501; i <= 700; i++ {
		v := fmt.Sprintf("%05d", i)
		// Just set itself in
		db.Set(v, v)
	}

	// Debug it
	fmt.Printf("Current DB:\n%s\n", db.DebugRanges())

	// Insert a ton to cause lots of range splits all over
	for i := 700; i < 2000; i++ {
		v := fmt.Sprintf("%05d", i)
		db.Set(v, v)
	}

	// Debug it
	fmt.Printf("Current DB:\n%s\n", db.DebugRanges())

	// Get some data just in case
	fmt.Println(*db.Get("01000"))

	// Let's make a crazy one
	mangleDB := database.NewDB(10_000) // o7 godspeed memory
	for i := 0; i < 50_000; i++ {
		mangleDB.Set(fmt.Sprintf("%d", i), "e")
	}
	fmt.Println(mangleDB.DebugRanges())
	fmt.Println(*mangleDB.Get("42000"))
	fmt.Println(*mangleDB.Get("20000"))
}
