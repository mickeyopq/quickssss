{{define "navbar"}}
<nav class="navbar navbar-inverse navbar-fixed-top">
	<div class="container">

		<a href="/welcome" class="navbar-brand">My blog</a>
	<div>
		<ul class="nav navbar-nav">
			<li {{if .IsHome}}class="active"{{end}}><a href="/">首頁</a></li>
			<li {{if .IsCategory}}class="active"{{end}}><a href="/category">分類</a></li>
			<li {{if .IsTopic}}class="active"{{end}}><a href="/topic">文章</a></li>
		</ul>
	</div>

<div class="pull-right">
	<ul class="nav navbar-nav">
	{{if .IsLogin}}
		<li><a href="/login?exit=true" onclick="clrcookie()">退出</a></li>
		{{else}}
		<li><a href="/login">管理員登錄</a></li>
		{{end}}
	</ul>
</div>
</div>
</nav>

	<!-- 新增一個jquery-cookie套件 -->
	{{template "js"}}
<script type="text/javascript" src="http://cdn.staticfile.org/jquery-cookie/1.4.1/jquery.cookie.min.js"></script>
<script type="text/javascript">
	function clrcookie () {
	$.cookie("uname", null);
	$.cookie("pwd", null);
	return
}
</script>

{{end}}