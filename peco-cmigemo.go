package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	args := []string{"cmigemo"}

	// dict is same location as executable
	// for example, you should put this executable into ~/.peco/
	p, err := exec.LookPath(os.Args[0])
	if err == nil {
		p, err = filepath.Abs(p)
		if err == nil {
			args = append(args, "-d", filepath.Join(filepath.Dir(p), "dict", "migemo-dict"))
		}
	}

	args = append(args, "-w", os.Args[1])

	cmd := exec.Command(args[0], args[1:]...)
	b, err := cmd.Output()
	if err != nil {
		os.Exit(1)
	}
	re, err := regexp.Compile(strings.TrimSpace(string(b)))
	if err != nil {
		os.Exit(1)
	}
	buf := bufio.NewReader(os.Stdin)
	for {
		b, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		line := string(b)
		if re.MatchString(line) {
			fmt.Println(line)
		}
	}
}
