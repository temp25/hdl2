package utils

import (
    "fmt"
    "strings"
    "regexp"
    "strconv"
)

func ParseM3u8Content(m3u8Content string,playbackUrl string, playbackUrlData string) map[string]map[string]string {
    
    var m3u8Info map[string]string
    var urlFormats = make(map[string]map[string]string)
    for _, line := range strings.Split(m3u8Content, "\n") {
        //fmt.Printf("\nindex %d, line %s\n",index, line)
        if strings.HasPrefix(line, "#EXT-X-STREAM-INF:") {
            //fmt.Println("inside if")
            if m3u8Info == nil {
                //fmt.Println("infoArray array is null initializing it")
                m3u8Info = make(map[string]string)
            }
            m3u8InfoCsv := strings.Replace(line, "#EXT-X-STREAM-INF:", "", -1)
            //fmt.Println("m3u8InfoCsv : ",m3u8InfoCsv)
            m3u8InfoRegex := regexp.MustCompile(`([\w\-]+)\=([\w\-]+|"[^"]*")`)
            
            for _, info := range m3u8InfoRegex.FindAllStringSubmatch(m3u8InfoCsv, -1) {
                //fmt.Printf("index : %d,\tm3u8Info : %s__%s\n", index, m3u8Info[1], m3u8Info[2]);
                m3u8Info[info[1]] = info[2]
            }
        } else if strings.Contains(line, ".m3u8") {
            //fmt.Println("inside else")
            //fmt.Println("infoArray len : ", len(infoArray));
            
            if m3u8Info != nil {
                
                averageBandwidthOrBandwidth := func() int {
    				var bw string
    				if m3u8Info["AVERAGE-BANDWIDTH"] != "" {
    					bw = m3u8Info["AVERAGE-BANDWIDTH"]
    				} else {
    					bw = m3u8Info["BANDWIDTH"]
    				}
    				var bwInt, _ = strconv.Atoi(bw)
    				return bwInt
    			}()
                
                kFactor := averageBandwidthOrBandwidth / 1000
                kForm := fmt.Sprintf("%dk", kFactor)
                
                m3u8Info["K-FORM"] = kForm
                m3u8Info["STREAM-URL"] = func() string {
                    streamUrl := func() string {
                        if strings.HasPrefix(line, "http") {
                            return line
                        }
                        
                        return strings.Replace(line, "master.m3u8", playbackUrl, -1)
                    }()
                    
                    if !strings.Contains(streamUrl, "~acl=/*~hmac") {
                        if !strings.Contains(streamUrl, "?") {
                            streamUrl += "?"
                        }
                        streamUrl += ("&" + playbackUrlData)
                    }
                    
                    return streamUrl
                }()
                
                //for key, value := range m3u8Info {
                //    fmt.Println("Key : ", key, "\t Value : ", value)
                //}
                
                urlFormats[fmt.Sprintf("hls-%d", kFactor)] = CopyMap(m3u8Info)
                
                //Reset m3u8InfoArray for next layer
                m3u8Info = nil
            }

            
        } else {
            //do nothing
        }
    }
    
    /*for key, value := range urlFormats {
        fmt.Println("\n\nKey : ", key)
        fmt.Println("Iterating values....")
        for infoKey, infoValue := range value {
            fmt.Println("infoKey : ", infoKey, "infoValue : ", infoValue)
        }
    }*/
    
    return urlFormats
}