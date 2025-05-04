package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Aman5681/notify/internal/orchestrator"
	"github.com/Aman5681/notify/internal/payload"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartWebSocketServer(service *orchestrator.Service) error {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("⚠️ WebSocket upgrade failed:", err)
			return
		}
		defer ws.Close()

		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				log.Println("⚠️ Read error:", err)
				break
			}

			fmt.Println(msg)

			convertedMessage, _ := convertToJson(msg)

			handler, ok := service.GetHandler(convertedMessage.Action)
			if !ok {
				ws.WriteMessage(websocket.TextMessage, []byte("Unknown action: "+convertedMessage.Action))
				continue
			}

			answer, err := handler(convertedMessage)
			if err != nil {
				ws.WriteMessage(websocket.TextMessage, []byte("Error: "+err.Error()))
			} else {
				ws.WriteMessage(websocket.TextMessage, []byte(answer))
			}
		}
	})

	log.Println("🌐 WebSocket server listening on :8080")
	return http.ListenAndServe(":8080", nil)
}

func convertToJson(msg []byte) (payload.Payload, error) {
	var result payload.Payload
	err := json.Unmarshal([]byte(msg), &result)
	return result, err

}
