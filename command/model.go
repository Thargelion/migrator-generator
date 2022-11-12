package command

import (
	"fmt"
	"strings"
)

type ExecutionBuilder interface {
	BuildExecutionString() string
}

type command struct {
	Action   string
	Name     string
	Project  string
	Context  string
	Solution string
}

func (comm *command) BuildExecutionString() string {
	return fmt.Sprintf("dotnet ef %s %s "+
		"--context %s "+
		"--project %s.Internal/%s.Internal.csproj "+
		"--startup-project %s",
		comm.Action,
		comm.Name,
		comm.Context,
		comm.Project,
		comm.Project,
		comm.Solution,
	)
}

func buildMigrationName(project string, version string) string {
	return strings.ToLower(project) + "_" + strings.ToLower(version)
}

func NewCommand(action string, version string, project string, solution string) ExecutionBuilder {
	return &command{
		Action:   action,
		Name:     buildMigrationName(project, version),
		Project:  project,
		Solution: solution,
		Context:  strings.ToLower(project) + "Context",
	}
}
