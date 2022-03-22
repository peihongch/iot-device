package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	data   = ""
	name   = ""
	remote = ""
	topic  = ""
)

var rootCmd = &cobra.Command{
	Use: "iot-device",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&data, "data", "", "device data file (e.g. ./data.csv)")
	rootCmd.PersistentFlags().StringVar(&name, "name", "", "device name")
}
