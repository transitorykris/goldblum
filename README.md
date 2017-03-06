# Goldblum

Your scientists were so preoccupied with whether or not they could, they didn’t stop to think if they should.

## Warning

Don't do this in production (or probably anywhere?). This is just a proof of concept. Seriously. It's not secure, it's not safe.

## What is it?

A golang webserver that allows live changes to its endpoints.

## Why?

To see if I could.

## How?

```bash
docker-compose up
```

Will give you an instance of Goldblum and a MySQL server.

Navigate in your browser to `http://localhost:8001/` to create your first endpoint.

Here's an easy one:

Set `Method` to `GET` and the endpoint to `/healthcheck` and the code to:

```golang
package main

import (
    "log"
    "net/http"

    gb "github.com/transitorykris/goldblum"
)

// Healthcheck is returned when the health of this service is requested
type Healthcheck struct {
	Status string `json:"status"`
}

// HealthcheckHandler returns the health of this service
func HealthcheckHandler() gb.HandlerFunc {
    return gb.HandlerFunc(func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
        log.Println("Healthcheck called")
        gb.Response(w, &Healthcheck{Status: "ok"}, http.StatusOK)
    })
}
```

Now navigate to `http://localhost:8002/healthcheck` to test out your new endpoint.

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