package req

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Streams struct {
	Streams []Stream `json:"streams"`
}
type Stream struct {
	Chan        Channel `json:"channel"`
	Viewers     int     `json:"viewers"`      // Amount of viewers on a Live stream
	VideoHeight int     `json:"video_height"` // Video height, ex. 720 -> 720p highest quality
}
type Channel struct {
	DisplayName     string `json:"display_name"`         // Stream name
	Game            string `json:"game"`                 // Stream currently played game
	Mature          bool   `json:"mature"`               // If stream is for mature audiences
	Title           string `json:"status"`               // Stream title
	Delay           int    `json:"delay"`                // Stream delay in seconds
	CreationDate    string `json:"created_at"`           // When the twitch account was created
	Url             string `json:"url"`                  // Stream URL, easily linkable into streamlink after
	Language        string `json:"language"`             // Language of the stream
	BroadcasterLang string `json:"broadcaster_language"` // Language of the broadcast, tends to differ in foreign coverages of esports
	Description     string `json:"description"`          // Channel description
}

const follow_url = "https://api.twitch.tv/kraken/streams/followed"
const get_user_url = "https://api.twitch.tv/kraken/user"
const get = "GET"

// Return list of LIVE streamers on follower list
func Live() Streams {
	url := follow_url
	reqT := get
	resp, err := Send(GenReq(&reqT, &url, nil))
	if err != nil {
		log.Fatal("Couldn't get followed list, quitting", err)
	}
	resp_data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could not parse response, quitting", err)
	}
	live := Streams{}
	err = json.Unmarshal([]byte(resp_data), &live)

	if err != nil {
		log.Fatal("Couldn't unmarshal json...", err)
	}
	return live
}


type User struct {
	Id     string `json:"_id"`         // Stream name
}
// Return list of all streamers user follows
func All() {
	//TODO
	//First we get the user id, then we get the follows for that channel
	//GET https://api.twitch.tv/kraken/users/<user ID>/follows/channels
	url := get_user_url
	reqT := get
	resp, err := Send(GenReq(&reqT, &url, nil))
	resp_data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could not parse all response")
	}
	user := User{}
	err = json.Unmarshal([]byte(resp_data), &user)

	url = "https://api.twitch.tv/kraken/users/" + user.Id + "/follows/channels"
	fmt.Println(url)
	resp, err = Send(GenReq(&reqT, &url, nil))
	if err != nil {
	    log.Fatal("Could not send request")
	}
	if resp.StatusCode > 299{
	    log.Fatal("Could not get channels", resp.Status)
	}
	resp_data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could not parse all response")
	}
	//TODO: Change into separate Struct
	all := Streams{}
	err = json.Unmarshal([]byte(resp_data), &all)
	if err != nil {
	    log.Fatalln("Couldn't unmarshal response", string(resp_data))
	}

}
