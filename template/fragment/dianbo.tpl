{{define "fragment/dianbo"}}
<div class="row page-header">
    <div class="col-sm-5">
        <blockquote><span class="glyphicon glyphicon-heart" aria-hidden="true"></span> 临时起意，即刻起卦，指点迷津。</blockquote>
    </div>
    <div class="col-sm-4 col-sm-offset-2">
        <button type="button" class="btn btn-warning btn-block center-block qigua">
            <div class="row"><div class="col-sm-6">起</div><div class="col-sm-6">卦</div></div>
        </button>
    </div>
</div>

<div class="panel panel-info jieke hidden">
    <div class="panel-heading">
        <h3 class="panel-title">
            <span class="h4">解 课</span> <span class="glyphicon glyphicon-play" aria-hidden="true"></span>

            <span class="lunar_time" style="margin-left: 15px;font-weight: 400;"></span>
            <small>（<span class="solar_time"></span>）</small>
        </h3>
    </div>
    <div class="panel-body">
        <table class="table table-bordered">
            <tr class="row">
                <td class="active text-center col-sm-2" style="vertical-align: middle;"><strong>落宫</strong></td>
                <td class="col-sm-4">
                    <h4><span class="label luogong_label"><b></b> <small style="color: white;"></small></span></h4>
                    <span class="luogong_shiyi"></span>
                </td>
                <td class="active text-center col-sm-2" style="vertical-align: middle;"><strong>机缘数</strong></td>
                <td class="col-sm-4">
                    <p>排位数 <mark class="paiweishu"></mark></p>
                    <p>宫位数 <mark class="gongweishu"></mark></p>
                </td>
            </tr>
            <tr class="row">
                <td class="active text-center col-sm-2" style="vertical-align: middle;"><strong>贵人冲犯</strong></td>
                <td class="col-sm-4 guirenchongfan"></td>
                <td class="active text-center col-sm-2" style="vertical-align: middle;"><strong>机缘方位</strong></td>
                <td class="col-sm-4 jiyuanfangwei"></td>
            </tr>
            <tr class="row">
                <td class="active text-center col-sm-2" style="vertical-align: middle;"><strong>五行/神煞</strong></td>
                <td class="col-sm-4"><span class="wuxin"></span> / <span class="shensha"></span></td>
                <td class="active text-center col-sm-2" style="vertical-align: middle;"><strong>八卦</strong></td>
                <td class="col-sm-4 bagua"></td>
            </tr>
        </table>

        <div class="row" style="margin: 0;">
            <div class="col-sm-6">
                <strong><span class="glyphicon glyphicon-list" aria-hidden="true"></span> 解惑</strong>
                <table class="table table-bordered table-hover jiehuo_list">
                </table>
            </div>
            <div class="col-sm-6">
                <strong><span class="glyphicon glyphicon-th" aria-hidden="true"></span> 断辞</strong>
                <table class="table table-bordered table-striped table-hover duanci_list">
                </table>
            </div>
        </div>
    </div>
</div>
<script type="text/javascript">
$(function () {
    $("#dianbo").delegate(".qigua", "click", function (){
        zhangjue.checkLiushen(1);
        zhangjue.loading();
        $("#dianbo .jieke").addClass("hidden");

        $.ajax({
            type: "GET",
            url: "/home/dianbo",
            dataType: "json"
        }).then(function (resp) {
            return zhangjue.fingerCount(resp.finger_count, resp)
        }).then(function (resp){
            $("#dianbo .lunar_time").text(resp.lunar_time);
            $("#dianbo .solar_time").text(resp.solar_time);

            var jixiang = /吉/;
            if (jixiang.test(resp.liushen.jixiong)) {
                $(".luogong_label").removeClass("label-success").addClass("label-danger");
            } else {
                $(".luogong_label").removeClass("label-danger").addClass("label-success");
            }
            $(".luogong_label b").text(resp.liushen.name);
            $(".luogong_label small").text(resp.liushen.jixiong);
            $(".luogong_shiyi").text(resp.liushen.shiyi);
            $(".guirenchongfan").text(resp.liushen.guirenchongfan);
            $(".paiweishu").text(resp.liushen.paiweishu);
            $(".gongweishu").text(resp.liushen.gongweishu);
            $(".jiyuanfangwei").text(resp.liushen.fangwei);
            $(".wuxin").text(resp.liushen.wuxin);
            $(".shensha").text(resp.liushen.shensha);
            $(".bagua").text(resp.liushen.bagua);

            $(".jiehuo_list").empty();
            resp.jiehuo_list.forEach(function(jiehuo){
                var tpl = '<tr class="row"><td class="active text-center col-sm-3" style="vertical-align: middle;"><strong>{type}</strong></td><td class="col-sm-9">{sentence}</td></tr>';
                var row = tpl.replace("{type}","问"+jiehuo.type).replace("{sentence}",jiehuo.sentence);
                $(".jiehuo_list").append(row);
            });

            $(".duanci_list").empty();
            for (var i = 0; i < resp.duanci_list.length; i+=4) {
                var row = [];
                resp.duanci_list.slice(i,i+4).forEach(function(one){
                    row.push('<td class="col-sm-3">'+one.sentence+'</td>');
                })
                $(".duanci_list").append('<tr class="row">'+row.join('')+'</tr>')
            }

            zhangjue.loaded();
            $("#dianbo .jieke").removeClass("hidden");
        });
    });
});
</script>
{{end}}