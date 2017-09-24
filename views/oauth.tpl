<html lang="ja">
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>.stamp</title>
    </head>
    <body>
<script>
if (window.location.hash) {
  var query = window.location.hash.replace(/^#/, "")
  var url =
    "exp://exp.host/@wheatandcat/dotstamp/+" +
    "?" +
    query +
    "&login={{.login}}&email={{.email}}"
  window.location = url
  document.write("Authenticating...")
} else {
  document.write("Uh oh something went wrong")
}
</script>
    </body>
</html>
