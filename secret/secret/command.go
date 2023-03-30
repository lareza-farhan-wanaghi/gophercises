package main

import (
	"fmt"

	"github.com/lareza-farhan-wanaghi/gophercises/secret"
	"github.com/spf13/cobra"
)

// init adds other commands to the root command. It is called by the cobra library before any other command functions
func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(setCmd)

	getCmd.PersistentFlags().StringP("key", "k", "6368616e676520746869732070617373", "specifies the encoding key")
	getCmd.PersistentFlags().StringP("filepath", "f", "./test", "specifies the filepath")
	setCmd.PersistentFlags().StringP("key", "k", "6368616e676520746869732070617373", "specifies the encoding key")
	setCmd.PersistentFlags().StringP("filepath", "f", "./test", "specifies the filepath")
}

var rootCmd = &cobra.Command{
	Use:   "secret",
	Short: "secret is a simple ecripted key-value pair data storage",
	Long:  `A simple key-value pair data storage built with the golang's built-in encryption algorithm`,
}

var getCmd = &cobra.Command{
	Use:   "get",
	Args:  cobra.ExactArgs(1),
	Short: "Get the value corresponding to the key from the encrypted data storage",
	Run: func(cob *cobra.Command, args []string) {
		key, err := cob.Flags().GetString("key")
		if err != nil {
			panic(err)
		}

		filepath, err := cob.Flags().GetString("filepath")
		if err != nil {
			panic(err)
		}

		fileFault := secret.FileFault(key, filepath)
		data, err := fileFault.Get(args[0])
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", data)
	},
}

var setCmd = &cobra.Command{
	Use:   "set",
	Args:  cobra.ExactArgs(2),
	Short: "Set the value of the given keyMap in the encrypted data",
	Run: func(cob *cobra.Command, args []string) {
		key, _ := cob.Flags().GetString("key")
		filepath, err := cob.Flags().GetString("filepath")
		if err != nil {
			panic(err)
		}

		fileFault := secret.FileFault(key, filepath)
		err = fileFault.Set(args[0], args[1])
		if err != nil {
			panic(err)
		}

		fmt.Printf("Successfully set the key '%s' as '%s' to the encrypted data storage\n", args[0], args[1])
	},
}
