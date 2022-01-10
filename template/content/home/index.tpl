{{template "fragment/header" .}}
<div class="container">
    <div class="row">
        <div class="col-sm-12">
            <h4 class="text-center">无事勿卦，卦后不复，心诚则灵 <small>道传小六壬</small></h4>
            <img src="/static/image/zhangjue.jpg" class="img-responsive center-block img-rounded" alt="Responsive image">
        </div>
    </div>
    <div class="row">
        <div class="col-sm-8 col-sm-offset-2 well">
            <!-- Nav tabs -->
            <ul class="nav nav-tabs" role="tablist">
                <li role="presentation" class="active"><a href="#jixiong" aria-controls="jixiong" role="tab" data-toggle="tab">断事吉凶</a></li>
                <li role="presentation"><a href="#dianbo" aria-controls="dianbo" role="tab" data-toggle="tab">即时点拨</a></li>
                <li role="presentation"><a href="#zeshi" aria-controls="zeshi" role="tab" data-toggle="tab">谋事择时</a></li>
            </ul>

            <!-- Tab panes -->
            <div class="tab-content">
                <div role="tabpanel" class="tab-pane active" id="jixiong">{{template "fragment/jixiong" .}}</div>
                <div role="tabpanel" class="tab-pane" id="dianbo">{{template "fragment/dianbo" .}}</div>
                <div role="tabpanel" class="tab-pane" id="zeshi">{{template "fragment/zeshi" .}}</div>
            </div>
        </div>
    </div>
</div>
{{template "fragment/footer" .}}