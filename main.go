package main

import (
	"fmt"
	"time"
	"strconv"
)

type Chunk struct {
	Count      uint64        // Rows in table.
	Delta      uint64        // Delta value.
	Total      uint64        // Count + delta.
	Index      uint64        // Number of chunks.
	Steps      uint64        // Numbers of chunks.
	Length     uint64        // Length of chunk.
	Percentage uint64        // Percentage overall process.
	Sleep      time.Duration // Wait between chunk and chunk.
	Start      uint64        // Start chunk number.
	End        uint64        // End chunk number. (Not use yet)
	Remain     uint64        // Remains chunks until finish.
	StartTime  time.Time     // Start each chunk.
	EndTime    time.Time     // End each chunk.
	Duration   time.Duration // Duration time between start and end duration.
	ETA        time.Duration // Estimated time of arrival to complete overall process.
}

func main() {
	c := Chunk{Count: 1000, Delta: 0, Length: 100, Start: 1, Sleep: 1 * time.Second}
	c.Loop(func(){
		time.Sleep(1 * time.Second)
	})
}

func (c *Chunk) Calculate() {
	c.Total      = c.Count + c.Delta
	c.Steps      = c.Total / c.Length
	c.Percentage = ((100 * c.Index) / c.Total)
	c.Duration   = c.EndTime.Sub(c.StartTime)
	c.Remain     = (c.Total - c.Index) / c.Length
	c.ETA        = time.Duration(c.Remain) * c.Duration
}

func (c *Chunk) Loop(f interface{}) {
	c.Calculate()
	for i := c.Start; i <= c.Total; i++ {
		if i % c.Length == 0 {
			c.StartTime = time.Now()

			f.(func())()

			time.Sleep(c.Sleep)

			c.EndTime = time.Now()
			c.Index   = i
			c.Calculate()
			c.Log()
		}
	}
}

func (c *Chunk) Log() {
	fmt.Printf("%*d %*d/%d %3d%% %.fs\n", IntLength(c.Steps - 1), c.Remain, IntLength(c.Total), c.Index, c.Total, c.Percentage, c.ETA.Seconds())
}

func IntLength(v uint64) int {
	return len(Int64ToString(v))
}

func Int64ToString(v uint64) string {
	return strconv.FormatUint(v, 10)
}
