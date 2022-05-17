package params

import "latihan/models"

func AddBiodataRequest(id, name, address, pekerjaan, alasan string) models.Biodata {
	return models.Biodata{
		Id:        id,
		Name:      name,
		Address:   address,
		Pekerjaan: pekerjaan,
		Alasan:    alasan,
	}
}
