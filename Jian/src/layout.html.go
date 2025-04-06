// DO NOT EDIT!
// Generate By Goh

package main

import (
	"bytes"
	"github.com/OblivionOcean/Daizen/model"
	"github.com/OblivionOcean/Daizen/utils"

	"github.com/OblivionOcean/Goh/utils"
)

func Root(site *model.SiteInfo, page *model.Page, body *bytes.Buffer, buf *bytes.Buffer) {
	buf.Grow(2354)
	buf.WriteString(`


<!DOCTYPE html>
<html lang="">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>`)
	buf.WriteString(`</title>
    <link rel="shortcut icon" href="`)
	Goh.EscapeHTML(site.Cfg.GetString("favicon", site.Theme.Cfg.GetString("favicon")), buf)
	buf.WriteString(`">
    <link rel="stylesheet" href="`)
	Goh.EscapeHTML(site.Theme.Cfg.GetString("cdn.domain"), buf)
	buf.WriteString(`/css/style.css">
    <link rel="stylesheet" href="`)
	Goh.EscapeHTML(site.Theme.Cfg.GetString("cdn.domain"), buf)
	buf.WriteString(`/css/grid.css">
    <script src="`)
	Goh.EscapeHTML(site.Theme.Cfg.GetString("cdn.domain"), buf)
	buf.WriteString(`/js/app.js"></script>
    <script src="`)
	Goh.EscapeHTML(site.Theme.Cfg.GetString("cdn.domain"), buf)
	buf.WriteString(`/js/srjs.js"></script>
    <meta http-equiv="X-UA-Compatible" content="IE=Edge,chrome=1">
    
        <meta name="keywords" content="">
        
    
        <meta name="keywords" content="`)
	Goh.EscapeHTML(utils.Slice(site.Cfg.GetString("keywords"), 0, 157), buf)
	buf.WriteString(`">
        
    
    <script>
        Jian.load.js('`)
	Goh.EscapeHTML(site.Theme.Cfg.GetString("cdn.npm_cdn"), buf)
	buf.WriteString(`/pandown@2.0.0', function () {
            pandown();
        });
    </script>
    <meta name="follow.it-verification-code" content="LlmddNY7dr0ZbV24UdUh"/>
    <meta name="360-site-verification" content="6e6223d4a72a81e520e67080db00fda6"/>
    <meta name="baidu-site-verification" content="codeva-Oz6pd5JcaV"/>
</head>
<body>

<div class="loading_page"><div class="loading_i center"></div><style>*{transition:none}.loading_i{visibility:visible}body{visibility:hidden}</style></div>
<header class="header card bg-`)
	buf.WriteString(site.Theme.Cfg.GetString("color.theme", "blue"))
	buf.WriteString(`" id="header">
    <div class="blog-logo">
        <a href="" class="logo">`)
	Goh.EscapeHTML(site.Cfg.GetString("title"), buf)
	buf.WriteString(`</a>
    </div>
    <nav class="navbar">
        <ul class="menu">
            
                <li class="menu-item">
                    
                        <a href="" class="menu-item-link"></a>
                    
                        <a href="" class="menu-item-link" aria-label="" title=""></a>
                    
                </li>
            
            <li class="menu-item">
                <a href="javascript:Jian.dark.change()" class="menu-item-link" aria-label="暗色模式" title="暗色模式"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-moon-fill" viewBox="0 0 16 16">
  <path d="M6 .278a.77.77 0 0 1 .08.858 7.2 7.2 0 0 0-.878 3.46c0 4.021 3.278 7.277 7.318 7.277q.792-.001 1.533-.16a.79.79 0 0 1 .81.316.73.73 0 0 1-.031.893A8.35 8.35 0 0 1 8.344 16C3.734 16 0 12.286 0 7.71 0 4.266 2.114 1.312 5.124.06A.75.75 0 0 1 6 .278"/>
</svg></a>
            </li>
        </ul>
    </nav>
</header>
<div class="msg" id="msg">
</div>
<main class="main w-full grid page-main no-side">
    `)
	buf.Write(body.Bytes())
	buf.WriteString(`
    
</main>

</body>
</html>
`)
}
