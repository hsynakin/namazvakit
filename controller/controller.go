package controller

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hsynakin/namazvakit/Models"
)

var xmlVakit Models.Envelope

func Controller(c *gin.Context) {

	iID := c.Params.ByName("id")

	url := "http://namazvakti.diyanet.gov.tr/wsNamazVakti.svc?singleWsdl="

	payload := strings.NewReader("<Envelope xmlns=\"http://schemas.xmlsoap.org/soap/envelope/\">\n\t<Body>\n\t\t<GunlukNamazVakti xmlns=\"http://tempuri.org/\">\n\t\t\t<IlceID>" + iID + "</IlceID>\n\t\t\t<username>namazuser</username>\n\t\t\t<password>NamVak!14</password>\n\t\t</GunlukNamazVakti>\n\t</Body>\n</Envelope>")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "text/xml; charset=\"UTF-8\"")
	req.Header.Add("SOAPAction", "http://tempuri.org/IwsNamazVakti/GunlukNamazVakti")
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Postman-Token", "e466ece3-fe9d-4fa7-a434-0ba7928b71e5")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	xml.Unmarshal(body, &xmlVakit)

	c.JSON(http.StatusOK, xmlVakit.Usr.USS.User.Users)

}
