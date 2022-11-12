package main

import (
	"bufio"
	"flag"
	"fmt"
	"migrator-generator/command"
	"migrator-generator/menu"
	"os"
	"os/exec"
	"strings"
)

const DefaultMenu = true
const DefaultProject = "Users"
const DefaultAction = "migrations add"
const DefaultVersion = "0_1_0"
const DefaultSolution = "Digidoc"

func endProgram() {
	fmt.Printf("Write exit to exit\n")
	for {
		consoleReader := bufio.NewReader(os.Stdin)
		fmt.Print(">")

		input, _ := consoleReader.ReadString('\n')

		input = strings.ToLower(input)

		if strings.HasPrefix(input, "exit") {
			fmt.Println("Good bye!")
			os.Exit(0)
		}
	}
}

func printSuccess() {
	fmt.Printf("" +
		"⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⠄⣀⣠⣤⣤⣤⣤⣤⣄⣀⡀⠄⠄⠄⠄⠄⠄⠄⠄\n" +
		"⠄⠄⠄⠄⠄⠄⠄⢀⣤⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣤⡀⠄⠄⠄⠄⠄\n" +
		"⠄⠄⠄⠄⠄⢀⣴⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢿⣿⣿⣿⣿⣆⠄⠄⠄⠄\n" +
		"⠄⠄⠄⠄⢠⣿⣿⣿⣿⣿⢻⣿⣿⣿⣿⣿⣿⣿⣿⣯⢻⣿⣿⣿⣿⣆⠄⠄⠄\n" +
		"⠄⠄⣼⢀⣿⣿⣿⣿⣏⡏⠄⠹⣿⣿⣿⣿⣿⣿⣿⣿⣧⢻⣿⣿⣿⣿⡆⠄⠄\n" +
		"⠄⠄⡟⣼⣿⣿⣿⣿⣿⠄⠄⠄⠈⠻⣿⣿⣿⣿⣿⣿⣿⣇⢻⣿⣿⣿⣿⠄⠄\n" +
		"⠄⢰⠃⣿⣿⠿⣿⣿⣿⠄⠄⠄⠄⠄⠄⠙⠿⣿⣿⣿⣿⣿⠄⢿⣿⣿⣿⡄⠄\n" +
		"⠄⢸⢠⣿⣿⣧⡙⣿⣿⡆⠄⠄⠄⠄⠄⠄⠄⠈⠛⢿⣿⣿⡇⠸⣿⡿⣸⡇⠄\n" +
		"⠄⠈⡆⣿⣿⣿⣿⣦⡙⠳⠄⠄⠄⠄⠄⠄⢀⣠⣤⣀⣈⠙⠃⠄⠿⢇⣿⡇⠄\n" +
		"⠄⠄⡇⢿⣿⣿⣿⣿⡇⠄⠄⠄⠄⠄⣠⣶⣿⣿⣿⣿⣿⣿⣷⣆⡀⣼⣿⡇⠄\n" +
		"⠄⠄⢹⡘⣿⣿⣿⢿⣷⡀⠄⢀⣴⣾⣟⠉⠉⠉⠉⣽⣿⣿⣿⣿⠇⢹⣿⠃⠄\n" +
		"⠄⠄⠄⢷⡘⢿⣿⣎⢻⣷⠰⣿⣿⣿⣿⣦⣀⣀⣴⣿⣿⣿⠟⢫⡾⢸⡟⠄⠄\n" +
		"⠄⠄⠄⠄⠻⣦⡙⠿⣧⠙⢷⠙⠻⠿⢿⡿⠿⠿⠛⠋⠉⠄⠂⠘⠁⠞⠄⠄⠄\n" +
		"⠄⠄⠄⠄⠄⠈⠙⠑⣠⣤⣴⡖⠄⠿⣋⣉⣉⡁⠄⢾⣦⠄⠄⠄⠄⠄⠄⠄⠄\n" +
		"⠄⠄⠄⠄⠄⠄⠄⠄⠛⠛⠋⠁⣠⣾⣿⣿⣿⣿⡆⠄⣿⠆⠄⠄⠄⠄⠄⠄⠄\n" +
		"⠄⠄⠄⠄Acción realizada!⠄⠄⠄⠄",
	)
	endProgram()
}

func executeMigrationBuilder(executionString string) error {
	fmt.Printf("Comando a ejecutar:\n %s \n", executionString)
	cmd := exec.Command(executionString)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var comm command.ExecutionBuilder
	var err error
	enableMenu := flag.Bool("enableMenu", DefaultMenu, "Use Menu")
	action := flag.String("action", DefaultAction, "Project Name")
	project := flag.String("project", DefaultProject, "Project Name")
	version := flag.String("version", DefaultVersion, "Project Version")
	solution := flag.String("solution", DefaultSolution, "Solution Name")

	flag.Parse()

	if *enableMenu == true {
		comm, err = menu.PrintAndGet()
	} else {
		comm = command.NewCommand(*action, *version, *project, *solution)
	}

	if err != nil {
		print(err.Error())
		endProgram()
		return
	}

	execErr := executeMigrationBuilder(comm.BuildExecutionString())

	if execErr != nil {
		fmt.Printf(execErr.Error() + "\n")
		endProgram()
	} else {
		printSuccess()
	}
}
