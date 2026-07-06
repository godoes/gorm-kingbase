package kingbase

import (
	"encoding/json"
	"errors"

	"github.com/godoes/gorm-kingbase/gokb"
	"gorm.io/gorm"
)

// The error codes to map KingbaseES errors to gorm errors,
// here is the KingbaseES error codes reference gokb.errorCodeNames
var errCodes = map[gokb.ErrorCode]error{
	"23505": gorm.ErrDuplicatedKey,
	"23503": gorm.ErrForeignKeyViolated,
	"42703": gorm.ErrInvalidField,
	"23514": gorm.ErrCheckConstraintViolated,
}

type ErrMessage struct {
	Code     gokb.ErrorCode
	Severity string
	Message  string
}

// Translate it will translate the error to native gorm errors.
func (dialector Dialector) Translate(err error) error {
	var kbErr *gokb.Error
	if errors.As(err, &kbErr) {
		if translatedErr, found := errCodes[kbErr.Code]; found {
			return translatedErr
		}
		return err
	}

	parsedErr, marshalErr := json.Marshal(err)
	if marshalErr != nil {
		return err
	}

	var errMsg ErrMessage
	unmarshalErr := json.Unmarshal(parsedErr, &errMsg)
	if unmarshalErr != nil {
		return err
	}

	if translatedErr, found := errCodes[errMsg.Code]; found {
		return translatedErr
	}
	return err
}
