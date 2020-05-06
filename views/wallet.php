
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>æKnow</title>
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
			<div class="col-md-3">

          <!-- Profile Image -->
          <div class="box box-primary">
			  <div class="box-header with-border">
              <h3 class="box-title">Summary </h3>
            </div>
            <div class="box-body box-profile">
			<center>
              <img class="" style="with:120px;" src="/views/qr_ak.png" alt="AK address">
			</center>
              <h3 class="profile-username text-center">QRcode</h3>

              <p class="text-muted text-center">of the address</p>

              <ul class="list-group list-group-unbordered">
                <li class="list-group-item">
                  <i class="fa fa-btc margin-r-5"></i><b>Balance:</b> <a class="pull-right"/>{{.Balance}} AE</a>
                </li>   
                
                <li class="list-group-item">
                  <i class="fa fa-calculator margin-r-5"></i><b>Nonce:</b> <a class="pull-right"/>{{.Nonce}} </a>
                 
                </li> 
                
               <center> <a href="https://www.aeknow.org/address/wallet/{{.Account}}" target="_blank">Transaction History</a></center>
                 
                 
              </ul>

            </div>
            
            <!-- /.box-body -->
          </div>
          <!-- /.box -->

        </div>
        <!-- /.col -->
       
       
       <div class="col-md-9">
		<!-- Horizontal Form -->
          <div class="box box-info">
            <div class="box-header with-border">
              <h3 class="box-title">Transaction</h3>
            </div>
            <!-- /.box-header -->
            <!-- form start -->
            <form class="form-horizontal" action="/transaction" method="POST">
              <div class="box-body">
				   <div class="form-group">
                  <label for="inputEmail3" class="col-sm-2 control-label"><font color=red>*</font>Sender ID</label>

                  <div class="col-sm-10">
                    <input type="text" class="form-control" name="sender_id" value="{{.Account}}" readonly="readonly">
                  </div>
                </div> 
				  
                <div class="form-group">
                  <label for="inputEmail3" class="col-sm-2 control-label"><font color=red>*</font>Recipient Address</label>

                  <div class="col-sm-10">
                    <input type="text" class="form-control" name="recipient_id" placeholder="address as ak_**** or AENS name as liuyang.chain">
                  </div>
                </div>
                
                <div class="form-group">
                  <label for="inputEmail3" class="col-sm-2 control-label"><font color=red>*</font>Amount</label>
                  <div class="col-sm-10">
					  <div class="input-group">
                    <input type="text" class="form-control" name="amount" placeholder="such as 7.89"> 
                    <span class="input-group-addon">
                      AE
                     </span>
                     </div>
                  </div>
                  
            
                </div>
                
                <div class="form-group">
                  <label for="inputPassword3" class="col-sm-2 control-label"><font color=red>*</font>Password</label>

                  <div class="col-sm-10">
                    <input type="password" class="form-control" name="password" placeholder="Password">
                  </div>
                </div>
                
                <div class="form-group">
                  <label for="inputEmail3" class="col-sm-2 control-label">Payload</label>
                  <div class="col-sm-10">
                    <div class="input-group">
                        
                    <textarea type="text" class="form-control" name="payload"></textarea>
                    <span class="input-group-addon">
                          Encrypt &nbsp;<input type="checkbox" name="isencrypt" disabled="disabled" >
                     </span>
                  </div>
                  </div>
                </div>
             
				<div class="form-group">
                  <label for="inputEmail3" class="col-sm-2 control-label">Gas</label>
                  <div class="col-sm-10">
					  <div class="input-group">
                    <input type="text" class="form-control" name="gas" placeholder="0.00005 AE by default"> 
                    <span class="input-group-addon">
                      AE
                     </span>
                     </div>
                  </div>
                </div>
                
                <div class="form-group">
                  <label for="inputEmail3" class="col-sm-2 control-label">Nonce</label>
                  <div class="col-sm-10">
					
                    <input type="text" class="form-control" name="nonce" placeholder="Nonce"> 
                    
                  </div>
                </div>
             
             <div class="form-group">
                  <label for="inputEmail3" class="col-sm-2 control-label">Node</label>
                  <div class="col-sm-10">
					
                    <input type="text" class="form-control" name="pubnode" placeholder="Default：PUBLIC_NODE"> 
                    
                  </div>
                </div>
                
                 <center>
                  <img class="" style="with:120px;" src="/views/qr_account.png"  alt="nonce" />
                  <br />Offline Nonce               
                  </center>
                  
		 <button type="submit" class="btn btn-info pull-left">Send</button>
		</div>
       
       
        </div>
        
        <div class="box-body">
			
             </div>
              <!-- /.box-body -->
              <div class="box-footer">
               
              </div>
              <!-- /.box-footer -->
            </form>
          </div>
          <!-- /.box -->
        
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
