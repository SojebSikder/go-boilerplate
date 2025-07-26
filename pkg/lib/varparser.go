package lib

import (
	"fmt"
	"regexp"
	"sync"
)

type VarParser struct {
	variables []map[string]interface{}
	mu        sync.RWMutex
}

var (
	instance         = &VarParser{}
	placeholderRegex = regexp.MustCompile(`\${\s*(.+?)\s*}`)
)

// Adds a new variable map to the variable store
func AddVariable(vars map[string]interface{}) {
	instance.mu.Lock()
	defer instance.mu.Unlock()
	instance.variables = append(instance.variables, vars)
}

// Returns a copy of the current variable list
func GetVariables() []map[string]interface{} {
	instance.mu.RLock()
	defer instance.mu.RUnlock()

	var copyVars []map[string]interface{}
	for _, v := range instance.variables {
		copy := make(map[string]interface{}, len(v))
		for key, val := range v {
			copy[key] = val
		}
		copyVars = append(copyVars, copy)
	}
	return copyVars
}

// Clears all stored variables
func ClearVariables() {
	instance.mu.Lock()
	defer instance.mu.Unlock()
	instance.variables = nil
}

// Replaces ${var} placeholders in text with their corresponding values
func Parse(text string) string {
	instance.mu.RLock()
	defer instance.mu.RUnlock()

	return placeholderRegex.ReplaceAllStringFunc(text, func(match string) string {

		submatches := placeholderRegex.FindStringSubmatch(match)
		if len(submatches) != 2 {
			return match
		}
		key := submatches[1]

		for _, variableSet := range instance.variables {
			if val, ok := variableSet[key]; ok {
				return fmt.Sprintf("%v", val)
			}
		}
		return match
	})
}
