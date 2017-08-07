package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("authorized")
	viper.AddConfigPath("./etc")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	config := oauth1.NewConfig(viper.GetString("consumerKey"), viper.GetString("consumerSecret"))
	token := oauth1.NewToken(viper.GetString("token"), viper.GetString("tokenSecret"))
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	if resp != nil {
		fmt.Printf("resp: %s", resp)
	}
	fmt.Printf("tweets: %s", tweets)
}