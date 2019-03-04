package utils

import (
	"encoding/json"
)

func GetMasterPlaybackUrl(playbackUriPageContents []byte) string {

	var masterPlaybackUrl string
	var result map[string]interface{}
	json.Unmarshal(playbackUriPageContents, &result)

	if int(result["statusCodeValue"].(float64)) == 200 {
		body := result["body"].(map[string]interface{})
		results := body["results"].(map[string]interface{})
		item := results["item"].(map[string]interface{})
		masterPlaybackUrl = item["playbackUrl"].(string)
	}

	return masterPlaybackUrl
}
