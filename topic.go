package main
//import "fmt"
import "strconv"
import "html/template"

type Topic struct
{
	ID int
	Link string
	Title string
	Content string
	CreatedBy int
	Is_Closed bool
	Sticky bool
	CreatedAt string
	LastReplyAt string
	//LastReplyBy int
	ParentID int
	Status string // Deprecated. Marked for removal.
	IpAddress string
	PostCount int
	LikeCount int
	ClassName string // CSS Class Name
	Data string // Used for report metadata
}

type TopicUser struct
{
	ID int
	Link string
	Title string
	Content string
	CreatedBy int
	Is_Closed bool
	Sticky bool
	CreatedAt string
	LastReplyAt string
	//LastReplyBy int
	ParentID int
	Status string // Deprecated. Marked for removal.
	IpAddress string
	PostCount int
	LikeCount int
	ClassName string
	Data string // Used for report metadata

	UserLink string
	CreatedByName string
	Group int
	Avatar string
	ContentLines int
	Tag string
	URL string
	URLPrefix string
	URLName string
	Level int
	Liked bool
}

type TopicsRow struct
{
	ID int
	Link string
	Title string
	Content string
	CreatedBy int
	Is_Closed bool
	Sticky bool
	CreatedAt string
	LastReplyAt string
	LastReplyBy int
	ParentID int
	Status string // Deprecated. Marked for removal. -Is there anything we could use it for?
	IpAddress string
	PostCount int
	LikeCount int
	ClassName string
	Data string // Used for report metadata

	Creator *User
	Css template.CSS
	ContentLines int
	LastUser *User

	ForumName string //TopicsRow
	ForumLink string
}

func get_topicuser(tid int) (TopicUser,error) {
	if config.CacheTopicUser != CACHE_SQL {
		topic, err := topics.Get(tid)
		if err == nil {
			user, err := users.CascadeGet(topic.CreatedBy)
			if err != nil {
				return TopicUser{ID:tid}, err
			}

			// We might be better off just passing seperate topic and user structs to the caller?
			return copy_topic_to_topicuser(topic, user), nil
		} else if users.GetLength() < users.GetCapacity() {
			topic, err = topics.CascadeGet(tid)
			if err != nil {
				return TopicUser{ID:tid}, err
			}
			user, err := users.CascadeGet(topic.CreatedBy)
			if err != nil {
				return TopicUser{ID:tid}, err
			}
			return copy_topic_to_topicuser(topic, user), nil
		}
	}

	tu := TopicUser{ID:tid}
	err := get_topic_user_stmt.QueryRow(tid).Scan(&tu.Title, &tu.Content, &tu.CreatedBy, &tu.CreatedAt, &tu.Is_Closed, &tu.Sticky, &tu.ParentID, &tu.IpAddress, &tu.PostCount, &tu.LikeCount, &tu.CreatedByName, &tu.Avatar, &tu.Group, &tu.URLPrefix, &tu.URLName, &tu.Level)
	tu.Link = build_topic_url(name_to_slug(tu.Title),tu.ID)
	tu.UserLink = build_profile_url(name_to_slug(tu.CreatedByName),tu.CreatedBy)

	the_topic := Topic{ID:tu.ID, Link:tu.Link, Title:tu.Title, Content:tu.Content, CreatedBy:tu.CreatedBy, Is_Closed:tu.Is_Closed, Sticky:tu.Sticky, CreatedAt:tu.CreatedAt, LastReplyAt:tu.LastReplyAt, ParentID:tu.ParentID, IpAddress:tu.IpAddress, PostCount:tu.PostCount, LikeCount:tu.LikeCount}
	//log.Printf("the_topic: %+v\n", the_topic)
	tu.Tag = groups[tu.Group].Tag
	topics.Add(&the_topic)
	return tu, err
}

func copy_topic_to_topicuser(topic *Topic, user *User) (tu TopicUser) {
	tu.UserLink = user.Link
	tu.CreatedByName = user.Name
	tu.Group = user.Group
	tu.Avatar = user.Avatar
	tu.URLPrefix = user.URLPrefix
	tu.URLName = user.URLName
	tu.Level = user.Level

	tu.ID = topic.ID
	tu.Link = topic.Link
	tu.Title = topic.Title
	tu.Content = topic.Content
	tu.CreatedBy = topic.CreatedBy
	tu.Is_Closed = topic.Is_Closed
	tu.Sticky = topic.Sticky
	tu.CreatedAt = topic.CreatedAt
	tu.LastReplyAt = topic.LastReplyAt
	tu.ParentID = topic.ParentID
	tu.IpAddress = topic.IpAddress
	tu.PostCount = topic.PostCount
	tu.LikeCount = topic.LikeCount
	tu.Data = topic.Data

	return tu
}

func get_topic_by_reply(rid int) (*Topic, error) {
	topic := Topic{ID:0}
	err := get_topic_by_reply_stmt.QueryRow(rid).Scan(&topic.ID, &topic.Title, &topic.Content, &topic.CreatedBy, &topic.CreatedAt, &topic.Is_Closed, &topic.Sticky, &topic.ParentID, &topic.IpAddress, &topic.PostCount, &topic.LikeCount, &topic.Data)
	topic.Link = build_topic_url(name_to_slug(topic.Title),topic.ID)
	return &topic, err
}

func build_topic_url(slug string, tid int) string {
	if slug == "" {
		return "/topic/" + strconv.Itoa(tid)
	}
	return "/topic/" + slug + "." + strconv.Itoa(tid)
}

func get_topic_url_prefix() string {
	return "/topic/"
}
