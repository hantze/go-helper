package formatter

import (
	"github.com/pkg/errors"
	"html"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

// GetQueryParams ...
func GetQueryParams(c *gin.Context, param []string, term string) (map[string]string, error) {
	results := make(map[string]string)
	for _, v := range param {
		temp := GetURLParam(c, v)
		switch term {
		case "required":
			if temp != nil && temp != "" {
				results[v] = html.EscapeString(temp.(string))
			} else {
				return nil, errors.New("failed")
			}
		default:
			if temp == nil || temp == "" {
				results[v] = ""
			} else {
				results[v] = html.EscapeString(temp.(string))
			}
		}
	}
	return results, nil
}

// GetQueryParam ...
func GetQueryParam(c *gin.Context, param string) string {
	var result string
	temp := GetURLParam(c, param)
	if temp != nil {
		result = html.EscapeString(temp.(string))
	} else {
		result = ""
	}
	return result
}

// CleanString ...
func CleanString(param *string) string {
	if param == nil || *param == "" {
		return ""
	}
	var replacer = strings.NewReplacer("exec", "", "--", "", "DROP", "", "EXEC", "", "drop", "", "'", "", ";", "")
	cleanParam := replacer.Replace(*param)
	p := bluemonday.UGCPolicy()
	return p.Sanitize(cleanParam)
}
