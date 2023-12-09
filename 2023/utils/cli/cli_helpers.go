package cli

import (
	"flag"
)

func GetArgs() map[string]any {
	argsMap := make(map[string]any)

	loggerEnabled := flag.Bool("debug", false, "Flag to enable logs for debugging")
	flag.Parse()
	argsMap["loggerEnabled"] = *loggerEnabled

	return argsMap
}
