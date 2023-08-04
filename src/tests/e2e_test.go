package tests

import (
	"os/exec"
	"stellar/util"
	"testing"
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func TestMainFunction(t *testing.T) {
	log.Info("Building executable file...")
	buildCommand := exec.Command("go", "build", "main.go")
	buildCommand.Dir = "../"
	util.RunCommandAndLog(buildCommand)

	log.Info("Running binary...")
	mainCommand := exec.Command("./main", "-a", "887565851781", "-o", "latency-samples", "-g", "endpoints", "-c", "../experiments/tests/aws/small-burst.json")
	mainCommand.Dir = "../"
	stdout, _ := mainCommand.StdoutPipe()
	util.RunCommandAndLog(mainCommand)

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())	
	}
}
