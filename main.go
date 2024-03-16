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

	// Current DB:
	// 	Range:
	// 	Low: ''
	// 	High: 'Inf'
	// 	Size: 500

	// Now, let's insert again to exceed the range
	for i := 501; i <= 700; i++ {
		v := fmt.Sprintf("%05d", i)
		// Just set itself in
		db.Set(v, v)
	}

	// Debug it
	fmt.Printf("Current DB:\n%s\n", db.DebugRanges())

	// Current DB:
	// 	Range:
	// 	Low: '00300'
	// 	High: 'Inf'
	// 	Size: 400

	// 	Range:
	// 	Low: ''
	// 	High: '00300'
	// 	Size: 300

	// Insert a ton to cause lots of range splits all over
	for i := 700; i < 2000; i++ {
		v := fmt.Sprintf("%05d", i)
		db.Set(v, v)
	}

	// Debug it
	fmt.Printf("Current DB:\n%s\n", db.DebugRanges())

	// Current DB:
	// 	Range:
	// 	Low: '01500'
	// 	High: 'Inf'
	// 	Size: 499

	// 	Range:
	// 	Low: ''
	// 	High: '00300'
	// 	Size: 300

	// 	Range:
	// 	Low: '00300'
	// 	High: '00600'
	// 	Size: 300

	// 	Range:
	// 	Low: '00600'
	// 	High: '00900'
	// 	Size: 300

	// 	Range:
	// 	Low: '00900'
	// 	High: '01200'
	// 	Size: 300

	// 	Range:
	// 	Low: '01200'
	// 	High: '01500'
	// 	Size: 300

	// Get some data just in case
	fmt.Println(*db.Get("01000"))
	// Getting key 01000 from range with low 00900
	// 01000
}
