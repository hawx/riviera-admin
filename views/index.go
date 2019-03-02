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
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Riviera Admin</title>
    <style>
      ` + styles + `
    </style>
  </head>
  <body>
    <header>
      <h1>riviera-admin</h1>
      <a href="javascript:location.href='{{.Url}}/subscribe?url='+encodeURIComponent(location.href)+'&redirect=origin;'">bookmarklet</a>
      <a href="{{.Url}}/sign-out">sign-out</a>
    </header>

    <div class="container">
      <ul class="feeds">
        {{range .Feeds}}
        <li>
          <h1><a href="{{.WebsiteUrl}}">{{.FeedTitle}}</a></h1>
          <p>{{.FeedDescription}}</p>
          <div class="buttons">
            <a href="{{.FeedUrl}}">feed</a>
            <a href="{{$.Url}}/unsubscribe?url={{.FeedUrl}}">unsubscribe</a>
          </div>
        </li>
        {{end}}
      </ul>
    </div>
  </body>
</html>
`
