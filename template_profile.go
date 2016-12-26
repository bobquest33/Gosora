package main
import "strconv"
import "io"

func init() {
template_profile_handle = template_profile
}

func template_profile(tmpl_profile_vars ProfilePage, w io.Writer) {
w.Write([]byte(`<!doctype html>
<html lang="en">
	<head>
		<title>` + tmpl_profile_vars.Title + `</title>
		<link href="/static/main.css" rel="stylesheet" type="text/css">
		<script type="text/javascript" src="/static/jquery-1.12.3.min.js"></script>
		<script type="text/javascript">
		var session = "` + tmpl_profile_vars.CurrentUser.Session + `";
		</script>
		<script type="text/javascript" src="/static/global.js"></script>
		<meta name="viewport" content="width=device-width,initial-scale = 1.0, maximum-scale=1.0,user-scalable=no" />
	</head>
	<body>
		<div class="container">
<div class="nav">
	<div class="move_left">
	<div class="move_right">
	<ul>
		<li class="menu_left menu_overview"><a href="/">Overview</a></li>
		<li class="menu_left menu_forums"><a href="/forums/">Forums</a></li>
		<li class="menu_left menu_topics"><a href="/">Topics</a></li>
		<li class="menu_left menu_create_topic"><a href="/topics/create/">Create Topic</a></li>
		`))
if tmpl_profile_vars.CurrentUser.Loggedin {
w.Write([]byte(`
		<li class="menu_left menu_account"><a href="/user/edit/critical/">Account</a></li>
		<li class="menu_left menu_profile"><a href="/user/` + strconv.Itoa(tmpl_profile_vars.CurrentUser.ID) + `">Profile</a></li>
		`))
if tmpl_profile_vars.CurrentUser.Is_Super_Mod {
w.Write([]byte(`<li class="menu_left menu_account"><a href="/panel/">Panel</a></li>`))
}
w.Write([]byte(`
		<li class="menu_left menu_logout"><a href="/accounts/logout?session=` + tmpl_profile_vars.CurrentUser.Session + `">Logout</a></li>
		`))
} else {
w.Write([]byte(`
		<li class="menu_left menu_register"><a href="/accounts/create/">Register</a></li>
		<li class="menu_left menu_login"><a href="/accounts/login/">Login</a></li>
		`))
}
w.Write([]byte(`
	</ul>
	</div>
	</div>
	<div style="clear: both;"></div>
</div>
`))
if len(tmpl_profile_vars.NoticeList) != 0 {
for _, item := range tmpl_profile_vars.NoticeList {
w.Write([]byte(`<div class="alert">` + item + `</div>`))
}
}
w.Write([]byte(`
<div class="colblock_left" style="max-width: 220px;">
	<div class="rowitem" style="padding: 0;"><img src="` + tmpl_profile_vars.ProfileOwner.Avatar + `" style="max-width: 100%;margin: 0;"/></div>
	<div class="rowitem" style="text-transform: capitalize;">
	<span style="font-size: 18px;">` + tmpl_profile_vars.ProfileOwner.Name + `</span>`))
if tmpl_profile_vars.ProfileOwner.Tag != "" {
w.Write([]byte(`<span class="username" style="float: right;font-weight: normal;">` + tmpl_profile_vars.ProfileOwner.Tag + `</span>`))
}
w.Write([]byte(`
	</div>
	<div class="rowitem passive">
		<a class="username">Add Friend</a>
		`))
if tmpl_profile_vars.CurrentUser.Is_Super_Mod && !tmpl_profile_vars.ProfileOwner.Is_Super_Mod {
w.Write([]byte(`
		`))
if tmpl_profile_vars.ProfileOwner.Is_Banned {
w.Write([]byte(`<a href="/users/unban/` + strconv.Itoa(tmpl_profile_vars.ProfileOwner.ID) + `?session=` + tmpl_profile_vars.CurrentUser.Session + `" class="username">Unban</a>`))
} else {
w.Write([]byte(`<a href="/users/ban/` + strconv.Itoa(tmpl_profile_vars.ProfileOwner.ID) + `?session=` + tmpl_profile_vars.CurrentUser.Session + `" class="username">Ban</a>`))
}
w.Write([]byte(`
		`))
}
w.Write([]byte(`
		<a href="/report/submit/` + strconv.Itoa(tmpl_profile_vars.ProfileOwner.ID) + `?session=` + tmpl_profile_vars.CurrentUser.Session + `&type=user" class="username report_item">Report</a>
	</div>
</div>
<div class="colblock_right">
	<div class="rowitem"><a>Comments</a></div>
</div>
<div class="colblock_right" style="overflow: hidden;">
	`))
if len(tmpl_profile_vars.ItemList) != 0 {
for _, item := range tmpl_profile_vars.ItemList {
w.Write([]byte(`
	<div class="rowitem passive deletable_block editable_parent" style="`))
if item.Avatar != "" {
w.Write([]byte(`background-image: url(` + item.Avatar + `), url(/static/white-dot.jpg);background-position: 0px `))
if item.ContentLines <= 5 {
w.Write([]byte(`-1`))
}
w.Write([]byte(`0px;background-repeat: no-repeat, repeat-y;background-size: 128px;padding-left: 136px;` + string(item.Css)))
}
w.Write([]byte(`">
		<span class="editable_block user_content">` + string(item.ContentHtml) + `</span>
		<br /><br />
		<a href="/user/` + strconv.Itoa(item.CreatedBy) + `" class="username">` + item.CreatedByName + `</a>
		`))
if tmpl_profile_vars.CurrentUser.Is_Mod {
w.Write([]byte(`<a href="/profile/reply/edit/submit/` + strconv.Itoa(item.ID) + `"><button class="username edit_item">Edit</button></a>
		<a href="/profile/reply/delete/submit/` + strconv.Itoa(item.ID) + `"><button class="username delete_item">Delete</button></a>`))
}
w.Write([]byte(`
		<a href="/report/submit/` + strconv.Itoa(item.ID) + `?session=` + tmpl_profile_vars.CurrentUser.Session + `&type=user-reply"><button class="username report_item">Report</button></a>
		`))
if item.Tag != "" {
w.Write([]byte(`<a class="username hide_on_mobile" style="float: right;">` + item.Tag + `</a>`))
}
w.Write([]byte(`
	</div>`))
}
}
w.Write([]byte(`
</div>
`))
if !tmpl_profile_vars.CurrentUser.Is_Banned {
w.Write([]byte(`
<div class="colblock_right">
	<form action="/profile/reply/create/" method="post">
		<input name="uid" value='` + strconv.Itoa(tmpl_profile_vars.ProfileOwner.ID) + `' type="hidden" />
		<div class="formrow">
			<div class="formitem"><textarea name="reply-content" placeholder="Insert reply here"></textarea></div>
		</div>
		<div class="formrow">
			<div class="formitem"><button name="reply-button" class="formbutton">Create Reply</button></div>
		</div>
	</form>
</div>
`))
}
w.Write([]byte(`
			<!--<link rel="stylesheet" href="https://use.fontawesome.com/8670aa03ca.css">-->
		</div>
	</body>
</html>`))
}
