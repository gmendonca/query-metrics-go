package cmd

import (
	"github.com/gmendonca/tapper/pkg/datadog"
	"github.com/gmendonca/tapper/pkg/elasticsearch"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "logs command line interface",
	Long:  `logs, the command line interface`,
	Run: func(cmd *cobra.Command, args []string) {
		e := &elasticsearch.Elasticsearch{
			Host:     viper.GetString("elasticsearch.host"),
			Port:     viper.GetInt("elasticsearch.port"),
			Username: viper.GetString("elasticsearch.username"),
			Password: viper.GetString("elasticsearch.password"),
			SSL:      viper.GetBool("elasticsearch.ssl"),
		}

		d := &datadog.Datadog{
			APIKey:        viper.GetString("datadog.api_key"),
			ApplicationID: viper.GetString("datadog.application_id"),
		}

		e.SendMetrics(d)
	},
}

func init() {
	rootCmd.AddCommand(logsCmd)
}
