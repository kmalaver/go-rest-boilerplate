package errors

// request body errors
// codes: 400_xx

var (
	ErrDefaultHTTPError = NewDynamic(400_00, "HTTPError")
	ErrUnmarhsal        = NewDynamic(400_01, "UnmarshalError")
	ErrValidation       = NewDynamic(400_02, "ValidationError")
)
