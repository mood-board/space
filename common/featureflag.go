package common

import "fmt"

var ErrFeatureDisabled = fmt.Errorf("feature disabled")

var flags = make(map[string]bool)

func EnableFlag(flag string) {
	flags[flag] = true
}

func DisableFlag(flag string) {
	flags[flag] = false
}

func IsFlagEnabled(flag string) bool {
	enabled, present := flags[flag]
	return present && enabled
}
