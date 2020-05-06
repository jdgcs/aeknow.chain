
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>Chaet</title>
    <style type="text/css">
        .talk_con {
            width: 100%;
            height: 100%;
            border: 1px solid #666;
            margin: 50px auto 0;
            background: #f9f9f9;
        }

        .talk_show {
            width: 100%;
            height: 420px;
            border: 1px solid #666;
            background: #fff;
            margin: 10px auto 0;
            overflow: auto;
        }

        .talk_input {
            width: 100%;
        }

        .talk_word {
            width: 90%;
            height: 26px;
            float: left;
            text-indent: 10px;
            margin: 2% 5%;
        }

        .talk_sub {
            width: 100%;
            height: 30px;
            float: left;
        }

        .atalk {
            margin: 10px;
        }

        .atalk span {
            display: inline-block;
            background: #0181cc;
            border-radius: 10px;
            color: #fff;
            padding: 5px 10px;
        }

        .btalk {
            margin: 10px;
            text-align: right;
        }

        .btalk span {
            display: inline-block;
            background: #ef8201;
            border-radius: 10px;
            color: #fff;
            padding: 5px 10px;
        }
    </style>
<link rel="stylesheet" href="/views/layim/dist/css/layui.css">
  <!-- Tell the browser to be responsive to screen width -->
  <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
  <!-- Bootstrap 3.3.7 -->
  <link rel="stylesheet" href="/views/static/bower_components/bootstrap/dist/css/bootstrap.min.css">
  <!-- Font Awesome -->
  <link rel="stylesheet" href="/views/static/bower_components/font-awesome/css/font-awesome.min.css">
  <!-- Ionicons -->
  <link rel="stylesheet" href="/views/static/bower_components/Ionicons/css/ionicons.min.css">
  <!-- Theme style -->
  <link rel="stylesheet" href="/views/static/dist/css/AdminLTE.css">
  <!-- AdminLTE Skins. Choose a skin from the css/skins
       folder instead of downloading all of them to reduce the load. -->
 <link rel="stylesheet" href="/views/static/dist/css/skins/skin.css">
  <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
  <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
  <!--[if lt IE 9]>
  <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
  <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
  <![endif]-->

  
</head>
<body class="hold-transition skin-blue sidebar-mini">
<!-- Site wrapper -->
<div class="wrapper">

<header class="main-header">
    <!-- Logo -->
    <a href="/" class="logo"  style="background:#f7296e">
      <!-- mini logo for sidebar mini 50x50 pixels -->
      <span class="logo-mini"><b>a</b>K</span>
      <!-- logo for regular state and mobile devices -->
      <span class="logo-lg"><b>ae</b>Know</span>
    </a>
    <!-- Header Navbar: style can be found in header.less -->
    <nav class="navbar navbar-static-top"  style="background:#f7296e">
      <!-- Sidebar toggle button-->
      <a href="#" class="sidebar-toggle" data-toggle="push-menu" role="button">
        <span class="sr-only">Toggle navigation</span>
      </a>

      <div class="navbar-custom-menu" >
        <ul class="nav navbar-nav">
       <!-- User Account: style can be found in dropdown.less -->
          <li class="dropdown user user-menu">
            <a href="#">
              <img src="/views/static/dist/img/ae.png" class="user-image" alt="User Image">
              <span class="hidden-xs">{{.Account}}</span>
            </a>
           
          </li>

        </ul>
      </div>
    </nav>
  </header>
  

  <!-- Left side column. contains the sidebar -->
  {{ template "sidebar" .}}

  <!-- Content Wrapper. Contains page content -->
  <div class="content-wrapper">
    <!-- Content Header (Page header) -->
    <section class="content-header">
      <h1>
       Chat
       
      </h1>
     <ol class="breadcrumb">
  <li>
    <a href="#">
    <i class="fa fa-dashboard"></i>Home</a>
  </li>
  <li>
    <a href="#">Knode</a>
  </li>
  <li class="active">chat</li>
</ol>

    </section>

    <!-- Main content -->
    <section class="content">

<script src="/views/layim/dist/layui.js"></script>
<script>
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
				socket.send("{{.Account}} Online")       
            };
            
           
            //收到消息推送
            function doWithMsg(msg) {
                //getdjxh()//这个函数是业务上面申请列表的函数 可以忽略
                window.external.CallFun('receiveMsg');//这个也是
            }
     //   }

 /* 
 function reconnect(url) {
            if (lockReconnect) return;
            lockReconnect = true;
            //没连接上会一直重连，设置延迟避免请求过多
            setTimeout(function() {
                getwebsocket();
                lockReconnect = false;
            }, 2000);
        }
*/
//心跳检测
/*
         var heartCheck = {
            timeout: 6000, //6秒
            timeoutObj: null,
            serverTimeoutObj: null,
            reset: function() {
                clearTimeout(this.timeoutObj);
                clearTimeout(this.serverTimeoutObj);
                return this;
            },
            start: function() {
                var self = this;
                this.timeoutObj = setTimeout(function() {
                    //这里发送一个心跳，后端收到后，返回一个心跳消息，
                    //onmessage拿到返回的心跳就说明连接正常
                    socket.send("ping");
                    self.serverTimeoutObj = setTimeout(function() { //如果超过一定时间还没重置，说明后端主动断开了
                        socket.close(); //如果onclose会执行reconnect，我们执行ws.close()就行了.如果直接执行reconnect 会触发onclose导致重连两次
                    }, self.timeout)
                }, this.timeout)
            }
        }
*/



