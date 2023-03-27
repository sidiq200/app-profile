package url

import (
	"github.com/sidiq200/app-profile/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.GetMusik) //syntax ujicoba
	page.Get("/faisal/:username", controller.GetProfileByUsername) //username
	page.Get("/faisal/:status", controller.GetdataFaisal) //status
	page.Get("/presensi", controller.GetPresensi) //Presensi

}
