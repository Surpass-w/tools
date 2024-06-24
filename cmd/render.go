package cmd

import (
	"encoding/json"
	"fmt"
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
	renderCmd.Flags().StringVarP(&internal.RenderOptions.ToPath, "to", "t", "", "yml path")
	renderCmd.Flags().StringVarP(&internal.RenderOptions.MetaData, "data", "d", "", "render meta data")

	renderCmd.RunE = func(cmd *cobra.Command, args []string) error {
		tplPath, _ := cmd.Flags().GetString("file")
		ymlPath := internal.RenderOptions.ToPath
		data := make(map[string]interface{})
		err := json.Unmarshal([]byte(internal.RenderOptions.MetaData), &data)
		if err != nil {
			return errors.New("json unmarshal data failed: " + err.Error())
		}
		iFDebug, _ := cmd.Flags().GetBool("debug")
		if iFDebug {
			params := make(map[string]interface{})
			params["from"] = tplPath
			params["to"] = ymlPath
			params["data"] = data
			info, _ := json.Marshal(params)
			fmt.Println(string(info))
		}
		err = internal.RenderFile(tplPath, ymlPath, data)
		if err != nil {
			return errors.New("generate yml failed: " + err.Error())
		}
		return err
	}
	rootCmd.AddCommand(renderCmd)
}
