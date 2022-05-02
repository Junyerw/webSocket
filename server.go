package main

import (
	"net/http"
	"github.com/gorilla/websocket"
)

//定义转换器
//使用websocket.Upgrader完成协议握手，得到WebSocket长连接
//操作websocket api, 读取客户端消息，然后原样发送回去
var(
	upgrader=websocket.Upgrader{
		//websocket的服务端一般独立部署，当我们从那个浏览器的直播页面里，想发起到websocket服务的连接的时候
		//本身是跨域的，比如从直播.com 跨域到 websocket.com
		//当Upgrader请求跨域的时候，我们需要告诉他，允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true;
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var(
		wsConn *websocket.Conn
		err error
		//msgType int
		//data []byte
		conn *
	)

	//底层完成HTTP的应答
	//Upgrade:websocket
	if wsConn, err = upgrader.Upgrade(w,r,nil);err != nil{
		return
	}

	//websocket.Conn  做一个数据的收发
	//for{
	//	// 数据手法格式 Text，Binary
	//	if _,data,err = conn.ReadMessage(); err!= nil{
	//		goto ERR
	//	}
	//	if err = conn.WriteMessage(websocket.TextMessage, data); err !=nil{
	//		goto ERR
	//	}
	//
	//}
//ERR:
//	conn.Close()


}

func main(){
	//http:localhost:7777/ws
	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe("0.0.0.0:7778", nil)
}


/*封装WebSocket
	隐藏细节，封装API
		封装Connection结构，隐藏WebSocket底层
		封装Connetcion的API，提供Send/Read/Close等线程安全接口

API原理
	SendMessage将消息投递到out channel
	ReadMessage从in channel读取消息

内部原理
	启动读协程，循环读取WebSocket,将消息投递到 in channel
	启动写协程，循环读取out channel,将消息写给Websocket
*/