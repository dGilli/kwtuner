// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package home

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/dgilli/kwtuner/views/layouts"
)

func Index() templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<form class=\"w-fit\" action=\"/generate\" method=\"POST\"><div class=\"space-y-12\"><div class=\"border-b border-white/10 pb-12\"><h2 class=\"text-base/7 font-semibold text-white\">Keyword Tuner</h2><p class=\"mt-1 text-sm/6 text-gray-400\">Optimize your text with AI-driven keyword integration.</p><div class=\"mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6\"><div class=\"sm:col-span-4\"><label for=\"platform\" class=\"block text-sm/6 font-medium text-white\">API Key</label><div class=\"mt-2\"><div class=\"flex items-center rounded-md bg-white/5 pl-3 outline outline-1 -outline-offset-1 outline-white/10 focus-within:outline focus-within:outline-2 focus-within:-outline-offset-2 focus-within:outline-indigo-500\"><select id=\"platform\" name=\"platform\" class=\"pl-0 text-base text-gray-500 bg-transparent border-0 sm:text-sm/6\"><option>claude</option></select> <input type=\"text\" name=\"apikey\" id=\"apikey\" class=\"block min-w-0 grow bg-transparent py-1.5 pl-1 pr-3 text-base text-white rounded-r-md placeholder:text-gray-500 focus:outline focus:outline-0 sm:text-sm/6\" placeholder=\"skd-######\"></div></div></div><div class=\"sm:col-span-6\" x-data=\"{ keywords: [], input: &#39;&#39;, colors: [\n                            &#39;bg-red-400/10 text-red-400 ring-red-400/20&#39;,\n                            &#39;bg-yellow-400/10 text-yellow-500 ring-yellow-400/20&#39;,\n                            &#39;bg-green-500/10 text-green-400 ring-green-500/20&#39;,\n                            &#39;bg-blue-400/10 text-blue-400 ring-blue-400/30&#39;,\n                            &#39;bg-indigo-400/10 text-indigo-400 ring-indigo-400/30&#39;,\n                            &#39;bg-purple-400/10 text-purple-400 ring-purple-400/30&#39;,\n                            &#39;bg-pink-400/10 text-pink-400 ring-pink-400/20&#39;\n                            ] }\" x-init=\"colors.sort((a, b) =&gt; 0.5 - Math.random())\"><label for=\"keywords\" class=\"block text-sm/6 font-medium text-white\">Keywords</label><div class=\"mt-2\" x-data=\"{ handleAddKeyword() { keywords.push(input); input = &#39;&#39;; $refs.input.focus() } }\"><input id=\"keywords\" name=\"keywords\" type=\"text\" x-model=\"input\" x-ref=\"input\" class=\"rounded-md bg-white/5 mr-3 px-3 py-1.5 text-base text-white outline outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-500 focus:outline focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6\"> <button type=\"button\" class=\"text-sm/6 font-semibold text-white\" @click=\"handleAddKeyword\">Add</button></div><div class=\"flex flex-wrap gap-2 max-w-md\" :class=\"keywords.length ? &#39;mt-4&#39; : &#39;&#39;\"><template x-for=\"(keyword, index) in keywords\" :key=\"index\"><span :class=\"colors[index % colors.length]\" class=\" inline-flex items-center gap-x-0.5 first:mt-2 rounded-md bg-badge-1 bg-gray-400/10 px-2 py-1 text-xs font-medium text-gray-400 ring-1 ring-inset ring-gray-400/20\"><span x-text=\"keyword\"></span> <button type=\"button\" @click=\"keywords.splice(index, 1)\" class=\"group relative -mr-1 size-3.5 rounded-sm hover:bg-gray-500/20\"><span class=\"sr-only\">Remove</span> <svg viewBox=\"0 0 14 14\" class=\"size-3.5 stroke-gray-600/50 group-hover:stroke-gray-600/75\"><path d=\"M4 4l6 6m0-6l-6 6\"></path></svg> <span class=\"absolute -inset-1\"></span></button></span></template></div></div><div class=\"col-span-full\"><label for=\"text\" class=\"block text-sm/6 font-medium text-white\">Your text</label><div class=\"mt-2\"><textarea name=\"text\" id=\"text\" rows=\"3\" class=\"block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white outline outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-500 focus:outline focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6\"></textarea></div></div></div></div><div class=\"border-b border-white/10 pb-12\"><div class=\"mt-10 grid grid-cols-1 gap-x-6 gap-y-8 sm:grid-cols-6\"><div class=\"sm:col-span-4\"><label for=\"email\" class=\"block text-sm/6 font-medium text-white\">Email address</label><div class=\"mt-2\"><input id=\"email\" name=\"email\" type=\"email\" autocomplete=\"email\" class=\"block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white outline outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-500 focus:outline focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6\"></div></div><div class=\"flex gap-3 sm:col-span-5\"><div class=\"flex h-6 shrink-0 items-center\"><div class=\"group grid size-4 grid-cols-1\"><input id=\"candidates\" aria-describedby=\"candidates-description\" name=\"candidates\" type=\"checkbox\" class=\"col-start-1 row-start-1 appearance-none rounded border border-white/10 bg-white/5 checked:border-indigo-600 checked:bg-indigo-600 indeterminate:border-indigo-600 indeterminate:bg-indigo-600 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 disabled:border-white/10 disabled:bg-transparent forced-colors:appearance-auto\"> <svg class=\"pointer-events-none col-start-1 row-start-1 size-3.5 self-center justify-self-center stroke-white group-has-[:disabled]:stroke-white/25\" viewBox=\"0 0 14 14\" fill=\"none\"><path class=\"opacity-0 group-has-[:checked]:opacity-100\" d=\"M3 8L6 11L11 3.5\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path> <path class=\"opacity-0 group-has-[:indeterminate]:opacity-100\" d=\"M3 7H11\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></svg></div></div><div class=\"max-w-sm text-sm/6\" hx-boost=\"true\" hx-history=\"false\"><label for=\"candidates\" class=\"text-white\">By proceeding, you confirm that you own or have permission to use this content and agree to our <a href=\"/terms.txt\" class=\"underline\">Terms of Service</a>.</label></div></div></div></div></div><div class=\"mt-6 flex items-center justify-end gap-x-6\"><a href=\"/\" class=\"text-sm/6 font-semibold text-white\">Reset</a> <button type=\"submit\" class=\"rounded-md bg-indigo-500 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-400 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500\">Run</button></div></form>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = layouts.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func Result(generated string) templ.Component {
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
		templ_7745c5c3_Var4 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "Results: ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(generated)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/home/index.templ`, Line: 104, Col: 19}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = layouts.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var4), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
