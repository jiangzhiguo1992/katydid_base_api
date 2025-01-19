package init

import (
	"katydid_base_api/configs"
	"katydid_base_api/tools"
)

func init() {
	// configs
	cloud, prod := tools.InitConfigStarts(configs.Files, configs.CloudKey, configs.ProdKey)
	tools.InitConfigEnds(configs.FilesGetByCloud(cloud))

	// logger
	tools.InitLogger(prod)
}
