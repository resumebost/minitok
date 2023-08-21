package constant

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

const ETCDAddress = "127.0.0.1:2379"

var AllConstants Constant

func InitConstant() {
	v := viper.New()

	if err := v.AddRemoteProvider("etcd3", ETCDAddress, "/minitok/config"); err != nil {
		klog.Fatalf("Viper can not add remote provider: %v", err)
	}

	v.SetConfigType("yaml")

	if err := v.ReadRemoteConfig(); err != nil {
		klog.Fatalf("Viper can not read remote config: %v", err)
	}

	if err := v.Unmarshal(&AllConstants); err != nil {
		klog.Fatalf("Viper can not convert constants into a Go struct: %v", err)
	}
}
