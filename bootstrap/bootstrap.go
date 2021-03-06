package bootstrap

import "Dict/config"

// Setup 初始化指定的服务，例如：Redis MySQL Logger 等模块.
func Setup() {
	autoLoader(
		config.Initialize, // 配置文件
		SetupMySQL,
	)
}

// autoLoader 自动加载初始化.
func autoLoader(funcName ...func()) {
	for _, v := range funcName {
		v()
	}
}
