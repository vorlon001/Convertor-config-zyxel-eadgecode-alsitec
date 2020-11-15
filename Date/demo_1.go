package main

import (

        "fmt"
        "time"
        "log"
)


func rangeDate(start, end time.Time) func() time.Time {
        y, m, d := start.Date()
        start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
        y, m, d = end.Date()
        end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

        return func() time.Time {
                if start.After(end) {
                        return time.Time{}
                }
                date := start
                start = start.AddDate(0, 0, 1)
                return date
        }
}

func main() {

	interval  := 10
	end       := time.Now()
	end        = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, time.Local)
	start := end.AddDate(0, 0, -interval )
	log.Println(start.Format("2006-01-02"), "-", end.Format("2006-01-02"))

	for rd := rangeDate(start, end); ; {
		now := rd()
		date := fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
		if now.IsZero() {
			break
		}
		
		fmt.Printf(" %v \n",date)
	}
}


