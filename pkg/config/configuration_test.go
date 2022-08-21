package config

import (
	"config_con/pkg/pipe/consumer"
	"config_con/pkg/pipe/consumer/api"
	"config_con/pkg/pipe/consumer/twitch"
	"os"
	"reflect"
	"testing"
)

func TestReadConfiguration(t *testing.T) {
	os.Setenv("CONFIG_FILE_PATH", "test_data/test_config.yaml")
	tests := []struct {
		name    string
		want    YamlConfiguration
		wantErr bool
	}{
		{
			name: "ReadConfiguration",
			want: YamlConfiguration{
				Consumers: []consumer.ConsumerConfig{
					{
						Name: "test_consumer",
						Api: api.ApiConfiguration{
							TwitchConsumers: []twitch.TwitchEventConfig{
								{
									EventSecret: "test_consumer_secret",
									Url: "test_consumer_url",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfiguration()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}
