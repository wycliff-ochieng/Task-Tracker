/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wycliff-ochieng/task-tracker/internal/task"
)

// listCmd represents the list command
func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list the tasks of the user",
		Long:  `List tasks according to either ID or status or rather list all tasks when no specification is provided`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListTaskCmd(args)
		},
	}
	return cmd
}

func RunListTaskCmd(args []string) {
	if len(args) < 0 {
		status := task.TaskStatus(args[0])
		return task.ListTask(status)
	}
	return task.ListTask("all")
}
