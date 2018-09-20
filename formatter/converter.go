package formatter

import (
	"strconv"
	"strings"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

// GetURLParam ...
func GetURLParam(c *gin.Context, key string) interface{} {
	urlQuery := c.Request.URL.Query()
	if len(urlQuery[key]) > 0 {
		return urlQuery[key][0]
	}
	return nil
}

// StringToIntegerValidator ...
func StringToIntegerValidator(param string, field string) (int, error) {
	val, err := strconv.Atoi(param)
	if err != nil {

		return 0, err
	}
	return val, nil
}

// StringToInteger ...
func StringToInteger(param string) int {
	val, _ := strconv.Atoi(param)
	return val
}

// StringToFloat ...
func StringToFloat(param string) float64 {
	val, _ := strconv.ParseFloat(param, 10)
	return val
}

// IntegerToString ...
func IntegerToString(param int) string {
	val := strconv.Itoa(param)
	return val
}

// ReplaceBrackets ...
func ReplaceBrackets(param string) string {
	var replacer = strings.NewReplacer("{", "", "}", "")
	return replacer.Replace(param)
}

// GetStringOrEmpty ...
func GetStringOrEmpty(param interface{}) string {
	if param != nil {
		return param.(string)
	}
	return ""
}

// LeftPad2Len ...
func LeftPad2Len(str string, padStr string, overallLen int) string {
	var padCountInt = 1 + ((overallLen - len(padStr)) / len(padStr))
	var retStr = strings.Repeat(padStr, padCountInt) + str
	return retStr[(len(retStr) - overallLen):]
}

// GenerateStringUUID ...
func GenerateStringUUID() string {
	return fmt.Sprintf("%s", uuid.NewV4())
}
