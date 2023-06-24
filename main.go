package main

import(
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5501988572288-5471685189094-NLfHBD03G8Z0uRMFi2bfRy4d")
	os.Setenv("CHANNEL_TO", "C05ERV46MPA")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_TO")}

	fileArr := []string{"benisme.gif"}

	for i:= 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}