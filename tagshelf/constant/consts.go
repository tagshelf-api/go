package constant

const (
	// Tagshelf API base url
	BaseURL = "https://app.tagshelf.com"

	// Endpoints
	EpToken      = BaseURL + "/token"
	EpStatus     = BaseURL + "/api/tagshelf/status"
	EpWhoAmI     = BaseURL + "/api/tagshelf/who-am-i"
	EpPing       = BaseURL + "/api/tagshelf/ping"
	EpFileUpload = BaseURL + "/api/file/upload"
	EpFileDetail = BaseURL + "/api/file/detail/%s"
	EpJobDetail  = BaseURL + "/api/job/detail/%s"

	// Methods
	MethodGET     = "GET"
	MethodPOST    = "POST"
	EpTokenM      = MethodPOST
	EpStatusM     = MethodGET
	EpWhoAmIM     = MethodGET
	EpPingM       = MethodGET
	EpFileUploadM = MethodPOST
	EpFileDetailM = MethodGET
	EpJobDetailM  = MethodGET

	// ContentTypes
	CtJSON  = "application/json"
	CtXForm = "application/x-www-form-urlencoded"

	// Auth Header
	AuthHBearer = "Bearer %s"
	AuthHAmx    = "amx %s"

	// User Agent
	UAHeader = "Tagshelf-Go-Client"
)

const (
	OAuthGT  = "grant_type"
	OAuthUSR = "username"
	OAuthPWD = "password"
)
