{{define "fragment/jixiong"}}
<form class="page-header" style="margin-top: 10px;">
    <div class="row row-no-gutters">
        <div class="col-sm-5">
            <span class="help-block">六神起课：所卦事项类型</span>
            <div class="form-group row row-no-gutters qike_list" style="padding-top: 0;line-height: 2;">
                {{range $k,$v := .qikeList}}
                <div class="radio">
                    <label>
                        <input type="radio" class="qike" name="qike" value="{{$v.ID}}">
                        <span class="label label-primary">{{$v.Name}}</span> {{$v.Suozhu}} <small class="text-info">({{$v.Shiergong}})</small>
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
                        <input type="radio" class="dizhi" name="dizhi" value="{{$k}}">
                        <span class="label label-default">{{index $v 0}}</span> {{index $v 1}}
                    </label>
                </div>
                {{end}}
            </div>
        </div>
    </div>
    <div class="row" style="margin-top: 30px;margin-bottom: 20px;">
        <div class="col-sm-6 col-sm-offset-3">
            <button type="button" class="btn btn-warning btn-block center-block qigua">
                <div class="row"><div class="col-sm-6">起</div><div class="col-sm-6">卦</div></div>
            </button>
        </div>
    </div>
</form>

<div class="panel panel-info jieke hidden">
    <div class="panel-heading">
        <h3 class="panel-title">
            <span class="h4">解 课</span> <span class="glyphicon glyphicon-play" aria-hidden="true"></span>
            <span class="lunar_time" style="margin-left: 15px;font-weight: 400;"></span>
            <small>（<span class="solar_time"></span>）</small>
        </h3>
    </div>
    <div class="panel-body">
        <div class="row">
            <div class="col-sm-5" style="position: relative;">
                <img src="/static/image/sixiangbagua.jpg" class="img-responsive center-block img-thumbnail" alt="Responsive image">
                <h4 class="yueke_mark" style="position: absolute;margin: 0px;"><span class="label label-danger"><b>月 课</b></span></h4>
                <h4 class="rike_mark" style="position: absolute;margin: 0px;"><span class="label label-danger"><b>日 课</b></span></h4>
                <h4 class="shike_mark" style="position: absolute;margin: 0px;"><span class="label label-danger"><b>时 课</b></span></h4>
            </div>
            <div class="col-sm-7">
                <table class="table table-bordered table-hover" style="padding-left: 20px;">
                    <tr class="row yueke">
                        <td class="active text-center col-sm-4" style="vertical-align: middle;"><strong>月课</strong> <small>(天盘:事项初期)</small></td>
                        <td class="col-sm-8">
                            <h4><span class="label liushen_label"><b></b> <small style="color: white;"></small></span></h4>
                            <span class="liushen_shiyi"></span>
                        </td>
                    </tr>
                    <tr class="row rike">
                        <td class="active text-center col-sm-4" style="vertical-align: middle;"><strong>日课</strong> <small>(地盘:事项过程)</small></td>
                        <td class="col-sm-8">
                            <h4><span class="label liushen_label"><b></b> <small style="color: white;"></small></span></h4>
                            <span class="liushen_shiyi"></span>
                        </td>
                    </tr>
                    <tr class="row shike">
                        <td class="active text-center col-sm-4" style="vertical-align: middle;"><strong>时课</strong> <small>(人盘:事项结果)</small></td>
                        <td class="col-sm-8">
                            <h4><span class="label liushen_label"><b></b> <small style="color: white;"></small></span></h4>
                            <span class="liushen_shiyi"></span>
                        </td>
                    </tr>
                </table>
            </div>
        </div>
    </div>
