package enum

type MeasurementEventEnum struct{}

func (e MeasurementEventEnum) List() map[string]string {
	return map[string]string{
		"JUMLAH_PESERTA": "JUMLAH_PESERTA",
		"GOOGLE_FORM":    "GOOGLE_FORM",
	}
}
