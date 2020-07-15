
<!DOCTYPE html>
<html lang="zh">
    <head>
        <meta charset="utf-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
        <meta http-equiv="Cache-Control" content="no-siteapp"/>
		<meta name="renderer" content="webkit" />
		<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" />
       
        <title>Editor</title>
       <link rel="stylesheet" href="/views/editor.md/examples/css/style.css" />
       <link rel="stylesheet" href="/views/editor.md/css/editormd.css" />
    </head>
    <body>
<style type="text/css">	
	input[type=text] {
    width: 500px;
    padding: 0 25px;
    height: 48px;
    border: 1px solid #f2f2f2;
    background: #f6f6f6;
    color: #202124;
    font-size: 14px;
    line-height: 48px;
    border-radius: 25px;
}

	textarea {
    width: 500px;
    padding: 0 10px;   
    border: 1px solid #f2f2f2;
    background: #f6f6f6;
    color: #202124;
    font-size: 14px;   
    border-radius: 25px;
}
</style>

<form class="form-horizontal" action="/saveblog" method="POST" >
<table style="margin-left:10px;;">	
<tr>
<td width="60px" align="left" border="10">
<b>Title:</b>
</td>
<td>
<input type="text" name="title" placeholder="title" style="margin-left:10px;" value="{{.PageTitle}}">
</td>
</tr>


<tr>
<td width="60px" align="left">
<b>Keywords:</b>
</td>
<td>
<input type="text"  name="tags" placeholder="Keywords of the article" style="margin-left:10px;" value="{{.PageTags}}">
 <div class="checkbox">
                  <label>
                    <input type="checkbox"  name="draft"> Draft
                  </label>
                </div>
</td>
</tr>


<tr>
<td width="60px" align="left" >
<b>Categories:</b>
</td>
<td>
<input type="text" name="categories" placeholder="Categories"  style="margin-left:10px;" value="{{.PageCategories}}">
</td>
</tr>


<tr>
<td width="60px" align="left"  valign="top">
<b>Abstract:</b>
</td>
<td align="left">
<textarea rows="6" name="description" style="margin-left:10px;width:100%">{{.PageDescription}}</textarea>
</td>
</tr>


</table>
 <input type="hidden" name="editpath" value="{{.EditPath}}">
		 <div id="test-editormd">
                <textarea style="display:none;" name="content">{{.PageContent}}</textarea>
            </div>
        </div>
        <button type="submit" class="btn btn-success pull-left" style="background-color:green;color:white">Save</button>

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
        </body>
</html>
