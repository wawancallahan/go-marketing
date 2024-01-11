package enum

type ChannelEventEnum struct{}

type ChannelEventValue string

const (
	ONLINE  ChannelEventValue = "ONLINE"
	OFFLINE ChannelEventValue = "OFFLINE"
)

func (e ChannelEventEnum) List() map[string]string {
	return map[string]string{
		"ONLINE":  "ONLINE",
		"OFFLINE": "OFFLINE",
	}
}
