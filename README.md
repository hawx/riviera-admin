# riviera-admin

An admin panel for [riviera][].

- Lists feeds subscribed to
- Provides a bookmarklet to subscribe to the current page's feed
- Can unsubscribe from existing feeds

``` bash
$ go get hawx.me/code/riviera-admin
$ riviera-admin subscriptions.xml
...
```

Then visit <http://localhost:8081> and login.

See `riviera-admin --help` for the other options that will be required when
running properly.

[riviera]: https://github.com/hawx/riviera
