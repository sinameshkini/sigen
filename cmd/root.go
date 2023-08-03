package cmd

import (
	"github.com/sinameshkini/sigen/template"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	templateName string
	out          string
	// Used for flags.
	cfgFile string
	// userLicense string

	rootCmd = &cobra.Command{
		Use:   "sigen",
		Short: "code generator based on template",
		Long:  `sigen is a simple tools for developers to auto generate codes`,
		Run: func(cmd *cobra.Command, args []string) {
			var (
				variables = make(map[string]string)
			)

			for _, arg := range args {
				if strings.HasPrefix(arg, "_") {
					kv := strings.Split(arg, "=")
					variables[kv[0]] = kv[1]
				}
			}

			if err := template.Make(templateName, out, variables); err != nil {
				logrus.Errorln(err)
			}
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	_ = logrus.New()
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $COMMAND_HOME/config.yml)")
	rootCmd.Flags().StringVarP(&templateName, "template", "t", "repository", "template key on config")
	rootCmd.Flags().StringVarP(&out, "out", "o", "./out", "output path")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			logrus.Errorln(err.Error())
		}

		configPath := homeDir + "/.sigen/."

		viper.AddConfigPath(configPath)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		logrus.Infoln("Using config file:", viper.ConfigFileUsed())
	} else {
		logrus.Errorln("reading config file error:", err.Error())
	}
}
