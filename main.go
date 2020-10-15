package main

func main()
{
	ws, id := slackConnect(token)

	for
	{
		msg, err := getMessage(ws)
		postMessage(ws, reply)
	}
}