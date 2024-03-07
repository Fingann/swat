// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package validation

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "fmt"

func ValidatePage() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\r\n async function cartUpdate(){\r\n    this.cartValue = 0;\r\n    document.querySelectorAll('input[data-price]').forEach(input => {\r\n        const price = parseFloat(input.getAttribute('data-price'));\r\n        const quantity = parseInt(input.value) || 0;\r\n        this.cartValue += price * quantity;\r\n    });\r\n    document.getElementById(\"checkout-sum\").innerText = cartValue;\r\n }\r\n window.cartUpdate = cartUpdate;\r\n\r\n</script><div id=\"root\"><div class=\"bg-gray-100 mt-10\"><div class=\"max-w-4xl mx-auto px-4\"><!-- Header Section for Total Cash and Cart Value --><div class=\"flex flex-col space-y-4 mb-2\"><div class=\"bg-gray-200 p-4 rounded-lg shadow\"><h2 class=\"text-2xl font-bold text-gray-800\">Available Balance</h2><p class=\"text-lg text-gray-600 mt-2\">Your total cash available:</p><p class=\"text-3xl font-semibold text-blue-600\">$<span id=\"total-cash\" hx-get=\"/validation/api/available\" hx-trigger=\"load, update-total-cash\" hc-swap=\"innerHTML\">0</span></p></div><div class=\"flex items-center justify-between bg-gray-200 p-4 rounded-lg shadow\"><div><h2 class=\"text-2xl font-bold text-gray-800\">Cart Total</h2><p class=\"text-lg text-gray-600\">Current value of your shopping cart:</p><p class=\"text-3xl font-semibold text-green-600\">$<span id=\"checkout-sum\">0</span></p></div><button id=\"buy-button\" class=\"bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded transition duration-150 ease-in-out\">Buy Now\r</button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = BuyButtonClick().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><!-- Products Display Section --><div id=\"items\" class=\"grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-6\" hx-get=\"/validation/products\" hx-trigger=\"load\" hc-swap=\"innerHTML\"><!-- Items will be displayed here --></div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

type ProductData struct {
	Id          string
	Name        string
	ImageBase64 string
	Price       string
	Description string
}

func Products(products []ProductData) templ.Component {
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
		for _, product := range products {
			templ_7745c5c3_Err = Product(product).Render(ctx, templ_7745c5c3_Buffer)
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

func Product(product ProductData) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"bg-white rounded-lg shadow-md p-6 max-w-md mx-auto mb-6\"><h3 class=\"text-lg font-semibold text-gray-900\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(product.Name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `tasks\validation\index.templ`, Line: 71, Col: 69}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3><img src=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString("data:image/png;base64, " + product.ImageBase64))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" alt=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(product.Name))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"rounded-lg mt-4 mb-4\"><p class=\"text-lg font-bold text-gray-900 mb-4\">Price: $")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(product.Price)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `tasks\validation\index.templ`, Line: 73, Col: 78}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><p class=\"text-gray-700 mb-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(product.Description)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `tasks\validation\index.templ`, Line: 74, Col: 58}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><!-- Quantity Incrementer --><div class=\"max-w-xs mx-auto\"><label for=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(product.Id))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block mb-2 text-sm font-medium text-gray-900 dark:text-white\">Choose quantity:</label><div class=\"relative flex items-center max-w-[8rem]\"><button id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("decrement-%s", product.Id)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"button\" class=\"decrement-button bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 border border-gray-300 rounded-l-lg p-3 h-11 focus:ring-gray-100 focus:ring-2 focus:outline-none\"><!-- SVG for minus --><svg class=\"w-3 h-3 text-gray-900 dark:text-white\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 18 2\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M1 1h16\"></path></svg></button> <input id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(product.Id))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"text\" data-price=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(product.Price))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" data-product-id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(product.Id))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"bg-gray-50 border-x-0 border-gray-300 h-11 text-center text-gray-900 block w-full dark:bg-gray-700\" placeholder=\"1\" value=\"0\" required> <button id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(fmt.Sprintf("increment-%s", product.Id)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" type=\"button\" class=\"increment-button bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 border border-gray-300 rounded-r-lg p-3 h-11 focus:ring-gray-100 focus:ring-2 focus:outline-none\"><!-- SVG for plus --><svg class=\"w-3 h-3 text-gray-900 dark:text-white\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 18 18\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M9 1v16M1 9h16\"></path></svg></button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = TriggerCartUpdate(product.Id).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = AddItemCount(product.Id, fmt.Sprintf("increment-%s", product.Id)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = SubtractItemCount(product.Id, fmt.Sprintf("decrement-%s", product.Id)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func TriggerCartUpdate(inputId string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_TriggerCartUpdate_22d9`,
		Function: `function __templ_TriggerCartUpdate_22d9(inputId){document.getElementById(inputId).addEventListener('change', function(event) {
        event.target.value = Math.max(0, event.target.value);
        cartUpdate()
    });
}`,
		Call:       templ.SafeScript(`__templ_TriggerCartUpdate_22d9`, inputId),
		CallInline: templ.SafeScriptInline(`__templ_TriggerCartUpdate_22d9`, inputId),
	}
}

