
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>Contracts - AEKnow</title>
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
  

  <!-- Left side column. contains the sidebar -->
  {{ template "sidebar" .}}

  <!-- Content Wrapper. Contains page content -->
  <div class="content-wrapper">
    <!-- Content Header (Page header) -->
    <section class="content-header">
      <h1>
       Know
        <small>{{.Account}}</small>
      </h1>
      <ol class="breadcrumb">
        <li><a href="#"><i class="fa fa-dashboard"></i> Home</a></li>
        <li><a href="#">Knode</a></li>
        <li class="active">aeternity</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">

      <!-- Default box -->
      <div class="box">
        <div class="box-header with-border">	     
   
		<!-- Horizontal Form -->  
            <div class="box">
             <div class="box box-info">
            <div class="box-header with-border">
              <h3 class="box-title">Contracts</h3>
            </div>
            <!-- /.box-header -->
            <div class="box-body">
              <table class="table table-bordered">
                <tr>
                  <th style="width: 25%">
                 <a href=/deploy><button type="button" class="btn btn-block btn-primary btn-lg">Deploy</button></a> 
                  </th>
                  <th style="width: 25%">
                  <a href=/call> <button type="button" class="btn btn-block btn-success btn-lg">Call</button></a>
                  </th>                  
                  <th style="width: 25%">
                   <a href=/decodecall> <button type="button" class="btn btn-block btn-warning btn-lg">Decode</button></a> 
                  </th>
                  <th style="width: 25%">
                   <a href=/token><button type="button" class="btn btn-block btn-info btn-lg">Token</button></a> 
                  </th>
                </tr>
                <tr></tr>
                <tr>
                <td colspan=4><center><h2>Smart Contract(Sophia) Resources</h2></center></td>
                </tr>
                <tr>
                <th>Protocol & Documents</th>               
                <th>Tools</th>
                <th>Examples</th>
                 <th>SDKs</th>
                </tr>
                <tr>
                <td>
                <li><a href=https://github.com/aeternity/protocol/tree/master/contracts>Smart Contracts</a></li>
                <li><a href=https://github.com/aeternity/AEXs>AEXs</a></li>
                <li><a href=https://github.com/aeternity/aepp-sdk-docs>SDK Docs</a></li>
                
                </td>
                
                <td>
					<li><a href=https://github.com/aeternity/aepp-aeproject-js>aeproject-js</a></li>
					<li><a href=https://github.com/aeternity/aesophia>aesophia</a></li>
					<li><a href=https://github.com/aeternity/aesophia_http>aesophia_http</a></li>
					<li><a href=https://github.com/aeternity/aesophia_cli>aesophia_cli</a></li>	
                </td>
                <td>
                <li><a href=https://github.com/aeternity/aepp-sophia-examples>Sophia Examples</a></li>
                <li><a href=https://github.com/aeternity/AEXs/blob/master/AEXS/aex-9.md>Fungible Tokens(AEX-9)</a></li>
                
                </td>
                <td>
					<li><a href=https://github.com/aeternity/aepp-sdk-js>Nodejs SDK</a></li>
					<li><a href="https://github.com/aeternity/aepp-sdk-python">Python SDK</li>
					<li><a href="https://github.com/aeternity/aepp-sdk-go">Golang SDK</a></li>
					<li><a href="https://github.com/aeternity/aepp-sdk-elixir">Elixir SDK</a></li>
					<li><a href="https://github.com/kryptokrauts/aepp-sdk-java">Java SDK</a></li>
                </td>
                </tr>
            </table>
            </div>
         
         </div>
             
       
       
        </div>
        

    
        
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
