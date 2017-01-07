/* This file was automatically generated by the software. Please don't edit it as your changes may be overwritten at any moment. */
package main
import "html/template"
import "io"
import "strconv"

func init() {
	template_topic_handle = template_topic
	//o_template_topic_handle = template_topic
	ctemplates = append(ctemplates,"topic")
	tmpl_ptr_map["topic"] = &template_topic_handle
	tmpl_ptr_map["o_topic"] = template_topic
}

func template_topic(tmpl_topic_vars TopicPage, w io.Writer) {
w.Write([]byte(`<!doctype html>
<html lang="en">
	<head>
		<title>` + tmpl_topic_vars.Title + `</title>
		<link href="/static/main.css" rel="stylesheet" type="text/css">
		<script type="text/javascript" src="/static/jquery-1.12.3.min.js"></script>
		<script type="text/javascript">
		var session = "` + tmpl_topic_vars.CurrentUser.Session + `";
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
if tmpl_topic_vars.CurrentUser.Loggedin {
w.Write([]byte(`
		<li class="menu_left menu_account"><a href="/user/edit/critical/">Account</a></li>
		<li class="menu_left menu_profile"><a href="/user/` + strconv.Itoa(tmpl_topic_vars.CurrentUser.ID) + `">Profile</a></li>
		`))
if tmpl_topic_vars.CurrentUser.Is_Super_Mod {
w.Write([]byte(`<li class="menu_left menu_account"><a href="/panel/">Panel</a></li>`))
}
w.Write([]byte(`
		<li class="menu_left menu_logout"><a href="/accounts/logout?session=` + tmpl_topic_vars.CurrentUser.Session + `">Logout</a></li>
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
<div id="back"><div id="main">`))
if len(tmpl_topic_vars.NoticeList) != 0 {
for _, item := range tmpl_topic_vars.NoticeList {
w.Write([]byte(`<div class="alert">` + item + `</div>`))
}
}
w.Write([]byte(`
<div class="rowblock">
	<form action='/topic/edit/submit/` + strconv.Itoa(tmpl_topic_vars.Topic.ID) + `' method="post">
		<div class="rowitem"`))
if tmpl_topic_vars.Topic.Sticky {
w.Write([]byte(` style="background-color: #FFFFEA;"`))
} else {
if tmpl_topic_vars.Topic.Is_Closed {
w.Write([]byte(` style="background-color: #eaeaea;"`))
}
}
w.Write([]byte(`>
			<a class='topic_name hide_on_edit'>` + tmpl_topic_vars.Topic.Title + `</a> 
			<span class='username hide_on_micro topic_status_e topic_status_` + tmpl_topic_vars.Topic.Status + ` hide_on_edit' style="font-weight:normal;float: right;">` + tmpl_topic_vars.Topic.Status + `</span> 
			<span class="username hide_on_micro" style="border-right: 0;font-weight: normal;float: right;">Status</span>
			`))
if tmpl_topic_vars.CurrentUser.Is_Mod {
w.Write([]byte(`
			<a href='/topic/edit/` + strconv.Itoa(tmpl_topic_vars.Topic.ID) + `' class="username hide_on_edit open_edit" style="font-weight: normal;margin-left: 6px;">Edit</a>
			<a href='/topic/delete/submit/` + strconv.Itoa(tmpl_topic_vars.Topic.ID) + `' class="username" style="font-weight: normal;">Delete</a>
			`))
if tmpl_topic_vars.Topic.Sticky {
w.Write([]byte(`<a href='/topic/unstick/submit/` + strconv.Itoa(tmpl_topic_vars.Topic.ID) + `' class="username" style="font-weight: normal;">Unpin</a>`))
} else {
w.Write([]byte(`<a href='/topic/stick/submit/` + strconv.Itoa(tmpl_topic_vars.Topic.ID) + `' class="username" style="font-weight: normal;">Pin</a>`))
}
w.Write([]byte(`
			
			<input class='show_on_edit topic_name_input' name="topic_name" value='` + tmpl_topic_vars.Topic.Title + `' type="text" />
			<select name="topic_status" class='show_on_edit topic_status_input' style='float: right;'>
				<option>open</option>
				<option>shut</option>
			</select>
			<button name="topic-button" class="formbutton show_on_edit submit_edit">Update</button>
			`))
}
w.Write([]byte(`
			<a href="/report/submit/` + strconv.Itoa(tmpl_topic_vars.Topic.ID) + `?session=` + tmpl_topic_vars.CurrentUser.Session + `&type=topic" class="username report_item" style="font-weight: normal;">Report</a>
		</div>
	</form>
</div>
<div class="rowblock post_container">
	<div class="rowitem passive editable_parent post_item" style="border-bottom: none;`))
if tmpl_topic_vars.Topic.Avatar != "" {
w.Write([]byte(`background-image: url(` + tmpl_topic_vars.Topic.Avatar + `), url(/static/white-dot.jpg);background-position: 0px `))
if tmpl_topic_vars.Topic.ContentLines <= 5 {
w.Write([]byte(`-1`))
}
w.Write([]byte(`0px;background-repeat: no-repeat, repeat-y;background-size: 128px;padding-left: 136px;` + string(tmpl_topic_vars.Topic.Css)))
}
w.Write([]byte(`">
		<p class="hide_on_edit topic_content user_content" style="margin: 0;padding: 0;">` + string(tmpl_topic_vars.Topic.Content.(template.HTML)) + `</p>
		<textarea name="topic_content" class="show_on_edit topic_content_input">` + string(tmpl_topic_vars.Topic.Content.(template.HTML)) + `</textarea>
		<br /><br />
		<a href="/user/` + strconv.Itoa(tmpl_topic_vars.Topic.CreatedBy) + `" class="username real_username">` + tmpl_topic_vars.Topic.CreatedByName + `</a>
		`))
if tmpl_topic_vars.Topic.Tag != "" {
w.Write([]byte(`<a class="username hide_on_micro" style="float: right;">` + tmpl_topic_vars.Topic.Tag + `</a>`))
} else {
if tmpl_topic_vars.Topic.URLName != "" {
w.Write([]byte(`<a href="` + tmpl_topic_vars.Topic.URL + `" class="username" style="color: #505050;float: right;">` + tmpl_topic_vars.Topic.URLName + `</a>
		<a class="username" style="color: #505050;float: right;border-right: 0;">` + tmpl_topic_vars.Topic.URLPrefix + `</a>`))
}
}
w.Write([]byte(`
	</div>
</div><br />
<div class="rowblock post_container" style="overflow: hidden;">`))
if len(tmpl_topic_vars.ItemList) != 0 {
for _, item := range tmpl_topic_vars.ItemList {
w.Write([]byte(`
	<div class="rowitem rowhead passive deletable_block editable_parent post_item" style="`))
if item.Avatar != "" {
w.Write([]byte(`background-image: url(` + item.Avatar + `), url(/static/white-dot.jpg);background-position: 0px `))
if item.ContentLines <= 5 {
w.Write([]byte(`-1`))
}
w.Write([]byte(`0px;background-repeat: no-repeat, repeat-y;background-size: 128px;padding-left: 136px;` + string(item.Css)))
}
w.Write([]byte(`">
		<p class="editable_block user_content" style="margin: 0;padding: 0;">` + string(item.ContentHtml) + `</p><br /><br />
		<a href="/user/` + strconv.Itoa(item.CreatedBy) + `" class="username real_username">` + item.CreatedByName + `</a>
		`))
if tmpl_topic_vars.CurrentUser.Perms.EditReply {
w.Write([]byte(`<a href="/reply/edit/submit/` + strconv.Itoa(item.ID) + `" class="mod_button"><button class="username edit_item">Edit</button></a>`))
}
w.Write([]byte(`
		`))
if tmpl_topic_vars.CurrentUser.Perms.DeleteReply {
w.Write([]byte(`<a href="/reply/delete/submit/` + strconv.Itoa(item.ID) + `" class="mod_button"><button class="username delete_item">Delete</button></a>`))
}
w.Write([]byte(`
		<a href="/report/submit/` + strconv.Itoa(item.ID) + `?session=` + tmpl_topic_vars.CurrentUser.Session + `&type=reply" class="mod_button"><button class="username report_item">Report</button></a>
		`))
if item.Tag != "" {
w.Write([]byte(`<a class="username hide_on_micro" style="float: right;">` + item.Tag + `</a>`))
} else {
if item.URLName != "" {
w.Write([]byte(`<a href="` + item.URL + `" class="username hide_on_mobile" style="color: #505050;float: right;" rel="nofollow">` + item.URLName + `</a>
		<a class="username hide_on_mobile" style="color: #505050;float: right;border-right: 0;">` + item.URLPrefix + `</a>`))
}
}
w.Write([]byte(`
	</div>
`))
}
}
w.Write([]byte(`</div>
`))
if tmpl_topic_vars.CurrentUser.Perms.CreateReply {
w.Write([]byte(`
<div class="rowblock">
	<form action="/reply/create/" method="post">
		<input name="tid" value='` + strconv.Itoa(tmpl_topic_vars.Topic.ID) + `' type="hidden" />
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
		</div><div style="clear: both;"></div></div></div>
	</body>
</html>`))
}
