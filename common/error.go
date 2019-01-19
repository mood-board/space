package common

import "fmt"

type UserError struct {
	err string
}

func NewUserError(text string, args ...interface{}) *UserError {
	return &UserError{fmt.Sprintf(text, args...)}
}

func (e *UserError) Error() string {
	return e.err
}

var (
	ErrUnauthorized       = NewUserError("unauthorized access")
	ErrInvalidCredentials = NewUserError("E-mail or password is incorrect")
	ErrInvalidUserType    = NewUserError("invalid user type")
	ErrInvalidObjectID    = NewUserError("given string is not a valid object ID")
	ErrInvalidToken       = NewUserError("invalid token")
	ErrNoBusiness         = NewUserError("no business found for user")
	ErrSelfRequest        = NewUserError("requesting money from yourself")
	ErrUnsupported        = NewUserError("unsupported method")
	ErrValidationFailed   = NewUserError("validation of object failed")
	ErrUnknownInstrument  = NewUserError("unknown instrument type")
	ErrPhoneUnreachable   = NewUserError("phone unreachable for transfer")

	ErrNoMessages  = NewUserError("failed to get messages")
	ErrTextTooLong = func(maxLength int) *UserError { return NewUserError("text cannot be longer than %d", maxLength) }

	ErrInvalidAmount = NewUserError("invalid amount")
	ErrZeroAmount    = NewUserError("zero amount")
	ErrDecimalValue  = NewUserError("cannot get decimal value")
	ErrCurrencyValue = NewUserError("cannot get currency value")

	ErrNoCashback    = NewUserError("failed to get cashback")
	ErrNoBalance     = NewUserError("failed to get balance")
	ErrNoWallet      = NewUserError("failed to get wallet")
	ErrNoInstruments = NewUserError("failed to get instruments")

	ErrInvalidArguments = NewUserError("invalid arguments")
	ErrInvalidArgument  = func(argument string) *UserError { return NewUserError("invalid argument: %s", argument) }
	ErrMissingArgument  = func(argument string) *UserError { return NewUserError("missing argument: %s", argument) }

	ErrDecodeData        = NewUserError("cannot decode data")
	ErrDecryptData       = NewUserError("cannot decrypt data")
	ErrUnmarshalData     = NewUserError("cannot unmarshal data")
	ErrNoTransactionData = NewUserError("no transactions found")

	ErrSoundOfThePolice = NewUserError("the request has been forwarded to the international police association")

	ErrSomethingWentWrong = NewUserError("something went wrong")
)
