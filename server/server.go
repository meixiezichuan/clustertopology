package server

import (
	"context"
	"encoding/json"
	"fmt"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"

	edgev1 "github.com/meixiezichuan/clustertopology/api/edge/v1"
	"github.com/meixiezichuan/clustertopology/server/websocket"
)

var clusterTopolgy edgev1.ClusterTopology
var updated bool = false
var k8sclient client.Client

const updateInterval time.Duration = time.Second * 10

type Message struct {
	Type    string            `json:"type"`
	Time    string            `json:"time"`
	Content edgev1.OriginInfo `json:"content"`
}

func Start(cli client.Client) {
	k8sclient = cli

	s := websocket.New("127.0.0.1:9999", websocket.ServerConfig{})
	s.RegisterEventHandler("/publish", websocket.EventMessage, onMessage)
	s.RegisterEventHandler("/publish", websocket.EventOpen, onOpen)
	s.RegisterEventHandler("/publish", websocket.EventClose, onClose)
	s.RegisterEventHandler("/publish", websocket.EventError, onError)
	go UpdateClusterTopology()
	s.Start()
}

func onMessage(ctx *websocket.EventContext) {
	var m Message
	err := json.Unmarshal(ctx.Msg, &m)
	if err != nil {
		return
	}
	handleMessage(&m)
}

func onClose(ctx *websocket.EventContext) {
	fmt.Printf("%v: on close at %v\n", ctx.Conn.ConnInfo, time.Now().String())
}

func onError(ctx *websocket.EventContext) {
	fmt.Printf("%v: on error at %v\n", ctx.Conn.ConnInfo, time.Now().String())
}

func onOpen(ctx *websocket.EventContext) {
	fmt.Printf("%v: on Open at %v\n", ctx.Conn.ConnInfo, time.Now().String())
}

func handleMessage(m *Message) {
	if m.Type == "NodeInfo" {
		clusterTopolgy.SetNetOriginInfo(&m.Content)
		updated = true
	}
}

func UpdateClusterTopology() {
	ticker := time.NewTicker(updateInterval)
	defer ticker.Stop()
	for range ticker.C {
		if updated {
			k8sclient.Update(context.TODO(), &clusterTopolgy, nil)
			updated = false
		}
	}
}
