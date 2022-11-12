package menu

import (
	"errors"
	"fmt"
	"migrator-generator/command"
)

func PrintAndGet() (command.ExecutionBuilder, error) {
	fmt.Printf("Welcome to Entity Framework command Helper!\n")
	action, err := printAndGetAction()

	if err != nil {
		return nil, err
	}

	version, errVer := printAndGetString("version")
	if errVer != nil {
		return nil, err
	}

	project, errProj := printAndGetString("project")
	if errProj != nil {
		return nil, err
	}

	solution, errSol := printAndGetString("solution")
	if errSol != nil {
		return nil, err
	}

	return command.NewCommand(action, version, project, solution), nil
}

func actionDetector(actionNumber int) (string, error) {
	switch actionNumber {
	case 1:
		return "add migration", nil
	case 2:
		return "update database", nil
	default:
		return "", errors.New("wrong option")
	}
}

func printAndGetAction() (string, error) {
	var action int
	fmt.Printf("Chose Action.\n1) Add Migration\n2) Update Databases\n")
	_, err := fmt.Scan(&action)
	if err != nil {
		return "", err
	}

	return actionDetector(action)
}

func printAndGetString(stringName string) (string, error) {
	var stringValue string
	fmt.Printf("Write %s.\n", stringName)
	_, err := fmt.Scan(&stringValue)
	if err != nil {
		return "", err
	}

	return stringValue, err
}
