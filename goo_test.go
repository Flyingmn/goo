package goo_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/Flyingmn/goo"
)

func TestDurationToChinese(t *testing.T) {
	fmt.Println(goo.DurationToChinese(2400*time.Hour + 0*time.Minute + 60*time.Second))
}
