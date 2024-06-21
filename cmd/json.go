package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"tool/internal"
)

var jsonCmd = &cobra.Command{
	Use:   "json [option]",
	Short: "use for handle json file",
}

func init() {
	jsonCmd.PersistentFlags().SortFlags = false
	jsonCmd.PersistentFlags().StringVarP(&internal.JsonOptions.JsonPath, "file", "f", "", "json file path")
	jsonCmd.PersistentFlags().StringVarP(&internal.JsonOptions.K, "key", "k", "", "json key, example: file.engine.engine_cpu_num")
	jsonCmd.PersistentFlags().StringVarP(&internal.JsonOptions.V, "value", "v", "", "value")
	jsonCmd.PersistentFlags().Int64VarP(&internal.JsonOptions.T, "type", "t", 2, "type of value")

	setCmd := &cobra.Command{
		Use:   "set",
		Short: "set json file value",
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath, _ := cmd.Flags().GetString("file")
			key, _ := cmd.Flags().GetString("key")
			value, _ := cmd.Flags().GetString("value")
			Type, _ := cmd.Flags().GetInt64("type")
			var v interface{}
			switch Type {
			case 1:
				// int类型
				vTmp, _ := strconv.Atoi(value)
				v = vTmp
			case 2:
				// string类型
				v = value
			case 3:
				// bool类型
				vTmp, _ := strconv.ParseBool(value)
				v = vTmp
			}
			err := internal.Set(filePath, key, v)
			if err != nil {
				return errors.New("set json file value failed: " + err.Error())
			}
			return err
		},
	}
	getCmd := &cobra.Command{
		Use:   "get",
		Short: "get json file value",
		RunE: func(cmd *cobra.Command, args []string) error {
			filePath, _ := cmd.Flags().GetString("file")
			key, _ := cmd.Flags().GetString("key")
			value, err := internal.Get(filePath, key)
			if err != nil {
				return errors.New("get json file value failed: " + err.Error())
			}
			os.Stdout.WriteString(value)
			return err
		},
	}
	jsonCmd.AddCommand(setCmd, getCmd)
	rootCmd.AddCommand(jsonCmd)
}
