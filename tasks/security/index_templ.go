// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package security

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SecurityHintPage() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col justify-center items-center px-4 sm:px-6 lg:px-8\"><div class=\"max-w-md w-full space-y-8\"><div class=\"bg-white p-8 rounded-lg shadow-lg\"><h1 class=\"text-3xl font-extrabold text-gray-900 mb-6\">Cybersecurity Hub</h1><p class=\"text-md text-gray-700 mb-4\">Embark on a journey into the depths of web security, exploring interactive challenges and enriching learning resources crafted for aspiring cybersecurity enthusiasts.</p><div class=\"mt-6\"><h2 class=\"text-left text-lg font-semibold text-gray-900\">Begin Your Adventure</h2><p class=\"text-left text-md text-gray-600 mt-2\">Unlock the secrets of secure web practices and their application in the digital realm.</p></div><div class=\"mt-4\"><h2 class=\"text-left text-lg font-semibold text-gray-900\">Uncover Hidden Knowledge</h2><p class=\"text-left text-md text-gray-600 mt-2\">Delve into the ways websites safeguard their domains and signal trust to visitors. A well-known file holds the key.</p></div><div class=\"mt-8\"><a href=\"#\" class=\"w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500\">Explore Challenges\r</a></div></div><div class=\"mt-8 bg-indigo-100 p-4 rounded-lg\"><p class=\"text-md text-indigo-800\">Insight: The path to mastery is often hidden in plain sight. Seek out the well-known to discover the secrets.</p></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}