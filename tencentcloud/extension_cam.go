package tencentcloud

const PAGE_ITEM = 200

const (
	CAM_POLICY_CREATE_STRATEGY_CUSTOM = "User"
	CAM_POLICY_CREATE_STRATEGY_PRESET = "QCS"
	CAM_POLICY_CREATE_STRATEGY_NULL   = ""
)

var CAM_POLICY_CREATE_STRATEGY = []string{
	CAM_POLICY_CREATE_STRATEGY_CUSTOM,
	CAM_POLICY_CREATE_STRATEGY_PRESET,
	CAM_POLICY_CREATE_STRATEGY_NULL,
}
