package validation

import (
	"encoding/json"
	"errors"
	"github.com/Mateus003/user-api/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validation = validator.New()
	transl	ut.Translator
)

func init(){
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok{
		en:= en.New()
		unt:= ut.New(en,en)
		transl,_ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserErrors(validation_err error) *rest_err.Rest_Err{
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr){
		return rest_err.BadRequestError("Invalid field type")
	}else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors){
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field: e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	}else{
		return rest_err.BadRequestError("Error trying to convert")
	} 
}