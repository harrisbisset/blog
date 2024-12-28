module github.com/harrisbisset/blog/backend/pre-render

go 1.23.4

replace github.com/harrisbisset/blog/internal/posts => ../../internal/posts/

require (
	github.com/harrisbisset/blog/internal/posts v0.0.0-00010101000000-000000000000
	github.com/yuin/goldmark v1.7.8
	github.com/yuin/goldmark-meta v1.1.0
)

require gopkg.in/yaml.v2 v2.3.0 // indirect
