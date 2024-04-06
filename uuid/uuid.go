package uuid

import (
	"fmt"
	"time"
)

const formatStr = "20060102150405"

const IDLength = len(formatStr)

func New() string {
	return fmt.Sprintf("%s", time.Now().Format(formatStr))
}
