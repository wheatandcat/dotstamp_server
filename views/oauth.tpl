<!doctype html>
<html lang="ja">
    <head>
        <meta charset="utf-8">
        <title>.stamp</title>
    </head>
    <body>
<script>
if (window.location.hash) {
  var query = window.location.hash.replace(/^#/, "");
  console.log(query);

  document.write(query);
  var url = window.location.origin + "/redirect" + "?" + query;
  window.location = url;
  document.write("Authenticating...");
} else {
  document.write("Uh oh something went wrong");
}
</script>
      <div>test</div>
    </body>
</html>
