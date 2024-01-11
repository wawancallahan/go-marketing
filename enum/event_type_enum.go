package enum

type EventTypeEnum struct{}

func (e EventTypeEnum) List() map[string]string {
	return map[string]string{
		"OFFLINE_EVENT": "OFFLINE_EVENT",
		"IG_LIVE":       "IG_LIVE",
		"WEBINAR":       "WEBINAR",
	}
}
