package tasks

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"commitr/misc"
)

func justAsk(q string, reader *bufio.Reader) string {
	fmt.Print(q)
	s, _ := reader.ReadString('\n')

	return s
}

// AskRequired reads from command line and answer is required
func AskRequired(q string, reader *bufio.Reader) string {
	s := misc.TrimWhitespace(justAsk(q, reader))

	if s == "" {
		fmt.Println("This field can't be blank. Please try again.")
		AskRequired(q, reader)
	}

	return s
}

// Ask reads from command line
func Ask(q string, reader *bufio.Reader) string {
	return misc.TrimWhitespace(justAsk(q, reader))
}

// ExecuteSync synchronously executes command
func ExecuteSync(cmd string) {
	fmt.Printf("\nRunning command: %v\n", cmd)

	if strings.HasPrefix(cmd, "cd") {
		err := os.Chdir(cmd[3:len(cmd)])
		misc.HandleFatalError(err, "Change dir error")
		return
	}

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

// LoadCommands loads a list of OS commands from file
func LoadCommands(fs, m, c string) []string {
	re := strings.NewReplacer("{{ message }}", m, "{{ comment }}", c)
	strs := strings.Split(re.Replace(fs), "\n")
	fmt.Println(strs)

	return misc.Filter(strs, misc.RemoveEmpty)
}
