package connectionwatcher

import (
	pb "github.com/turbot/steampipe/pkg/pluginmanager/grpc/proto"
	"github.com/turbot/steampipe/pkg/steampipeconfig/modconfig"
)

func NewConnectionConfigMap(connectionMap map[string]*modconfig.Connection) map[string]*pb.ConnectionConfig {
	configMap := make(map[string]*pb.ConnectionConfig)
	for k, v := range connectionMap {
		configMap[k] = &pb.ConnectionConfig{
			Plugin:          v.Plugin,
			PluginShortName: v.PluginShortName,
			Config:          v.Config,
		}
	}
	return configMap
}
