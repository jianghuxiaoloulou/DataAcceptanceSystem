package setting

// 配置文件

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")   // 配置文件名 (不带扩展格式)
	vp.AddConfigPath("configs/") // 配置文件的路径
	vp.SetConfigType("yaml")     // 如果你的配置文件没有写扩展名，那么这里需要声明你的配置文件属于什么格式
	err := vp.ReadInConfig()     //找到并读取配置文件
	if err != nil {
		return nil, err
	}
	return &Setting{vp: vp}, nil
}
