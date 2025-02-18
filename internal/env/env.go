package env

type Env struct {
	Port             int
	DataBaseURL      string
	DefaultAdminName string
	DefaultAdminPass string
	UrlPrefix        string
	DisableSwagger   bool
	Version          string
	RefreshHours     int
}
