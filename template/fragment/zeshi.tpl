{{define "fragment/zeshi"}}
<form class="page-header" style="margin-top: 10px;">
    <div class="row row-no-gutters">
        <div class="col-sm-7">
            <span class="help-block">六神起课：所卦事项类型</span>
            <div class="form-group row input-lg qike_list" style="padding-top: 0;">
                {{range .qikeList}}
                <div class="radio">
                    <label>
                        <input type="radio" class="qike" name="qike" value="{{.ID}}">
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
            <button type="button" class="btn btn-warning btn-block center-block zeshi">
                <div class="row">
                    <div class="col-sm-6">择</div>
                    <div class="col-sm-6">时</div>
                </div>
            </button>
        </div>
    </div>
</form>

<div class="panel panel-info hidden jieke">
    <div class="panel-heading">
        <h3 class="panel-title">
            <span class="h4">日 期</span> <span class="glyphicon glyphicon-play" aria-hidden="true"></span>

            <span class="lunar" style="margin-left: 15px;font-weight: 400;"></span>
            <small>（<span class="solar"></span>）</small>
        </h3>
    </div>
    <div class="panel-body">
        <table class="table table-bordered table-hover table-condensed">
            <tr class="row">
                <td class="col-sm-3 text-center danger" style="vertical-align: middle;">
                    <h3><span class="label label-danger">大安 <small style="color: white;">大吉</small></span></h3>
                </td>
                <td class="col-sm-9">
                    <div class="row daan_row">
                    </div>
                </td>
            </tr>
            <tr class="row">
                <td class="col-sm-3 text-center danger" style="vertical-align: middle;">
                    <h3><span class="label label-danger">速喜 <small style="color: white;">中吉</small></span></h3>
                </td>
                <td class="col-sm-9">
                    <div class="row suxi_row">
                    </div>
                </td>
            </tr>
            <tr class="row">
                <td class="col-sm-3 text-center danger" style="vertical-align: middle;">
                    <h3><span class="label label-danger">小吉 <small style="color: white;">小吉</small></span></h3>
                </td>
                <td class="col-sm-9">
                    <div class="row xiaoji_row">
                    </div>
                </td>
            </tr>
            <tr class="row">
                <td class="col-sm-3 text-center success" style="vertical-align: middle;">
                    <h3><span class="label label-success">留连 <small style="color: white;">小凶</small></span></h3>
                </td>
                <td class="col-sm-9">
                    <div class="row liulian_row">
                    </div>
                </td>
            </tr>
            <tr class="row">
                <td class="col-sm-3 text-center success" style="vertical-align: middle;">
                    <h3><span class="label label-success">赤口 <small style="color: white;">中凶</small></span></h3>
                </td>
                <td class="col-sm-9">
                    <div class="row chikou_row">
                    </div>
                </td>
            </tr>
            <tr class="row">
                <td class="col-sm-3 text-center success" style="vertical-align: middle;">
                    <h3><span class="label label-success">空亡 <small style="color: white;">大凶</small></span></h3>
                </td>
                <td class="col-sm-9">
                    <div class="row kongwang_row">
                    </div>
                </td>
            </tr>
        </table>
    </div>
</div>
<script type="text/javascript">
$(function () {
    $("#zeshi").delegate(".qike_list input", "change", function () {
        var qike = $("#zeshi .qike:checked").val();
        zhangjue.checkLiushen(qike);
    });

    $("#zeshi").delegate(".zeshi", "click", function () {
        var qike = $("#zeshi .qike:checked").val();
        if (!qike) {return alert("请选择六神起课");}

        var date = $("#zeshi .date").val();
        if (!date) {return alert("请选择日期");}

        $.ajax({
            type: "GET",
            url: "/home/zeshi",
            data: {
                qike: qike,
                date: date,
            },
            dataType: "json",
        }).then(function (resp) {
            $("#zeshi .lunar").text(resp.lunar);
            $("#zeshi .solar").text(resp.solar);

            $(".daan_row").empty();
            resp.liushen.daan.forEach(function (one) {
                var tpl = '<p class="col-sm-4 text-center"><span class="glyphicon glyphicon-time" aria-hidden="true"></span> <strong>{dizhi_name}</strong></p><p class="col-sm-8 text-center">{solar_time}</p>';
                var row = tpl.replace("{dizhi_name}",one.dizhi_name).replace("{solar_time}",one.solar_time);
                $(".daan_row").append(row);
            });
            $(".liulian_row").empty();
            resp.liushen.liulian.forEach(function (one) {
                var tpl = '<p class="col-sm-4 text-center"><span class="glyphicon glyphicon-time" aria-hidden="true"></span> <strong>{dizhi_name}</strong></p><p class="col-sm-8 text-center">{solar_time}</p>';
                var row = tpl.replace("{dizhi_name}",one.dizhi_name).replace("{solar_time}",one.solar_time);
                $(".liulian_row").append(row);
            });
            $(".suxi_row").empty();
            resp.liushen.suxi.forEach(function (one) {
                var tpl = '<p class="col-sm-4 text-center"><span class="glyphicon glyphicon-time" aria-hidden="true"></span> <strong>{dizhi_name}</strong></p><p class="col-sm-8 text-center">{solar_time}</p>';
                var row = tpl.replace("{dizhi_name}",one.dizhi_name).replace("{solar_time}",one.solar_time);
                $(".suxi_row").append(row);
            });
            $(".chikou_row").empty();
            resp.liushen.chikou.forEach(function (one) {
                var tpl = '<p class="col-sm-4 text-center"><span class="glyphicon glyphicon-time" aria-hidden="true"></span> <strong>{dizhi_name}</strong></p><p class="col-sm-8 text-center">{solar_time}</p>';
                var row = tpl.replace("{dizhi_name}",one.dizhi_name).replace("{solar_time}",one.solar_time);
                $(".chikou_row").append(row);
            });
            $(".xiaoji_row").empty();
            resp.liushen.xiaoji.forEach(function (one) {
                var tpl = '<p class="col-sm-4 text-center"><span class="glyphicon glyphicon-time" aria-hidden="true"></span> <strong>{dizhi_name}</strong></p><p class="col-sm-8 text-center">{solar_time}</p>';
                var row = tpl.replace("{dizhi_name}",one.dizhi_name).replace("{solar_time}",one.solar_time);
                $(".xiaoji_row").append(row);
            });
            $(".kongwang_row").empty();
            resp.liushen.kongwang.forEach(function (one) {
                var tpl = '<p class="col-sm-4 text-center"><span class="glyphicon glyphicon-time" aria-hidden="true"></span> <strong>{dizhi_name}</strong></p><p class="col-sm-8 text-center">{solar_time}</p>';
                var row = tpl.replace("{dizhi_name}",one.dizhi_name).replace("{solar_time}",one.solar_time);
                $(".kongwang_row").append(row);
            });

            $("#zeshi .jieke").removeClass("hidden");
        });
    });
});
</script>
{{end}}