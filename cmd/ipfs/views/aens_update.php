
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>æKnow -AENS</title>
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
       AENS
        <small>Aeternity blockchain based distributed naming system</small>
      </h1>
      <ol class="breadcrumb">
        <li><a href="#"><i class="fa fa-dashboard"></i> Home</a></li>
        <li><a href="#">Know</a></li>
        <li class="active">AENS</li>
      </ol>
    </section>

    <!-- Main content -->
    <section class="content">

      <!-- Default box -->
      <div class="box">
        <div class="box-header with-border"> 
		
			
       <div class="col-md-7">
          <!-- /.box -->
        <div class="box box-success">
            <div class="box-header with-border">
              <h3 class="box-title">Update(Point) AENS name: {{.AENSName}} </h3>
            </div>
            <div class="box-body">
				 <form action="/updatenamepointer" method="post">
					 <div class="row">
				<input type="hidden" name="aensname" class="form-control"  value="{{.AENSName}}">
				</div>	 
              <div class="row">
                <div class="col-xs-3">
                  <input type="text" class="form-control" value="AE address" disabled>
                </div>
                <div class="col-xs-8">
                  <input type="text" name="aeaddress" class="form-control" placeholder="ak_address"  value="{{.AEAddress}}">
                </div>                
              </div><br/>
              
               <div class="row">
                <div class="col-xs-3">
                  <input type="text" class="form-control"  value="BTC address" disabled>
                </div>
                <div class="col-xs-8">
                  <input type="text" name="btcaddress" class="form-control" placeholder="" value="{{.BTCAddress}}">
                </div>                
              </div><br/>
              
              <div class="row">
                <div class="col-xs-3">
                  <input type="text" class="form-control"  value="Contract address" disabled>
                </div>
                <div class="col-xs-8">
                  <input type="text" name="contractaddress" class="form-control" placeholder="ct_address" value="{{.ContractAddress}}">
                </div>                
              </div><br/>
              
              
              <div class="row">
                <div class="col-xs-3">
                  <input type="text" class="form-control"  value="Email address" disabled>
                </div>
                <div class="col-xs-8">
                  <input type="text" name="emailaddress" placeholder="aaa@bbb.ccc" class="form-control" value="{{.EmailAddress}}">
                </div>                
              </div><br/>
              
                <div class="row">
                <div class="col-xs-3">
                  <input type="text" class="form-control"  value="ETH address" disabled>
                </div>
                <div class="col-xs-8">
                  <input type="text" name="ethaddress" class="form-control" placeholder="0x address" value="{{.ETHAddress}}">
                </div>                
              </div><br/>
              
              
               <div class="row">
                <div class="col-xs-3">
                  <input type="text" class="form-control" placeholder="" value="IPFS address" disabled>
                </div>
                <div class="col-xs-8">
                  <input type="text" name="ipfsaddress" class="form-control" value="{{.IPFSAddress}}">
                </div>                
              </div><br/>
                             
              <div class="row">
                <div class="col-xs-3">
                  <input type="text" class="form-control" placeholder="" value="IPNS address" disabled>
                </div>
                <div class="col-xs-8">
                  <input type="text" name="ipnsaddress" class="form-control" value="{{.IPNSAddress}}">
                </div>                
              </div><br/>
              
              
               <div class="row">
                <div class="col-xs-3">
                  <input type="text" class="form-control" value="Oracle address" disabled>
                </div>
                <div class="col-xs-8">
                  <input type="text" name="oracleaddress"  class="form-control" placeholder="ok_address"  value="{{.OracleAddress}}">
                </div>                
              </div><br/>
            
            
                <div class="row">
                <div class="col-xs-3">
                  <input type="text" class="form-control"  value="Web address" disabled>
                </div>
                <div class="col-xs-8">
                  <input type="text" name="webaddress" class="form-control" placeholder="http://...such as blog, facebook, qzone, tweet" value="{{.WebAddress}}">
                </div>                
              </div><br/>
             
               <span class="input-group-btn">
                            <button type="submit" class="btn btn-success btn-flat">Update(Point)</button>
                          </span>
            </form>
            </div>
            <!-- /.box-body -->
          </div>
          <!-- /.box -->
        
        
         <div class="box box-danger">
            <div class="box-header with-border">
              <h3 class="box-title">Update(Point) AENS name Expert Mode</h3>
            </div>
            <div class="box-body">
           <form action="/expertupdatenamepointer" method="post">
            <div class="form-group">
				<input type="hidden" name="aensname" class="form-control"  value="{{.AENSName}}">
                  <label>Pointers json:</label>
                  <textarea class="form-control" rows="6" name="pointerjson" placeholder="Put pointers json here ..."></textarea>
                </div>
           <span class="input-group-btn">
                            <button type="submit" class="btn btn-danger btn-flat">Update(Point)</button>
                          </span>
           </form>
           
            </div>
            
            
            <!-- /.box-body -->
          </div>
          <!-- /.box -->
       </div>
      
      <div class="col-md-5">
		   <div class="box box-info">
		  <div class="box-header with-border">
              <h3 class="box-title">Summary of {{.AENSName}}</h3>
            </div>
		<ul class="list-group list-group-unbordered">
			
			<li class="list-group-item">
                  <i class="glyphicon glyphicon-font"></i><b>Address:</b> <a class="pull-right"/>{{.Account}}</a>
                </li>    
                <li class="list-group-item">
                  <i class="fa fa-btc margin-r-5"></i><b>ID:</b> <a class="pull-right"/>{{.NameID}}</a>
                </li>                                
                
                <li class="list-group-item">
                  <i class="fa fa-calculator margin-r-5"></i><b>Expire height:</b> <a class="pull-right"/>{{.NameTTL}}</a>
                 
                </li>   
              </ul>
              
          <div class="form-group">
                  <label>Name json:</label>
                  <textarea class="form-control" rows="6" placeholder="" disabled>{{.NameJson}}</textarea>
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
