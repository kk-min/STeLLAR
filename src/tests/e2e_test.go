package tests

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"stellar/util"
	"testing"
)

func TestMainFunction(t *testing.T) {
	log.Info("Building executable file...")
	buildCommand := exec.Command("go", "build", "main.go")
	buildCommand.Dir = "../"
	util.RunCommandAndLog(buildCommand)

	mainCommand := exec.Command("./main", "-a", "887565851781", "-o", "latency-samples", "-g", "endpoints", "-c", "../experiments/tests/aws/small-burst.json")
	mainCommand.Dir = "../"

	stdout, _ := mainCommand.StdoutPipe()
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	log.Info("Running binary...")
	util.RunCommandAndLog(mainCommand)

}
