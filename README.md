# Goldblum

Your scientists were so preoccupied with whether or not they could, they didn’t stop to think if they should.

## Warning

Don't do this in production (or probably anywhere?). This is just a proof of concept. Seriously. It's not secure, it's not safe.

## What is it?

A golang webserver that allows live changes to its endpoints by hot reloading!

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

import {
    "fmt"
    "net/http"
}

const page = "<html><body><h1>Hello, World!</h1></body></html>"

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, page)
}
```

Now navigate to `http://localhost:8001/helloworld` to test out your new endpoint.

You can go back to the editor and modify this function, all without stopping the webserver!

## License

Copyright 2017 Ahead by a Century, LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.