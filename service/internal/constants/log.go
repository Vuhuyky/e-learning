package constant

type LOGGER_TYPE string

const (
	INFO_LOG  LOGGER_TYPE = "info"
	ERROR_LOG LOGGER_TYPE = "error"
)

type LOGGER_DEFAULT_TITLE string

const (
	TITLE_GET_PAYLOAD = "get-payload"
	TITLE_GET         = "get"
	TITLE_CREATE      = "create"
	TITLE_UPDATE      = "update"
	TITLE_DELETE      = "delete"
)