func AddItemCount(inputId string, buttonId string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_AddItemCount_e27c`,
		Function: `function __templ_AddItemCount_e27c(inputId, buttonId){document.getElementById(buttonId).addEventListener('click', function() {
        let input =document.getElementById(inputId)
        console.log(input.value)
        input.value = parseInt(input.value) + 1
        cartUpdate()
    });
}`,
		Call:       templ.SafeScript(`__templ_AddItemCount_e27c`, inputId, buttonId),
		CallInline: templ.SafeScriptInline(`__templ_AddItemCount_e27c`, inputId, buttonId),
	}
}

func SubtractItemCount(inputId string, buttonId string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_SubtractItemCount_9aab`,
		Function: `function __templ_SubtractItemCount_9aab(inputId, buttonId){document.getElementById(buttonId).addEventListener('click', function() {
        let input =document.getElementById(inputId)
        console.log(input.value)
        if (parseInt(input.value) > 0) {
            input.value = parseInt(input.value) - 1
        }
        cartUpdate()
    });
}`,
		Call:       templ.SafeScript(`__templ_SubtractItemCount_9aab`, inputId, buttonId),
		CallInline: templ.SafeScriptInline(`__templ_SubtractItemCount_9aab`, inputId, buttonId),
	}
}

func BuyButtonClick() templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_BuyButtonClick_56ef`,
		Function: `function __templ_BuyButtonClick_56ef(){document.getElementById("buy-button").addEventListener('click', function() {
            let productQuantityList =[]
            document.querySelectorAll('input[data-price]').forEach(input => {
                const price = parseFloat(input.getAttribute('data-price'));
                const quantity = parseInt(input.value) || 0;
                productQuantityList.push({id: input.getAttribute('data-product-id'), quantity: quantity});
            });
            let totalCash = document.getElementById("total-cash").innerText;
            let cartSum = document.getElementById("checkout-sum").innerText
            // check if the total cash available is less than the shopping cart value
            if (parseInt(totalCash) <= parseInt(cartSum)) {
                // if it is not, then display an error message
                Swal.fire({
                    icon: 'error',
                    title: 'Oops...',
                    text: 'You do not have enough cash to checkout!',
                })
                return 
            }
                // if it is, then proceed to checkout
                Swal.fire({
                    title: 'Are you sure?',
                    text: "You are about to checkout! The items will cost you $" + cartSum + ". Are you sure you want to proceed?",
                    icon: 'warning',
                    showCancelButton: true,
                    confirmButtonColor: '#3085d6',
                    cancelButtonColor: '#d33',
                    confirmButtonText: 'Yes, checkout!'
                }).then((result) => {
                    if (!result.isConfirmed) {
                        return 
                    }
                    fetch("/validation/api/purchase", {
                            method: "POST",
                            headers: {
                                "Content-Type": "application/json",
                            },
                            body: JSON.stringify(productQuantityList),
                    }).then((response) => {
                            console.log(response)
                            if (!response.ok) {
                                throw new Error("Network response was not ok.");
                            }
                            return response.json();
                    }).then(response =>{
                            console.log("response", response)

                            Swal.fire({
                            title: 'Success!',
                            text: response.message || 'Your purchase was successful!',
                            icon: 'success',
                            confirmButtonColor: '#3085d6',
                            }).then((result) => {
                            
                                // update the total cash available
                                htmx.trigger("#total-cash", "update-total-cash", {});
                                // and reset the shopping cart value
                                document.getElementById("checkout-sum").innerText = 0;
                                // Reset the quantity of the items
                                document.getElementById("id1").value = 0;
                                document.getElementById("id2").value = 0;
                            })
                    }).catch((error) => {
                            Swal.fire({
                                icon: 'error',
                                title: 'Oops...',
                                text: 'There was an error processing your request!',
                            })
                            console.error("There was an error!", error);
                             
                    });
                        
                })
        });
}`,
		Call:       templ.SafeScript(`__templ_BuyButtonClick_56ef`),
		CallInline: templ.SafeScriptInline(`__templ_BuyButtonClick_56ef`),
	}
}
