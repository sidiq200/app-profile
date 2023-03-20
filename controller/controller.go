package controller

import (
	
	"github.com/gofiber/fiber/v2"
	"github.com/aiteung/musik"
	"github.com/gofiber/websocket/v2"
	"github.com/sidiq200/faisal"
	"github.com/whatsauth/whatsauth"
	"github.com/sidiq200/app-profile/config"
)

var Dataser = "username"
var Dataprof = "profil"

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

// func GetHome(c *fiber.Ctx) error {
// 	getip := musik.GetIPaddress()
// 	return c.JSON(getip)
// }

func GetdataFaisal(c *fiber.Ctx) error{
	gedata := faisal.GetDataProfFromStatus("Active", config.MongoConn, Dataprof)
	return c.JSON(gedata)
}

func GetdataUsername(c *fiber.Ctx) error{
	name := c.Params("username")
	gedata := faisal.GetDataProfFromStatus(name, config.MongoConn, Dataser)
	return c.JSON(gedata)
}



// func GetdataFaisal(c *fiber.Ctx) error{
// 	getip := faisal.GetDataAllbyStatus("active", config.MongoConn, "data_compllain")
// 	return c.JSON(getip)
// }

//uploaded
