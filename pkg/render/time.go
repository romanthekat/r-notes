package render

import "time"

func FormatIdAsIsoDate(zkId string) string {
	date, err := time.Parse("200601021504", zkId)
	if err != nil {
		panic(err)
	}

	return date.Format("2006-01-02 15:04")
}
