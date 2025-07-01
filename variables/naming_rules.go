package variables

import (
	"fmt"
)

// Multi Word Variable Naming Rules
// Naming rules for multi-word variables in Go:

func MultiWordVariableNamingRules() {
	// 1. Use CamelCase for multi-word variable names.
	var multiWordVariableName string = "This is a multi-word variable name"

	// 2. Avoid underscores in variable names.
	var multi_word_variable_name string = "This is not recommended"

	// 3. Start with a lowercase letter for package-level variables.
	var packageLevelVariable string = "Package level variable"

	// 4. Start with an uppercase letter for exported variables.
	var ExportedVariable string = "Exported variable"

	// Print the variables
	fmt.Println(multiWordVariableName)
	fmt.Println(multi_word_variable_name)
	fmt.Println(packageLevelVariable)
	fmt.Println(ExportedVariable)
}
