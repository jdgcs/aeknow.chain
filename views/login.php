<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>login</title>
  <!-- Tell the browser to be responsive to screen width -->
  <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
  <!-- Bootstrap 3.3.7 -->
  <link rel="stylesheet" href="/views/static/bower_components/bootstrap/dist/css/bootstrap.min.css">
  <!-- Font Awesome -->
  <link rel="stylesheet" href="/views/static/bower_components/font-awesome/css/font-awesome.min.css">
  <!-- Ionicons -->
  <link rel="stylesheet" href="/views/static/bower_components/Ionicons/css/ionicons.min.css">
  <!-- Theme style -->
  <link rel="stylesheet" href="/views/static/dist/css/AdminLTE.min.css">
  <!-- iCheck -->
  <link rel="stylesheet" href="/views/static/plugins/iCheck/square/blue.css">

  <!-- HTML5 Shim and Respond.js IE8 support of HTML5 elements and media queries -->
  <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
  <!--[if lt IE 9]>
  <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
  <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
  <![endif]-->

  <!-- Google Font -->
</head>
<body class="hold-transition register-page">
<div class="register-box">
  <div class="register-logo">
    <a href="/"><b>æKnow </b></a>
    <div><small><h3>An Eternity Knowledge System</h3></small></div>
  </div>

  <div class="register-box-body" style="width:500px">
    <p class="login-box-msg">Login</p>

    <form action="/login" method="post">
            
      <div class="form-group">
	  <label>Select</label>
	  <select class="form-control" name="accountname">
		{{.Options}}
	  </select>
	</div>
      <div class="form-group has-feedback">
        <input type="password" name="password" class="form-control" placeholder="Password">
        <span class="glyphicon glyphicon-lock form-control-feedback"></span>
      </div>
      
           
      <div class="row">
      
        <!-- /.col -->
        <div class="col-xs-4">
          <button type="submit" class="btn btn-primary btn-block btn-flat">OK</button>
        </div>
        <div class="pull-right">
			
			<a href=/registernew>Register New</a> -
			
			<a href=/import>Import</a></h3>
		</div>
        <!-- /.col -->       
      </div>
    </form>


 
  </div>
  <!-- /.form-box -->
  
</div>
<center><div><small><h3>login</h3></small></div>
Powered by aeternity, ipfs
</center>
<!-- /.register-box -->

<!-- jQuery 3 -->
<script src="/views/static/bower_components/jquery/dist/jquery.min.js"></script>
<!-- Bootstrap 3.3.7 -->
<script src="/views/static/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>


</body>
</html>
