package main

import (
	"flag"
	"fmt"
	"github.com/temp25/hdl2/utils"
	"log"
	"net/url"
	"os"
	"strings"
)

//flag descriptions
var helpFlagDesc = "Prints this help and exit"
var listFormatsFlagDesc = "List available video formats for given url"
var formatFlagDesc = "Video format to download video in specified resolution"
var ffmpegPathFlagDesc = "Location of the ffmpeg binary(absolute path)"
var metadataFlagDesc = "Add metadata to the video file"
var outputFileNameFlagDesc = "Output file name"

//flag declarations
var helpFlag = flag.Bool("help", false, helpFlagDesc)
var listFormatsFlag = flag.Bool("list", false, listFormatsFlagDesc)
var formatFlag = flag.String("format", "", formatFlagDesc)
var ffmpegPathFlag = flag.String("ffmpeg-location", "", ffmpegPathFlagDesc)
var metadataFlag = flag.Bool("add-metadata", false, metadataFlagDesc)
var outputFileNameFlag = flag.String("output", "", outputFileNameFlagDesc)

func init() {
	//shorthand notations
	flag.BoolVar(helpFlag, "h", false, helpFlagDesc)
	flag.BoolVar(listFormatsFlag, "l", false, listFormatsFlagDesc)
	flag.StringVar(formatFlag, "f", "", formatFlagDesc)
	flag.BoolVar(metadataFlag, "m", false, metadataFlagDesc)
	flag.StringVar(outputFileNameFlag, "o", "", outputFileNameFlagDesc)

	//custom flag usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Usage: %s [OPTIONS] URL\n\n", os.Args[0])
		fmt.Println("Options:")
		fmt.Fprintf(os.Stdout, "-h, --help\t\t%s\n", helpFlagDesc)
		fmt.Fprintf(os.Stdout, "-l, --list\t\t%s\n", listFormatsFlagDesc)
		fmt.Fprintf(os.Stdout, "-f, --format\t\t%s\n", formatFlagDesc)
		fmt.Fprintf(os.Stdout, "--ffmpeg-location\t%s\n", ffmpegPathFlagDesc)
		fmt.Fprintf(os.Stdout, "-m, --add-metadata\t%s\n", metadataFlagDesc)
		fmt.Fprintf(os.Stdout, "-o, --output\t\t%s\n", outputFileNameFlagDesc)
		os.Exit(0)
		//flag.PrintDefaults()
	}
}

func main() {
	/*
	   contentBytes, err := utils.Make_Get_Request("https://hssouthsp-vh.akamaihd.net/i/videos/vijay_hd/chinnathambi/149/master_,106,180,400,800,1300,2000,3000,4500,kbps.mp4.csmil/master.m3u8?hdnea=st=1551575624~exp=1551577424~acl=/*~hmac=3d89f2aab02315ee100156209746e0e9f3bc70b0b52c17573300b5caa517cfd6")//"http://www.google.com")
	   if err != nil {
	       log.Fatal(fmt.Errorf("%s. Error msg : %s", err, contentBytes))
	   }

	   content := fmt.Sprintf("%s", contentBytes)

	   fmt.Println(content)
	*/

	flag.Parse()
	flagCount := len(flag.Args())
	if *helpFlag {
		flag.Usage()
	} else if flagCount == 0 {
		fmt.Println("Must provide atleast one url at end")
		flag.Usage()
		os.Exit(-1)
	} else if flagCount > 1 {
		fmt.Println("Url must be provided at end before options")
		flag.Usage()
		os.Exit(-1)
	} else if videoUrl := flag.Args()[0]; videoUrl != "" {
		parsedUrl, err := url.Parse(videoUrl)
		if err != nil {
			log.Fatal(err)
		}
		switch parsedUrl.Scheme {
		case "":
			fmt.Println("Replacing empty url scheme with https")
			parsedUrl.Scheme = "https"
		case "https":
			//do nothing
		case "http":
			fmt.Println("Replacing http url scheme with https")
			parsedUrl.Scheme = "https"
		default:
			fmt.Println("Invalid url scheme please enter valid one")
			os.Exit(-1)
		}

		videoUrl = fmt.Sprintf("%v", parsedUrl)

		fmt.Println("Parsed video url is", parsedUrl)

		isValidUrl, videoId := utils.IsValidHotstarUrl(videoUrl)
		if isValidUrl {
			if *listFormatsFlag {
				//list video formats
				fmt.Println("Listing video formats for video id, ", videoId)
				/*videoFormats := */utils.GetVideoFormats(videoUrl, videoId)//, *formatFlag, *ffmpegPathFlag, *outputFileNameFlag, *metadataFlag)
				//fmt.Printf("\nmasterPlaybackPageContentsBytes : \n%s\n\n", videoFormats)
				//utils.ParseM3u8Content(videoFormats)
			} else if *formatFlag != "" {
				if !strings.HasPrefix(*formatFlag, "hls-") {
					fmt.Println("Invalid format specified")
					os.Exit(-1)
				} else {
					//TODO: add code for download
				}
			} else {
				//TODO: Check for other flags if associated with url if any
			}
		} else {
			fmt.Println("Invalid hotstar url. Please enter a valid one")
			os.Exit(-1)
		}

	} else {
		fmt.Println("Invalid args specified")
		flag.Usage()
	}

}
