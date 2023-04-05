package main

import (
	"fmt"
	"sort"

	"github.com/lareza-farhan-wanaghi/gophercises/secret"
	"github.com/spf13/cobra"
)

// init adds other commands to the root command. It is called by the cobra library before any other command functions
func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(listCmd)

	rootCmd.PersistentFlags().StringP("key", "k", "6368616e676520746869732070617373", "specifies the encoding key")
	rootCmd.PersistentFlags().StringP("filepath", "f", "./test", "specifies the filepath")
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
		encodingkey, err := cob.Flags().GetString("key")
		if err != nil {
			panic(err)
		}

		filepath, err := cob.Flags().GetString("filepath")
		if err != nil {
			panic(err)
		}

		fileFault := secret.FileFault(encodingkey, filepath)
		data, err := fileFault.Get(args[0])
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", data)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all data from the encrypted storage in the json format",
	Run: func(cob *cobra.Command, args []string) {
		encodingkey, err := cob.Flags().GetString("key")
		if err != nil {
			panic(err)
		}

		filepath, err := cob.Flags().GetString("filepath")
		if err != nil {
			panic(err)
		}

		fileFault := secret.FileFault(encodingkey, filepath)
		data, err := fileFault.GetAll()
		if err != nil {
			panic(err)
		}

		if len(data) <= 0 {
			fmt.Printf("You have no data in the storage\n")
			return
		}

		keys := []string{}
		for k, _ := range data {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

		fmt.Printf("You have the following key-value pairs:\n")
		for i, k := range keys {
			fmt.Printf("%d. key: \"%s\", value: \"%s\"\n", i+1, k, data[k])
		}
	},
}

var setCmd = &cobra.Command{
	Use:   "set",
	Args:  cobra.ExactArgs(2),
	Short: "Set the value of the given keyMap in the encrypted data",
	Run: func(cob *cobra.Command, args []string) {
		encodingkey, _ := cob.Flags().GetString("key")
		filepath, err := cob.Flags().GetString("filepath")
		if err != nil {
			panic(err)
		}

		fileFault := secret.FileFault(encodingkey, filepath)
		err = fileFault.Set(args[0], args[1])
		if err != nil {
			panic(err)
		}

		fmt.Printf("Successfully associate the key '%s' with a value of '%s'\n", args[0], args[1])
	},
}
