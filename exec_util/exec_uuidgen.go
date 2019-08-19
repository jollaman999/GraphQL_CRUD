package exec_util

import (
	"../logger"
	"os/exec"
)

func Exec_uuidgen() (string, error) {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		logger.Logger.Println(err)
		return "", err
	}

	return string(out), nil
}