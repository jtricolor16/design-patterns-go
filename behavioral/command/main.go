package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("command")
	rootCmd := cobra.Command{}
	rootCmd.AddCommand(Volume())
	rootCmd.AddCommand(Turn())
	rootCmd.Execute()
}

// sender
func Volume() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "volume [value]",
		Short: "Change volume value",
		Run: func(cmd *cobra.Command, args []string) {
			value, _ := cmd.Flags().GetBool("value")
			commandVolume := NewCommandVolume(value)
			commandVolume.Execute()
		},
	}
	cmd.Flags().BoolP("value", "v", false, "value bool")
	return cmd
}

// sender
func Turn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "turn [value]",
		Short: "Turn on or turn off the TV",
		Run: func(cmd *cobra.Command, args []string) {
			value, _ := cmd.Flags().GetBool("value")
			commandTurn := NewCommandTurn(value)
			commandTurn.Execute()
		},
	}
	cmd.Flags().BoolP("value", "v", false, "value bool")
	return cmd
}
