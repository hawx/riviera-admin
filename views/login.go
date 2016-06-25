package views

import "html/template"

var Login, _ = template.New("login").Parse(login)

const login = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <title>Riviera Admin</title>
    <style>
      body {
          font: 16px/1.3em Helvetica;
          width: 100%;
          margin: 0;
          padding: 0;
      }

      #cover {
          top: 0;
          left: 0;
          z-index: 1000;
          position: absolute;
          height: 100%;
          width: 100%;
          background: rgba(0, 255, 255, .7);
          display: block
          padding: 0;
          margin: 0;
      }

      #cover a {
          position: relative;
          display: block;
          left: 50%;
          top: 50%;
          text-align: center;
          width: 100px;
          margin-left: -50px;
          height: 50px;
          line-height: 50px;
          margin-top: -25px;
          font-size: 16px;
          font-weight: bold;
          border: 1px solid;
      }
    </style>
  </head>
  <body>
    <div id="cover">
      <a href="{{.Url}}{{.PathPrefix}}/sign-in" title="Sign-in with Persona">Sign-in</a>
    </div>
  </body>
</html>
`
