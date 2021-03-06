// This file is automatically generated by qtc from "simple.qtpl".
// See https://github.com/valyala/quicktemplate for details.

//line quicktemplate/simple.qtpl:1
package quicktemplate

//line quicktemplate/simple.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line quicktemplate/simple.qtpl:1
import "github.com/SlinSo/goTemplateBenchmark/model"

//line quicktemplate/simple.qtpl:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line quicktemplate/simple.qtpl:2
func StreamSimpleQtc(qw422016 *qt422016.Writer, u *model.User) {
	//line quicktemplate/simple.qtpl:2
	qw422016.N().S(`
<html>
    <body>
        <h1>`)
	//line quicktemplate/simple.qtpl:5
	qw422016.E().S(u.FirstName)
	//line quicktemplate/simple.qtpl:5
	qw422016.N().S(`</h1>

        <p>Here's a list of your favorite colors:</p>
        <ul>
        `)
	//line quicktemplate/simple.qtpl:9
	for _, colorName := range u.FavoriteColors {
		//line quicktemplate/simple.qtpl:9
		qw422016.N().S(`
            <li>`)
		//line quicktemplate/simple.qtpl:10
		qw422016.E().S(colorName)
		//line quicktemplate/simple.qtpl:10
		qw422016.N().S(`</li>
        `)
		//line quicktemplate/simple.qtpl:11
	}
	//line quicktemplate/simple.qtpl:11
	qw422016.N().S(`
        </ul>
    </body>
</html>
`)
//line quicktemplate/simple.qtpl:15
}

//line quicktemplate/simple.qtpl:15
func WriteSimpleQtc(qq422016 qtio422016.Writer, u *model.User) {
	//line quicktemplate/simple.qtpl:15
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line quicktemplate/simple.qtpl:15
	StreamSimpleQtc(qw422016, u)
	//line quicktemplate/simple.qtpl:15
	qt422016.ReleaseWriter(qw422016)
//line quicktemplate/simple.qtpl:15
}

//line quicktemplate/simple.qtpl:15
func SimpleQtc(u *model.User) string {
	//line quicktemplate/simple.qtpl:15
	qb422016 := qt422016.AcquireByteBuffer()
	//line quicktemplate/simple.qtpl:15
	WriteSimpleQtc(qb422016, u)
	//line quicktemplate/simple.qtpl:15
	qs422016 := string(qb422016.B)
	//line quicktemplate/simple.qtpl:15
	qt422016.ReleaseByteBuffer(qb422016)
	//line quicktemplate/simple.qtpl:15
	return qs422016
//line quicktemplate/simple.qtpl:15
}
