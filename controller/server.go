package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/atakanozceviz/gorient.v3/model"
	"github.com/googollee/go-socket.io"
)

var tpl = template.Must(template.ParseGlob("./view/templates/*"))
var conn = model.Conn{}

func StartServer(tcp string) error {

	ws, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	ws.On("connection", func(so socketio.Socket) {
		conn.Add(so.Id(), so)

		LogErr(so.Join(so.Id()))
		LogErr(so.Join("all"))

		so.Emit("connection", so.Id())
		log.Println(so.Id() + "=> on connection")

		so.On("connectedto", func(to string) {
			LogErr(so.BroadcastTo(to, "client", so.Id()))
		})

		so.On("disconnection", func() {
			LogErr(so.Leave("all"))
			LogErr(so.Leave(so.Id()))
			LogErr(so.BroadcastTo("all", "DC", so.Id()))
			log.Println(so.Id() + "=> on disconnect")
			conn.Remove(so.Id())
		})

		//Send device orientation to server which connected to.
		so.On("orient", func(data string) {
			LogErr(so.BroadcastTo(data[7:27], "orient", data))

			o := model.Orient{}
			LogErr(json.Unmarshal([]byte(data), &o))

			x, y := Coord(o.Gamma, o.Beta)
			LogErr(so.BroadcastTo(o.ID, "coord", "{\"x\":"+FloatToString(x)+",\"y\":"+FloatToString(y)+"}"))

		})

	})
	ws.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", ws)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./view/static/"))))
	http.HandleFunc("/", serverHandler)
	http.HandleFunc("/client", clientHandler)

	log.Println("Serving at:" + tcp)
	return http.ListenAndServe(tcp, nil)
}

func serverHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "server.html", model.ServerData{Title: "Server"})
}

func clientHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if _, ok := conn[id]; ok {
		tpl.ExecuteTemplate(w, "client.html", model.ClientData{Title: "Client", ID: id})
	} else {
		w.Write([]byte("Wrong ID!"))
	}

}

func LogErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
