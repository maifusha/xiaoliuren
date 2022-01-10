{{define "fragment/date"}}
<input type="text" id="date" name="date">

<link rel="stylesheet" type="text/css" href="/static/lib/datetimepicker/jquery.datetimepicker.min.css"/>
<script type="text/javascript" src="/static/lib/datetimepicker/jquery.datetimepicker.full.min.js"></script>
<script type="text/javascript">
$(function(){
    $.datetimepicker.setLocale('ch');
    $("#date").datetimepicker({
        timepicker: false,
        format: "Y-m-d",
        inline: true,
        defaultSelect: false,
    });
});
</script>
{{end}}