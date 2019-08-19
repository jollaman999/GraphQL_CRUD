package root_util

import (
	"fmt"
	"os"
)

func Check_root() bool {
	if os.Geteuid() != 0 {
		fmt.Println("Please run as root!")
		return false
	}

	return true
}
