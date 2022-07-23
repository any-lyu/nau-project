package apierror

import "family-joint-school/common/errors/errorcode"

// 通用错误
var (
	ErrCanceled              error = New(errorcode.Code_CANCELED, "canceled")
	ErrUnknown               error = New(errorcode.Code_UNKNOWN, "unknown")
	ErrInvalidArgument       error = New(errorcode.Code_INVALID_ARGUMENT, "invalid argument")
	ErrDeadlineExceeded      error = New(errorcode.Code_DEADLINE_EXCEEDED, "deadline exceeded")
	ErrNotFound              error = New(errorcode.Code_NOT_FOUND, "not found")
	ErrAlreadyExists         error = New(errorcode.Code_ALREADY_EXISTS, "already exists")
	ErrPermissionDenied      error = New(errorcode.Code_PERMISSION_DENIED, "permission denied")
	ErrResourceExhausted     error = New(errorcode.Code_RESOURCE_EXHAUSTED, "resource exhausted")
	ErrFailedPrecondition    error = New(errorcode.Code_FAILED_PRECONDITION, "failed precondition")
	ErrAborted               error = New(errorcode.Code_ABORTED, "aborted")
	ErrOutOfRange            error = New(errorcode.Code_OUT_OF_RANGE, "out of range")
	ErrUnimplemented         error = New(errorcode.Code_UNIMPLEMENTED, "unimplemented")
	ErrInternal              error = New(errorcode.Code_INTERNAL, "internal server error")
	ErrUnavailable           error = New(errorcode.Code_UNAVAILABLE, "unavailable")
	ErrDataLoss              error = New(errorcode.Code_DATA_LOSS, "data loss")
	ErrUnauthenticated       error = New(errorcode.Code_UNAUTHENTICATED, "unauthenticated")
	ErrRequestEntityTooLarge error = New(errorcode.Code_REQUEST_ENTITY_TOO_LARGE, "request entity too large")
	ErrRateLimitReached      error = New(errorcode.Code_RATE_LIMIT_REACHED, "rate limit reached")
	ErrDuplicateRequest      error = New(errorcode.Code_DUPLICATE_REQUEST, "duplicate request")
)
