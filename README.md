# Afterstack using martini
http://martini.codegangsta.io/
http://blog.gopheracademy.com/day-11-martini
http://0value.com/build-a-restful-API-with-Martini
http://www.vinaysahni.com/best-practices-for-a-pragmatic-restful-api


    cf push --buildpack=git://github.com/hmalphettes/heroku-buildpack-go.git

## Differences with the original buildpack

* Some bits and pieces related to python were removed.
* The '.godir' file is named '_godir': cf will omit to upload files that start with '.' by default.

Follow the PR here as we improve things: https://github.com/vito/heroku-buildpack-go/pull/1
