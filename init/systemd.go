package init

import (
	"katydid_base_api/tools"
)

func init() {
	// configs
	_, _, prod := tools.InitConfigs()

	// logger
	tools.InitLogger(prod)
}
