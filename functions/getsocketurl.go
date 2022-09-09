package lib

import (
	"bitbucket.org/taubyte/go-sdk/event"
	"bitbucket.org/taubyte/go-sdk/pubsub"
)

func getChannel(h event.HttpEvent) string {
	room, _ := h.Query().Get("room")

	channelName := "someChannel"
	if len(room) > 0 {
		channelName += "/" + room
	}

	return channelName
}

//export getsocketurl
func getsocketurl(e event.Event) uint32 {
	h, err := e.HTTP()
	if err != nil {
		return 1
	}

	url, err := func() (url url.URL, err error) {
		channel, err := pubsub.Channel(getChannel(h))
		if err != nil {
			return
		}

		return channel.WebSocket().Url()
	}()
	if err != nil {
		h, err := e.HTTP()
		if err != nil {
			return 1
		}

		h.Write([]byte(err.Error()))
		return 1
	}

	h.Write([]byte(url.Path))

	return 0
}
