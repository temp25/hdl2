package tests

import (
	"github.com/temp25/hdl2/utils"
	"io/ioutil"
	"log"
	"fmt"
	"testing"
)

func TestParseM3u8Content1(t *testing.T) {
	
	playbackUrl := "https://hssouthsp-vh.akamaihd.net/i/videos/vijay_hd/chinnathambi/149/master_,106,180,400,800,1300,2000,3000,4500,kbps.mp4.csmil/master.m3u8?hdnea=st=1551575624~exp=1551577424~acl=/*~hmac=3d89f2aab02315ee100156209746e0e9f3bc70b0b52c17573300b5caa517cfd6"
    playbackUrlData := "hdnea=st=1551575624~exp=1551577424~acl=/*~hmac=3d89f2aab02315ee100156209746e0e9f3bc70b0b52c17573300b5caa517cfd6"
    
    m3u8Content, err := ioutil.ReadFile("m3u8Content1.txt")
	if err != nil {
		log.Fatal(err)
	}
    
    parsedContent := utils.ParseM3u8Content(fmt.Sprintf("%s", m3u8Content), playbackUrl, playbackUrlData)
    
    /*for key, value := range parsedContent {
        fmt.Println("\n\nKey : ", key)
        fmt.Println("Iterating values....")
        for infoKey, infoValue := range value {
            fmt.Println("infoKey : ", infoKey, "infoValue : ", infoValue)
        }
    }*/
   
   var expectedVideoFormats = make(map[string]map[string]string)
   var expectedFormats map[string]string

expectedFormats = make(map[string]string)
expectedFormats["PROGRAM-ID"] = "1"
expectedFormats["BANDWIDTH"] = "167271"
expectedFormats["RESOLUTION"] = "320x180"
expectedFormats["CODECS"] = "\"avc1.66.30, mp4a.40.2\""
expectedFormats["CLOSED-CAPTIONS"] = "NONE"
expectedFormats["STREAM-URL"] = "https://hssouthsp-vh.akamaihd.net/i/videos/vijay_hd/chinnathambi/149/master_,106,180,400,800,1300,2000,3000,4500,kbps.mp4.csmil/index_0_av.m3u8?null=0&id=AgCdLeTnMSxxookre1yyOZrUVjGsAjTrI2jaZKKjKzRKekEWQ81I2j3HSzMs2ZZcxJTgLWz%2f4cRk1A%3d%3d&hdnea=st=1551575624~exp=1551577424~acl=/*~hmac=3d89f2aab02315ee100156209746e0e9f3bc70b0b52c17573300b5caa517cfd6"
expectedFormats["K-FORM"] = "167k"
expectedVideoFormats["hls-167"] = utils.CopyMap(expectedFormats)

expectedFormats = make(map[string]string)
expectedFormats["PROGRAM-ID"] = "1"
expectedFormats["BANDWIDTH"] = "327344"
expectedFormats["RESOLUTION"] = "320x180"
expectedFormats["CODECS"] = "\"avc1.66.30, mp4a.40.2\""
expectedFormats["CLOSED-CAPTIONS"] = "NONE"
expectedFormats["STREAM-URL"] = "https://hssouthsp-vh.akamaihd.net/i/videos/vijay_hd/chinnathambi/149/master_,106,180,400,800,1300,2000,3000,4500,kbps.mp4.csmil/index_1_av.m3u8?null=0&id=AgCdLeTnMSxxookre1yyOZrUVjGsAjTrI2jaZKKjKzRKekEWQ81I2j3HSzMs2ZZcxJTgLWz%2f4cRk1A%3d%3d&hdnea=st=1551575624~exp=1551577424~acl=/*~hmac=3d89f2aab02315ee100156209746e0e9f3bc70b0b52c17573300b5caa517cfd6"
expectedFormats["K-FORM"] = "327k"
expectedVideoFormats["hls-327"] = utils.CopyMap(expectedFormats)

expectedFormats = make(map[string]string)
expectedFormats["PROGRAM-ID"] = "1"
expectedFormats["BANDWIDTH"] = "552127"
expectedFormats["RESOLUTION"] = "416x234"
expectedFormats["CODECS"] = "\"avc1.66.30, mp4a.40.2\""
expectedFormats["CLOSED-CAPTIONS"] = "NONE"
expectedFormats["STREAM-URL"] = "https://hssouthsp-vh.akamaihd.net/i/videos/vijay_hd/chinnathambi/149/master_,106,180,400,800,1300,2000,3000,4500,kbps.mp4.csmil/index_2_av.m3u8?null=0&id=AgCdLeTnMSxxookre1yyOZrUVjGsAjTrI2jaZKKjKzRKekEWQ81I2j3HSzMs2ZZcxJTgLWz%2f4cRk1A%3d%3d&hdnea=st=1551575624~exp=1551577424~acl=/*~hmac=3d89f2aab02315ee100156209746e0e9f3bc70b0b52c17573300b5caa517cfd6"
expectedFormats["K-FORM"] = "552k"
expectedVideoFormats["hls-552"] = utils.CopyMap(expectedFormats)

expectedFormats = make(map[string]string)
expectedFormats["PROGRAM-ID"] = "1"
expectedFormats["BANDWIDTH"] = "960823"
expectedFormats["RESOLUTION"] = "640x360"
expectedFormats["CODECS"] = "\"avc1.66.30, mp4a.40.2\""
expectedFormats["CLOSED-CAPTIONS"] = "NONE"
expectedFormats["STREAM-URL"] = "https://hssouthsp-vh.akamaihd.net/i/videos/vijay_hd/chinnathambi/149/master_,106,180,400,800,1300,2000,3000,4500,kbps.mp4.csmil/index_3_av.m3u8?null=0&id=AgCdLeTnMSxxookre1yyOZrUVjGsAjTrI2jaZKKjKzRKekEWQ81I2j3HSzMs2ZZcxJTgLWz%2f4cRk1A%3d%3d&hdnea=st=1551575624~exp=1551577424~acl=/*~hmac=3d89f2aab02315ee100156209746e0e9f3bc70b0b52c17573300b5caa517cfd6"
expectedFormats["K-FORM"] = "960k"
expectedVideoFormats["hls-960"] = utils.CopyMap(expectedFormats)

expectedFormats = make(map[string]string)
expectedFormats["PROGRAM-ID"] = "1"
expectedFormats["BANDWIDTH"] = "1472714"
expectedFormats["RESOLUTION"] = "720x404"
expectedFormats["CODECS"] = "\"avc1.66.30, mp4a.40.2\""
expectedFormats["CLOSED-CAPTIONS"] = "NONE"
expectedFormats["STREAM-URL"] = "https://hssouthsp-vh.akamaihd.net/i/videos/vijay_hd/chinnathambi/149/master_,106,180,400,800,1300,2000,3000,4500,kbps.mp4.csmil/index_4_av.m3u8?null=0&id=AgCdLeTnMSxxookre1yyOZrUVjGsAjTrI2jaZKKjKzRKekEWQ81I2j3HSzMs2ZZcxJTgLWz%2f4cRk1A%3d%3d&hdnea=st=1551575624~exp=1551577424~acl=/*~hmac=3d89f2aab02315ee100156209746e0e9f3bc70b0b52c17573300b5caa517cfd6"
expectedFormats["K-FORM"] = "1472k"
expectedVideoFormats["hls-1472"] = utils.CopyMap(expectedFormats)

expectedFormats = make(map[string]string)
expectedFormats["PROGRAM-ID"] = "1"
expectedFormats["BANDWIDTH"] = "2188953"
expectedFormats["RESOLUTION"] = "1280x720"
expectedFormats["CODECS"] = "\"avc1.66.30, mp4a.40.2\""
expectedFormats["CLOSED-CAPTIONS"] = "NONE"
expectedFormats["STREAM-URL"] = "https://hssouthsp-vh.akamaihd.net/i/videos/vijay_hd/chinnathambi/149/master_,106,180,400,800,1300,2000,3000,4500,kbps.mp4.csmil/index_5_av.m3u8?null=0&id=AgCdLeTnMSxxookre1yyOZrUVjGsAjTrI2jaZKKjKzRKekEWQ81I2j3HSzMs2ZZcxJTgLWz%2f4cRk1A%3d%3d&hdnea=st=1551575624~exp=1551577424~acl=/*~hmac=3d89f2aab02315ee100156209746e0e9f3bc70b0b52c17573300b5caa517cfd6"
expectedFormats["K-FORM"] = "2188k"
expectedVideoFormats["hls-2188"] = utils.CopyMap(expectedFormats)

    
	if parsedContent == nil {
		t.Error("Expected originalMap  but got clonedMap")
	}


}