package response

type (
	Code int
)

// generic success code
const (
	CodeOK      Code = 200000
	CodeCreated Code = 201000
)

// generic error code
const (
	CodeBadRequest          Code = 400000
	CodeUnauthorized        Code = 401000
	CodeForbidden           Code = 403000
	CodeNotFound            Code = 404000
	CodeInternalServerError Code = 500000
)
