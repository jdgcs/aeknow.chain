
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
    <script src="/views/jquery.min.js"></script>
    <script type="text/javascript">
        $(function () {
            // 询问框获取用户昵称
            let username ={{.Account}}
            //localStorage.getItem("username") ? 
             //   localStorage.getItem("username") : disp_prompt();
           
            let mywords= $("#mywords");
            let talkWords = $("#talkwords");
            let talkSubmit = $("#talksub");
            // webSocket
            let wsURL = "ws://127.0.0.1:8888/websocket";
            ws = new WebSocket(wsURL);
            try {
                // 监听连接服务器
                ws.onopen = function () {
                    console.log("已连接服务器")
                    //ws.send("{\"username\":\"{{.Account}}\", \"message\":\"hello\"}")
                };

                // 监听关闭服务器
                ws.onclose = function () {
                    if (ws) {
                        ws.close();
                        ws = null
                    }
                    console.log("关闭服务器连接")
                };

                // 监听信息
                ws.onmessage = function (result) {
                    let data = JSON.parse(result.data);
                    let className = "direct-chat-msg";
                    let myuserclassName = "direct-chat-name";
                    let user = data.username
                    let myuser = data.username
                    // 如果是本人,放在右边 不是本人 放在左边
                    if (data.username === username){
                        className = "direct-chat-msg right";
                        myuserclassName = "direct-chat-name pull-right";
                        myuser = "me";
                    }
                                        
                    str1 = mywords.html() + 
                        '<div class=" + className + ">'+user+'=><span>' 
                        + data.message + '</span></div>';
                    
                    str=mywords.html() +'<div class="'+className+'"><div class="direct-chat-info clearfix"><span class="'+myuserclassName+'">'+myuser+'</span></div>'
                    +'<img class="direct-chat-img" src="http://127.0.0.1:8080/ipfs/QmWSzmtsH19CrYXo6sGv3ocpyEZgDUYbAJJXofUzFe2BKz" alt="message user image"><div class="direct-chat-text">'
                     + data.message+'</div></div>'
                    
                    
                    mywords.html(str);
                    
                   // var scrollHeight = words.prop("scrollHeight");
                    var scrollHeight = mywords.prop("scrollHeight");
                   // words.scrollTop(scrollHeight);
                    mywords.scrollTop(scrollHeight);
                };

                // 监听错误
                ws.onerror = function () {
                    if (ws) {
                        ws.close();
                        ws = null;
                    }
                    console.log("服务器连接失败")
                }
            } catch (e) {
                console.log(e.message)
            }

            document.onkeydown = function (event) {
                let e = event || window.event;
                if (e && e.keyCode === 13) { //回车键的键值为13
                    talkSubmit.click()
                }
            };

            talkSubmit.click(function () {
                // 获取输入框内容
                let content = talkWords.val();
                if (content === "") {
                    // 消息为空时弹窗
                    alert("消息不能为空");
                    return;
                }

                // 发送数据
                if (ws == null){
                    alert("连接服务器失败,请刷新页面");
                    window.location.reload();
                    return
                }
                let request = {"username":username, "message":content,"topic":{{.ChatTopic}}};
                ws.send(JSON.stringify(request));
                // 清空输入框
                talkWords.val("")
            })
        });
        
        function disp_prompt() {
            let username = prompt("请输入昵称");
            if (username == null || username === "") {
                disp_prompt()
            }else {
                localStorage.setItem("username", username);
                return username;
            }
        }
    </script>
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

  <!-- Google Font -->
  <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:300,400,600,700,300italic,400italic,600italic">
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
       Chatroom
       
      </h1>
      <ol class="breadcrumb">
        <li><a href="#"><i class="fa fa-dashboard"></i> Home</a></li>
        <li><a href="#">Knode</a></li>
        <li class="active">aeternity</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">

    
              <div class="row">
            <div class="col-md-6">
              <!-- DIRECT CHAT -->
              <div class="box box-warning direct-chat direct-chat-warning">
                <div class="box-header with-border">
                  <h3 class="box-title">Direct Chat</h3>

                  <div class="box-tools pull-right">
                    <span data-toggle="tooltip" title="3 New Messages" class="badge bg-yellow">3</span>
                    <button type="button" class="btn btn-box-tool" data-widget="collapse"><i class="fa fa-minus"></i>
                    </button>
                    <button type="button" class="btn btn-box-tool" data-toggle="tooltip" title="Contacts"
                            data-widget="chat-pane-toggle">
                      <i class="fa fa-comments"></i></button>
                    <button type="button" class="btn btn-box-tool" data-widget="remove"><i class="fa fa-times"></i>
                    </button>
                  </div>
                </div>
                <!-- /.box-header -->
                <div class="box-body">
                  <!-- Conversations are loaded here -->
                  <div class="direct-chat-messages" style="height:400px" id="mywords">
                    <!-- Message. Default to the left -->
                    <div class="direct-chat-msg" >
                      <div class="direct-chat-info clearfix">
                        <span class="direct-chat-name pull-left">Alexander Pierce</span>
                        <span class="direct-chat-timestamp pull-right">23 Jan 2:00 pm</span>
                      </div>
                      <!-- /.direct-chat-info -->
                      <img class="direct-chat-img" src="/views/static/dist/img/user1-128x128.jpg" alt="message user image">
                      <!-- /.direct-chat-img -->
                      <div class="direct-chat-text">
                        Is this template really for free? That's unbelievable!
                      </div>
                      <!-- /.direct-chat-text -->
                    </div>
                    <!-- /.direct-chat-msg -->

                    <!-- Message to the right -->
                    <div class="direct-chat-msg right">
                      <div class="direct-chat-info clearfix">
                        <span class="direct-chat-name pull-right">Sarah Bullock</span>
                        <span class="direct-chat-timestamp pull-left">23 Jan 2:05 pm</span>
                      </div>
                      <!-- /.direct-chat-info -->
                      <img class="direct-chat-img" src="/views/static/dist/img/user3-128x128.jpg" alt="message user image">
                      <!-- /.direct-chat-img -->
                      <div class="direct-chat-text">
                        You better believe it!
                      </div>
                      <!-- /.direct-chat-text -->
                    </div>
                    <!-- /.direct-chat-msg -->

                    <!-- Message. Default to the left -->
                    <div class="direct-chat-msg">
                      <div class="direct-chat-info clearfix">
                        <span class="direct-chat-name pull-left">Alexander Pierce</span>
                        <span class="direct-chat-timestamp pull-right">23 Jan 5:37 pm</span>
                      </div>
                      <!-- /.direct-chat-info -->
                      <img class="direct-chat-img" src="/views/static/dist/img/user1-128x128.jpg" alt="message user image">
                      <!-- /.direct-chat-img -->
                      <div class="direct-chat-text">
                        Working with AdminLTE on a great new app! Wanna join?
                      </div>
                      <!-- /.direct-chat-text -->
                    </div>
                    <!-- /.direct-chat-msg -->

                    <!-- Message to the right -->
                    <div class="direct-chat-msg right">
                      <div class="direct-chat-info clearfix">
                        <span class="direct-chat-name pull-right">Sarah Bullock</span>
                        <span class="direct-chat-timestamp pull-left">23 Jan 6:10 pm</span>
                      </div>
                      <!-- /.direct-chat-info -->
                      <img class="direct-chat-img" src="/views/static/dist/img/user3-128x128.jpg" alt="message user image">
                      <!-- /.direct-chat-img -->
                      <div class="direct-chat-text">
                        I would love to.
                      </div>
                      <!-- /.direct-chat-text -->
                    </div>
                    <!-- /.direct-chat-msg -->

                  </div>
                  <!--/.direct-chat-messages-->

                  <!-- Contacts are loaded here -->
                  <div class="direct-chat-contacts">
                    <ul class="contacts-list">
                      <li>
                        <a href="#">
                          <img class="contacts-list-img" src="/views/static/dist/img/user1-128x128.jpg" alt="User Image">

                          <div class="contacts-list-info">
                                <span class="contacts-list-name">
                                  Count Dracula
                                  <small class="contacts-list-date pull-right">2/28/2015</small>
                                </span>
                            <span class="contacts-list-msg">How have you been? I was...</span>
                          </div>
                          <!-- /.contacts-list-info -->
                        </a>
                      </li>
                      <!-- End Contact Item -->
                      <li>
                        <a href="#">
                          <img class="contacts-list-img" src="/views/static/dist/img/user7-128x128.jpg" alt="User Image">

                          <div class="contacts-list-info">
                                <span class="contacts-list-name">
                                  Sarah Doe
                                  <small class="contacts-list-date pull-right">2/23/2015</small>
                                </span>
                            <span class="contacts-list-msg">I will be waiting for...</span>
                          </div>
                          <!-- /.contacts-list-info -->
                        </a>
                      </li>
                      <!-- End Contact Item -->
                      <li>
                        <a href="#">
                          <img class="contacts-list-img" src="/views/static/dist/img/user3-128x128.jpg" alt="User Image">

                          <div class="contacts-list-info">
                                <span class="contacts-list-name">
                                  Nadia Jolie
                                  <small class="contacts-list-date pull-right">2/20/2015</small>
                                </span>
                            <span class="contacts-list-msg">I'll call you back at...</span>
                          </div>
                          <!-- /.contacts-list-info -->
                        </a>
                      </li>
                      <!-- End Contact Item -->
                      <li>
                        <a href="#">
                          <img class="contacts-list-img" src="/views/static/dist/img/user5-128x128.jpg" alt="User Image">

                          <div class="contacts-list-info">
                                <span class="contacts-list-name">
                                  Nora S. Vans
                                  <small class="contacts-list-date pull-right">2/10/2015</small>
                                </span>
                            <span class="contacts-list-msg">Where is your new...</span>
                          </div>
                          <!-- /.contacts-list-info -->
                        </a>
                      </li>
                      <!-- End Contact Item -->
                      <li>
                        <a href="#">
                          <img class="contacts-list-img" src="/views/static/dist/img/user6-128x128.jpg" alt="User Image">

                          <div class="contacts-list-info">
                                <span class="contacts-list-name">
                                  John K.
                                  <small class="contacts-list-date pull-right">1/27/2015</small>
                                </span>
                            <span class="contacts-list-msg">Can I take a look at...</span>
                          </div>
                          <!-- /.contacts-list-info -->
                        </a>
                      </li>
                      <!-- End Contact Item -->
                      <li>
                        <a href="#">
                          <img class="contacts-list-img" src="/views/static/dist/img/user8-128x128.jpg" alt="User Image">

                          <div class="contacts-list-info">
                                <span class="contacts-list-name">
                                  Kenneth M.
                                  <small class="contacts-list-date pull-right">1/4/2015</small>
                                </span>
                            <span class="contacts-list-msg">Never mind I found...</span>
                          </div>
                          <!-- /.contacts-list-info -->
                        </a>
                      </li>
                      <!-- End Contact Item -->
                    </ul>
                    <!-- /.contatcts-list -->
                  </div>
                  <!-- /.direct-chat-pane -->
                </div>
                <!-- /.box-body -->
                <div class="box-footer">
                 
                   <div class="input-group">
						<input type="text"  class="form-control" id="talkwords" placeholder="输入聊天内容">
						<span class="input-group-btn">
							<input type="button" class="btn btn-warning btn-flat" value="发送"  id="talksub">
						</span>
					</div>
                  
                </div>
                <!-- /.box-footer-->
              </div>
              <!--/.direct-chat -->
            </div>
            <!-- /.col -->

            <div class="col-md-4">
              <!-- USERS LIST -->
              <div class="box box-danger">
                <div class="box-header with-border">
                  <h3 class="box-title">Latest Members</h3>

                  <div class="box-tools pull-right">
                    <span class="label label-danger">8 New Members</span>
                    <button type="button" class="btn btn-box-tool" data-widget="collapse"><i class="fa fa-minus"></i>
                    </button>
                    <button type="button" class="btn btn-box-tool" data-widget="remove"><i class="fa fa-times"></i>
                    </button>
                  </div>
                </div>
                <!-- /.box-header -->
                <div class="box-body no-padding">
                  <ul class="users-list clearfix">
                    <li>
                      <img src="/views/static/dist/img/user1-128x128.jpg" alt="User Image">
                      <a class="users-list-name" href="#">Alexander Pierce</a>
                      <span class="users-list-date">Today</span>
                    </li>
                    <li>
                      <img src="/views/static/dist/img/user8-128x128.jpg" alt="User Image">
                      <a class="users-list-name" href="#">Norman</a>
                      <span class="users-list-date">Yesterday</span>
                    </li>
                    <li>
                      <img src="/views/static/dist/img/user7-128x128.jpg" alt="User Image">
                      <a class="users-list-name" href="#">Jane</a>
                      <span class="users-list-date">12 Jan</span>
                    </li>
                    <li>
                      <img src="/views/static/dist/img/user6-128x128.jpg" alt="User Image">
                      <a class="users-list-name" href="#">John</a>
                      <span class="users-list-date">12 Jan</span>
                    </li>
                    <li>
                      <img src="/views/static/dist/img/user2-160x160.jpg" alt="User Image">
                      <a class="users-list-name" href="#">Alexander</a>
                      <span class="users-list-date">13 Jan</span>
                    </li>
                    <li>
                      <img src="/views/static/dist/img/user5-128x128.jpg" alt="User Image">
                      <a class="users-list-name" href="#">Sarah</a>
                      <span class="users-list-date">14 Jan</span>
                    </li>
                    <li>
                      <img src="/views/static/dist/img/user4-128x128.jpg" alt="User Image">
                      <a class="users-list-name" href="#">Nora</a>
                      <span class="users-list-date">15 Jan</span>
                    </li>
                    <li>
                      <img src="/views/static/dist/img/user3-128x128.jpg" alt="User Image">
                      <a class="users-list-name" href="#">Nadia</a>
                      <span class="users-list-date">15 Jan</span>
                    </li>
                  </ul>
                  <!-- /.users-list -->
                </div>
                <!-- /.box-body -->
                <div class="box-footer text-center">
                  <a href="javascript:void(0)" class="uppercase">View All Users</a>
                </div>
                <!-- /.box-footer -->
              </div>
              <!--/.box -->
            </div>
            <!-- /.col -->
          </div>
          <!-- /.row -->
       </div>
        <!-- /.box-body -->
        <div class="box-footer">

        </div>
        <!-- /.box-footer-->
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
<script src="/views/static/bower_components/jquery/dist/jquery.min.js"></script>
<!-- Bootstrap 3.3.7 -->
<script src="/views/static/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>
<!-- SlimScroll -->
<script src="/views/static/bower_components/jquery-slimscroll/jquery.slimscroll.min.js"></script>
<!-- FastClick -->
<script src="/views/static/bower_components/fastclick/lib/fastclick.js"></script>
<!-- AdminLTE App -->
<script src="/views/static/dist/js/adminlte.min.js"></script>
<!-- AdminLTE for demo purposes -->
<script src="/views/static/dist/js/demo.js"></script>
<script>
  $(document).ready(function () {
    $('.sidebar-menu').tree()
  })
</script>
</body>
</html>
