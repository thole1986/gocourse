package intermediate

import (
	"fmt"
	"time"
)

func main() {
	// 00:00:00 UTC on Jan 1, 1970

	now := time.Now()

	unixTime := now.Unix()

	// Return a big number
	fmt.Println("Current Unix Time:", unixTime)
	// Convert to human readable format.
	t := time.Unix(unixTime, 0)
	fmt.Println(t)
	fmt.Println("Time: ", t.Format("2006-01-02"))

}
