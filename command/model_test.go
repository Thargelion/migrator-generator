package command

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestCommandBuildExecutionString(t *testing.T) {
	// Given
	expected := "dotnet ef lele users_0_1_0 --context usersContext --project Users.Internal/Users.Internal.csproj" +
		" --startup-project Digidoc"
	comm := NewCommand("lele", "0_1_0", "Users", "Digidoc")
	// When
	actual := comm.BuildExecutionString()
	// Then
	assert.Equal(t, actual, expected)
}
