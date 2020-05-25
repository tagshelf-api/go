package constant

const (
	// Tagshelf API base url
	BaseURL = "https://app.tagshelf.com"

	// Endpoints
	EndpointToken      = BaseURL + "/token"
	EndpointStatus     = BaseURL + "/api/tagshelf/status"
	EndpointWhoAmI     = BaseURL + "/api/tagshelf/who-am-i"
	EndpointPing       = BaseURL + "/api/tagshelf/ping"
	EndpointFileUpload = BaseURL + "/api/file/upload"
	EndpointFileDetail = BaseURL + "/api/file/detail/%s"
	EndpointJobDetail  = BaseURL + "/api/job/detail/%s"

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
