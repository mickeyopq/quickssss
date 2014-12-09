{{define "navbar"}}
		<a href="/" class="navbar-brand">My blog</a>
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
		<li><a href="/login?exit=true">退出</a></li>
		{{else}}
		<li><a href="/login">管理員登錄</a></li>
		{{end}}
	</ul>
</div>

{{end}}