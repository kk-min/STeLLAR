package tests

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"testing"
)

func TestMainFunction(t *testing.T) {
	log.Info("Building executable file...")
	buildCommand := exec.Command("go", "build", "main.go")
	buildCommand.Dir = "../"
	err := buildCommand.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Wait() returns nil if the command successfully executes and exits with status code 0
	err = buildCommand.Wait()
	if err != nil {
		log.Fatal(err)
	}

	mainCommand := exec.Command("./main", "-a", "887565851781", "-o", "latency-samples", "-g", "endpoints", "-c", "../experiments/tests/aws/small-burst.json")
	mainCommand.Dir = "../"

	// Create scanner and scan stdout in a goroutine to get real-time printing
	stdout, _ := mainCommand.StdoutPipe()
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	go func() {
		for scanner.Scan() {
			log.Info(scanner.Text())
		}
	}()

	log.Info("Running binary...")
	err = mainCommand.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = mainCommand.Wait()
	if err != nil {
		log.Fatal(err)
	}

}
