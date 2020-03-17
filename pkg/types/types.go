package types

const (
	ButtonsType         = "buttons"
	RequestErrorMessage = "Looks like there was a hick-up in my though process. Could you please try again?"
)

// Replies
type Replies struct {
	Replies []interface{} `json:"replies"`
}

// TextMessage defines a response of type text message.
// Example:
// {
// 	"type": "text",
//	"delay": 2,
//	"content": "MY_TEXT",
// }
type TextMessage struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Delay   int    `json:"delay,omitempty"`
}

// {"cases":188433,"deaths":7500,"recovered":80873,"updated":1584454250388}
type Response struct {
	Cases     int `json:"cases"`
	Deaths    int `json:"deaths"`
	Recovered int `json:"recovered"`
	Updated   int `json:"updated"`
}

// {"country":"Germany","cases":8084,"todayCases":812,"deaths":20,"todayDeaths":3,"recovered":67,"critical":2}
type CountryResponse struct {
	Country     string `json:"country"`
	Cases       int `json:"cases"`
	Deaths      int `json:"deaths"`
	Recovered   int `json:"recovered"`
	Critical    int `json:"critical"`
	TodayCases  int `json:"todayCases"`
	TodayDeaths int `json:"todayDeaths"`
}
