// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package xss

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

// XssPage serves as the main container for the XSS demonstration, including a blog entry and comment sections.
func Page() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"root\" class=\"mt-10 flex flex-col items-center justify-center bg-gray-100 px-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = BlogEntry().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CommentSection().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = CreateCommentSection().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

// BlogEntry represents a single blog post, including a title, content, and an illustrative image.
func BlogEntry() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"blog-entry w-4/5\"><article class=\"max-w-4xl mx-auto bg-white p-6 rounded-lg shadow-lg\"><h1 class=\"text-3xl font-bold text-gray-800 mb-4\">Eventyret om Kodekrigeren og Skriptskurken</h1><img src=\"/static/dist/images/xss-blog-story.webp\" alt=\"XSS Illustrasjon\" class=\"mt-4 w-4/5 rounded-lg justify-center mx-auto mb-4\"><p class=\"text-gray-600 mb-4\">I et land langt, langt borte, i en verden styrt av kode og digitale eventyr, bodde en tapper kodekriger kjent som SkikkreSiv. SkikkreSivs oppdrag var enkelt men av ytterste viktighet: å verne Innbyggerne mot SkriptKåres ondsinnede planer i Secure-Land.</p><p class=\"text-gray-600 mb-4\">SkriptKåre, den beryktede skurken, hadde mestret kunsten å hviske skadelige skript inn i de mest uskyldige av websider. Hans siste påfunn? Et skript så lurt at det kunne få enhver side til å utbasunere \"Muhaha!\" til intetanendes forbauselse.</p><p class=\"text-gray-600 mb-4\">Men SkikkreSiv var ikke en som lot seg skremme så lett. Med et arsenal av Content-Security-Policy skjold og Sanitization-formler ved sin side, gikk hun modig til verks. Hver gang SkriptKåre forsøkte seg på et nytt triks, var SkikkreSiv der for å avverge angrepet, beskytte nettstedene og holde Innbyggernes data trygg.</p><p class=\"text-gray-600 mb-4\">Til slutt, etter mange trefninger hvor SkriptKåre ble overlistet av SkikkreSivs snarrådighet, måtte han trekke seg tilbake, beseiret og uten makt til å utføre sine skumle planer. Secure-Land var igjen et trygt sted, takket være SkikkreSivs urokkelige vilje og beskyttelse.</p><p class=\"text-gray-600\">Så, la oss ta lærdom fra SkikkreSivs eventyr: Den beste beskyttelsen er å være forberedt. Ved å anvende sikkerhetsbestemmelser og holde våre nettsteder trygt forseglet mot inntrengere, kan vi alle bidra til å holde det digitale riket sikkert for alle.</p></article></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

type Comments struct {
	User    string
	Comment string
	Date    string
}

// CommentSection displays comments in a table format, dynamically loaded with HTMX.
func CommentSection() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"comment-section mt-8 w-4/5\" hx-sse=\"connect:/xss/sse\"><h3 class=\"text-2xl font-semibold text-gray-700 mb-4\">Comments</h3><div hx-get=\"/xss/comments\" hx-trigger=\"load, sse:update, every 10s\" hx-target=\"#comments\" hx-swap=\"innerHTML\" class=\"table-container\"><table class=\"min-w-full leading-normal\"><thead><tr><th class=\"px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider\">User\r</th><th class=\"px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider\">Comment\r</th><th class=\"px-5 py-3 border-b-2 border-gray-200 bg-gray-100 text-left text-xs font-semibold text-gray-700 uppercase tracking-wider\">Posted At\r</th></tr></thead> <tbody id=\"comments\"><!-- Dynamic HTMX content will be loaded here --></tbody></table></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

// CommentRows processes a slice of Comments and generates HTML table rows.
func CommentRows(comments []Comments) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		for _, comment := range comments {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr class=\"hover:bg-gray-100\"><td class=\"px-5 py-5 border-b border-gray-200 bg-white text-sm\"><div class=\"flex items-center\"><div class=\"ml-3\"><p class=\"text-gray-900 whitespace-no-wrap\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.Raw(comment.User).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div></div></td><td class=\"px-5 py-5 border-b border-gray-200 bg-white text-sm\"><p class=\"text-gray-900 whitespace-pre-wrap\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(comment.Comment)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `tasks\xss\index.templ`, Line: 76, Col: 77}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></td><td class=\"px-5 py-5 border-b border-gray-200 bg-white text-sm\"><p class=\"text-gray-900 whitespace-no-wrap\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(comment.Date)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `tasks\xss\index.templ`, Line: 79, Col: 73}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></td></tr>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func CreateCommentSection() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"create-comment-section mt-8 mb-8 w-4/5\"><h3 class=\"text-xl font-semibold text-gray-800 mb-4\">Leave a Comment</h3><form hx-post=\"/xss/comments/create\" hx-swap=\"innerHTML\" hx-target=\"#comments\" x-on:htmx:after-on-load.window=\"commentText = &#39;&#39;;\" class=\"w-full\"><div class=\"mb-6\"><label for=\"user-name\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300\">Your Name</label> <input hx-preserve type=\"text\" id=\"user-name\" name=\"user\" class=\"bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5\" placeholder=\"John Doe\" required></div><div class=\"mb-6\"><label for=\"comment-text\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300\">Your Comment</label> <textarea hx-preserve id=\"comment-text\" name=\"comment\" rows=\"4\" class=\"bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5\" placeholder=\"Write something...\" required></textarea></div><button type=\"submit\" class=\"text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center\">Post Comment</button></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
