<!doctype html>
<html lang="ja">
    <head>
        <meta charset="utf-8">
        <link rel="shortcut icon" href="{{.StaticUrl}}images/common/favicon.ico" >
        <link rel="stylesheet" type="text/css" href="{{.StaticUrl}}css/bootstrap.min.css{{.Version}}">
        <link rel="stylesheet" type="text/css" href="{{.StaticUrl}}css/sticky-footer.css{{.Version}}" />
        <link rel="stylesheet" type="text/css" href="{{.StaticUrl}}css/app.min.css{{.Version}}" />
        <link rel="stylesheet" type="text/css" href="{{.StaticUrl}}css/app.min.css{{.Version}}" />
        <link rel="stylesheet" type="text/css" href="{{.StaticUrl}}css/font-awesome.min.css" integrity="sha384-wvfXpqpZZVQGK6TAh5PVlGOfQNHSoD2xbE+QkPxCAFlNEevoEH3Sl0sibVcOQVnN" crossorigin="anonymous">
        <title>.stamp</title>
    </head>
    <body>
        <input type="hidden" name="uploadToken" id="uploadToken" value="uploadToken"/>

        <div id="root"></div>
        <script src="{{.StaticUrl}}js/bundle.js{{.Version}}"></script>
    </body>
</html>
