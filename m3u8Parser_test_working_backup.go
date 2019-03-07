package tests

import (
	"github.com/temp25/hdl2/utils"
	"io/ioutil"
	"log"
	"fmt"
	"testing"
	//"strings"
	"reflect"
)

func TestParseM3u8Content2(t *testing.T) {
	
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
            fmt.Println("infoKey : ", infoKey, "infoValue : ", infoValue, "\t infoValue len : ",len(infoValue)) //, "\tinfoValue has new_line", strings.Contains(infoValue, "\n")
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
fmt.Println("expectedFormats[\"STREAM-URL\"] len : ", len(expectedFormats["STREAM-URL"]))
expectedFormats["K-FORM"] = "167k"
expectedVideoFormats["hls-167"] = utils.CopyMap(expectedFormats)



    //fmt.Println("expectedVideoFormats len : ", len(expectedVideoFormats))
    //fmt.Println("parsedContent len : ", len(parsedContent))
    
    fmt.Printf("\n\nDumping expectedVideoFormats streamUrl\n")
    for pos, ch := range expectedFormats["STREAM-URL"] {
		fmt.Printf("pos : %d\tchar : |%c|\tascii_code : %d\n", pos, ch, int(ch))
	}
    
    fmt.Printf("\n\nDumping parsedContent streamUrl\n")
    for pos, ch := range parsedContent["hls-167"]["STREAM-URL"] {
		fmt.Printf("pos : %d\tchar : |%c|\tascii_code : %d\n", pos, ch, int(ch))
	}
    /*for pos, char := range  {
        fmt.Printf("character %c starts at byte position %d\n", char, pos)
    }*/
   
    
    /*for pos, ch := range "Hello, 世界" {
		fmt.Printf("%d: %c\n", pos, ch)
	}*/
    
	if !reflect.DeepEqual(expectedVideoFormats, parsedContent) {
		t.Error("Expected \n", expectedVideoFormats, "\n\n\nbut got \n", parsedContent)
	}

}