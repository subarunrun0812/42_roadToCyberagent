// reference: https://github.com/CyberAgentHack/42tokyo-road-to-dojo-go/blob/master/pkg/server/handler/setting.go

package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	// ガチャ1回あたりのコイン消費量
	GachaCoinConsumption = 100
)

// HandleSettingGet ゲーム設定情報取得処理
func HandleSettingGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, err := json.Marshal(&settingGetResponse{
			GachaCoinConsumption: GachaCoinConsumption,
		})
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Write(data)
	}
}

type settingGetResponse struct {
	GachaCoinConsumption int32 `json:"gachaCoinConsumption"`
}