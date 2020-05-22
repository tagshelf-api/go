package constant

const (
	// Tagshelf API base url
	BaseURL = "https://tagshelf.com"

	// Endpoints
	EpStatus     = BaseURL + "/api/tagshelf/status"
	EpWhoAmI     = BaseURL + "/api/tagshelf/who-am-i"
	EpPing       = BaseURL + "/api/tagshelf/ping"
	EpFileUpload = BaseURL + "/api/file/upload"
	EpFileDetail = BaseURL + "/api/file/detail/%s"
	EpJobDetail  = BaseURL + "/api/job/detail/%s"
)
