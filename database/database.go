package database

import (
	"fmt"
	"sort"
	"strings"
)

type Database struct {
	Ranges       []*Range
	maxRangeSize int
}

func NewDB(maxRangeSize int) Database {
	return Database{
		Ranges: []*Range{{
			Low: "", // We'll assume the low is always included

			High: Inf, // We'll assume high never includes
			KV:   map[string]string{},
		}},
		maxRangeSize: maxRangeSize,
	}
}

func (d *Database) Get(key string) *string {
	// range range range... lol
	for _, rng := range d.Ranges {
		if rng.OwnsKey(key) {
			fmt.Println("Getting key", key, "from range with low", rng.Low)
			return rng.Get(key)
		}
	}
	return nil
}

func (d *Database) Set(key, value string) {
	for _, rng := range d.Ranges {
		if rng.OwnsKey(key) {
			// Insert the value
			rng.Set(key, value)

			// Check if the range is too large
			if len(rng.KV) > d.maxRangeSize {
				fmt.Println("\n### Range too large, splitting! ###\n")
				d.SplitRange(rng)
			}
			return
		}
	}
}

func (d *Database) SplitRange(rng *Range) {
	// Let's just split it ~half-way... rip performance
	var keys []string
	for k, _ := range rng.KV {
		keys = append(keys, k)
	}

	// sort that boi (rip performance)
	sort.Strings(keys)

	// We only need to make one new range, and delete from the old range
	newRangeKeys := keys[:d.maxRangeSize/2]
	newRangeKV := map[string]string{}
	for _, k := range newRangeKeys {
		newRangeKV[k] = *rng.Get(k)
		rng.Delete(k)
	}

	newRange := Range{
		Low: rng.Low,

		High: newRangeKeys[(d.maxRangeSize/2)-1],
		KV:   newRangeKV,
	}

	// Update the old range low
	rng.Low = newRangeKeys[(d.maxRangeSize/2)-1]

	// Add the new range to our DB
	d.Ranges = append(d.Ranges, &newRange)
}

func (d *Database) Delete(key string) {
	for _, rng := range d.Ranges {
		if rng.OwnsKey(key) {
			rng.Delete(key)
			return
		}
	}
}

// DebugRanges is a pretty print for the ranges
func (d *Database) DebugRanges() string {
	str := strings.Builder{}
	for _, rng := range d.Ranges {
		str.WriteString("Range:\n")
		str.WriteString(fmt.Sprintf("\tLow: '%s'\n", rng.Low))
		if rng.High == Inf {
			str.WriteString("\tHigh: 'Inf'\n")
		} else {
			str.WriteString(fmt.Sprintf("\tHigh: '%s'\n", rng.High))
		}
		str.WriteString(fmt.Sprintf("\tSize: %d\n", len(rng.KV)))
	}
	return str.String()
}
