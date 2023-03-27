package controller

import (
	
	"github.com/aiteung/musik"//link percobaan

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/sidiq200/faisal"
	"github.com/whatsauth/whatsauth"
	"github.com/sidiq200/app-profile/config"
)

var Dataser = "username"
var Dataprof = "profil"
var pend = "pendidikan"

func WsWhatsAuthQR(c *websocket.Conn) {
	whatsauth.RunSocket(c, config.PublicKey, config.Usertables[:], config.Ulbimariaconn)
}

func PostWhatsAuthRequest(c *fiber.Ctx) error {
	if string(c.Request().Host()) == config.Internalhost {
		var req whatsauth.WhatsauthRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		ntfbtn := whatsauth.RunModuleLegacy(req, config.PrivateKey, config.Usertables[:], config.Ulbimariaconn)
		return c.JSON(ntfbtn)
	} else {
		var ws whatsauth.WhatsauthStatus
		ws.Status = string(c.Request().Host())
		return c.JSON(ws)
	}

}

func GetMusik(c *fiber.Ctx) error {
	geturl := musik.GetIPaddress()
	return c.JSON(geturl)
}

func GetdataFaisal(c *fiber.Ctx) error{
	gedata := faisal.GetDataProfFromStatus("Active", config.MongoConn, Dataprof)
	return c.JSON(gedata)
}

func GetProfileByUsername(c *fiber.Ctx) error{
	name := c.Params("username")
	getdata := faisal.GetDataProfFromStatus(name, config.MongoConn, Dataser)
	return c.JSON(getdata)
}

// func GetPresensi(c *fiber.Ctx) error {
// 	presi := presensi.GetPresensiCurrentMonth(config.Ulbimongoconn)
// 	return c.JSON(presi)
// }

func InsertProfile(c *fiber.Ctx) error{
	model := c.Params(faisal.ListData)
	Data := faisal.InsertProfile(config.MongoConn,
			model.pendidikan,
			model.Bio,
			model.Username,
			model.Checkin,
			model.Biodata,
	)
	return c.JSON(Data)
}




// func GetdataFaisal(c *fiber.Ctx) error{
// 	getip := faisal.GetDataAllbyStatus("active", config.MongoConn, "data_compllain")
// 	return c.JSON(getip)
// }

//uploaded
