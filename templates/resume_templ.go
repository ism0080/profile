// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/ism0080/profile/models"
)

func Resume(data models.ResumePage) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"prose prose-headings:text-black container mx-auto px-4 text-black\"><h1>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(data.Resume.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 10, Col: 26}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><p>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(data.Resume.Blurb)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 11, Col: 25}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div><h2>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(data.Resume.Experience.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 13, Col: 38}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><div class=\"ml-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, job := range data.Resume.Experience.Jobs {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"mb-0 font-bold\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(job.Role)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 16, Col: 42}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><p class=\"my-0 font-extralight\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(job.Company)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 17, Col: 51}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><p class=\"mt-0\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var8 string
				templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(job.Description)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 18, Col: 39}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				for _, example := range job.Examples {
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"mb-0 underline\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var9 string
					templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(example.Role)
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 21, Col: 48}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><ul>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					for _, ex := range example.Experiences {
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><span class=\"font-bold\">")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var10 string
						templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(ex.Project)
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 25, Col: 47}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(": </span> <span>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var11 string
						templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(ex.Description)
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 26, Col: 33}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						if ex.Description != "" && ex.Technology != "" {
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span>- </span> ")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"italic\">")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var12 string
						templ_7745c5c3_Var12, templ_7745c5c3_Err = templ.JoinStringErrs(ex.Technology)
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 30, Col: 47}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var12))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></li>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</ul>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div><h2>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var13 string
			templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(data.Resume.Skills.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 40, Col: 34}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><div class=\"ml-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, skill := range data.Resume.Skills.Skills {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p><span class=\"font-bold\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var14 string
				templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(skill.Key)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 43, Col: 44}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var15 string
				templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(skill.Value)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 43, Col: 67}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div><h2>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var16 string
			templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(data.Resume.Skills.Other.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 48, Col: 40}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h2><div class=\"ml-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, skill := range data.Resume.Skills.Other.Skills {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p><span class=\"font-bold\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var17 string
				templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(skill.Key)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 51, Col: 44}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> ")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var18 string
				templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(skill.Value)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/resume.templ`, Line: 51, Col: 67}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><button>Button Resume</button></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = Base(data.Base).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
