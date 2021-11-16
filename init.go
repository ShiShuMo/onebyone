package onebyone

import (
	"encoding/json"
	"strings"

	"github.com/cdle/sillyGirl/core"
	"github.com/gin-gonic/gin"
)

func init() {

	core.Server.POST("/wx/push", func(c *gin.Context) {
		data, _ := c.GetRawData()
		type AutoGenerated struct {
			ptPin   string `json:"pt_pin"`
			message string `json:"message"`
		}
		ag := &AutoGenerated{}
		json.Unmarshal(data, ag)
		ptPin := ag.ptPin
		message := ag.message
		tp := "wx"
		core.Bucket("pin" + strings.ToUpper(tp)).Foreach(func(k, v []byte) error {
			if string(k) == ptPin && ptPin != "" {
				if push, ok := core.Pushs[tp]; ok {
					push(string(v), message, tp)
				}
			}
			return nil
		})
	})
}
