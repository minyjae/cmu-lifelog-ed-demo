package utils

// Error codes are machine-readable identifiers returned in the "code" field of
// a response. Unlike "message" (human-readable, may be translated), a code is
// stable and meant to be consumed by the frontend for branching / i18n.
//
// Convention:
//   - UPPER_SNAKE_CASE
//   - describe WHAT went wrong, not the HTTP status (e.g. INVALID_CREDENTIALS,
//     not UNAUTHORIZED)
//   - keep them stable; change the message, never the code
const (
	// --- Generic ---
	CodeInvalidRequest  = "INVALID_REQUEST"  // malformed body / params
	CodeInvalidID       = "INVALID_ID"       // bad path/query id
	CodeValidationFailed = "VALIDATION_FAILED" // failed field validation
	CodeNotFound        = "NOT_FOUND"        // resource does not exist
	CodeForbidden       = "FORBIDDEN"        // authenticated but no permission
	CodeUnauthorized    = "UNAUTHORIZED"     // missing / invalid auth
	CodeInternalError   = "INTERNAL_SERVER_ERROR"
	CodeConflict        = "CONFLICT" // resource state conflict (e.g. duplicate)

	// --- Auth ---
	CodeMissingCredentials = "MISSING_CREDENTIALS"     // email/password not provided
	CodeInvalidCredentials = "INVALID_CREDENTIALS"     // wrong email or password
	CodeTokenGenFailed     = "TOKEN_GENERATION_FAILED" // JWT signing failed
	CodeInvalidToken       = "INVALID_TOKEN"           // bad / expired JWT
	CodeEmailAlreadyExists = "EMAIL_ALREADY_EXISTS"    // duplicate email on register

	// --- User ---
	CodeUserNotFound = "USER_NOT_FOUND"

	// --- Faculty ---
	CodeFacultyNotFound = "FACULTY_NOT_FOUND"
)
