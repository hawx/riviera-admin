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

      header {
        margin: 2em 0;
      }

      header a {
        float: right;
        margin-left: 1em;
      }

      header h1 {
        font-size: 23px;
      }

      form {
        margin-bottom: 2em;
      }

      table {
        width: 100%;
        border-collapse: collapse;
      }

      tr {
        border-bottom: 5px solid transparent;
      }

      td:last-child {
        text-align: right;
      }
    </style>
  </head>
  <body>
    <header>
      <a href="/sign-out">Sign-out</a>
      <a href="javascript:location.href='http://localhost:8081/subscribe?url='+encodeURIComponent(location.href)+'&redirect=origin;'">Bookmarklet</a>
      <h1>Riviera Admin</h1>
    </header>

    <table>
      {{range .}}
        <tr>
          <td>{{.}}</td>
          <td><a href="/unsubscribe?url={{.}}">unsubscribe</a></td>
        </tr>
      {{end}}
    </table>
  </body>
</html>
`
