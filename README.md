# Goldblum

Your scientists were so preoccupied with whether or not they could, they didn’t stop to think if they should.

## Warning

Don't do this in production (or probably anywhere?). This is just a proof of concept. Seriously. It's not secure, it's not safe.

## What is it?

A golang webserver that allows adding new endpoints and modifying the source code for old endpoints live!

![Example screenshot](https://raw.githubusercontent.com/transitorykris/goldblum/master/images/example.png)

## Why?

To see if I could.

## How?

```bash
docker-compose up
```

This will give you an instance of Goldblum and a MySQL server. It may take a few minutes the first time to come up.

Navigate in your browser to `http://localhost:8001/` to create your first endpoint.

Here's an easy one:

Set `Method` to `GET` and the endpoint to `/helloworld` and the code to:

```golang
package main

import (
    "fmt"
    "net/http"

    gb "github.com/transitorykris/goldblum"
)

const page = "<html><body><h1>Hello, World!</h1></body></html>"

func Handler(g *gb.Goldblum, w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, page)
}
```

Now navigate to `http://localhost:8001/helloworld` to test out your new endpoint.

Go ahead and head back to the editor and make a change to the endpoints source!

## Tell me more

### *gb.Goldblum

Your handler gets a sweet struct containing a connection to the database as well as a structured logger.

```golang
type Goldblum struct {
	DB  *sqlx.DB
	Log *logrus.Logger
}
```

## License

Copyright 2017 Kris Foster

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

![Jeff Goldblum](https://raw.githubusercontent.com/transitorykris/goldblum/master/images/goldblum.jpg)
