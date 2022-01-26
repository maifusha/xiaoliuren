{{template "fragment/header" .}}
<div class="container" style="padding: 0;">
    <div class="row" style="margin-top: 30px;">
        <div class="col-sm-3 text-center">
            <h4>无事勿卦，卦后不复，心意相诚</h4>
            <small>道门小六壬</small>
            <div style="position: relative;">
                <img src="/static/image/zhangjue.jpg" class="img-responsive img-rounded center-block" alt="Responsive image"
                     style="margin-top: 50px;">
                <span id="thumb" class="glyphicon glyphicon-hand-up" aria-hidden="true" style="font-size: 25px;color: red;position: absolute;"></span>
            </div>
        </div>

        <div class="col-sm-9 well content">
            <!-- Nav tabs -->
            <ul class="nav nav-tabs nav-justified" role="tablist">
                <li role="presentation" class="active"><a href="#jixiong" aria-controls="jixiong" role="tab" data-toggle="tab"><span class="glyphicon glyphicon-question-sign" aria-hidden="true"></span> 后事吉凶</a></li>
                <li role="presentation"><a href="#dianbo" aria-controls="dianbo" role="tab" data-toggle="tab"><span class="glyphicon glyphicon-list-alt" aria-hidden="true"></span> 当下点拨</a></li>
                <li role="presentation"><a href="#zeshi" aria-controls="zeshi" role="tab" data-toggle="tab"><span class="glyphicon glyphicon-time" aria-hidden="true"></span> 谋事择时</a></li>
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
<script type="text/javascript">
$(function () {
    window.zhangjue = new Object({
        currentLiushen: null,
        position: {
            1: {top:"60%",left:"39%"},
            2: {top:"18%",left:"43%"},
            3: {top:"14%",left:"56%"},
            4: {top:"16%",left:"69%"},
            5: {top:"59%",left:"67%"},
            6: {top:"58%",left:"53%"},
        },

        fingerCount: function (count, resp) {
            var self = this;
            var i = 1;
            var start = self.currentLiushen

            var defer = $.Deferred()
            setInterval(function () {
                if (i > count) {
                    return defer.resolve(resp);
                }

                self.checkLiushen(start+i-1)

                i++;
            }, 100)

            return defer.promise();
        },

        checkLiushen: function (index) {
            var self = this;
            self.currentLiushen = index%6;
            if (self.currentLiushen == 0) {
                self.currentLiushen = 6;
            }

            $("#thumb").css("top", self.position[self.currentLiushen].top).css("left", self.position[self.currentLiushen].left);
        },

        loading: function () {
            $(".content").loading({
                theme: 'dark',
                message: '掐指中。。。',
                onStart: function(loading) {
                    loading.overlay.slideDown(400);
                },
                onStop: function(loading) {
                    loading.overlay.slideUp(400);
                },
            });
        },

        loaded: function () {
            $(".content").loading("stop");
        }
    });

    zhangjue.checkLiushen(1);

    $(document).delegate('a[data-toggle="tab"]', "shown.bs.tab", function () {
        zhangjue.checkLiushen(1);
    });
});
</script>
<style type="text/css">
    .loading-overlay {z-index: 9999 !important;}
</style>
{{template "fragment/footer" .}}