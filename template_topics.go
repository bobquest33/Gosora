// +build !no_templategen

// Code generated by Gosora. More below:
/* This file was automatically generated by the software. Please don't edit it as your changes may be overwritten at any moment. */
package main
import "net/http"
import "strconv"

func init() {
	template_topics_handle = template_topics
	//o_template_topics_handle = template_topics
	ctemplates = append(ctemplates,"topics")
	tmpl_ptr_map["topics"] = &template_topics_handle
	tmpl_ptr_map["o_topics"] = template_topics
}

func template_topics(tmpl_topics_vars TopicsPage, w http.ResponseWriter) {
w.Write(header_0)
w.Write([]byte(tmpl_topics_vars.Title))
w.Write(header_1)
if len(tmpl_topics_vars.Header.Stylesheets) != 0 {
for _, item := range tmpl_topics_vars.Header.Stylesheets {
w.Write(header_2)
w.Write([]byte(item))
w.Write(header_3)
}
}
w.Write(header_4)
if len(tmpl_topics_vars.Header.Scripts) != 0 {
for _, item := range tmpl_topics_vars.Header.Scripts {
w.Write(header_5)
w.Write([]byte(item))
w.Write(header_6)
}
}
w.Write(header_7)
w.Write([]byte(tmpl_topics_vars.CurrentUser.Session))
w.Write(header_8)
if !tmpl_topics_vars.CurrentUser.Is_Super_Mod {
w.Write(header_9)
}
w.Write(header_10)
w.Write(menu_0)
w.Write([]byte(tmpl_topics_vars.Header.Site.Name))
w.Write(menu_1)
if tmpl_topics_vars.CurrentUser.Loggedin {
w.Write(menu_2)
w.Write([]byte(tmpl_topics_vars.CurrentUser.Link))
w.Write(menu_3)
w.Write([]byte(tmpl_topics_vars.CurrentUser.Session))
w.Write(menu_4)
} else {
w.Write(menu_5)
}
w.Write(menu_6)
w.Write(header_11)
if tmpl_topics_vars.Header.Widgets.RightSidebar != "" {
w.Write(header_12)
}
w.Write(header_13)
if len(tmpl_topics_vars.Header.NoticeList) != 0 {
for _, item := range tmpl_topics_vars.Header.NoticeList {
w.Write(header_14)
w.Write([]byte(item))
w.Write(header_15)
}
}
w.Write(topics_0)
if len(tmpl_topics_vars.ItemList) != 0 {
for _, item := range tmpl_topics_vars.ItemList {
w.Write(topics_1)
if item.Sticky {
w.Write(topics_2)
} else {
if item.Is_Closed {
w.Write(topics_3)
}
}
w.Write(topics_4)
if item.Creator.Avatar != "" {
w.Write(topics_5)
w.Write([]byte(item.Creator.Avatar))
w.Write(topics_6)
}
w.Write(topics_7)
w.Write([]byte(strconv.Itoa(item.PostCount)))
w.Write(topics_8)
w.Write([]byte(item.LastReplyAt))
w.Write(topics_9)
w.Write([]byte(item.Link))
w.Write(topics_10)
w.Write([]byte(item.Title))
w.Write(topics_11)
if item.ForumName != "" {
w.Write(topics_12)
w.Write([]byte(item.ForumLink))
w.Write(topics_13)
w.Write([]byte(item.ForumName))
w.Write(topics_14)
}
w.Write(topics_15)
w.Write([]byte(item.Creator.Link))
w.Write(topics_16)
w.Write([]byte(item.Creator.Name))
w.Write(topics_17)
if item.Is_Closed {
w.Write(topics_18)
}
if item.Sticky {
w.Write(topics_19)
}
w.Write(topics_20)
if item.LastUser.Avatar != "" {
w.Write(topics_21)
w.Write([]byte(item.LastUser.Avatar))
w.Write(topics_22)
}
w.Write(topics_23)
w.Write([]byte(item.LastUser.Link))
w.Write(topics_24)
w.Write([]byte(item.LastUser.Name))
w.Write(topics_25)
w.Write([]byte(item.LastReplyAt))
w.Write(topics_26)
}
} else {
w.Write(topics_27)
if tmpl_topics_vars.CurrentUser.Perms.CreateTopic {
w.Write(topics_28)
}
w.Write(topics_29)
}
w.Write(topics_30)
w.Write(footer_0)
if tmpl_topics_vars.Header.Widgets.RightSidebar != "" {
w.Write(footer_1)
w.Write([]byte(string(tmpl_topics_vars.Header.Widgets.RightSidebar)))
w.Write(footer_2)
}
w.Write(footer_3)
}
