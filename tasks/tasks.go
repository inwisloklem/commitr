package tasks

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"

	"commitr/misc"
)

// Ask reads from command line
func Ask(q string, reader *bufio.Reader, required bool) string {
	fmt.Print(q)

	s, _ := reader.ReadString('\n')
	res := regexp.MustCompile(`\r|\n|^\s+|\s+$`).ReplaceAllString(s, "")

	if required == true && res == "" {
		fmt.Println("This field can't be blank. Please try again.")
		Ask(q, reader, required)
	}

	return res
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

// LoadCommands loads a list of OS commands from file
func LoadCommands(fs, m, c string) []string {
	re := strings.NewReplacer(`{{\s*message\s*}}`, m, `{{\s*commit\s*}}`, c)
	strs := strings.Split(re.Replace(fs), "\n")

	return misc.Filter(strs, misc.RemoveEmpty)
}
