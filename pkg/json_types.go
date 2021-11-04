package pkg

import (
	"fmt"
	"strings"
	"time"
)

var (
	shortMonthToNumber = map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}
)

type JSONNode struct {
	Executors    uint64 `json:"executors,omitempty"`
	JVMName      string `json:"jvm-name,omitempty"`
	JVMVendor    string `json:"jvm-vendor,omitempty"`
	JVMVersion   string `json:"jvm-version,omitempty"`
	IsController bool   `json:"master"`
	OS           string `json:"os,omitempty"`
}

type JSONPlugin struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type JSONReport struct {
	Install          string            `json:"install"`
	Jobs             map[string]uint64 `json:"jobs"`
	Nodes            []JSONNode        `json:"nodes"`
	Plugins          []JSONPlugin      `json:"plugins"`
	ServletContainer string            `json:"servletContainer,omitempty"`
	TimestampString  string            `json:"timestamp"`
	Version          string            `json:"version"`
}

func (j *JSONReport) Timestamp() (time.Time, error) {
	return time.Parse(time.RFC3339, JSONTimestampToRFC3339(j.TimestampString))
}

func JSONTimestampToRFC3339(ts string) string {
	withoutZone := strings.TrimSuffix(ts, " +0000")
	splitDateAndTime := strings.SplitN(withoutZone, ":", 2)
	dayMonthYear := strings.Split(splitDateAndTime[0], "/")
	return fmt.Sprintf("%s-%s-%sT%sZ", dayMonthYear[2], shortMonthToNumber[dayMonthYear[1]], dayMonthYear[0], splitDateAndTime[1])
}