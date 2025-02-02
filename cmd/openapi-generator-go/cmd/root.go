/*
Copyright © 2020 Tino Rusch

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "openapi-generator-go",
	Short: "generate openapi based go code!",
	Long:  `generate openapi based go code!.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(
		initConfig,
		initLogging,
		initAllowedStringFormats,
	)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringP("spec", "s", "api.yaml", "filepath to the openapi spec file")
	rootCmd.PersistentFlags().StringP("output", "o", "./models", "output directory")
	rootCmd.PersistentFlags().String("package-name", "api", "package name")
	rootCmd.PersistentFlags().String("log-level", "info", "log level")
	rootCmd.PersistentFlags().StringArrayP("path", "p", []string{}, "path to include in the filtered output spec, can be repeated")
	rootCmd.PersistentFlags().String("path-file", "", "file with list of paths to include in the filtered output spec, one path per line")

	rootCmd.PersistentFlags().StringArray("formats", []string{"uuid"}, "additional string formats to allow during schema validation, only required when using the filter flags/command")
}

func initLogging() {
	level, _ := rootCmd.PersistentFlags().GetString("log-level")
	lvl, err := zerolog.ParseLevel(strings.ToLower(level))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse --log-level")
	}
	zerolog.SetGlobalLevel(lvl)
}

func initAllowedStringFormats() {
	additionalFormats, err := rootCmd.PersistentFlags().GetStringArray("formats")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse --formats")
	}

	for _, value := range additionalFormats {
		// the openapi validation is very strict about the allowed string formats
		// this method allows us to register additional values for it to accept.
		// This is only required when we filter the spec.
		openapi3.DefineStringFormat(value, `.*`)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".openapi-generator-go" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".openapi-generator-go")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func getAllowedPaths(flags *pflag.FlagSet) []string {
	allowedPaths, err := flags.GetStringArray("path")
	if err != nil {
		log.Fatal().Err(err).Msg("can't read path flag(s)")
	}

	allowedPathsFile, err := flags.GetString("path-file")
	if err != nil {
		log.Fatal().Err(err).Msg("can't read path-file flag")
	}
	if allowedPathsFile != "" {
		additionalPathsFile, err := os.Open(allowedPathsFile)
		if err != nil {
			log.Fatal().Str("path-file", allowedPathsFile).Err(err).Msg("can't read path-file")
		}

		var line string
		reader := bufio.NewReader(additionalPathsFile)
		for {
			line, err = reader.ReadString('\n')
			if err != nil && err != io.EOF {
				break
			}
			allowedPaths = append(allowedPaths, strings.TrimSpace(line))
			if err != nil {
				break
			}
		}
		if err != io.EOF {
			log.Fatal().Str("path-file", allowedPathsFile).Err(err).Msg("can't read paths in path file")
		}
	}

	return allowedPaths
}
