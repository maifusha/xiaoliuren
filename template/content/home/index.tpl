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
                <li role="presentation" class="active"><a href="#dsjx" aria-controls="dsjx" role="tab" data-toggle="tab">断事吉凶</a></li>
                <li role="presentation"><a href="#jsmj" aria-controls="jsmj" role="tab" data-toggle="tab">即时迷津</a></li>
                <li role="presentation"><a href="#mszs" aria-controls="mszs" role="tab" data-toggle="tab">谋事择时</a></li>
            </ul>

            <!-- Tab panes -->
            <div class="tab-content">
                <div role="tabpanel" class="tab-pane active" id="dsjx">{{template "fragment/dsjx" .}}</div>
                <div role="tabpanel" class="tab-pane" id="jsmj">{{template "fragment/jsmj" .}}</div>
                <div role="tabpanel" class="tab-pane" id="mszs">{{template "fragment/jsmj" .}}</div>
            </div>
        </div>
    </div>
</div>
{{template "fragment/footer" .}}