{{define "fragment/jixiong"}}
<form class="page-header" style="margin-top: 10px;">
    <div class="row row-no-gutters">
        <div class="col-sm-5">
            <span class="help-block">六神起课：所卦事项类型</span>
            <div class="form-group row row-no-gutters" style="padding-top: 0;line-height: 2;">
                {{range .qikeList}}
                <div class="radio">
                    <label>
                        <input type="radio" name="qike" value="{{.ID}}">
                        <span class="label label-primary">{{.Name}}</span> {{.Suozhu}} <small class="text-info">({{.Shiergong}})</small>
                    </label>
                </div>
                {{end}}
            </div>
        </div>

        <div class="col-sm-3 col-sm-pull-1" style="padding-left: 5%;">
            <span class="help-block">事发日期</span>
            <div class="form-group">
                {{template "fragment/date" .}}
            </div>
        </div>

        <div class="col-sm-4 form-inline" style="padding-left: 10px;">
            <span class="help-block">事发时辰</span>
            <div class="form-group row row-no-gutters" style="line-height: 2.5;">
                {{range $k,$v := .dizhiList}}
                <div class="radio col-sm-6">
                    <label>
                        <input type="radio" name="dizhi" value="{{$k}}">
                        <span class="label label-default">{{index $v 0}}</span> {{index $v 1}}
                    </label>
                </div>
                {{end}}
            </div>
        </div>
    </div>
    <div class="row" style="margin-top: 30px;margin-bottom: 20px;">
        <div class="col-sm-6 col-sm-offset-3">
            <button type="button" class="btn btn-warning btn-block center-block">
                <div class="row"><div class="col-sm-6">起</div><div class="col-sm-6">卦</div></div>
            </button>
        </div>
    </div>
</form>

<div class="panel panel-info">
    <div class="panel-heading">
        <h3 class="panel-title">
            <span class="h4">解 课 > </span>
            <span style="margin-left: 15px;font-weight: 400;">甲子年正月初二子时</span>
            <small>（<span>2011-01-01 13:22</span>）</small>
        </h3>
    </div>
    <div class="panel-body">
        <div class="row">
            <div class="col-sm-5">
                <img src="/static/image/sixiangbagua.jpg" class="img-responsive center-block img-thumbnail" alt="Responsive image">
            </div>
            <div class="col-sm-7">
                <table class="table table-bordered table-hover" style="padding-left: 20px;">
                    <tr class="row">
                        <td class="active text-center col-sm-3" style="vertical-align: middle;"><strong>月盘</strong> <small>(事项初期)</small></td>
                        <td class="col-sm-9">
                            <h4><span class="label label-danger">大安 <small style="color: white;">大吉</small></span></h4>
                            <span>xxxxxxxxxxxxx</span>
                        </td>
                    </tr>
                    <tr class="row">
                        <td class="active text-center col-sm-3" style="vertical-align: middle;"><strong>日盘</strong> <small>(事项过程)</small></td>
                        <td class="col-sm-9">
                            <h4><span class="label label-success">留连 <small style="color: white;">小凶</small></span></h4>
                            <span>xxxxxxxxxxxxx</span>
                        </td>
                    </tr>
                    <tr class="row">
                        <td class="active text-center col-sm-3" style="vertical-align: middle;"><strong>时盘</strong> <small>(事项结果)</small></td>
                        <td class="col-sm-9">
                            <h4><span class="label label-danger">大安 <small style="color: white;">大吉</small></span></h4>
                            <span>xxxxxxxxxxxxx</span>
                        </td>
                    </tr>
                </table>
            </div>
        </div>
    </div>
</div>
{{end}}