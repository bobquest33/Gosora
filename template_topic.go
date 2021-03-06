// +build !no_templategen

// Code generated by Gosora. More below:
/* This file was automatically generated by the software. Please don't edit it as your changes may be overwritten at any moment. */
package main
import "net/http"
import "strconv"

func init() {
	template_topic_handle = template_topic
	//o_template_topic_handle = template_topic
	ctemplates = append(ctemplates,"topic")
	tmpl_ptr_map["topic"] = &template_topic_handle
	tmpl_ptr_map["o_topic"] = template_topic
}

func template_topic(tmpl_topic_vars TopicPage, w http.ResponseWriter) {
w.Write(header_0)
w.Write([]byte(tmpl_topic_vars.Title))
w.Write(header_1)
if len(tmpl_topic_vars.Header.Stylesheets) != 0 {
for _, item := range tmpl_topic_vars.Header.Stylesheets {
w.Write(header_2)
w.Write([]byte(item))
w.Write(header_3)
}
}
w.Write(header_4)
if len(tmpl_topic_vars.Header.Scripts) != 0 {
for _, item := range tmpl_topic_vars.Header.Scripts {
w.Write(header_5)
w.Write([]byte(item))
w.Write(header_6)
}
}
w.Write(header_7)
w.Write([]byte(tmpl_topic_vars.CurrentUser.Session))
w.Write(header_8)
if !tmpl_topic_vars.CurrentUser.Is_Super_Mod {
w.Write(header_9)
}
w.Write(header_10)
w.Write(menu_0)
w.Write([]byte(tmpl_topic_vars.Header.Site.Name))
w.Write(menu_1)
if tmpl_topic_vars.CurrentUser.Loggedin {
w.Write(menu_2)
w.Write([]byte(tmpl_topic_vars.CurrentUser.Link))
w.Write(menu_3)
w.Write([]byte(tmpl_topic_vars.CurrentUser.Session))
w.Write(menu_4)
} else {
w.Write(menu_5)
}
w.Write(menu_6)
w.Write(header_11)
if tmpl_topic_vars.Header.Widgets.RightSidebar != "" {
w.Write(header_12)
}
w.Write(header_13)
if len(tmpl_topic_vars.Header.NoticeList) != 0 {
for _, item := range tmpl_topic_vars.Header.NoticeList {
w.Write(header_14)
w.Write([]byte(item))
w.Write(header_15)
}
}
w.Write(topic_0)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_1)
if tmpl_topic_vars.Page > 1 {
w.Write(topic_2)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_3)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Page - 1)))
w.Write(topic_4)
}
if tmpl_topic_vars.LastPage != tmpl_topic_vars.Page {
w.Write(topic_5)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_6)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Page + 1)))
w.Write(topic_7)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_8)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Page + 1)))
w.Write(topic_9)
}
w.Write(topic_10)
if tmpl_topic_vars.Topic.Sticky {
w.Write(topic_11)
} else {
if tmpl_topic_vars.Topic.Is_Closed {
w.Write(topic_12)
}
}
w.Write(topic_13)
w.Write([]byte(tmpl_topic_vars.Topic.Title))
w.Write(topic_14)
if tmpl_topic_vars.Topic.Is_Closed {
w.Write(topic_15)
}
if tmpl_topic_vars.CurrentUser.Perms.EditTopic {
w.Write(topic_16)
w.Write([]byte(tmpl_topic_vars.Topic.Title))
w.Write(topic_17)
if tmpl_topic_vars.CurrentUser.Perms.CloseTopic {
w.Write(topic_18)
}
w.Write(topic_19)
}
w.Write(topic_20)
w.Write([]byte(tmpl_topic_vars.Topic.ClassName))
w.Write(topic_21)
if tmpl_topic_vars.Topic.Avatar != "" {
w.Write(topic_22)
w.Write([]byte(tmpl_topic_vars.Topic.Avatar))
w.Write(topic_23)
if tmpl_topic_vars.Topic.ContentLines <= 5 {
w.Write(topic_24)
}
w.Write(topic_25)
}
w.Write(topic_26)
w.Write([]byte(tmpl_topic_vars.Topic.Content))
w.Write(topic_27)
w.Write([]byte(tmpl_topic_vars.Topic.Content))
w.Write(topic_28)
w.Write([]byte(tmpl_topic_vars.Topic.UserLink))
w.Write(topic_29)
w.Write([]byte(tmpl_topic_vars.Topic.CreatedByName))
w.Write(topic_30)
if tmpl_topic_vars.CurrentUser.Perms.LikeItem {
w.Write(topic_31)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_32)
if tmpl_topic_vars.Topic.Liked {
w.Write(topic_33)
}
w.Write(topic_34)
}
if tmpl_topic_vars.CurrentUser.Perms.EditTopic {
w.Write(topic_35)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_36)
}
if tmpl_topic_vars.CurrentUser.Perms.DeleteTopic {
w.Write(topic_37)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_38)
}
if tmpl_topic_vars.CurrentUser.Perms.PinTopic {
if tmpl_topic_vars.Topic.Sticky {
w.Write(topic_39)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_40)
} else {
w.Write(topic_41)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_42)
}
}
w.Write(topic_43)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_44)
w.Write([]byte(tmpl_topic_vars.CurrentUser.Session))
w.Write(topic_45)
if tmpl_topic_vars.Topic.LikeCount > 0 {
w.Write(topic_46)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.LikeCount)))
w.Write(topic_47)
}
if tmpl_topic_vars.Topic.Tag != "" {
w.Write(topic_48)
w.Write([]byte(tmpl_topic_vars.Topic.Tag))
w.Write(topic_49)
} else {
w.Write(topic_50)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.Level)))
w.Write(topic_51)
}
w.Write(topic_52)
if len(tmpl_topic_vars.ItemList) != 0 {
for _, item := range tmpl_topic_vars.ItemList {
if item.ActionType != "" {
w.Write(topic_53)
w.Write([]byte(item.ActionIcon))
w.Write(topic_54)
w.Write([]byte(item.ActionType))
w.Write(topic_55)
} else {
w.Write(topic_56)
w.Write([]byte(item.ClassName))
w.Write(topic_57)
if item.Avatar != "" {
w.Write(topic_58)
w.Write([]byte(item.Avatar))
w.Write(topic_59)
if item.ContentLines <= 5 {
w.Write(topic_60)
}
w.Write(topic_61)
}
w.Write(topic_62)
w.Write([]byte(item.ContentHtml))
w.Write(topic_63)
w.Write([]byte(item.UserLink))
w.Write(topic_64)
w.Write([]byte(item.CreatedByName))
w.Write(topic_65)
if tmpl_topic_vars.CurrentUser.Perms.LikeItem {
w.Write(topic_66)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_67)
if item.Liked {
w.Write(topic_68)
}
w.Write(topic_69)
}
if tmpl_topic_vars.CurrentUser.Perms.EditReply {
w.Write(topic_70)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_71)
}
if tmpl_topic_vars.CurrentUser.Perms.DeleteReply {
w.Write(topic_72)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_73)
}
w.Write(topic_74)
w.Write([]byte(strconv.Itoa(item.ID)))
w.Write(topic_75)
w.Write([]byte(tmpl_topic_vars.CurrentUser.Session))
w.Write(topic_76)
if item.LikeCount > 0 {
w.Write(topic_77)
w.Write([]byte(strconv.Itoa(item.LikeCount)))
w.Write(topic_78)
}
if item.Tag != "" {
w.Write(topic_79)
w.Write([]byte(item.Tag))
w.Write(topic_80)
} else {
w.Write(topic_81)
w.Write([]byte(strconv.Itoa(item.Level)))
w.Write(topic_82)
}
w.Write(topic_83)
}
}
}
w.Write(topic_84)
if tmpl_topic_vars.CurrentUser.Perms.CreateReply {
w.Write(topic_85)
w.Write([]byte(strconv.Itoa(tmpl_topic_vars.Topic.ID)))
w.Write(topic_86)
}
w.Write(topic_87)
w.Write(footer_0)
if tmpl_topic_vars.Header.Widgets.RightSidebar != "" {
w.Write(footer_1)
w.Write([]byte(string(tmpl_topic_vars.Header.Widgets.RightSidebar)))
w.Write(footer_2)
}
w.Write(footer_3)
}
