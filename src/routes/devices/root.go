package devices

import (
	"applicant-backend/src/controller/devices"
	"fmt"
	"net/http"
)

type Route struct {
	controller    *devices.DeviceController
	path          string
	MethodHandler MethodHandler
}

func (r *Route) RegisterToPath() {
	fmt.Println("registering2:", r.path, r.MethodHandler)
	http.HandleFunc(r.path, r.MethodHandler.getHandler())
}

func RootRoute(controller *devices.DeviceController) Route {
	handler := MethodHandler{}
	handler.InitiateDefaults()
	handler.GET = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "GET ROOT Works!!!")
	}
	handler.POST = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "POST ROOT Works!!!")
	}
	return Route{MethodHandler: handler, controller: controller, path: "/"}
}

func DetailRoute(controller *devices.DeviceController) Route {
	detailHandler := MethodHandler{}
	detailHandler.InitiateDefaults()
	detailHandler.GET = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "GET DETAIL Works!!!")
	}
	detailHandler.POST = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "POST DETAIL Works!!!")
	}
	return Route{MethodHandler: detailHandler, controller: controller, path: "/detail"}
}
