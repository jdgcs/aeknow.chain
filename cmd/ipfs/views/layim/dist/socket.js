if(!/^http(s*):\/\//.test(location.href)){
alert('请部署到localhost');
}	
	
var lockReconnect = false;	
var socket;
var wsUrl = 'ws://localhost:8888/websocket';
//getwebsocket();
self.setInterval("heart()",60000);
function heart()
{
	socket.send('ping');
	console.log('ping')
}


//TODO: test stability
 //function getwebsocket() { //新建websocket的函数 页面初始化 断开连接时重新调用
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
				socket.send(myAccount+" Online")       
            };
            
           
            //收到消息推送
            function doWithMsg(msg) {
                //getdjxh()//这个函数是业务上面申请列表的函数 可以忽略
                window.external.CallFun('receiveMsg');//这个也是
            }
     //   }

layui.config({
  base: '/views/layim/dist/js/' //扩展 JS 所在目录
}).extend({
  //socket: 'socket'
});

layui.define(['jquery','contextMenu'], function (exports) {
    var contextMenu = layui.contextMenu;
    var $ = layui.jquery;
    var ext = {
        init : function(){//定义右键操作
            $(".layim-list-friend >li > ul > li").contextMenu({
                width: 140, // width
                itemHeight: 30, // 菜单项height
                bgColor: "#fff", // 背景颜色
                color: "#333", // 字体颜色
                fontSize: 15, // 字体大小
                hoverBgColor: "#009bdd", // hover背景颜色
                hoverColor: "#fff", // hover背景颜色
                target: function(ele) { // 当前元素
                    $(".ul-context-menu").attr("data-id",ele[0].id);
                    $(".ul-context-menu").attr("data-name",ele.find("span").html());
                    $(".ul-context-menu").attr("data-img",ele.find("img").attr('src'));
                },
                menu: [
                    { // 菜单项
                        text: "发送消息",
                        icon: "",
                        callback: function(ele) {
                            var othis = ele.parent(),
                                friend_id = othis[0].dataset.id.replace(/^layim-friend/g, ''),
                                friend_name = othis[0].dataset.name,
                                friend_avatar = othis[0].dataset.img;
                            conf.layim.chat({
                                name: friend_name
                                ,type: 'friend'
                                ,avatar: friend_avatar
                                ,id: friend_id
                            });
                        }
                    },                
                    {
                        text: "查看资料",
                        icon: "",
                        callback: function(ele) {
                            var othis = ele.parent(),friend_id = othis[0].dataset.id.replace(/^layim-friend/g, '');
                            im.getInformation({
                                id: friend_id,
                                type:'friend'
                            });                        
                        }
                    },
                    {
                        text: "聊天记录",
                        icon: "",
                        callback: function(ele) {
                            var othis = ele.parent(),
                                friend_id = othis[0].dataset.id.replace(/^layim-friend/g, ''),
                                friend_name = othis[0].dataset.name;
                            im.getChatLog({
                                name: friend_name,
                                id: friend_id,
                                type:'friend'
                            });    
                        }
                    }                                                    
                ]
            });
        }
    }
  exports('ext', ext);
}); 

layui.use(['layim','contextMenu'], function(layim){ 
  //基础配置
  layim.config({

    //初始化接口
    init: {
      url: '/views/layim/sample/json/getList.json'
      ,data: {}
    }
    
    //查看群员接口
    ,members: {
      url: '/views/layim/sample/json/getMembers.json'
      ,data: {}
    }
    
    //上传图片接口
    ,uploadImage: {
      url: '/uploadimage' //（返回的数据格式见下文）
      ,type: '' //默认post
    } 
    
    //上传文件接口
    ,uploadFile: {
      url: '/uploadfile' //（返回的数据格式见下文）
      ,type: '' //默认post
    }
    
    ,isAudio: true //开启聊天工具栏音频
    ,isVideo: true //开启聊天工具栏视频
    
    //扩展工具栏
    ,tool: [{
      alias: 'code'
      ,title: '代码'
      ,icon: '&#xe64e;'
    }]
    
    //,brief: true //是否简约模式（若开启则不显示主面板）
    
    //,title: 'WebIM' //自定义主面板最小化时的标题
    ,right: '10px' //主面板相对浏览器右侧距离
    //,minRight: '90px' //聊天面板最小化时相对浏览器右侧距离
    ,initSkin: '1.jpg' //1-5 设置初始背景
    //,skin: ['aaa.jpg'] //新增皮肤
    //,isfriend: false //是否开启好友
    //,isgroup: false //是否开启群组
    //,min: true //是否始终最小化主面板，默认false
    ,notice: true //是否开启桌面消息提醒，默认false
    //,voice: false //声音提醒，默认开启，声音文件为：default.mp3
    
    ,msgbox:  '/views/layim/dist/css/modules/layim/html/msgbox.html' //消息盒子页面地址，若不开启，剔除该项即可
    ,find:  '/views/layim/dist/css/modules/layim/html/find.html' //发现页面地址，若不开启，剔除该项即可
    ,chatLog: '/views/layim/dist/css/modules/layim/html/chatlog.html' //聊天记录页面地址，若不开启，剔除该项即可
    
  });

 
  //监听在线状态的切换事件
  layim.on('online', function(data){
    //console.log(data);
  });
  
  //监听签名修改
  layim.on('sign', function(value){
    //console.log(value);
  });

  //监听自定义工具栏点击，以添加代码为例
  layim.on('tool(code)', function(insert){
    layer.prompt({
      title: '插入代码'
      ,formType: 2
      ,shade: 0
    }, function(text, index){
      layer.close(index);
      insert('[pre class=layui-code]' + text + '[/pre]'); //将内容插入到编辑器
    });
  });
  
  //监听layim建立就绪
  layim.on('ready', function(res){
	   layui.ext.init(); 
    //console.log(res.mine);    
   // layim.msgbox(5); //模拟消息盒子有新消息，实际使用时，一般是动态获得
 
  });

  //监听发送消息
  layim.on('sendMessage', function(data){
    var To = data.to;
    console.log(data);
    //socket.send('{"From":"{{.Account}}","Body":"'+JSON.stringify(data)+'"}');
    socket.send(JSON.stringify(data));    
    //if(To.type === 'friend'){
   //   layim.setChatStatus('<span style="color:#FF5722;">对方正在输入。。。</span>');
   // }
    
   
  });

  //监听查看群员
  layim.on('members', function(data){
    //console.log(data);
  });
  
  //监听聊天窗口的切换
  layim.on('chatChange', function(res){
    var type = res.data.type;
    console.log(res.data.id)
    if(type === 'friend'){
      //模拟标注好友状态
      //layim.setChatStatus('<span style="color:#FF5722;">在线</span>');
    } else if(type === 'group'){
      //模拟系统消息
      /*
      layim.getMessage({
        system: true
        ,id: res.data.id
        ,type: "group"
        ,content: '模拟群员'+(Math.random()*100|0) + '加入群聊'
      });*/
    }
  });
  
  
 socket.onmessage = function(res){  				
				console.log(res.data)
				//heartCheck.reset().start();
				
					//if(res.data != 'pong'){		
						layim.getMessage(JSON.parse(res.data)); //res.data即你发送消息传递的数据（阅读：监听发送的消息）
					//}
				
			};


});


