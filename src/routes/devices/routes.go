package devices

import (
	"applicant-backend/src/controller/devices"
	"fmt"
	"net/http"
)

type MethodHandler struct {
	GET    func(http.ResponseWriter, *http.Request)
	POST   func(http.ResponseWriter, *http.Request)
	PUT    func(http.ResponseWriter, *http.Request)
	PATCH  func(http.ResponseWriter, *http.Request)
	DELETE func(http.ResponseWriter, *http.Request)
}

type DeviceRouter struct {
	controller *devices.DeviceController
	Routes     []Route
}

func (r *DeviceRouter) RegisterRoutes() {
	for _, v := range r.Routes {
		route := v
		fmt.Println("registering:", v, v.path, v.MethodHandler)
		route.RegisterToPath()
	}
}

func NewDeviceRouter(controller *devices.DeviceController) *DeviceRouter {
	routes := []Route{
		RootRoute(controller),
		DetailRoute(controller),
	}
	return &DeviceRouter{controller: controller, Routes: routes}
}

func (m *MethodHandler) InitiateDefaults() {
	m.GET = func(writer http.ResponseWriter, w *http.Request) {
		writer.WriteHeader(403)
		fmt.Fprint(writer, "GET is not supported")
	}
	m.PUT = func(writer http.ResponseWriter, w *http.Request) {
		writer.WriteHeader(403)
		fmt.Fprint(writer, "PUT is not supported")
	}
	m.PATCH = func(writer http.ResponseWriter, w *http.Request) {
		writer.WriteHeader(403)
		fmt.Fprint(writer, "PATCH is not supported")
	}
	m.POST = func(writer http.ResponseWriter, w *http.Request) {
		writer.WriteHeader(403)
		fmt.Fprint(writer, "POST is not supported")
	}
	m.DELETE = func(writer http.ResponseWriter, w *http.Request) {
		writer.WriteHeader(403)
		fmt.Fprint(writer, "DELETE is not supported")
	}
}
func (m *MethodHandler) getHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			m.GET(writer, request)
		case "POST":
			m.POST(writer, request)
		case "PATCH":
			m.PATCH(writer, request)
		case "PUT":
			m.PUT(writer, request)
		case "DELETE":
			m.DELETE(writer, request)
		}
	}
}
