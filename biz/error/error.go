package biz_error

// System level error codes (1000-1999)
const (
	ServerError      = 1000 // Internal server error
	ParamInvalid     = 1001 // Invalid parameters
	Unauthorized     = 1002 // Unauthorized access
	Forbidden        = 1003 // Forbidden access
	NotFound         = 1004 // Resource not found
	MethodNotAllowed = 1005 // Method not allowed
	Timeout          = 1006 // Request timeout
)

// Business level error codes (2000-2999)
const (
	UserNotExistOrPassword = 2000 // User does not exist
	UserAlreadyExist       = 2001 // User already exists
	PasswordError          = 2002 // Incorrect password
	TokenInvalid           = 2003 // Invalid token
	TokenExpired           = 2004 // Token has expired
)

// Database error codes (3000-3999)
const (
	DBError          = 3000 // Database error
	DBRecordNotFound = 3001 // Record not found
	DBDuplicate      = 3002 // Duplicate record
)

// Third-party service error codes (4000-4999)
const (
	ThirdPartyError = 4000 // Third-party service error
	APICallFailed   = 4001 // API call failed
)

// Error defines the error structure
type Error struct {
	Code    int    `json:"code"`    // Error code
	Message string `json:"message"` // Error message
}

// Implement error interface
func (e *Error) Error() string {
	return e.Message
}

// Error code to message mapping
var errorMessages = map[int]string{
	// System level errors
	ServerError:      "Internal server error",
	ParamInvalid:     "Invalid parameters",
	Unauthorized:     "Unauthorized access",
	Forbidden:        "Forbidden access",
	NotFound:         "Resource not found",
	MethodNotAllowed: "Method not allowed",
	Timeout:          "Request timeout",

	// Business level errors
	UserNotExistOrPassword: "User does not exist or password is incorrect",
	UserAlreadyExist:       "User already exists",
	PasswordError:          "Incorrect password",
	TokenInvalid:           "Invalid token",
	TokenExpired:           "Token has expired",

	// Database errors
	DBError:          "Database error",
	DBRecordNotFound: "Record not found",
	DBDuplicate:      "Duplicate record",

	// Third-party service errors
	ThirdPartyError: "Third-party service error",
	APICallFailed:   "API call failed",
}

// New creates a new error with predefined message
func New(code int) *Error {
	return &Error{
		Code:    code,
		Message: errorMessages[code],
	}
}

// NewWithMessage creates a new error with custom message
func NewWithMessage(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// IsError checks if the error matches the specified error code
func IsError(err error, code int) bool {
	if err == nil {
		return false
	}
	if e, ok := err.(*Error); ok {
		return e.Code == code
	}
	return false
}

// GetCode retrieves the error code
func GetCode(err error) int {
	if err == nil {
		return 0
	}
	if e, ok := err.(*Error); ok {
		return e.Code
	}
	return ServerError
}

// GetMessage retrieves the error message
func GetMessage(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(*Error); ok {
		return e.Message
	}
	return err.Error()
}

// Global error variables for direct use
var (
	ErrServer           = New(ServerError)
	ErrParamInvalid     = New(ParamInvalid)
	ErrUnauthorized     = New(Unauthorized)
	ErrForbidden        = New(Forbidden)
	ErrNotFound         = New(NotFound)
	ErrMethodNotAllowed = New(MethodNotAllowed)
	ErrTimeout          = New(Timeout)

	ErrUserNotExistOrPassword = New(UserNotExistOrPassword)
	ErrUserAlreadyExist       = New(UserAlreadyExist)
	ErrPassword               = New(PasswordError)
	ErrTokenInvalid           = New(TokenInvalid)
	ErrTokenExpired           = New(TokenExpired)

	ErrDB               = New(DBError)
	ErrDBRecordNotFound = New(DBRecordNotFound)
	ErrDBDuplicate      = New(DBDuplicate)

	ErrThirdParty    = New(ThirdPartyError)
	ErrAPICallFailed = New(APICallFailed)
)
