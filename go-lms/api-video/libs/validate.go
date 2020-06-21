package libs

import (
	"regexp"

	"repo.nefrosovet.ru/go-lms/api-video/api"
)

const (
	ValidationSipRequiredMessage = `The sip field is required`
	ValidationSipInvalidMessage  = `Bad sipUrl format or invalid domain`
)

func ValidateSip(sip *api.Sip) (*map[string]string, error) {
	validation := make(map[string]string)

	if sip == nil {
		validation["required"] = ValidationSipRequiredMessage
		return &validation, nil
	}

	startsWithSip, err := regexp.MatchString(`^sip:`, string(*sip))
	if err != nil {
		return nil, err
	}

	re, err := regexp.Compile(`[<]*(?:(?:sip\:)?)((\w+)@([^>]*))`)
	if err != nil {
		return nil, err
	}

	match := re.FindStringSubmatch(string(*sip))

	if !startsWithSip || len(match) != 4 {
		validation["sipValid"] = ValidationSipInvalidMessage
	}

	return &validation, nil
}
