package invalid

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}

func ValidationError[s any](validate *validator.Validate, params s, err error) map[string]interface{} {
	errorList := make(map[string]interface{})
	if err != nil {
		errors := err.(validator.ValidationErrors)

		for _, er := range errors {
			var errMsg string
			field, _ := reflect.TypeOf(params).FieldByName(er.StructField())
			fieldName := field.Tag.Get("json")
			switch er.Tag() {
			case "required":
				errMsg = fmt.Sprintf("%s form ini wajib diisi", fieldName)
			case "email":
				errMsg = fmt.Sprintf("%s bukan format email yang benar", fieldName)
			case "oneof":
				errMsg = fmt.Sprintf("%s harus berupa %s", fieldName, er.Param())
			case "min":
				errMsg = fmt.Sprintf("%s minimal %s", fieldName, er.Param())
			case "max":
				errMsg = fmt.Sprintf("%s minimal %s", fieldName, er.Param())

			}
			errorList[fieldName] = errMsg
		}
	}
	return errorList
}

var (
	ErrNotFound   = errors.New("requested item not found")       //not found produk
	ErrWrongInput = errors.New("username atau password salah")   //wrong password or username
	ErrNoSesFound = errors.New("login required for this action") //session login invalid ||
	ErrEmptyField = errors.New("semua form wajib di isi")        //empty field invalid for html form
)
