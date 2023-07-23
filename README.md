# session
An easy and useful cookie based authentication library for golang fiber framework. It uses JWT tokens to start a session. You can set the lifetime of token.
Let's start with example: 

Login :

<!-- Make your login controls here. Check if e-mail and password are correct. -->

<!-- Created new session client. -->
<span style="color: #569cd6;">session</span> := <span style="color: #9cdcfe;">session</span>.<span style="color: #dcdcaa;">New</span>()<br>
<!-- secretKey can be any string. It is used to make the JWT token understandable only by your server. -->
<!-- Send fiber.context, issuer as a string, and secretKey. Issuer is the data that we want to hide inside JWT token. It is the user ID in this example. -->
<span style="color: #9cdcfe;">err</span> := <span style="color: #dcdcaa;">session.Set</span>(c, <span style="color: #ce9178;">user.ID</span>, <span style="color: #ce9178;">secretKey</span>)


Now our user is logged in. Only you should do is taking auth cookie in every request and controll if it is valid. Our library will do this for you. 


Checking user validation :

session := session.New()
issuerString, err := session.Get(c, secretKey)
if err != nil {
  return 0, err
}

//The issuer string is id of the user that client claims to be. So, you can consider this is the user with that id, and do your staff. 


Logout :

logging out is also so easy. 

s := session.New()
s.Delete(c)
//The auth cookie will be deleted authomaticly.


Using flash messages : 
Flash messages are used for transfering message from a page to another page. We basicly write our message inside a cookie named flashmessage. And read it in another page.
By this method you can send successfull or failed messages to the mainpage for example.

Saving Flash Message : 

//We will just send the fiber.context and the message we want to show.
session.SetFlash(c, "Loged out Successfully")


Reading Flash Message :

//We will take the message which inside the cookie. And you can do whatever you want with it. I will send to the frontend. And i will handle there. 
alert := session.GetFlash(c)

c.Render("swiperslide_index", fiber.Map{
		"Alert":  alert,
})

