### Quickstart Steps
1) Install [Go 1.8+](https://golang.org/dl/)

2) Install kudzu CLI:
```bash
$ go get github.com/kudzu-cms/kudzu/…
```

3) Create a new project (path is created in your GOPATH):
```bash
$ kudzu new github.com/nilslice/reviews
```

4) Enter your new project directory:
```bash
$ cd $GOPATH/src/github.com/nilslice/reviews
```

5) Generate content type file and boilerplate code (creates `content/review.go`):
```bash
$ kudzu generate content review title:"string" author:"string" rating:"float64" body:"string":richtext website_url:"string" items:"[]string" photo:string:file
```

6) Build your project:
```bash
$ kudzu build
```

7) Run your project with defaults:
```bash
$ kudzu run
```

8) Open browser to [`http://localhost:8080/admin`](http://localhost:8080/admin)

### Notes
- One-time initialization to set configuration
- All fields can be changed in Configuration afterward
