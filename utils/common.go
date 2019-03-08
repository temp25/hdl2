package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CopyMap(m map[string]string) map[string]string {
	cp := make(map[string]string)
	for k, v := range m {
		cp[k] = v
	}
	return cp
}

func PadZeroRight(num int64) int64 {
	tmp := fmt.Sprintf("%-13d", num)
	tmp = strings.Replace(tmp, " ", "0", -1)
	paddedNum, err := strconv.ParseInt(tmp, 10, 64)
	if err != nil {
		panic(err)
	}
	return paddedNum
}

func CountDigits(i int64) (count int64) {
	for i != 0 {
		i /= 10
		count = count + 1
	}
	return count
}

func GetDateStr(timeFloat64 float64) string {
	timeMillis := int64(timeFloat64)
	if CountDigits(timeMillis) == 13 {
		timeMillis = PadZeroRight(timeMillis)
	}
	location, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}
	return time.Unix(0, timeMillis*int64(time.Millisecond)).In(location).String()
}

func FetchPairMatchingKey(array []string, item string) string {
	
	for _, value := range array {
		splitValues := strings.SplitN(value, "=", 2)
		if(len(splitValues) > 1){
		    
		    if strings.EqualFold(splitValues[0], "PATH") { //case insensitive search
		    
		        pathValues := strings.FieldsFunc(splitValues[1], SplitSemiColonOrColon)
		        
		        //Iterate PATH env values to see if a path exists
		        for _, pathValue := range pathValues {
		            if strings.Contains(pathValue, item) {
		                return pathValue
		            }
	            }

		        
		    }else if strings.Contains(splitValues[0], item) {
				return splitValues[1]
			}
		}
	}
	
    return ""
}

func SplitSemiColonOrColon(r rune) bool {
	return r == ';' || r == ':'
}