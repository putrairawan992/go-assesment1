package repositories

import (
	"latihan/models"
	"latihan/params"
)

func GetBiodataById(id string) models.Biodata {
	var biodata models.Biodata

	biodatas := GetBiodata()

	for _, val := range biodatas {
		if val.Id == id {
			biodata = val
			break
		}
	}
	return biodata
}

func GetBiodata() []models.Biodata {
	var biodata []models.Biodata

	initBio := params.AddBiodataRequest("1", "Putra Irawan", "Jalan Letnan Murod Lrg Akasia", "Front End Engginer", "Belajar Golang")
	biodata = append(biodata, initBio)

	initBio = params.AddBiodataRequest("2", "Istiqomah", "Jalan Letnan Murod Lrg Akasia", "Istri", "Memasak. Mengurus Rumah tangga")
	biodata = append(biodata, initBio)

	initBio = params.AddBiodataRequest("3", "Albirru Ashauqi Hakia", "Jalan Letnan Murod Lrg Akasia", "AKPOL", "Akademi Kepolisian")
	biodata = append(biodata, initBio)

	return biodata

}
