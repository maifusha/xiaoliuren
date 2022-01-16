{{define "fragment/zeshi"}}
<form class="page-header" style="margin-top: 10px;">
    <div class="row row-no-gutters">
        <div class="col-sm-7">
            <span class="help-block">六神起课：所卦事项类型</span>
            <div class="form-group row input-lg" style="padding-top: 0;">
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

        <div class="col-sm-5">
            <span class="help-block">事发日期</span>
            <div class="form-group">
                {{template "fragment/date" .}}
            </div>
        </div>
    </div>
    <div class="row">
        <div class="col-sm-6 col-sm-offset-3">
            <button type="button" class="btn btn-warning btn-block center-block">
                <div class="row">
                    <div class="col-sm-6">择</div>
                    <div class="col-sm-6">时</div>
                </div>
            </button>
        </div>
    </div>
</form>

<div class="panel panel-info">
    <div class="panel-heading">
        <h3 class="panel-title">
            <span class="h4">日 期 > </span>

            <span style="margin-left: 15px;font-weight: 400;">甲子年正月初二</span>
            <small>（<span>2011-01-01</span>）</small>
        </h3>
    </div>
    <div class="panel-body">
        <table class="table table-bordered table-hover table-condensed">
            <tr class="row">
                <td class="col-sm-3 text-center danger" style="vertical-align: middle;">
                    <h3><span class="label label-danger">大安 <small style="color: white;">大吉</small></span></h3>
                </td>
                <td class="col-sm-9">
                    <div class="row">
                        <p class="col-sm-4 text-center"><strong>亥时</strong></p><p class="col-sm-8 text-center">2022-11-01 11:00~13:00</p>
                        <p class="col-sm-4 text-center"><strong>子时</strong></p><p class="col-sm-8 text-center">2022-11-01 11:00~13:00</p>
                    </div>
                </td>
            </tr>
            <tr class="row">
                <td class="col-sm-3 text-center danger" style="vertical-align: middle;">
                    <h3><span class="label label-danger">速喜 <small style="color: white;">中吉</small></span></h3>
                </td>
                <td class="col-sm-9">
                    <div class="row">
                        <p class="col-sm-4 text-center"><strong>亥时</strong></p><p class="col-sm-8 text-center">2022-11-01 11:00~13:00</p>
                        <p class="col-sm-4 text-center"><strong>子时</strong></p><p class="col-sm-8 text-center">2022-11-01 11:00~13:00</p>
                    </div>
                </td>
            </tr>
            <tr class="row">
                <td class="col-sm-3 text-center danger" style="vertical-align: middle;">
                    <h3><span class="label label-danger">小吉 <small style="color: white;">小吉</small></span></h3>
                </td>
                <td class="col-sm-9">
                    <div class="row">
                        <p class="col-sm-4 text-center"><strong>亥时</strong></p><p class="col-sm-8 text-center">2022-11-01 11:00~13:00</p>
                        <p class="col-sm-4 text-center"><strong>子时</strong></p><p class="col-sm-8 text-center">2022-11-01 11:00~13:00</p>
                    </div>
                </td>
            </tr>
            <tr class="row">
                <td class="col-sm-3 text-center success" style="vertical-align: middle;">
                    <h3><span class="label label-success">留连 <small style="color: white;">小凶</small></span></h3>
                </td>
                <td class="col-sm-9">
                    <div class="row">
                        <p class="col-sm-4 text-center"><strong>亥时</strong></p><p class="col-sm-8 text-center">2022-11-01 11:00~13:00</p>
                        <p class="col-sm-4 text-center"><strong>子时</strong></p><p class="col-sm-8 text-center">2022-11-01 11:00~13:00</p>
                    </div>
                </td>
            </tr>
            <tr class="row">
                <td class="col-sm-3 text-center success" style="vertical-align: middle;">
                    <h3><span class="label label-success">赤口 <small style="color: white;">中凶</small></span></h3>
                </td>
                <td class="col-sm-9">
                    <div class="row">
                        <p class="col-sm-4 text-center"><strong>亥时</strong></p><p class="col-sm-8 text-center">2022-11-01 11:00~13:00</p>
                        <p class="col-sm-4 text-center"><strong>子时</strong></p><p class="col-sm-8 text-center">2022-11-01 11:00~13:00</p>
                    </div>
                </td>
            </tr>
            <tr class="row">
                <td class="col-sm-3 text-center success" style="vertical-align: middle;">
                    <h3><span class="label label-success">空亡 <small style="color: white;">大凶</small></span></h3>
                </td>
                <td class="col-sm-9">
                    <div class="row">
                        <p class="col-sm-4 text-center"><strong>亥时</strong></p><p class="col-sm-8 text-center">2022-11-01 11:00~13:00</p>
                        <p class="col-sm-4 text-center"><strong>子时</strong></p><p class="col-sm-8 text-center">2022-11-01 11:00~13:00</p>
                    </div>
                </td>
            </tr>
        </table>
    </div>
</div>
{{end}}