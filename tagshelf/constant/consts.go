package constant

import "os"

var (
	// BaseURL tagshelf api base url.
	BaseURL = getEnv("TAGSHELF_API_BASE_URL", DefaultBaseURL)

	// Endpoints
	EndpointToken        = BaseURL + "/token"
	EndpointStatus       = BaseURL + "/api/tagshelf/status"
	EndpointWhoAmI       = BaseURL + "/api/tagshelf/who-am-i"
	EndpointPing         = BaseURL + "/api/tagshelf/ping"
	EndpointFileUpload   = BaseURL + "/api/file/upload"
	EndpointFileDetail   = BaseURL + "/api/file/detail/%s"
	EndpointJobDetail    = BaseURL + "/api/job/detail/%s"
	EndpointCompanyInbox = BaseURL + "/api/company/inbox"
)

const (
	// Tagshelf API base url
	DefaultBaseURL = "https://app.tagshelf.com"

	// Methods
	MethodGET                = "GET"
	MethodPOST               = "POST"
	EndpointTokenMethod      = MethodPOST
	EndpointStatusMethod     = MethodGET
	EndpointWhoAmIMethod     = MethodGET
	EndpointPingMethod       = MethodGET
	EndpointFileUploadMethod = MethodPOST
	EndpointFileDetailMethod = MethodGET
	EndpointJobDetailMethod  = MethodGET

	// ContentTypes
	ContentTypeJSON  = "application/json"
	ContentTypeXForm = "application/x-www-form-urlencoded"

	// Auth Header
	AuthHBearer = "Bearer %s"
	AuthHAmx    = "amx %s"

	// User Agent
	UserAgentHeader = "Tagshelf-Go-Client"
)

const (
	OAuthGrantType = "grant_type"
	OAuthUSR       = "username"
	OAuthPWD       = "password"
)

func getEnv(env, def string) (val string) {
	if val = os.Getenv(env); val == "" {
		return def
	}

	return val
}
