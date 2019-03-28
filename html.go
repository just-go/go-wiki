package main

const Template = `<!doctype html>
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
    

<script>
Mousetrap.bind(['e'], function( e ) {
  e.preventDefault();
  window.location = "/edit" + window.location.pathname;
  return false;
});
</script>
<div id="wiki-wrapper" class="page">
<!--<div id="head">
   <h1>1</h1>

 <ul class="actions">
    <li class="minibutton">
      <div id="searchbar">
        <form action="/search" method="get" id="search-form">
        <div id="searchbar-fauxtext">
          <input type="text" name="q" id="search-query" value="Search&hellip;" autocomplete="off">
          <a href="#" id="search-submit" title="Search this wiki">
            <span>Search</span>
          </a>
        </div>
        </form>
      </div>    </li>
    <li class="minibutton"><a href="/"
       class="action-home-page">Home</a></li>
    <li class="minibutton"><a href="/pages"
      class="action-all-pages">All</a></li>
    <li class="minibutton"><a href="/fileview"
    class="action-fileview">Files</a></li>
      <li class="minibutton jaws">
        <a href="#" id="minibutton-new-page">New</a></li>
        <li class="minibutton jaws">
          <a href="#" id="minibutton-rename-page">Rename</a></li>
        <li class="minibutton"><a href="/edit/1"
           class="action-edit-page">Edit</a></li>
    <li class="minibutton"><a href="/history/1"
       class="action-page-history">History</a></li>
    <li class="minibutton"><a href="/latest_changes"
       class="action-page-history">Latest Changes</a></li>
  </ul>

</div>-->
<div id="wiki-content">
<div class="">
  <div id="wiki-body" class="gollum-markdown-content">
    <div class="markdown-body">
      
		{{ .Body }}
	
    </div>
  </div>
  </div>

</div>

<!--<div id="footer">
  <p id="last-edit"></p>
</div>-->

</div>


</body>
</html>`
