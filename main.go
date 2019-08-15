package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter port number you wish to kill:")
	port, _ := reader.ReadString('\n')

	port = ":" + strings.TrimRight(port, " \n")

	args := []string{
		"-i", port,
	}

	output, err := exec.Command("lsof", args...).CombinedOutput()

	if err != nil {
		fmt.Println(err)
		return
	}

	result := string(output)
	fmt.Println("\n" + result)
	pid := strings.Fields(result)[10]

	fmt.Printf("Are you sure you wish to kill process with PID: %v? Y or N\n", pid)

	confirm, _ := reader.ReadString('\n')

	if strings.TrimRight(confirm, " \n") != "Y" {
		return
	}

	if _, err := strconv.Atoi(pid); err != nil {
		fmt.Println("Problem getting the PID of the process. Exiting.")
		return
	}

	args = []string{
		pid,
	}

	_, err = exec.Command("kill", args...).CombinedOutput()

	if err != nil {
		fmt.Println(err)
		return
	}
}
