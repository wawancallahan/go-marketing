package enum

type ProductCategoryEnum struct{}

func (e ProductCategoryEnum) List() map[string]string {
	return map[string]string{
		"PINJAM_MODAL_INVENTORY": "PINJAM_MODAL_INVENTORY",
		"PINJAM_MODAL_USAHA":     "PINJAM_MODAL_USAHA",
		"PINJAM_MODAL_TOKO":      "PINJAM_MODAL_TOKO",
		"PINJAM_MODAL_KARYAWAN":  "PINJAM_MODAL_KARYAWAN",
	}
}
