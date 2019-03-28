package main

const IndexTemplate = `<!doctype html>
<html>
<head>
    <title>Wiki</title>
    <meta charset="utf-8" />

      <link rel="stylesheet" type="text/css" href="/css/gollum.css" media="all">
	  <link rel="stylesheet" type="text/css" href="/css/editor.css" media="all">
	  <link rel="stylesheet" type="text/css" href="/css/dialog.css" media="all">
	  <link rel="stylesheet" type="text/css" href="/css/template.css" media="all">
	  <link rel="stylesheet" type="text/css" href="/css/print.css" media="print">

    {{ if .CustomCSS }}
    <link rel="stylesheet" href="/css/custom.css" type="text/css">
    {{ end }}

</head>
<body>
    


<div id="wiki-wrapper" class="page">

<div id="wiki-content">
<div class="">
  <div id="wiki-body-index" class="gollum-markdown-content">
    <div class="markdown-body">
      
		{{ .Body }}
	
    </div>
  </div>
  </div>

</div>


</div>


</body>
</html>`
