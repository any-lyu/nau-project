package errorcode

// Code 是具体错误码的定义.
//
// 名字要求是 模块名+错误码描述, 通用错误码例外.
type Code int32

const (
	// OK is returned on success.
	Code_OK Code = 0
	// CANCELED indicates the operation was canceled (typically by the caller).
	Code_CANCELED Code = 1
	// UNKNOWN error. An example of where this error may be returned is
	// if a Status value received from another address space belongs to
	// an error-space that is not known in this address space. Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	Code_UNKNOWN Code = 2
	// INVALID_ARGUMENT indicates client specified an invalid argument.
	// Note that this differs from FAILED_PRECONDITION. It indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	Code_INVALID_ARGUMENT Code = 3
	// DEADLINE_EXCEEDED means operation expired before completion.
	// For operations that change the state of the system, this error may be
	// returned even if the operation has completed successfully. For
	// example, a successful response from a server could have been delayed
	// long enough for the deadline to expire.
	Code_DEADLINE_EXCEEDED Code = 4
	// NOT_FOUND means some requested entity (e.g., file or directory) was
	// not found.
	Code_NOT_FOUND Code = 5
	// ALREADY_EXISTS means an attempt to create an entity failed because one
	// already exists.
	Code_ALREADY_EXISTS Code = 6
	// PERMISSION_DENIED indicates the caller does not have permission to
	// execute the specified operation. It must not be used for rejections
	// caused by exhausting some resource (use RESOURCE_EXHAUSTED
	// instead for those errors). It must not be
	// used if the caller cannot be identified (use UNAUTHENTICATED
	// instead for those errors).
	Code_PERMISSION_DENIED Code = 7
	// RESOURCE_EXHAUSTED indicates some resource has been exhausted, perhaps
	// a per-user quota, or perhaps the entire file system is out of space.
	Code_RESOURCE_EXHAUSTED Code = 8
	// FAILED_PRECONDITION indicates operation was rejected because the
	// system is not in a state required for the operation's execution.
	// For example, directory to be deleted may be non-empty, an rmdir
	// operation is applied to a non-directory, etc.
	//
	// A litmus test that may help a service implementor in deciding
	// between FAILED_PRECONDITION, ABORTED, and UNAVAILABLE:
	//  (a) Use UNAVAILABLE if the client can retry just the failing call.
	//  (b) Use ABORTED if the client should retry at a higher-level
	//      (e.g., restarting a read-modify-write sequence).
	//  (c) Use FAILED_PRECONDITION if the client should not retry until
	//      the system state has been explicitly fixed. E.g., if an "rmdir"
	//      fails because the directory is non-empty, FAILED_PRECONDITION
	//      should be returned since the client should not retry unless
	//      they have first fixed up the directory by deleting files from it.
	//  (d) Use FAILED_PRECONDITION if the client performs conditional
	//      REST Get/Update/Delete on a resource and the resource on the
	//      server does not match the condition. E.g., conflicting
	//      read-modify-write on the same resource.
	Code_FAILED_PRECONDITION Code = 9
	// ABORTED indicates the operation was ABORTED, typically due to a
	// concurrency issue like sequencer check failures, transaction aborts,
	// etc.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Code_ABORTED Code = 10
	// OUT_OF_RANGE means operation was attempted past the valid range.
	// E.g., seeking or reading past end of file.
	//
	// Unlike INVALID_ARGUMENT, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate INVALID_ARGUMENT if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// OUT_OF_RANGE if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between FAILED_PRECONDITION and
	// OUT_OF_RANGE. We recommend using OUT_OF_RANGE (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an OUT_OF_RANGE error to detect when
	// they are done.
	Code_OUT_OF_RANGE Code = 11
	// UNIMPLEMENTED indicates operation is not implemented or not
	// supported/enabled in this service.
	Code_UNIMPLEMENTED Code = 12
	// INTERNAL errors. Means some invariants expected by underlying
	// system has been broken. If you see one of these errors,
	// something is very broken.
	Code_INTERNAL Code = 13
	// UNAVAILABLE indicates the service is currently unavailable.
	// This is a most likely a transient condition and may be corrected
	// by retrying with a backoff. Note that it is not always safe to retry
	// non-idempotent operations.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Code_UNAVAILABLE Code = 14
	// DATA_LOSS indicates unrecoverable data loss or corruption.
	Code_DATA_LOSS Code = 15
	// UNAUTHENTICATED indicates the request does not have valid
	// authentication credentials for the operation.
	Code_UNAUTHENTICATED Code = 16
	// REQUEST_ENTITY_TOO_LARGE indicates the request body is too large.
	Code_REQUEST_ENTITY_TOO_LARGE Code = 17
	// RATE_LIMIT_REACHED indicates that the frequency limit has been reached.
	Code_RATE_LIMIT_REACHED Code = 18
	// DUPLICATE_REQUEST indicates that the request is a duplicate request.
	Code_DUPLICATE_REQUEST Code = 19
)
