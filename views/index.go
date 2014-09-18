package views

import (
	"html/template"
)

var Index, _ = template.New("index").Parse(index)

const index = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>Riviera Admin</title>
    <style>
      body {
          font: 16px/1.3em Helvetica;
          max-width: 500px;
          margin: 0 auto;
      }

      ul {
          list-style: none;
          padding: 0;
      }

      ul li a {
          float: right;
      }

      #bookmarklet {
        padding: 3px;
        display: inline-block;
        border: 1px solid black;
      }
    </style>
  </head>
  <body>
    <h1>Riviera Admin</h1>
    <a id="bookmarklet" href="javascript:location.href='http://localhost:8081/subscribe?url='+encodeURIComponent(location.href)+'&redirect=origin;'">bookmarklet</a>

    <form action="/subscribe" method="GET">
      <input type="text" id="url" name="url" />
      <input type="submit" value="Subscribe" />
    </form>

    <ul>
      {{range .}}
        <li>
          <strong>{{.}}</strong>
          <a href="/unsubscribe?url={{.}}">unsubscribe</a>
        </li>
      {{end}}
    </ul>
  </body>
</html>
`
