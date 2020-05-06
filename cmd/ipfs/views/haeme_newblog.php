<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>æKnow - Dashboard</title>
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
  
  <link rel="stylesheet" href="/views/editor.md/examples/css/style.css" />
  <link rel="stylesheet" href="/views/editor.md/css/editormd.css" />

  <!-- Google Font -->
</head>
<body class="hold-transition skin-blue sidebar-mini">
<div class="wrapper">
	
 <header class="main-header">  
    <!-- Header Navbar: style can be found in header.less -->
    <nav class="navbar navbar-static-top"  style="background:#f7296e;margin-left:0px;"> 
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

  <!-- Content Wrapper. Contains page content -->
  <div class="content-wrapper" style="margin-left:0px;padding-left:0px;">
    <!-- Main content -->
    <div class="box">		
		 <form class="form-horizontal" action="/saveblog" method="POST">
		<div class="col-md-9" style="margin-left:0px;padding-left:0px;">	 
		 <div class="form-group">
                  <label class="col-sm-2 control-label" style="margin-left:0px;padding-left:0px;"><font color=red>*</font>Title</label>

                  <div class="col-sm-10">
                    <input type="text" class="form-control" name="title" placeholder="title">
                  </div>
                </div>
                       
          <div class="form-group">
                  <label class="col-sm-2 control-label">Keywords:</label>

                  <div class="col-sm-10">
                    <input type="text" class="form-control" name="tags" placeholder="Keywords of the article">
                  </div>
                </div>
          
          
          <div class="form-group">
                  <label class="col-sm-2 control-label">Categories</label>
                  <div class="col-sm-10">
                    <input type="text" class="form-control" name="categories" placeholder="Categories">
                  </div>
                </div> 
           <div class="form-group">
                 
          </div> 
               
         </div>
         
         <input type="hidden" name="editpath" value="{{.EditPath}}">
		 <div id="test-editormd">
                <textarea style="display:none;" name="content"></textarea>
            </div>
        </div>
        <button type="submit" class="btn btn-success pull-left" style="background-color:green;">Post</button>

        <script src="/views/editor.md/examples/js/jquery.min.js"></script>
        <script src="/views/editor.md/editormd.min.js"></script>      
        <script type="text/javascript">
			var testEditor;

            $(function() {
                testEditor = editormd("test-editormd", {
                    width   : "100%",
                    height  : 640,
                    syncScrolling : "single",
                    tex:true,
                    imageUpload : true,
                    imageUploadURL : "/uploadblogimage",
                    imageFormats      : ["jpg", "jpeg", "gif", "png", "bmp", "webp","mp4","avi"],
                    codeFold : true,
                    taskList : true,                   
                    placeholder : "Enjoy Knowledge! writing now...",
                    htmlDecode : true,
                    path    : "/views/editor.md/lib/"
                });
                
                /*
                // or
                testEditor = editormd({
                    id      : "test-editormd",
                    width   : "90%",
                    height  : 640,
                    path    : "/views/editor.md/lib/"
                });
                */
            });
        </script>
	
	 </form>
	 
	 </div>	
		<!-- /.content -->
</div>

</div>
<!-- ./wrapper -->

<!-- jQuery 3 -->

<!-- Bootstrap 3.3.7 -->
<script src="/views/static/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>

<!-- AdminLTE App -->
<script src="/views/static/dist/js/adminlte.min.js"></script>

</body>
</html>
