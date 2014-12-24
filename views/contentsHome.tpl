{{define "contentsHome"}}
<!--  -->
<div class="col-md-8">

    <h1 class="page-header">
    Blog Home
        <small>by Pike</small>
    </h1>

    <!-- First Blog Post -->
    {{range .Topics}}
        <h2><a href="/topic/view/{{.Id}}">{{.Title}} </a></h2>
        <!-- <a href="#">Blog Post Title</a> -->
    <hr>
        <p>
            {{.Content}}
        </p>
    <span class="pull-right">
    <p><span class="glyphicon glyphicon-time"></span>Posted on {{.Created}}</p>
    <!-- <h4> -->
    </span>
    <!-- <div class="pull-right"> -->
    <span class="text-right">
 <span class="badge">{{.Views}}</span>views| <span class="badge"> {{.ReplyCount}} </span>comments 
 <!-- </h4> -->
    </span>
 <!-- </div> -->
   <hr>
    <a class="btn btn-info" href="/topic/view/{{.Id}}">回覆<span class="glyphicon glyphicon-chevron-right"></span></a>
    <hr>

{{end}}
</div>
{{end}}