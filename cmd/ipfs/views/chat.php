
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
var myAccount='{{.Account}}';


if(!/^http(s*):\/\//.test(location.href)){
//alert('请部署到localhost');
}	
	
var lockReconnect = false;	
var socket;
var wsUrl = 'ws://127.0.0.1:8888/websocket';


self.setInterval("heart()",60000);
function heart()
{
	socket.send('ping');
	console.log('ping')
}

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


layui.config({
                base: '/views/layim/dist/js/'
            }).extend({
                //socket: 'socket',
               // contextmenu:'contextMenu',
            });

layui.use(['layim','contextmenu'], function(layim){
  var menu = layui.contextmenu ;
    
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
    
    //,title: 'Chaet' //自定义主面板最小化时的标题
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
    menu.init([{
        target: '.layim-list-friend',
        menu: [{
            text: "新增分组",
            callback: function(target) {
                layer.msg(target.find('span').text());
            }
        }]
    },
    {
        target: '.layim-list-friend >li>h5>span',
        menu: [{
            text: "重命名",
            callback: function(target) {
                layer.msg(target.find('span').text());
            }
        },
        {
            text: "删除分组",
            callback: function(target) {
                layer.msg(target.find('span').text());
            }
        }]
    }]);
    
    
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
