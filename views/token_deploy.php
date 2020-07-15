<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>æKnow - Transaction</title>
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
</head>
<body class="hold-transition skin-blue sidebar-mini">
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
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
              <img src="/views/static/dist/img/ae.png" class="user-image" alt="User Image">
              <span class="hidden-xs">{{.Account}}</span>
            </a>
           <ul class="dropdown-menu">
              <!-- User image -->
              <li class="user-header"  style="height:60px;background-color:wihte;">
                <a href="/logout">Logout</a>
              </li>
              </ul>
          </li>

        </ul>
      </div>
    </nav>
  </header>
 <!-- Left side column. contains the logo and sidebar -->
{{ template "sidebar" .}}

  <!-- Content Wrapper. Contains page content -->
  <div class="content-wrapper">
	   <br/><br/> <br/><br/>
	  <center><h1>Deployed <font color=red>{{.Recipient_id}}</font> SUCCESSFULLY! </h1></center><br />
	  
	   <br/><br/>
    <a href="https://www.aeknow.org/block/transaction/{{.PageTitle}}" target="_blank">{{.PageTitle}}</a>
    <br/><br/>
   <center><a href= {{.PageContent}}><button type="button" class="btn btn-block btn-success btn-lg">Publish the Token Information to AEKnow.org With AE</button></a></center>
  <br /> <br />
<!--<center><a href= {{.PageContent}}><button type="button" class="btn btn-block btn-success btn-lg">Publish the Token Information to  AEKnow.org With AEG</button></a></center>-->


</div>
  <!-- /.content-wrapper -->
 {{ template "footer" .}}
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
