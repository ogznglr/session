package session

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Flash struct {
}

//Sets flash message as a cookie in client's browser.
//Flash messages has much more lesser expires time than sessions
//This is because flash messages are immediate data. The are consumed rapidly.

func (flash Flash) SetFlash(c *fiber.Ctx, message string) {
	cookie := fiber.Cookie{
		Name:     "flashmessage",
		Value:    message,
		Expires:  time.Now().Add(time.Second * 15),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return
}

//Gets the cookie with key of flashmessage.
//Makes map with message and a boolean.
//You can use the boolean in front end while showing an alert. if there is message, then show the alert.

func (flash Flash) GetFlash(c *fiber.Ctx) *fiber.Map {
	var isMessage bool
	message := c.Cookies("flashmessage")
	if message != "" {
		isMessage = true
	}
	//This process deletes the flashmessage cookie that client has.
	deleteCookie := fiber.Cookie{
		Name:    "flashmessage",
		Expires: time.Now().Add(-10 * time.Minute),
	}
	c.Cookie(&deleteCookie)

	data := &fiber.Map{
		"message":  message,
		"is-alert": isMessage,
	}
	return data
}
