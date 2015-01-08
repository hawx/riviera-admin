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
      html, body {
          margin: 0;
          padding: 0;
      }

      body {
          font: 14px/1.3 Verdana, Geneva, sans-serif;
          color: #000;
          background: #fff;
      }

      a, a:visited {
          text-decoration: none;
          color: #365da9;
      }
      a:hover, a:focus, a:active {
          text-decoration: underline;
          color: #2a6497;
      }

      .container {
          max-width: 40em;
          margin: 0 auto;
          padding: 0 1em;
      }
      .container:before, .container:after {
          clear: both;
          content: " ";
          display: table;
      }

      header {
          margin: 0;
          background: #eee;
          font-size: 1em;
          border-bottom: 1px solid #ddd;
      }

      header > .container {
          display: flex;
          flex-direction: row;
          justify-content: space-between;
      }

      header h1, header > .container > a {
          margin: 0;
          padding: 1.3rem;
          height: 1.3rem;
          line-height: 1.3rem;
          display: inline-block;
      }

      header h1 {
          font-size: 1.5em;
          padding-left: 0;
          margin-left: .5rem;
          font-weight: bold;
          align-self: flex-start;
      }

      header > .container > a {
          font-size: 1.1em;
          text-decoration: none;
          margin-left: auto;
          color: #333;
      }

      header > .container > a + a {
          margin-left: 0;
      }

      ul {
          width: auto;
          list-style: none;
          padding: 0;
          margin: 2.6rem 0;
      }

      li {
         border-bottom: 1px solid #ddd;
         padding: 0 .5rem;
         width: auto;
         margin: 1.3rem 0;
         width: 100%;
         position: relative;
      }

      ul li:last-child {
          border-bottom: none;
      }

      li h2 {
          font-size: 1.2em;
          margin-bottom: .5rem;
      }

      li h3 {
          overflow: hidden;
          white-space: nowrap;
          text-overflow: ellipsis;
          margin-top: 0;
          font-size: 1em;
          height: 1.3rem;
          font-weight: normal;
      }

      .actions {
          display: flex;
          justify-content: flex-end;
          margin: 1em 0;
      }

      .action, .action:hover {
          color: #A97612;
      }

      .filter {
          padding: 1em 0px;
          background: #fefefe;
          border-bottom: 1px solid #eee;
      }

      .filter input[type=search] {
          display: block;
          border: none;
          width: 100%;
          padding: 0 0.5rem;
          font: 1em Verdana, Geneva, sans-serif;
          margin: 0;
          background: none repeat scroll 0% 0% transparent;
      }

    </style>
  </head>
  <body>
    <header>
      <div class="container">
        <h1>riviera-admin</h1>
        <a href="javascript:location.href='{{.Url}}{{.PathPrefix}}/subscribe?url='+encodeURIComponent(location.href)+'&redirect=origin;'">bookmarklet</a>
        <a href="{{.PathPrefix}}/sign-out">sign-out</a>
      </div>
    </header>

    <div class="filter">
      <div class="container">
        <input id="filter" type="search" placeholder="Filter..." tabindex="1" />
      </div>
    </div>

    <div class="container">
      <ul class="feeds">
        {{range .Feeds}}
        <li>
          <h2><a href="{{.FeedUrl}}">{{.FeedTitle}}</a></h2>
          <h3>&rarr; <a href="{{.WebsiteUrl}}">{{.WebsiteUrl}}</a></h3>
          <p>{{.FeedDescription}}</p>
          <div class="actions">
            <a class="action" href="{{$.PathPrefix}}/unsubscribe?url={{.FeedUrl}}">unsubscribe</a>
          </div>
        </li>
        {{end}}
      </ul>
    </div>

    <script>
      var filter = document.getElementById("filter");
      var list = document.querySelector(".feeds");

      filter.addEventListener("keyup", function() {
          var value = filter.value.toUpperCase();
          var items = list.getElementsByTagName("li");

          for (var i = 0; i < items.length; i++) {
              var item = items[i];
              var name = item.querySelector("h2 a");
              var link = item.querySelector("h3 a");

              if (name.innerHTML.toUpperCase().indexOf(value) != -1 ||
                 name.getAttribute("href").toUpperCase().indexOf(value) != -1 ||
                 link.innerHTML.toUpperCase().indexOf(value) != -1) {
                  item.style.display = "list-item";
              } else {
                  item.style.display = "none";
              }
          }
      });
    </script>
  </body>
</html>
`
