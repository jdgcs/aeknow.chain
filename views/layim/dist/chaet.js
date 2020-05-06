if(!/^http(s*):\/\//.test(location.href)){
alert('请部署到localhost');
}	
	
var lockReconnect = false;	
var socket;
var wsUrl = 'ws://localhost:8888/websocket';


self.setInterval("heart()",60000);
function heart()
{
	socket.send('ping');
	console.log('ping')
}

layui.define(['jquery'], function (exports) {	
			console.log(wsUrl);
            socket = new WebSocket(wsUrl);    
                    
            socket.onerror = function(event) {
                console.log('websocket服务出错了');
                alert("Websocket Server Error");
                //reconnect(wsUrl);
            };
            socket.onclose = function(event) {
                console.log('websocket服务关闭了');
                 window.alert("Websocket Server is Closed");
                //reconnect(wsUrl);
            };
            socket.onopen = function(event) {
                //heartCheck.reset().start(); //传递信息
                console.log("连接成功!"+new Date().toUTCString());
				socket.send("{{.Account}} Online")       
            };
            
           
            //收到消息推送
            function doWithMsg(msg) {
                //getdjxh()//这个函数是业务上面申请列表的函数 可以忽略
                window.external.CallFun('receiveMsg');//这个也是
            }
	 exports('socket', socket);   
});
