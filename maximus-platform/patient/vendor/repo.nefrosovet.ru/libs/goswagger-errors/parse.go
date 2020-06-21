package goswagger_errors

import (
	"regexp"

	oapiErrors "github.com/go-openapi/errors"
)

var (
	patternUnmarshalError = regexp.MustCompile(`cannot unmarshal \w+ into Go struct field \w+(?:\.(\w+))+ of type (\w+)`)
	patternEOF            = regexp.MustCompile(`parsing .* failed, because unexpected EOF`)
	patternRequired       = regexp.MustCompile(`(\w+) in \w+ is required`)
	patternType           = regexp.MustCompile(`(\w+) in \w+ should be one of .*`)
	patternFormat         = regexp.MustCompile(`(\w+) in \w+ must be of type .*`)

	patternParseError        = regexp.MustCompile(`parsing .* failed, because parse error: expected (\w+) .*`)
	patternTooShortError     = regexp.MustCompile(`(\w+) should be at least \w+ chars long`)
	patternTooShortBodyError = regexp.MustCompile(`(\w+) in \w+ should be at least \w+ chars long`)
)

func parseComposite(err *oapiErrors.CompositeError) *ValidationError {
	res := &ValidationError{
		Validation: map[string]interface{}{},
	}

	wrapCore := func(message string) {
		if res.Core != "" {
			res.Core += "\n"
		}

		res.Core += message
	}

	for _, subErr := range flattenComposite(err).Errors {
		switch apiErr := subErr.(type) {
		case *oapiErrors.ParseError:
			if m := patternUnmarshalError.FindStringSubmatch(apiErr.Error()); m != nil {
				res.Validation[m[1]] = m[2]
			} else if m := patternParseError.FindStringSubmatch(apiErr.Error()); m != nil {
				wrapCore("JSON parse error")
			} else if patternEOF.MatchString(apiErr.Error()) {
				res.JSON = "EOF"
			} else {
				wrapCore(apiErr.Error())
			}
		case *oapiErrors.Validation:
			if m := patternRequired.FindStringSubmatch(apiErr.Error()); m != nil {
				res.Validation[m[1]] = "required"
			} else if m := patternType.FindStringSubmatch(apiErr.Error()); m != nil {
				res.Validation[m[1]] = "enum"
			} else if m := patternFormat.FindStringSubmatch(apiErr.Error()); m != nil {
				res.Validation[m[1]] = "format"
			} else if m := patternTooShortBodyError.FindStringSubmatch(apiErr.Error()); m != nil {
				res.Validation[m[1]] = "min"
			} else if m := patternTooShortError.FindStringSubmatch(apiErr.Error()); m != nil {
				res.Validation[m[1]] = "min"
			} else {
				wrapCore(apiErr.Error())
			}
		default:
			wrapCore(apiErr.Error())
		}
	}

	return res
}

func flattenComposite(errs *oapiErrors.CompositeError) *oapiErrors.CompositeError {
	var res []error
	for _, er := range errs.Errors {
		switch e := er.(type) {
		case *oapiErrors.CompositeError:
			if len(e.Errors) > 0 {
				flat := flattenComposite(e)
				if len(flat.Errors) > 0 {
					res = append(res, flat.Errors...)
				}
			}
		default:
			if e != nil {
				res = append(res, e)
			}
		}
	}

	return oapiErrors.CompositeValidationError(res...)
}
