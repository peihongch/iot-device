package pkg

import "fmt"

func SprintRed(str string) string {
	return fmt.Sprintf("\\033[1;31;40m%s\\033[0m\n", str)
}

func SprintBlue(str string) string {
	return fmt.Sprintf("\\033[1;34;40m%s\\033[0m\n", str)
}

func SprintGreen(str string) string {
	return fmt.Sprintf("\\033[1;32;40m%s\\033[0m\n", str)
}

func SprintCyan(str string) string {
	return fmt.Sprintf("\\033[1;36;40m%s\\033[0m\n", str)
}
