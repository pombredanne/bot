package commands

import (
	"strings"
	"regexp"
	"log"
)

// Command is a struct which separates the user's input for easier handling of commands
type Command struct {
	Raw     		string   	// Raw is full string passed to the command
	IsCommand 	bool   		// Confirmation if this is a command or just a regular message
	Message 		string   	// Full string without the prefix
	Command 		string   	// Command is the first argument passed to the bot
	FullArg 		string 		// Full argument as a single string
	Args    		[]string 	// Arguments as array
}

// Parse the arguments returning the Command to execute and the arguments passed to it
func Parse(c string, prefix string) *Command {
	cmd := &Command{Raw: c}
	c = strings.TrimSpace(c)
	cmd.IsCommand = strings.HasPrefix(c, prefix)

	// we can stop here if no prefix is detected
	if !cmd.IsCommand {
		cmd.Message = c
		return cmd
	}

	// Trim the prefix and extra spaces
	cmd.Message = strings.TrimPrefix(c, prefix)
	cmd.Message = strings.TrimSpace(cmd.Message)

	// check if we have the command and not only the prefix
	cmd.IsCommand = cmd.Message != ""
	if !cmd.IsCommand {
		return cmd
	}

	// get the command
	pieces := strings.SplitN(cmd.Message, " ", 2)
	cmd.Command = pieces[0]

	// get the arguments and remove extra spaces
	if len(pieces) > 1 {
		reg, err := regexp.Compile("\\s+")
		if err != nil {
         log.Fatal(err)
      }
		cmd.FullArg = reg.ReplaceAllString(strings.TrimSpace(pieces[1]), " ")
		cmd.Args = strings.Split(cmd.FullArg, " ")
	}

	return cmd
}