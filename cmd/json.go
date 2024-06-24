package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strconv"
	"tool/internal"
)

var jsonCmd = &cobra.Command{
	Use:   "json [option]",
	Short: "use for handle json file",
}

func init() {
	jsonCmd.PersistentFlags().SortFlags = false
	jsonCmd.PersistentFlags().StringSliceVarP(&internal.JsonOptions.Paths, "path", "p", []string{}, "path list, example: 'engine,engine_cpu_num'")
	jsonCmd.PersistentFlags().StringVarP(&internal.JsonOptions.V, "value", "v", "", "value")
	jsonCmd.PersistentFlags().StringVarP(&internal.JsonOptions.T, "type", "t", "string", "type of value, example: int|string|bool")

	setCmd := &cobra.Command{
		Use:   "set",
		Short: "set json file value",
		RunE: func(cmd *cobra.Command, args []string) error {
			f, _ := cmd.Flags().GetString("file")
			paths, _ := cmd.Flags().GetStringSlice("path")
			value, _ := cmd.Flags().GetString("value")
			Type, _ := cmd.Flags().GetString("type")
			iFDebug, _ := cmd.Flags().GetBool("debug")
			if iFDebug {
				params := make(map[string]interface{})
				params["file"] = f
				params["path"] = paths
				params["value"] = value
				params["type"] = Type
				info, _ := json.Marshal(params)
				fmt.Println(string(info))
			}
			var v interface{}
			switch Type {
			case internal.TypeInt:
				// int类型
				vTmp, _ := strconv.Atoi(value)
				v = vTmp
			case internal.TypeString:
				// string类型
				v = value
			case internal.TypeBool:
				// bool类型
				vTmp, _ := strconv.ParseBool(value)
				v = vTmp
			}
			err := internal.Set(f, paths, v)
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
			f, _ := cmd.Flags().GetString("file")
			paths, _ := cmd.Flags().GetStringSlice("path")
			Type, _ := cmd.Flags().GetString("type")
			iFDebug, _ := cmd.Flags().GetBool("debug")
			if iFDebug {
				params := make(map[string]interface{})
				params["file"] = f
				params["path"] = paths
				info, _ := json.Marshal(params)
				fmt.Println(string(info))
			}
			jData, err := internal.Get(f, paths)
			if err != nil {
				return errors.New("get json file value failed: " + err.Error())
			}
			if jData != nil {
				switch Type {
				case internal.TypeInt:
					res, _ := jData.Int()
					fmt.Println(res)
				case internal.TypeString:
					res, _ := jData.String()
					fmt.Println(res)
				case internal.TypeBool:
					res, _ := jData.Bool()
					fmt.Println(res)
				}
			}
			return err
		},
	}
	jsonCmd.AddCommand(setCmd, getCmd)
	rootCmd.AddCommand(jsonCmd)
}
