<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>æKnow - Settings</title>
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

	<meta http-equiv="pragma" content="no-cache">
	<!-- HTTP 1.0 -->
	<meta http-equiv="cache-control" content="no-cache">
	<!-- Prevent caching at the proxy server -->
	<meta http-equiv="expires" content="0">

  <!-- Google Font -->
</head>
<body >
<div class="col-md-9">
<section class="content">		
		    <div class="box box-warning">
            <div class="box-header with-border">
              <h3 class="box-title">Haeme Settings</h3>
            </div>
            <!-- /.box-header -->
            <div class="box-body">
              <form role="form"  action="/savesitesetting" method="POST" >
                <!-- text input -->
                <div class="form-group">
                  <label>Title</label>
                  <input type="text" name="title" class="form-control" value="{{.Title}}">
                </div>
                
                <div class="form-group">
                  <label>Subtitle</label>
                  <input type="text" name="subtitle" class="form-control" value="{{.Subtitle}}">
                </div>
                
               <!-- textarea -->
                <div class="form-group">
                  <label>Site Description</label>
                  <textarea class="form-control" name="sitedescription" rows="3" >{{.Description}}</textarea>
                </div>
              
               <div class="form-group">
                  <label>Author</label>
                  <input type="text" name="author" class="form-control" value="{{.Author}}">
                </div>  
              
               <div class="form-group">
                  <label>Author Description</label>
                  <textarea class="form-control" name="authordescription" rows="3" >{{.AuthorDescription}}</textarea>
                </div>
                
              
                <!-- select -->
                <div class="form-group">
                  <label>Theme</label>
                  <select class="form-control" name="theme">
                    <option>aeknow</option>                   
                  </select>
                </div>
               
                 <div class="box-footer">
                <button type="submit" class="btn btn-primary">Save</button>
              </div>
              </form>
            </div>
            <!-- /.box-body -->
          </div>
          <!-- /.box --> 
 	</section>	
</div>
</body>
</html>
