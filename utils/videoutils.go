package utils

import (
	"fmt"
	"log"
	"strings"
	"os"
	"text/tabwriter"
)

var retryCount = 0

func GetVideoFormats(videoUrl string, videoId string) map[string]map[string]string {

	var requestHeaders = map[string]string{
		"Hotstarauth":     GenerateHotstarAuth(),
		"X-Country-Code":  "IN",
		"X-Platform-Code": "JIO",
	}

	videoUrlContentBytes, err := Make_Get_Request(videoUrl, requestHeaders)

	if err != nil {
	    if retryCount+1 < 10 {
	        //retry again for fetching formats
	        retryCount++
	        return GetVideoFormats(videoUrl, videoId)
	    }
		log.Fatal(fmt.Errorf("%s. Error msg : %s", err, videoUrlContentBytes))
	}
	
	videoUrlContent := fmt.Sprintf("%s", videoUrlContentBytes)

	//fmt.Println(content)

	playbackUri, _/*metaDataMap*/, err := GetPlaybackUri(videoUrlContent, videoUrl, videoId)

	if err != nil {
	    if retryCount+1 < 10 {
	        //retry again for fetching formats
	        retryCount++
	        return GetVideoFormats(videoUrl, videoId)
	    }
		log.Fatal(fmt.Errorf("Error occurred : %s", err))
	}
	
	playbackUriContentBytes, err := Make_Get_Request(playbackUri, requestHeaders)

	if err != nil {
	    if retryCount+1 < 10 {
	        //retry again for fetching formats
	        retryCount++
	        return GetVideoFormats(videoUrl, videoId)
	    }
		log.Fatal(fmt.Errorf("Error occurred : %s", err))
	}
	
	masterPlaybackUrl := GetMasterPlaybackUrl(playbackUriContentBytes)
	
	var queryParams string
	masterPlaybackUrlQueryParam := strings.Split(masterPlaybackUrl, "?")
	
	if len(masterPlaybackUrlQueryParam) > 1 {
	    queryParams = masterPlaybackUrlQueryParam[1]
	}

	masterPlaybackPageContentsBytes, err := Make_Get_Request(masterPlaybackUrl, requestHeaders)

	if err != nil {
	    if retryCount+1 < 10 {
	        //retry again for fetching formats
	        retryCount++
	        return GetVideoFormats(videoUrl, videoId)
	    }
		log.Fatal(fmt.Errorf("Error occurred : %s", err))
	}

	//fmt.Printf("\nmasterPlaybackPageContentsBytes : \n%s\n", masterPlaybackPageContentsBytes)
	
	//return fmt.Sprintf("%s", masterPlaybackPageContentsBytes)
	
	return ParseM3u8Content(fmt.Sprintf("%s", masterPlaybackPageContentsBytes), masterPlaybackUrl, queryParams)

}


func ListVideoFormats(videoUrl string, videoId string) {
    //fmt.Println("Listing video formats for video id, ", videoId)
	videoFormats := GetVideoFormats(videoUrl, videoId)//, *formatFlag, *ffmpegPathFlag, *outputFileNameFlag, *metadataFlag)
	
	//NewWriter(io.Writer, minWidth, tabWidth, padding, padchar, flags)
    tw := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0) //tabwriter.Debug
    fmt.Fprintln(tw, "format code\textension\tresolution\tbandwidth\tcodec & frame rate\t")

	for formateId, formatInfo := range videoFormats {
		if frameRate, isFrameRatePresent := formatInfo["FRAME-RATE"]; isFrameRatePresent {
			fmt.Fprintf(tw, "%s\tmp4\t%s\t%s\t%s  %s fps\n", formateId, formatInfo["RESOLUTION"], formatInfo["K-FORM"], formatInfo["CODECS"], frameRate)
		} else {
			fmt.Fprintf(tw, "%s\tmp4\t%s\t%s\t%s\n", formateId, formatInfo["RESOLUTION"], formatInfo["K-FORM"], formatInfo["CODECS"])
		}

	}
	tw.Flush()
	os.Exit(0)
}

func DownloadVideo() {
    ffmpeg, isAvailable := GetFfmpegPath()
    if isAvailable {
        ffmpegPath :
    }
}