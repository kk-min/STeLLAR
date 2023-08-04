package tests

import (
	"bufio"
	"fmt"
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

	err = buildCommand.Wait()
	if err != nil {
		log.Fatal(err)
	}

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
	err = mainCommand.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = mainCommand.Wait()
	if err != nil {
		log.Fatal(err)
	}

}
