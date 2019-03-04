package utils

import (
	"fmt"
	"log"
)

func GetVideoFormats(videoUrl string, videoId string) string {

	var requestHeaders = map[string]string{
		"Hotstarauth":     GenerateHotstarAuth(),
		"X-Country-Code":  "IN",
		"X-Platform-Code": "JIO",
	}

	videoUrlContentBytes, err := Make_Get_Request(videoUrl, requestHeaders)

	if err != nil {
		log.Fatal(fmt.Errorf("%s. Error msg : %s", err, videoUrlContentBytes))
	}

	videoUrlContent := fmt.Sprintf("%s", videoUrlContentBytes)

	//fmt.Println(content)

	playbackUri, _/*metaDataMap*/, err := GetPlaybackUri(videoUrlContent, videoUrl, videoId)

	if err != nil {
		log.Fatal(fmt.Errorf("Error occurred : %s", err))
	}

	playbackUriContentBytes, err := Make_Get_Request(playbackUri, requestHeaders)

	if err != nil {
		log.Fatal(fmt.Errorf("Error occurred : %s", err))
	}

	masterPlaybackUrl := GetMasterPlaybackUrl(playbackUriContentBytes)

	masterPlaybackPageContentsBytes, err := Make_Get_Request(masterPlaybackUrl, requestHeaders)

	if err != nil {
		log.Fatal(fmt.Errorf("Error occurred : %s", err))
	}

	//fmt.Printf("\nmasterPlaybackPageContentsBytes : \n%s\n", masterPlaybackPageContentsBytes)
	
	return fmt.Sprintf("%s", masterPlaybackPageContentsBytes)

}
