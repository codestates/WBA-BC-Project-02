package validator

import (
	"errors"
	"github.com/codestates/WBA-BC-Project-02/common/enum"
	"strings"
)

// CheckBlank string trim 한 값이 "" 인 경우 BadRequest
func CheckBlank(STR string) error {
	s := strings.Trim(STR, enum.SpaceSTR)
	if s == enum.BlankSTR {
		return errors.New("공백 문자입니다")
	}
	return nil
}
