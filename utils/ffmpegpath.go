package utils

import (
  "os"
)

func GetFfmpegPath() (string, bool) {
    
    path := FetchPairMatchingKey(os.Environ(), "ffmpeg")
    
    return path, func() bool {
        if path != "" {
            return true
        }
        return false
    }()
}