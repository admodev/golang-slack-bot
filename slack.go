package slack

url := fmt.Sprintf("https://slack.com/api/rtm.start?token=%s", token)
resp, err := http.Get(url)
//...
body, err := ioutil.ReadAll(resp.Body)
//...
var respObj responseRtmStart
err = json.Unmarshal(body, &respObj)