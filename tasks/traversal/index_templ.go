// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package traversal

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "fmt"

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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"root\" class=\"bg-gray-100\"><div class=\"bg-indigo-600\"><div class=\"max-w-7xl mx-auto py-4 px-4 sm:px-6 lg:px-8\"><h1 class=\"text-2xl font-bold text-white\">File Upload System</h1></div></div><div class=\"max-w-7xl mx-auto py-8 sm:px-6 lg:px-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = UploadForm().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = FileList().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func UploadForm() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button id=\"uploadButton\" class=\"inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500\">Select file to upload\r</button><div id=\"upload-status\"></div><script>\r\n    document.getElementById('uploadButton').addEventListener('click', async () => {\r\n        const { value: file } = await Swal.fire({\r\n            title: \"Select image\",\r\n            input: \"file\",\r\n            inputAttributes: {\r\n                \"aria-label\": \"Upload your profile picture\"\r\n            }\r\n        });\r\n\r\n        if (file) {\r\n            const reader = new FileReader();\r\n            reader.onload = (e) => { \r\n                // After displaying the image, proceed to upload\r\n                uploadFile(file);\r\n            };\r\n            reader.readAsDataURL(file);\r\n        }\r\n    });\r\n\r\n    function uploadFile(file) {\r\n        // Create FormData object and append the file\r\n        let formData = new FormData();\r\n        formData.append('file', file);\r\n\r\n        // Use fetch API to send the file to the server\r\n        fetch('/traversal/upload', {\r\n            method: 'POST',\r\n            body: formData,\r\n        })\r\n        .then(response => {\r\n            if (!response.ok) {\r\n                throw new Error('Network response was not ok');\r\n            }\r\n            //dispatch a custom event to update the file list\r\n            return response.text();\r\n        })\r\n        .then(result => {\r\n            // replace fileList with the updated content\r\n            fileList=document.getElementById('fileList')\r\n            fileList.innerHTML = result;\r\n            // Process the updated content\r\n        \r\n            htmx.process(fileList);\r\n            // Handle success. You might want to update the UI to reflect the successful upload.\r\n            Swal.fire('Success', 'File uploaded successfully', 'success');\r\n        })\r\n        .catch(error => {\r\n            // Handle errors, such as displaying a message to the user\r\n            console.error('Error during file upload:', error);\r\n            Swal.fire('Error', 'There was an error uploading the file', 'error');\r\n        });\r\n    }\r\n</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

type UploadedFile struct {
	Name string
	Size int64
}

func BytesToHumanReadable(size int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	var i int
	for i = 0; size >= 1024 && i < len(units)-1; i++ {
		size /= 1024
	}
	return fmt.Sprintf("%d %s", size, units[i])
}

func FileList() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div><div class=\"mt-5\"><h2 class=\"text-lg font-semibold text-lf text-gray-900 mb-4\">Opplastet Filer</h2><div class=\"grid grid-cols-1 gap-4\" hx-get=\"/traversal/list/items\" hx-trigger=\"load,sse:upload\" hx-swap=\"innerHTML\" id=\"fileList\"></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func ListItems(files []UploadedFile) templ.Component {
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
		for _, file := range files {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-white overflow-hidden shadow rounded-lg\"><div class=\"p-4 flex justify-between items-center\"><div class=\"flex-1 truncate\"><div class=\"flex items-center ml-5 space-x-3\"><h3 class=\"text-sm font-medium text-lg text-gray-900 truncate\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(file.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `tasks\traversal\index.templ`, Line: 114, Col: 97}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3></div></div><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 templ.SafeURL = templ.SafeURL(fmt.Sprintf("/traversal/download?file=%s", file.Name))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var6)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"ml-2 bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 text-white font-bold py-2 px-4 rounded\" download>Last Ned\r</a></div></div>")
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
