package Models

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"Envelope"`
	Usr     Body     `xml:"Body"`
}
type Body struct {
	XMLName xml.Name   `xml:"Body"`
	USS     NVResponse `xml:"GunlukNamazVaktiResponse"`
}
type NVResponse struct {
	XMLName xml.Name `xml:"GunlukNamazVaktiResponse"`
	User    NVresult `xml:"GunlukNamazVaktiResult"`
}

type NVresult struct {
	XMLName xml.Name   `xml:"GunlukNamazVaktiResult"`
	Users   NamazVakit `xml:"NamazVakti"`
}
type NamazVakit struct {
	XMLName                xml.Name `xml:"NamazVakti" json:"-"`
	AyinSekliURL           string   `xml:"AyinSekliURL"`
	Gunes                  string   `xml:"Gunes"`
	GunesDogus             string   `xml:"GunesDogus"`
	GunesBatis             string   `xml:"GunesBatis"`
	Ogle                   string   `xml:"Ogle"`
	Ikindi                 string   `xml:"Ikindi"`
	Aksam                  string   `xml:"Aksam"`
	Imsak                  string   `xml:"Imsak"`
	Yatsi                  string   `xml:"Yatsi"`
	MiladiTarihUzun        string   `xml:"MiladiTarihUzun"`
	MiladiTarihUzunIso8601 string   `xml:"MiladiTarihUzunIso8601"`
	MiladiTarihKisa        string   `xml:"MiladiTarihKisa"`
	MiladiTarihKisaIso8601 string   `xml:"MiladiTarihKisaIso8601"`
	KibleSaati             string   `xml:"KibleSaati"`
	HicriTarihUzun         string   `xml:"HicriTarihUzun"`
	HicriTarihUzunIso8601  string   `xml:"HicriTarihUzunIso8601"`
	HicriTarihKisa         string   `xml:"HicriTarihKisa"`
	HicriTarihKisaIso8601  string   `xml:"HicriTarihKisaIso8601"`
}
