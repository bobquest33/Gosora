package main

func init() {
// Site Info
site.Name = "TS" // Should be a setting in the database
site.Email = "" // Should be a setting in the database
site.Url = "localhost"
site.Port = "8080"
site.EnableSsl = false
site.EnableEmails = false
config.SslPrivkey = ""
config.SslFullchain = ""

// Database details
db_config.Host = "localhost"
db_config.Username = "root"
db_config.Password = "password"
db_config.Dbname = "gosora"
db_config.Port = "3306" // You probably won't need to change this

// Limiters
config.MaxRequestSize = 5 * megabyte

// Caching
config.CacheTopicUser = CACHE_STATIC
config.UserCacheCapacity = 100 // The max number of users held in memory
config.TopicCacheCapacity = 100 // The max number of topics held in memory

// Email
config.SmtpServer = ""
config.SmtpUsername = ""
config.SmtpPassword = ""
config.SmtpPort = "25"

// Misc
config.DefaultRoute = route_topics
config.DefaultGroup = 3 // Should be a setting in the database
config.ActivationGroup = 5 // Should be a setting in the database
config.StaffCss = "staff_post"
config.UncategorisedForumVisible = true
config.MinifyTemplates = false
config.MultiServer = false // Experimental: Enable Cross-Server Synchronisation and several other features

//config.Noavatar = "https://api.adorable.io/avatars/{width}/{id}@{site_url}.png"
config.Noavatar = "https://api.adorable.io/avatars/285/{id}@{site_url}.png"
config.ItemsPerPage = 25

// Developer flag
dev.DebugMode = true
//dev.SuperDebug = true
//dev.Profiling = true
}
