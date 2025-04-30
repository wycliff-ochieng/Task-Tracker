/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/wycliff-ochieng/task-tracker/internal/task"
)

// addCmd represents the add command
func NewAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "adding tasks to task list",
		Long:  "command for adding task to your tracker",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddTaskCmd(args)
		},
	}
	return cmd
}

func RunAddTaskCmd(args []string) error {
	if len(args) == 0 {
		return errors.New("Kindly input a description")
	}
	description := RunAddTaskCmd(args)
	return task.AddTask(description)
}
