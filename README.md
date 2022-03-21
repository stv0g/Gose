<p align="center" >
    <img style="width: 30%; margin: 4em 0" src="frontend/img/gose-logo.svg" alt="GoSƐ logo" />
    <h1 align="center">GoSƐ - A terascale file-uploader</h1>
</p>

<!-- [![Codacy coverage](https://img.shields.io/codacy/coverage/27eec133fcfd4459885d78f52d03daa8?style=flat-square)](https://app.codacy.com/gh/stv0g/gose/) -->
<!-- [![GitHub Workflow Status](https://img.shields.io/github/workflow/status/stv0g/wice/build?style=flat-square)](https://github.com/stv0g/wice/actions) -->
[![goreportcard](https://goreportcard.com/badge/github.com/stv0g/gose?style=flat-square)](https://goreportcard.com/report/github.com/stv0g/gose/)
[![Codacy grade](https://img.shields.io/codacy/grade/27eec133fcfd4459885d78f52d03daa8?style=flat-square)](https://app.codacy.com/gh/stv0g/gose/)
[![libraries.io](https://img.shields.io/librariesio/github/stv0g/gose?style=flat-square)](https://libraries.io/github/stv0g/gose)
[![License](https://img.shields.io/github/license/stv0g/gose?style=flat-square)](https://github.com/stv0g/gose/blob/master/LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/stv0g/gose?style=flat-square)
[![Go Reference](https://pkg.go.dev/badge/github.com/stv0g/gose.svg)](https://pkg.go.dev/github.com/stv0g/gose)

## [Demo](https://gose.0l.de)

## Features

-   Scalable to multiple replicas
    -   No other backend services apart from S3 storage are required
-   Upload progress-bar and transfer statistics
-   Direct upload to Amazon S3 via presigned-URLs
-   Direct download from Amazon S3
-   Drag & Drop
-   Multi-part / chunked upload
-   File integrity checks after finished upload via using MD5 checksum & ETags
-   Optional link shortening via an external service
-   Optional notification about new uploads via [shoutrrr](https://containrrr.dev/shoutrrr/v0.5/)

## Roadmap

Checkout the [Github issue tracker](https://github.com/stv0g/gose/issues?q=is%3Aissue+is%3Aopen+label%3Aenhancement).

## Installation

### Pre-compiled binaries from GitHub releases

Take the download link for your OS/Arch from the [Releases Page](https://github.com/stv0g/gose/releases/) and run:

```bash
export RELEASE_URL=https://github.com/stv0g/gose/releases/download/v0.0.2/gose_0.0.2_linux_amd64
wget "${RELEASE_URL}" -O gose
chmod +x gose
mv gose /usr/local/bin
```

### Kubernetes / Kustomize

1. Copy default configuration file: `cp config.yaml kustomize/config.yaml`
1. Adjust config: `nano kustomize/config.yaml`
1. Apply configuration: `kubectl apply -k kustomize`

### Docker

Via environment variables in `.env` file:

```bash
docker run --env-file=.env --publish=8080:8080 ghcr.io/stv0g/gose
```

or via a configuration file:

```bash
docker run -v$(pwd)/config.yaml:/config.yaml --publish=8080:8080 ghcr.io/stv0g/gose -config /config.yaml
```

## Configuration

Gose can be configured via a configuration file and/or environment variables

### File

For reference have a look at the [example configuration file](config.yaml).

### Environment variables

All settings from the configuration file can also be set via environment variables:

| Variable                              | Example Value                                                                 | Description                           |
| :--                                   | :--                                                                           | :--                                   |
| `GOSE_S3_BUCKET`                      | `gose-uploads`                                                                | Name of S3 bucket                     |
| `GOSE_S3_ENDPOINT`                    | `s3.0l.de`                                                                    | Hostname of S3 server                 |
| `GOSE_S3_REGION`                      | `s3`                                                                          | Region of S3 server                   |
| `GOSE_S3_PATH_STYLE`                  | `true`                                                                        | Prepend bucket name to path           |
| `GOSE_S3_NO_SSL`                      | `false`                                                                       | Disable SSL encryption for S3         |
| `GOSE_S3_ACCESS_KEY`                  | `""`                                                                          | S3 Access Key                         |
| `GOSE_S3_SECRET_KEY`                  | `""`                                                                          | S3 Secret Key                         |
| `AWS_ACCESS_KEY_ID`                   | ``                                                                            | alias for `GOSE_S3_ACCESS_KEY`        |
| `AWS_SECRET_ACCESS_KEY`               | ``                                                                            | alias for `AWS_SECRET_ACCESS_KEY`     |
| `GOSE_S3_MAX_UPLOAD_SIZE`             | `5TB`                                                                         | Maximum upload size                   |
| `GOSE_S3_PART_SIZE`                   | `5MB`                                                                         | Part-size for multi-part uploads      |
| `GOSE_S3_EXPIRATION_DEFAULT_CLASS`    | `1week # one of the tags below`                                               | Default expiration class              |
| `GOSE_SERVER_LISTEN`                  | `":8080"`                                                                     | Listen address and port of Gose       |
| `GOSE_SHORTENER_ENDPOINT`             | `"https://shlink-api/rest/v2/short-urls/shorten?apiKey=<your-api-token>&format=txt&longUrl={{.UrlEscaped}}"`  | API Endpoint of link shortener |
| `GOSE_SHORTENER_METHOD`               | `GET`                                                                         | HTTP method for link shortener        |
| `GOSE_SHORTENER_RESPONSE`             | `raw`                                                                         | Response type of link shortener       |
| `GOSE_NOTIFICATION_URLS`              | `pushover://shoutrrr:<api-token>@<user-key>?devices=laptop1&title=Upload`     | Service URLs for [shoutrrr notifications](https://containrrr.dev/shoutrrr/) |
| `GOSE_NOTIFICATION_TEMPLATE`          | `"New Upload: {{.URL}}"`                                                      | Notification message template         |

## Author

GoSƐ has been written by [Steffen Vogel](mailto:post@steffenvogel.de).

## License

GoSƐ is licensed under the [Apache 2.0 license](./LICENSE).
