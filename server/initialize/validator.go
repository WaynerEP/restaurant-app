package initialize

import (
	"github.com/WaynerEP/restaurant-app/server/utils"
	"log"
	"reflect"
	"strings"

	"github.com/WaynerEP/restaurant-app/server/global"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni *ut.UniversalTranslator
)
var (
	errRequired       = "Este campo es obligatorio."
	errEmail          = "El {0} debe ser una dirección de correo electrónico válida."
	errUniqueDBExists = "El valor del campo ya existe en la base de datos."
)

func SetupValidator() {
	translator := en.New()
	uni = ut.New(translator, translator)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	utils.Trans, _ = uni.GetTranslator("en")

	utils.CustomValidate = validator.New()
	if err := entranslations.RegisterDefaultTranslations(utils.CustomValidate, utils.Trans); err != nil {
		log.Fatal(err)
	}

	utils.CustomValidate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	_ = utils.CustomValidate.RegisterTranslation("required", utils.Trans, func(ut ut.Translator) error {
		return ut.Add("required", errRequired, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = utils.CustomValidate.RegisterTranslation("email", utils.Trans, func(ut ut.Translator) error {
		return ut.Add("email", errEmail, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	_ = utils.CustomValidate.RegisterTranslation("unique_db", utils.Trans, func(ut ut.Translator) error {
		return ut.Add("unique_db", errUniqueDBExists, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("unique_db", fe.Field())
		return t
	})

	/*
		validate:"unique_db=column:table:except:idColumn"
			column: Es el nombre de la columna a evaluar en bd, por defecto toma el tag json; ex. code (Ojo que en db es snake_case)
			table: Nombre de la tabla a evaluar, por defecto es la estructura - Gorm ya lo formatea a snake_case
			except: Es el nombre del campo PK de la estructura de GO del cual se quiere obtener el valor, por defecto es ID - Only UpdateOffice funciona
			idColum: Es el nombre de la columna PK en BD, por defecto es id - Only UpdateOffice funciona
	*/
	_ = utils.CustomValidate.RegisterValidation("unique_db", func(fl validator.FieldLevel) bool {
		var count int64
		var query string
		var args []interface{}
		var value interface{}
		field := fl.Field()
		switch field.Type().Kind() {
		case reflect.String:
			value = field.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value = field.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			value = field.Uint()
		case reflect.Float32, reflect.Float64:
			value = field.Float()
		default:
			log.Printf("validation failed: unsupported field type: %v", field.Type().Kind())
			return false
		}

		column, table, except, idColumn := getUniqueDBParams(fl.Param())
		idExcept := fl.Top().FieldByName(except).Interface()

		if column == "" {
			column = fl.FieldName()
		}
		query = column + " = ?"
		if reflect.ValueOf(idExcept).IsZero() {
			//CREATE
			args = append(args, value)
		} else {
			// UPDATE
			query = idColumn + " != ? AND " + query
			args = append(args, idExcept, value)
		}
		if table == "" {
			global.GVA_DB.Model(fl.Parent().Interface()).
				Where(query, args[0:]...).
				Count(&count)
		} else {
			global.GVA_DB.Table(table).
				Where(query, args[0:]...).
				Count(&count)
		}
		return count < 1
	})
}

func getUniqueDBParams(tag string) (column string, table string, except string, idColumn string) {
	// Si la etiqueta es vacía, se devuelven los valores predeterminados
	if tag == "" {
		return "", "", "ID", "id"
	}

	// Se divide la etiqueta en los parámetros separados por un espacio
	params := strings.Split(tag, ":")

	// Valores por default
	except = "ID"
	idColumn = "id"

	// Si solo hay un parámetro, se asume que es el nombre del campo en la estructura
	if len(params) == 1 {
		column = params[0]
	} else if len(params) == 2 {
		column = params[0]
		table = params[1]
	} else if len(params) == 3 {
		column = params[0]
		table = params[1]
		except = params[2]
	} else {
		column = params[0]
		table = params[1]
		except = params[2]
		idColumn = params[3]
	}

	// Se devuelven los valores de idField y primaryKey
	return column, table, except, idColumn
}
