package constant

// Environment type represents various kind of environments an application can run in.
type Environment string

const (
	Production  = Environment("prod")
	Development = Environment("dev")
	QA          = Environment("qa")
	Staging     = Environment("staging")
)

const (
	LogInfoFilePath   = "mobilecore.json"
	LogAccessFilePath = "access.log"
)

const (
	OS       = "os"
	BUSINESS = "BUSINESS"
	NO       = "NO"
	DB       = "DB"
	IN       = "IN"
)

const (
	POST = "POST"
	GET  = "GET"
)

const (
	ALL           = "ALL"
	MOBILECOREAPP = "MobileCoreApp"
)

const (
	PORT                      = 9007
	DIALER_TIME_OUT           = 1000
	MaxIdleConnectionsPerHost = 5
)

const (
	COVID_STATE_API_URL          = "https://data.covid19india.org/v4/min/data.min.json"
	COVID_STATE_API_TIMEOUT      = 1000
	REVERSE_GEO_CODE_API_URL     = "http://api.positionstack.com/v1/reverse?access_key=c3045fb4fc4c1d159316961b82c975da"
	REVERSE_GEO_CODE_API_TIMEOUT = 1000
	MONGO_DOC_NAME               = "Covid Data"
)
