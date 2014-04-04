What the hell?
--------------
Furious Mustard is intended to be a basic media index/display web app thing. Think Plex Media Server without the massive memory footprint, janky web interface and a bunch of streaming stuff I don't care about. This is my way of learning about Go/Redis and is designed for my rather specific use case, I'm not sure how much further effort I'll put in once it works.

Design stuff
------------
- Go!
- Use the http module to serve pages - [Martini][http://martini.codegangsta.io/]?
- Thin server, thick-ish client?
- Storage: Redis
- Cache TheMovieDB API responses in Redis
- Client: Angular/Backbone/Ember/other flavour of the month?
- Goroutines:
  + HTTP server
  + A thing to discover files and queue them for updating
  + A pool of workers to handle API access, caching, etc
    = TheMovieDB API has rate limits, honour them

Planned features
----------------
- Index your Movies/TV episodes using TheMovieDB and (something for TV).
- Gather media information using mediainfo (possibly libmediainfo if I can work out how Cgo works).
- Provide a pretty web interface to browse through your collection.
