package configs

// standard variables
var (
	DB_PATH     = "."
	DB_NAME     = "sqlite.db"
	DB_USERNAME = ""
	DB_PASSWORD = ""
)

const (
	DB_AUTHCRYPT = "SHA256"

	DBDriverName = "sqlite3"

	// db constants
	TimeFormatRFC1123 = "Mon, 02 Jan 2006 15:04:05 MST"
	TimeFormatRFC3339 = "2006-01-02T15:04:05Z07:00"
	DislikeValue      = 0
	LikeValue         = 1

	// frontend constants
	LimitQuestionPerPage = 10
	LimitAnswersPerPage  = 10
	LimitTagsPerPage     = 25
)