layui.use('layim', function(layim){
  
  //演示自动回复
  var autoReplay = [
    'AENS chat玩起来', 
    '你没发错吧？face[微笑] ',
    '洗澡中，请勿打扰，偷窥请购票，个体四十，团体八折，订票电话：一般人我不告诉他！face[哈哈] ',
    '你好，我是主人的美女秘书，有什么事就跟我说吧，等他回来我会转告他的。face[心] face[心] face[心] ',
    '表情ace[威武] face[威武] face[威武] face[威武] ',
    '<（@￣︶￣@）>',
    '你要和我说话？你真的要和我说话？你确定自己想说吗？你一定非说不可吗？那你说吧，这是自动回复。',
    'face[黑线]  你慢慢说，别急……',
    '(*^__^*) face[嘻嘻] ，启动个域名？'
  ];
  
  //基础配置
  layim.config({

    //初始化接口
    init: {
      url: '/views/layim/sample/json/getList.json'
      ,data: {}
    }
    
    //或采用以下方式初始化接口
    /*
    ,init: {
      mine: {
        "username": "LayIM体验者" //我的昵称
        ,"id": "100000123" //我的ID
        ,"status": "online" //在线状态 online：在线、hide：隐身
        ,"remark": "在深邃的编码世界，做一枚轻盈的纸飞机" //我的签名
        ,"avatar": "a.jpg" //我的头像
      }
      ,friend: []
      ,group: []
    }
    */
    

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

  /*
  layim.chat({
    name: '在线客服-小苍'
    ,type: 'kefu'
    ,avatar: 'http://tva3.sinaimg.cn/crop.0.0.180.180.180/7f5f6861jw1e8qgp5bmzyj2050050aa8.jpg'
    ,id: -1
  });
  layim.chat({
    name: '在线客服-心心'
    ,type: 'kefu'
    ,avatar: 'http://tva1.sinaimg.cn/crop.219.144.555.555.180/0068iARejw8esk724mra6j30rs0rstap.jpg'
    ,id: -2
  });
  layim.setChatMin();*/

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

    //console.log(res.mine);
    
    layim.msgbox(5); //模拟消息盒子有新消息，实际使用时，一般是动态获得
  
    //添加好友（如果检测到该socket）
    layim.addList({
      type: 'group'
      ,avatar: "http://tva3.sinaimg.cn/crop.64.106.361.361.50/7181dbb3jw8evfbtem8edj20ci0dpq3a.jpg"
      ,groupname: 'AENS开发'
      ,id: "AK_dsadasdas"
      ,members: 0
    });
    layim.addList({
      type: 'friend'
      ,avatar: "http://tp2.sinaimg.cn/2386568184/180/40050524279/0"
      ,username: '冲田杏梨'
      ,groupid: 2
      ,id: "1233333312121212"
      ,remark: "本人冲田杏梨将结束AV女优的工作"
    });
    
    setTimeout(function(){
      //接受消息（如果检测到该socket）
      layim.getMessage({
        username: "Hi"
        ,avatar: "http://qzapp.qlogo.cn/qzapp/100280987/56ADC83E78CEC046F8DF2C5D0DD63CDE/100"
        ,id: "10000111"
        ,type: "friend"
        ,content: "临时："+ new Date().getTime()
      });
      
      /*layim.getMessage({
        username: "贤心"
        ,avatar: "http://tp1.sinaimg.cn/1571889140/180/40030060651/1"
        ,id: "100001"
        ,type: "friend"
        ,content: "嗨，你好！欢迎体验LayIM。演示标记："+ new Date().getTime()
      });*/
      
    }, 3000);
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



</script>
        <!-- /.box-body -->
       
      </div>
      <!-- /.box -->

    </section>
    <!-- /.content -->
  </div>
  <!-- /.content-wrapper -->

{{ template "footer" .}}


  <!-- Add the sidebar's background. This div must be placed
       immediately after the control sidebar -->
  <div class="control-sidebar-bg"></div>
</div>
<!-- ./wrapper -->

<!-- jQuery 3 -->


</body>
</html>
