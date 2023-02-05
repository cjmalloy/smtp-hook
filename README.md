smtp-hook
========================
SMTP server to webhook.

Modified docker wrapper for [@alash3al/smtp2http](https://github.com/alash3al/smtp2http) that
adds support for environment variables.

Dev 
===
- `go mod vendor`
- `go build`

Dev with Docker
==============
Locally :
- `docker build -t smtp2http .`
- `docker run -p 25:30025 smtp2http --timeout.read=50 --timeout.write=50 --webhook=http://some.hook/api`

The `timeout` options are of course optional but make it easier to test in local with `telnet localhost 30025`
Here is a telnet example payload : 
```
HELO zeus
MAIL FROM:<email@from.com>
RCPT TO:<youremail@example.com>
DATA
Subject: Test message

your mail content.
.

```

Docker (production)
=====
- `docker run -e WEBHOOK="http://some.hook/api" -p 25:30025 ghcr.io/cjmalloy/smtp-hook:v1.0.0`

Native usage
=====
`smtp2http --listen=:25 --webhook=http://localhost:8080/api/smtp-hook`
`smtp2http --help`

Contribution
============
Forked from repo @alash3al/smtp2http
