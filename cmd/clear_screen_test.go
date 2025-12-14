package cmd

import (
	"testing"
)

func TestClearScreen(t *testing.T) {
	// Just call the function to get coverage
	// We don't check the output since it's a side-effect function
	ClearScreen()
	t.Log("ClearScreen executed")
}
