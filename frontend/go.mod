module github.com/harrisbisset/blog/frontend

go 1.23.4

replace github.com/harrisbisset/blog/internal/render => ../internal/render/

replace github.com/harrisbisset/blog/internal/posts => ../internal/posts/

require (
	github.com/a-h/templ v0.3.819
	github.com/gofiber/fiber/v2 v2.52.6
	github.com/harrisbisset/blog/internal/posts v0.0.0-00010101000000-000000000000
	github.com/harrisbisset/blog/internal/render v0.0.0-00010101000000-000000000000
	github.com/yuin/goldmark v1.7.8
	github.com/yuin/goldmark-meta v1.1.0
)

require (
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.58.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
