package progress

import "fmt"

// Progress is a struct to describe
// the progress bar
type Progress struct {
	percent int
	cur     int
	total   int
	rate    string
	graph   string
}

// New initializes a new instance of Progress
func (bar *Progress) New(start, total int) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = "█"
	}
	bar.percent = bar.getPercent()
}

func (bar *Progress) getPercent() int {
	return int(float32(bar.cur) / float32(bar.total) * 100)
}

// Play increments the progress bar
func (bar *Progress) Play(cur int) string {
	bar.cur = cur
	bar.rate = ""
	bar.percent = bar.getPercent()
	for i := 0; i < bar.percent; i+=2 {
		bar.rate += bar.graph
	}
	return fmt.Sprintf("\r[%-50s]%3d%%", bar.rate, bar.percent)
}

// Finish ends printing the bar graph
func (bar *Progress) Finish() {
	fmt.Println()
}