package route

import (
    "github.com/gorilla/mux"
    service "github.com/heaptracetechnology/zendesk/service"
    "log"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "CreateUser",
        "POST",
        "/createUser",
        service.CreateUser,
    },
    Route{
        "CreateTicket",
        "POST",
        "/createTicket",
        service.CreateTicket,
    },
    Route{
        "ListTicket",
        "POST",
        "/listTicket",
        service.ListTicket,
    },
    Route{
        "DeleteTicket",
        "POST",
        "/deleteTicket",
        service.DeleteTicket,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        log.Println(route.Name)
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
