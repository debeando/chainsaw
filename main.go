package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
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
	c := Chunk{Count: 500, Delta: 50, Length: 100, Start: 1, Sleep: 1 * time.Second}
	c.Loop(func(){
		// time.Sleep(1 * time.Second)
		c.Log()
	})
}

func (c *Chunk) SetTotal() {
	c.Total = c.Count + c.Delta
}

func (c *Chunk) SetSteps() {
	c.Steps = DivisionAndRound(c.Total, c.Length)
}

func (c *Chunk) SetProgress() {
	c.Percentage = ((100 * c.Index) / c.Total)
}

func (c *Chunk) SetRemain() {
	c.Remain = DivisionAndRound(c.Total - c.Index, c.Length)
}

func (c *Chunk) SetEnd() {
	c.End = c.Index
}

func (c *Chunk) SetStart() {
	c.Start = c.Index - (c.Length - 1)

	if c.Index >= c.Total {
		c.Start = c.Count + 1
	}
}

func (c *Chunk) SetDuration() {
	c.Duration = c.EndTime.Sub(c.StartTime)
}

func (c *Chunk) SetETA() {
	c.ETA = time.Duration(c.Remain) * c.Duration
}

func (c *Chunk) SetStartTime() {
	c.StartTime = time.Now()
}

func (c *Chunk) SetEndTime() {
	c.EndTime = time.Now()
}

func (c *Chunk) Wait() {
	time.Sleep(c.Sleep)
}

func (c *Chunk) SetIncrement() {
	c.Index = c.Index + c.Length

	if (c.Index + c.Length) > c.Total {
		c.Index = c.Total
	}
}

func (c *Chunk) Loop(f interface{}) {
	c.Index  = 0
	c.Start  = 0

	c.SetTotal()
	c.SetSteps()

	for c.Index = c.Start; c.Index < c.Total; {
		c.SetIncrement()
		c.SetProgress()
		c.SetRemain()
		c.SetEnd()
		c.SetStart()
		c.SetDuration()
		c.SetETA()
		c.SetStartTime()
		f.(func())()
		c.Wait()
		c.SetEndTime()
	}
}

func (c *Chunk) Log() {
	fmt.Printf("%*d %*d/%d [%*d-%*d] %3d%% %.fs\n",
		IntLength(c.Steps),
		c.Remain,
		IntLength(c.Total),
		c.Index,
		c.Total,
		IntLength(c.Total),
		c.Start,
		IntLength(c.Total),
		c.End,
		c.Percentage,
		c.ETA.Seconds(),
	)
}

func IntLength(v uint64) int {
	return len(Int64ToString(v))
}

func Int64ToString(v uint64) string {
	return strconv.FormatUint(v, 10)
}

func DivisionAndRound(x uint64, y uint64) uint64 {
	return uint64(math.Round(float64(x) / float64(y)))
}

// los ficheros a subir al S3 podrian tener este formato:
// ddbb_table_date_chunk
