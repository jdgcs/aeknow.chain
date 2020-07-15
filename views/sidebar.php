 {{ define "sidebar" }}
  <aside class="main-sidebar">
    <!-- sidebar: style can be found in sidebar.less -->
    <section class="sidebar">
      <!-- sidebar menu: : style can be found in sidebar.less -->
      <ul class="sidebar-menu" data-widget="tree">
        <li>
          <a href="/" >
            <i class="fa fa-dashboard"></i> <span>Dashboard</span>          
          </a>
        </li>
               
        <!--  <li class="active treeview menu-open"> -->
		 <li class="active treeview menu-open">
          <a href="/wallet">
            <i class="glyphicon glyphicon-briefcase"></i>
            <span>Assets</span>
            <span class="pull-right-container">
              <i class="fa fa-angle-left pull-right"></i>
            </span>
          </a>
          <ul class="treeview-menu">
            <li><a href="/wallet"><i class="glyphicon glyphicon-font"></i> AE</a></li>
            <li><a href="/aens"><i class="glyphicon glyphicon-retweet"></i> AENS</a></li>
            <li><a href="/token"><i class="glyphicon glyphicon-briefcase"></i> AEX-9 Tokens</a></li>
            <li><a href="/contracts"><i class="glyphicon glyphicon-list-alt"></i>Contracts</a></li>
             <!--<li><a href="pages/charts/inline.html"><i class="fa fa-object-group"></i> GA</a></li>
            <li><a href="pages/charts/inline.html"><i class="glyphicon glyphicon-link"></i> Orcales</a></li>-->
            
          </ul>
        </li>
        <li><a href="/chat?topic=ak_fCCw1JEkvXdztZxk8FRGNAkvmArhVeow89e64yX4AxbCPrVh5"><i class="fa fa-users"></i> <span>Chaet</span></a></li>
        
        
        <li><a href="/views/haeme.html"><i class="glyphicon glyphicon-pencil"></i> <span>Haeme</span></a></li>
                  
         <li><a href="/files"><i class="fa fa-folder-open"></i> <span>Files</span></a></li>       
         <li>
          <a href="#" >
            <i class="glyphicon glyphicon-envelope"></i> <span>Messages </span>  
             <span class="pull-right-container">
              <small class="label pull-right bg-yellow">12</small>
              <small class="label pull-right bg-green">16</small>
              <small class="label pull-right bg-red">5</small>
            </span>        
          </a>
        </li>
        
      </ul>
      
      <ul class="sidebar-menu" data-widget="tree">
		   <li class="header"> <center><h4>Navigation</h4></center></li>  
		   <li><a href="/views/haeme.html"><i class="fa fa-list"></i> <span>Browser</span></a></li>
		   <li><a href="/views/haeme.html"><i class="glyphicon glyphicon-th"></i> <span>Nodes Category</span></a></li>
		    <li><a href="/views/haeme.html"><i class="fa fa-search-plus"></i> <span>Search</span></a></li>
		    <li><a href="/views/haeme.html"><i class="fa fa-connectdevelop"></i> <span>DEX</span></a></li>
      </ul>
      
       <ul class="sidebar-menu" data-widget="tree">
		   <li class="header"> <center><h4>Settings</h4></center></li>  
      
        
        <li>
          <a href="/views/haeme.html" >
            <i class="glyphicon glyphicon-cog"></i><span>Local</span>          
          </a>
        </li>
        <li><a href="/export"><i class="fa fa-exchange"></i> <span>Backup/Restore</span></a></li>
       </ul>  
        
    </section>
    <!-- /.sidebar -->
  </aside>
{{ end }}
