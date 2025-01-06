// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package view_blog

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/harrisbisset/blog/frontend/models"
	"github.com/harrisbisset/blog/frontend/render/views/view_layout"
	posts "github.com/harrisbisset/blog/internal/posts"
)

func BlogPost(p posts.Post) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = BlogStyle().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, " <div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = models.GetContent(p).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div><div class=\"my-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = view_layout.MobileNavDisplay().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = view_layout.Main().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func BlogStyle() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<style>\r\n        h1 {\r\n            color: #d164a9;\r\n            font-size: 2.25rem;\r\n            line-height: 3rem;\r\n            width: 90%;\r\n        }\r\n\r\n        h2 {\r\n            color: #4b4599;\r\n            font-size: 2rem;\r\n            line-height: 2.5rem;\r\n        }\r\n\r\n        h1, h2 {\r\n            margin-top: 10px;\r\n            margin-bottom: 7px;\r\n        }\r\n\r\n        h3 {\r\n            color: #7600d7;\r\n            font-size: 1.5rem;\r\n            line-height: 2rem;\r\n        }\r\n\r\n        h4 {\r\n            color: #7600d7;\r\n            font-size: 1.25rem;\r\n            line-height: 1.7rem;\r\n        }\r\n\r\n        h1, h2, h3, h4 {\r\n            font-weight: 600;\r\n        }\r\n\r\n        p, li {\r\n            font-family: \"Open Sans\";\r\n            margin-top: 2px;\r\n            margin-bottom: 2px;\r\n        }\r\n\r\n        blockquote {\r\n            width: 90%;\r\n            background-color: #eee8f4;\r\n            padding: 0px;\r\n        }\r\n\r\n        blockquote ~ blockquote {\r\n            width: 100%;\r\n            background-color: #eee8f4;\r\n            padding: 0px;\r\n        }\r\n\r\n\r\n        p {\r\n            padding: 8px;\r\n        }\r\n\r\n        p, li {\r\n            padding: 5px;\r\n            color: #0f172a;\r\n            line-height: 1.4;\r\n            text-align: left;\r\n        }\r\n\r\n        ul, ol {\r\n            padding-left: 45px;\r\n        }\r\n\r\n        ul {\r\n            list-style-type: disc;\r\n        }\r\n\r\n        ol {\r\n            list-style-type: decimal;\r\n        }\r\n\r\n        hr {\r\n            border-color: #1e293b;\r\n        }\r\n\r\n        code {\r\n            display: block;\r\n            background-color: #fbf2fe;\r\n            border-radius: 25px;\r\n            padding-top: 5px;\r\n            padding-right: 5px;\r\n            padding-bottom: 5px;\r\n            padding-left: 12px;\r\n            border-left-width: 15px;\r\n            border-left-color: #dedddd;\r\n            overflow-x: scroll;\r\n            width: 100%;\r\n        }\r\n\r\n        em {\r\n            padding-right: 2px;\r\n        }\r\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
