{{define "contentsSidebar"}}
<div class="col-md-3">

    <!-- Blog Categories Well -->
    <div class="well">
        <h4>Blog Categories</h4>
        <div class="row">
            <!-- <div class="col-lg-6"> -->
                <!-- <ul class="list-unstyled"> -->
                <ul>
                {{range .Categories}}
            <li><a href="/?cate={{.Title}}">{{.Title}} </a></li>
                    </li>
                {{end}}
                </ul>
        </div>
    </div>
<div class="well">
    <a href="">
    <!-- 廣告： -->
         <h4><<桶心>>純天然食用油</h4>
        <p>關心您的健康，請用純天然、 純物理提煉製造再生食用油</p>
    </a>
    <small>一天一杯，健康活力</small>
</div>
</div>

{{end}}