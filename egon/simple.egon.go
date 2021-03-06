// Generated by egon.
// 🚫Edit at your own risk.

package egon
import (
"github.com/commondream/egon"
"fmt"
"html"
"io"
"github.com/SlinSo/goTemplateBenchmark/model"
)

func SimpleView(//line simple.egon:2
u *model.User) *egon.View {
	packageName := "egon"
	name := "Simple"
	templatePath := "simple.egon"
	renderFunc := func(w io.Writer) error {
		return SimpleTemplate(w, u)
	}
	return &egon.View{PackageName: packageName, Name: name, TemplatePath: templatePath, RenderFunc: renderFunc}
}

func SimpleTemplate(w io.Writer, //line simple.egon:2
u *model.User) error {//line simple.egon:2
_, _ = fmt.Fprint(w, "\n")
//line simple.egon:3
_, _ = fmt.Fprint(w, "\n<html>\n    <body>\n        <h1>")
//line simple.egon:5
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  u.FirstName )))
//line simple.egon:5
_, _ = fmt.Fprint(w, "</h1>\n\n        <p>Here's a list of your favorite colors:</p>\n        <ul>\n        ")
//line simple.egon:9
 for _, colorName := range u.FavoriteColors { 
//line simple.egon:10
_, _ = fmt.Fprint(w, "\n            <li>")
//line simple.egon:10
_, _ = fmt.Fprint(w, html.EscapeString(fmt.Sprintf("%v",  colorName )))
//line simple.egon:10
_, _ = fmt.Fprint(w, "</li>")
//line simple.egon:10
 } 
//line simple.egon:11
_, _ = fmt.Fprint(w, "\n        </ul>\n    </body>\n</html>")
return nil
}
