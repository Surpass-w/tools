package cmd

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"tool/internal"
)

var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "use for render tpl file",
}

func init() {
	renderCmd.Flags().SortFlags = false
	renderCmd.Flags().StringVarP(&internal.RenderOptions.FromPath, "from", "f", "", "tpl path")
	renderCmd.Flags().StringVarP(&internal.RenderOptions.ToPath, "to", "t", "", "yml path")
	renderCmd.Flags().StringVarP(&internal.RenderOptions.MetaData, "data", "d", "", "render meta data")

	renderCmd.RunE = func(cmd *cobra.Command, args []string) error {
		tplPath := internal.RenderOptions.FromPath
		ymlPath := internal.RenderOptions.FromPath
		data := make(map[string]interface{})
		err := json.Unmarshal([]byte(internal.RenderOptions.MetaData), &data)
		if err != nil {
			return errors.New("json unmarshal data failed: " + err.Error())
		}
		return internal.RenderFile(tplPath, ymlPath, data)
	}
	rootCmd.AddCommand(renderCmd)
}
