package routers

import (
	"StudentManage/controller"
	"net/http"
)

func Router() {
	http.HandleFunc("/create", controller.Create)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/update", controller.Updata)
	http.HandleFunc("/queryone/", controller.QueryOne)
	http.HandleFunc("/queryall", controller.QueryAll)
}
