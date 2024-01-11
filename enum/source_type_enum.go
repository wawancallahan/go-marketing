package enum

type SourceTypeEnum struct{}

func (e SourceTypeEnum) List() map[string]string {
	return map[string]string{
		"OFFLINE_EVENT":    "OFFLINE_EVENT",
		"EVENT_ONLINE":     "EVENT_ONLINE",
		"REGIONAL":         "REGIONAL",
		"OFFICIAL_WEBSITE": "OFFICIAL_WEBSITE",
	}
}
