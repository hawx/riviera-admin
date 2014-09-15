package views

import "html/template"

var Index, _ = template.New("index").Parse(index)

const index = `
<!DOCTYPE html>
<html>
  <head>
    <title>Riviera Admin</title>
  </head>
  <body>
    <h1>Riviera Admin</h1>

    <form action="{{$.Riviera}}-/subscribe" method="GET">
      <label for="url">URL:</label>
      <input type="text" id="url" name="url" />

      <input type="submit" value="Subscribe" />
    </form>

    <ul>
      {{range .Urls}}
        <li>
          <strong>{{.}}</strong>
          <a href="{{$.Riviera}}-/unsubscribe?url={{.}}">unsubscribe</a>
        </li>
      {{end}}
    </ul>
  </body>
</html>
`
