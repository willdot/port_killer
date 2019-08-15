package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var err error

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter port number you wish to kill:")

	port, err := readFromUserAndTrim(reader)

	if err != nil {
		fmt.Println(err)
		return
	}

	port = ":" + port

	if err != nil {
		fmt.Println(err)
		return
	}

	pid, err := getPid(port)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Check the PID is the correct part of the process details by trying to convert it to int
	if _, err := strconv.Atoi(pid); err != nil {
		fmt.Println("Problem getting the PID of the process. Exiting.")
		return
	}

	fmt.Printf("Are you sure you wish to kill process with PID: %v? Y or N\n", pid)

	confirm, err := readFromUserAndTrim(reader)

	if err != nil {
		fmt.Println(err)
		return
	}

	if confirm != "Y" {
		fmt.Println("Exiting without killing port")
		return
	}

	_, err = killProcess(pid)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func killProcess(pid string) ([]byte, error) {
	args := []string{
		pid,
	}

	return exec.Command("kill", args...).CombinedOutput()
}

func getPid(port string) (string, error) {
	args := []string{
		"-i", port,
	}

	procssDetails, err := exec.Command("lsof", args...).CombinedOutput()

	if err != nil {
		return "", err
	}

	processString := string(procssDetails)
	// Print the process details to the user so they can see what the result is before they confirm the kill
	fmt.Println("\n" + processString)

	return strings.Fields(processString)[10], nil
}

func readFromUserAndTrim(reader *bufio.Reader) (string, error) {

	result, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	result = strings.TrimRight(result, "\n")

	return result, nil
}
