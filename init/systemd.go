package init

import (
	"katydid_base_api/configs"
	"katydid_base_api/tools"
)

func init() {
	// configs
	cloud, prod := tools.InitConfigStarts(configs.FilesGet(), configs.CloudKey, configs.ProdKey)
	tools.InitConfigEnds(configs.FilesGetByCloud(cloud))

	// logger
	tools.InitLogger(prod)
}
