package tasks

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"

	"commitr/misc"
)

// Ask reads from command line
func Ask(q string, reader *bufio.Reader) string {
	fmt.Print(q)
	s, _ := reader.ReadString('\n')

	return regexp.MustCompile(`\r|\n|^\s+|\s+$`).ReplaceAllString(s, "")
}

// ExecuteSync synchronously executes command
func ExecuteSync(cmd string) {
	fmt.Printf("\nRunning command: %v\n", cmd)

	c := exec.Command("script", "-qfc", cmd, "/dev/null")

	reader, err := c.StdoutPipe()
	misc.HandleFatalError(err, "Create StdoutPipe error")

	scanner := bufio.NewScanner(reader)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	err = c.Start()
	misc.HandleFatalError(err, "Start command error")

	err = c.Wait()
	misc.HandleFatalError(err, "Wait command error")
}

// ExecCommands executes a list of OS commands
func ExecCommands(cmds []string) {
	for _, cmd := range cmds {
		ExecuteSync(cmd)
	}
}
