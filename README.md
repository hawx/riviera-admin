# riviera-admin

An admin panel for [riviera][].

- Lists feeds subscribed to
- Can subscribe to new feeds
- Can unsubscribe from existing feeds
- Provides a bookmarklet to subscribe to the current page's feed

``` bash
$ go get github.com/hawx/riviera-admin
$ riviera-admin --user john@doe.com
...
```

Then visit <http://localhost:8081> and login.

See `riviera-admin --help` for the other options that will be required when
running properly.

[riviera]: https://github.com/hawx/riviera
