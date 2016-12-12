package main

import (
	"log"
	"mail/models"
	"mail/modules"
	_ "mail/routers"
	"time"

	"github.com/astaxie/beego"
)

func main() {
	go GetData()
	go func() {
		sock := new(service.Sock)
		sock.Start()
	}()
	beego.Run()
}

func GetData() {
	// var l sync.Mutex
	for {
		t := time.Now().Unix()
		if t%10 == 0 {

			l := new(models.List)
			emails := l.Alive()
			for _, v := range emails {
				// fmt.Println(v.Id)
				// fmt.Println(v.Title)
				go send(v)
				time.Sleep(5000 * time.Millisecond)
			}

		}
		time.Sleep(1000 * time.Millisecond)
	}
}

func send(v models.Email) {
	serv := new(service.Service)
	issend := serv.SendToMail("mail.guodu.com.hk", "Kuaiex@guodu.com.hk", "1550520500@qq.com", "Kuaiex@2556", v.Title, v.Content, "html")
	list := new(models.List)

	if issend {
		go func() {
			log.Println("insert chan")
			service.Chant <- v.Id
		}()
		list.Tagsend(v.Id, 1)
	} else {
		list.Tagsend(v.Id, -1)
	}
}