</div>
<script type="text/javascript">
$(function () {
    $("#jixiong").delegate(".qike_list input", "change", function () {
        var qike = $("#jixiong .qike:checked").val();
        zhangjue.checkLiushen(qike);
    });

    $("#jixiong").delegate(".qigua", "click", function (){
        var qike = $("#jixiong .qike:checked").val();
        if (!qike) {return alert("请选择六神起课");}

        var date = $("#jixiong .date").val();
        if (!date) {return alert("请选择日期");}

        var dizhi = $("#jixiong .dizhi:checked").val();
        if (!dizhi) {return alert("请选择时辰");}

        zhangjue.checkLiushen(qike);
        zhangjue.loading();
        $("#jixiong .jieke").addClass("hidden");

        $.ajax({
            type: "GET",
            url: "/home/jixiong",
            data: {
                qike: qike,
                date: date,
                dizhi: dizhi,
            },
            dataType: "json",
        }).then(function (resp) {
            return zhangjue.fingerCount(resp.finger_count, resp)
        }).then(function(resp){
            $("#jixiong .lunar_time").text(resp.lunar_time);
            $("#jixiong .solar_time").text(resp.solar_time);

            $(".yueke .liushen_label b").text(resp.sangong.yueke.name);
            $(".yueke .liushen_label small").text(resp.sangong.yueke.jixiong);
            $(".yueke .liushen_shiyi").text(resp.sangong.yueke.shiyi);

            $(".rike .liushen_label b").text(resp.sangong.rike.name);
            $(".rike .liushen_label small").text(resp.sangong.rike.jixiong);
            $(".rike .liushen_shiyi").text(resp.sangong.rike.shiyi);

            $(".shike .liushen_label b").text(resp.sangong.shike.name);
            $(".shike .liushen_label small").text(resp.sangong.shike.jixiong);
            $(".shike .liushen_shiyi").text(resp.sangong.shike.shiyi);

            markSangong(resp.sangong.yueke.index,resp.sangong.rike.index,resp.sangong.shike.index);

            var jixiang = /吉/;
            if (jixiang.test(resp.sangong.yueke.jixiong)) {
                $(".yueke .liushen_label").removeClass("label-success").addClass("label-danger");
            } else {
                $(".yueke .liushen_label").removeClass("label-danger").addClass("label-success");
            }
            if (jixiang.test(resp.sangong.rike.jixiong)) {
                $(".rike .liushen_label").removeClass("label-success").addClass("label-danger");
            } else {
                $(".rike .liushen_label").removeClass("label-danger").addClass("label-success");
            }
            if (jixiang.test(resp.sangong.shike.jixiong)) {
                $(".shike .liushen_label").removeClass("label-success").addClass("label-danger");
            } else {
                $(".shike .liushen_label").removeClass("label-danger").addClass("label-success");
            }

            zhangjue.loaded();
            $("#jixiong .jieke").removeClass("hidden");
        });
    });

    function markSangong(yuekeIndex, rikeIndex, shikeIndex) {
        var position = {
            1: {top:"34%",left:"6%"},
            2: {top:"3%",left:"6%"},
            3: {top:"3%",left:"35%"},
            4: {top:"34%",left:"65%"},
            5: {top:"67%",left:"35%"},
            6: {top:"34%",left:"35%"},
        };

        $(".yueke_mark").css("top",position[yuekeIndex].top).css("left",position[yuekeIndex].left);
        $(".rike_mark").css("top",position[rikeIndex].top).css("left",position[rikeIndex].left);
        $(".shike_mark").css("top",position[shikeIndex].top).css("left",position[shikeIndex].left);

        $(".rike_mark").css("margin-top","0px");
        $(".shike_mark").css("margin-top","0px");

        if (rikeIndex==yuekeIndex) {
            $(".rike_mark").css("margin-top","24px");
        }

        if (shikeIndex==yuekeIndex && shikeIndex==rikeIndex) {
            $(".shike_mark").css("margin-top","48px")
        } else {
            if (shikeIndex==yuekeIndex || shikeIndex==rikeIndex) {
                $(".shike_mark").css("margin-top","24px");
            }
        }
    }
});
</script>
{{end}}