package cmd

import (
	"encoding/json"
	"io"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func ExitOnError(cmd *cobra.Command, err error) {
	if err == nil {
		return
	}

	_, _ = io.WriteString(cmd.ErrOrStderr(), err.Error()+"\n")
	_ = cmd.Usage()
	os.Exit(1)
}

func WriteString(cmd *cobra.Command, text string) {
	_, _ = io.WriteString(cmd.OutOrStdout(), text+"\n")
}

func WriteAsYAML(cmd *cobra.Command, obj interface{}, writer io.Writer) {
	encoder := yaml.NewEncoder(writer)
	defer encoder.Close()

	encoder.SetIndent(2)
	err := encoder.Encode(obj)
	ExitOnError(cmd, err)
}

func WriteAsJSON(cmd *cobra.Command, obj interface{}, writer io.Writer) {
	b, err := json.MarshalIndent(obj, "", "  ")
	ExitOnError(cmd, err)
	_, _ = writer.Write(b)
}