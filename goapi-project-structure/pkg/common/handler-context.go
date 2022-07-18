package common

type HandleFunc func(ctx HContext)

type HContext interface {
	Bind(interface{}) error
	BindQuery(interface{}) error
	Query(string) (string, bool)
	DefaultQuery(string, string) string
	Param(string) string
	Header(string) string
	Authorization() string
	RequestId() string
	ResponseJSON(int, interface{})
}
