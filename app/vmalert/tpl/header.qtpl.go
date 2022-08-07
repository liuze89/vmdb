// Code generated by qtc from "header.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmalert/tpl/header.qtpl:1
package tpl

//line app/vmalert/tpl/header.qtpl:1
import (
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/VictoriaMetrics/VictoriaMetrics/app/vmalert/utils"
)

//line app/vmalert/tpl/header.qtpl:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmalert/tpl/header.qtpl:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmalert/tpl/header.qtpl:10
func StreamHeader(qw422016 *qt422016.Writer, r *http.Request, navItems []NavItem, title string) {
//line app/vmalert/tpl/header.qtpl:10
	qw422016.N().S(`
    `)
//line app/vmalert/tpl/header.qtpl:11
	prefix := utils.Prefix(r.URL.Path)

//line app/vmalert/tpl/header.qtpl:11
	qw422016.N().S(`
<!DOCTYPE html>
<html lang="en">
<head>
    <title>vmalert`)
//line app/vmalert/tpl/header.qtpl:15
	if title != "" {
//line app/vmalert/tpl/header.qtpl:15
		qw422016.N().S(` - `)
//line app/vmalert/tpl/header.qtpl:15
		qw422016.E().S(title)
//line app/vmalert/tpl/header.qtpl:15
	}
//line app/vmalert/tpl/header.qtpl:15
	qw422016.N().S(`</title>
    <link href="`)
//line app/vmalert/tpl/header.qtpl:16
	qw422016.E().S(prefix)
//line app/vmalert/tpl/header.qtpl:16
	qw422016.N().S(`static/css/bootstrap.min.css" rel="stylesheet" />
    <style>
        body{
          min-height: 75rem;
          padding-top: 4.5rem;
        }
        .group-heading {
            cursor: pointer;
            padding: 5px;
            margin-top: 5px;
            position: relative;
        }
        .group-heading .anchor {
            position:absolute;
            top:-60px;
        }
        .group-heading span {
            float: right;
            margin-left: 5px;
            margin-right: 5px;
        }
         .group-heading:hover {
            background-color: #f8f9fa!important;
        }
        .table {
            table-layout: fixed;
        }
        .table .error-cell{
            word-break: break-word;
            font-size: 14px;
        }
        pre {
            overflow: scroll;
            min-height: 30px;
            max-width: 100%;
        }
        pre::-webkit-scrollbar {
          -webkit-appearance: none;
          width: 0px;
          height: 5px;
        }
        pre::-webkit-scrollbar-thumb {
          border-radius: 5px;
          background-color: rgba(0,0,0,.5);
          -webkit-box-shadow: 0 0 1px rgba(255,255,255,.5);
        }
    </style>
</head>
<body>
    `)
//line app/vmalert/tpl/header.qtpl:65
	streamprintNavItems(qw422016, r, title, navItems)
//line app/vmalert/tpl/header.qtpl:65
	qw422016.N().S(`
    <main class="px-2">
`)
//line app/vmalert/tpl/header.qtpl:67
}

//line app/vmalert/tpl/header.qtpl:67
func WriteHeader(qq422016 qtio422016.Writer, r *http.Request, navItems []NavItem, title string) {
//line app/vmalert/tpl/header.qtpl:67
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/tpl/header.qtpl:67
	StreamHeader(qw422016, r, navItems, title)
//line app/vmalert/tpl/header.qtpl:67
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/tpl/header.qtpl:67
}

//line app/vmalert/tpl/header.qtpl:67
func Header(r *http.Request, navItems []NavItem, title string) string {
//line app/vmalert/tpl/header.qtpl:67
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/tpl/header.qtpl:67
	WriteHeader(qb422016, r, navItems, title)
//line app/vmalert/tpl/header.qtpl:67
	qs422016 := string(qb422016.B)
//line app/vmalert/tpl/header.qtpl:67
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/tpl/header.qtpl:67
	return qs422016
//line app/vmalert/tpl/header.qtpl:67
}

//line app/vmalert/tpl/header.qtpl:71
type NavItem struct {
	Name string
	Url  string
}

//line app/vmalert/tpl/header.qtpl:77
func streamprintNavItems(qw422016 *qt422016.Writer, r *http.Request, current string, items []NavItem) {
//line app/vmalert/tpl/header.qtpl:77
	qw422016.N().S(`
`)
//line app/vmalert/tpl/header.qtpl:79
	prefix := "/vmalert/"
	if strings.HasPrefix(r.URL.Path, prefix) {
		prefix = ""
	}

//line app/vmalert/tpl/header.qtpl:83
	qw422016.N().S(`
<nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark">
  <div class="container-fluid">
    <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul class="navbar-nav me-auto mb-2 mb-md-0">
            `)
//line app/vmalert/tpl/header.qtpl:88
	for _, item := range items {
//line app/vmalert/tpl/header.qtpl:88
		qw422016.N().S(`
                <li class="nav-item">
                    `)
//line app/vmalert/tpl/header.qtpl:91
		u, _ := url.Parse(item.Url)

//line app/vmalert/tpl/header.qtpl:92
		qw422016.N().S(`
                    <a class="nav-link`)
//line app/vmalert/tpl/header.qtpl:93
		if current == item.Name {
//line app/vmalert/tpl/header.qtpl:93
			qw422016.N().S(` active`)
//line app/vmalert/tpl/header.qtpl:93
		}
//line app/vmalert/tpl/header.qtpl:93
		qw422016.N().S(`"
                       href="`)
//line app/vmalert/tpl/header.qtpl:94
		if u.IsAbs() {
//line app/vmalert/tpl/header.qtpl:94
			qw422016.E().S(item.Url)
//line app/vmalert/tpl/header.qtpl:94
		} else {
//line app/vmalert/tpl/header.qtpl:94
			qw422016.E().S(path.Join(prefix, item.Url))
//line app/vmalert/tpl/header.qtpl:94
		}
//line app/vmalert/tpl/header.qtpl:94
		qw422016.N().S(`">
                        `)
//line app/vmalert/tpl/header.qtpl:95
		qw422016.E().S(item.Name)
//line app/vmalert/tpl/header.qtpl:95
		qw422016.N().S(`
                    </a>
                </li>
            `)
//line app/vmalert/tpl/header.qtpl:98
	}
//line app/vmalert/tpl/header.qtpl:98
	qw422016.N().S(`
        </ul>
  </div>
</nav>
`)
//line app/vmalert/tpl/header.qtpl:102
}

//line app/vmalert/tpl/header.qtpl:102
func writeprintNavItems(qq422016 qtio422016.Writer, r *http.Request, current string, items []NavItem) {
//line app/vmalert/tpl/header.qtpl:102
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmalert/tpl/header.qtpl:102
	streamprintNavItems(qw422016, r, current, items)
//line app/vmalert/tpl/header.qtpl:102
	qt422016.ReleaseWriter(qw422016)
//line app/vmalert/tpl/header.qtpl:102
}

//line app/vmalert/tpl/header.qtpl:102
func printNavItems(r *http.Request, current string, items []NavItem) string {
//line app/vmalert/tpl/header.qtpl:102
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmalert/tpl/header.qtpl:102
	writeprintNavItems(qb422016, r, current, items)
//line app/vmalert/tpl/header.qtpl:102
	qs422016 := string(qb422016.B)
//line app/vmalert/tpl/header.qtpl:102
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmalert/tpl/header.qtpl:102
	return qs422016
//line app/vmalert/tpl/header.qtpl:102
}
