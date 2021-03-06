/* Copyright Azareal 2016 - 2018 */
package main
import "log"
import "net/http"

var plugins map[string]*Plugin = make(map[string]*Plugin)

// Hooks with a single argument. Is this redundant? Might be useful for inlining, as variadics aren't inlined? Are closures even inlined to begin with?
var hooks map[string][]func(interface{})interface{} = map[string][]func(interface{})interface{}{
	"forums_frow_assign": nil,
	"topic_create_frow_assign": nil,
	"rrow_assign": nil, // TO-DO: Rename this hook to topic_rrow_assign
}

// Hooks with a variable number of arguments
var vhooks map[string]func(...interface{})interface{} = map[string]func(...interface{})interface{}{
	"simple_forum_check_pre_perms": nil,
	"forum_check_pre_perms": nil,
	"intercept_build_widgets": nil,
	"forum_trow_assign": nil,
	"topics_topic_row_assign": nil,
	//"topics_user_row_assign": nil,
	"create_group_preappend": nil, // What is this? Investigate!
	"topic_create_pre_loop": nil,
}

// Hooks which take in and spit out a string. This is usually used for parser components
var sshooks map[string][]func(string)string = map[string][]func(string)string{
	"preparse_preassign": nil,
	"parse_assign": nil,
}

// The hooks which run before the template is rendered for a route
var pre_render_hooks map[string][]func(http.ResponseWriter, *http.Request, *User, interface{})bool = map[string][]func(http.ResponseWriter, *http.Request, *User, interface{})bool{
	"pre_render": nil,

	"pre_render_forum_list": nil,
	"pre_render_view_forum": nil,
	"pre_render_topic_list": nil,
	"pre_render_view_topic": nil,
	"pre_render_profile": nil,
	"pre_render_custom_page": nil,
	"pre_render_overview": nil,
	"pre_render_create_topic": nil,

	"pre_render_account_own_edit_critical": nil,
	"pre_render_account_own_edit_avatar": nil,
	"pre_render_account_own_edit_username": nil,
	"pre_render_account_own_edit_email": nil,
	"pre_render_login": nil,
	"pre_render_register": nil,
	"pre_render_ban": nil,

	"pre_render_panel_dashboard": nil,
	"pre_render_panel_forums": nil,
	"pre_render_panel_delete_forum": nil,
	"pre_render_panel_edit_forum": nil,
	"pre_render_panel_settings": nil,
	"pre_render_panel_setting": nil,
	"pre_render_panel_plugins": nil,
	"pre_render_panel_users": nil,
	"pre_render_panel_edit_user": nil,
	"pre_render_panel_groups": nil,
	"pre_render_panel_edit_group": nil,
	"pre_render_panel_edit_group_perms": nil,
	"pre_render_panel_themes": nil,
	"pre_render_panel_mod_log": nil,

	"pre_render_error": nil, // Note: This hook isn't run for a few errors whose templates are computed at startup and reused, such as InternalError. This hook is also not available in JS mode.
	"pre_render_security_error": nil,
}

type Plugin struct
{
	UName string
	Name string
	Author string
	URL string
	Settings string
	Active bool
	Tag string
	Type string
	Installable bool
	Installed bool

	Init func()error
	Activate func()error
	Deactivate func()
	Install func()error
	Uninstall func()error

	Hooks map[string]int
}

func LoadPlugins() error {
	rows, err := get_plugins_stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	var uname string
	var active bool
	var installed bool
	for rows.Next() {
		err = rows.Scan(&uname, &active, &installed)
		if err != nil {
			return err
		}

		// Was the plugin deleted at some point?
		plugin, ok := plugins[uname]
		if !ok {
			continue
		}
		plugin.Active = active
		plugin.Installed = installed
		plugins[uname] = plugin
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

func NewPlugin(uname string, name string, author string, url string, settings string, tag string, ptype string, init func()error, activate func()error, deactivate func(), install func()error, uninstall func()error) *Plugin {
	return &Plugin{
		UName: uname,
		Name: name,
		Author: author,
		URL: url,
		Settings: settings,
		Tag: tag,
		Type: ptype,
		Installable: (install != nil),
		Init: init,
		Activate: activate,
		Deactivate: deactivate,
		Install: install,
		//Uninstall: uninstall,

		/*
		The Active field should never be altered by a plugin. It's used internally by the software to determine whether an admin has enabled a plugin or not and whether to run it. This will be overwritten by the user's preference.
		*/
		Active: false,
		Installed: false,
		Hooks: make(map[string]int),
	}
}

func (plugin *Plugin) AddHook(name string, handler interface{}) {
	switch h := handler.(type) {
		case func(interface{})interface{}:
			if len(hooks[name]) == 0 {
				var hookSlice []func(interface{})interface{}
				hookSlice = append(hookSlice, h)
				hooks[name] = hookSlice
			} else {
				hooks[name] = append(hooks[name], h)
			}
			plugin.Hooks[name] = len(hooks[name])
		case func(string)string:
			if len(sshooks[name]) == 0 {
				var hookSlice []func(string)string
				hookSlice = append(hookSlice, h)
				sshooks[name] = hookSlice
			} else {
				sshooks[name] = append(sshooks[name], h)
			}
			plugin.Hooks[name] = len(sshooks[name])
		case func(http.ResponseWriter, *http.Request, *User, interface{})bool:
			if len(pre_render_hooks[name]) == 0 {
				var hookSlice []func(http.ResponseWriter, *http.Request, *User, interface{})bool
				hookSlice = append(hookSlice, h)
				pre_render_hooks[name] = hookSlice
			} else {
				pre_render_hooks[name] = append(pre_render_hooks[name], h)
			}
			plugin.Hooks[name] = len(pre_render_hooks[name])
		case func(...interface{}) interface{}:
			vhooks[name] = h
			plugin.Hooks[name] = 0
		default:
			panic("I don't recognise this kind of handler!") // Should this be an error for the plugin instead of a panic()?
	}
}

func (plugin *Plugin) RemoveHook(name string, handler interface{}) {
	switch handler.(type) {
		case func(interface{})interface{}:
			key := plugin.Hooks[name]
			hook := hooks[name]
			if len(hook) == 1 {
				hook = []func(interface{})interface{}{}
			} else {
				hook = append(hook[:key], hook[key + 1:]...)
			}
			hooks[name] = hook
		case func(string)string:
			key := plugin.Hooks[name]
			hook := sshooks[name]
			if len(hook) == 1 {
				hook = []func(string)string{}
			} else {
				hook = append(hook[:key], hook[key + 1:]...)
			}
			sshooks[name] = hook
		case func(http.ResponseWriter, *http.Request, *User, interface{})bool:
			key := plugin.Hooks[name]
			hook := pre_render_hooks[name]
			if len(hook) == 1 {
				hook = []func(http.ResponseWriter, *http.Request, *User, interface{})bool{}
			} else {
				hook = append(hook[:key], hook[key + 1:]...)
			}
			pre_render_hooks[name] = hook
		case func(...interface{}) interface{}:
			delete(vhooks, name)
		default:
			panic("I don't recognise this kind of handler!") // Should this be an error for the plugin instead of a panic()?
	}
	delete(plugin.Hooks, name)
}

var plugins_inited bool = false
func init_plugins() {
	for name, body := range plugins {
		log.Print("Added plugin " + name)
		if body.Active {
			log.Print("Initialised plugin " + name)
			if plugins[name].Init != nil {
				err := plugins[name].Init()
				if err != nil {
					log.Print(err)
				}
			} else {
				log.Print("Plugin " + name + " doesn't have an initialiser.")
			}
		}
	}
	plugins_inited = true
}

func run_hook(name string, data interface{}) interface{} {
	for _, hook := range hooks[name] {
		data = hook(data)
	}
	return data
}

func run_hook_noreturn(name string, data interface{}) {
	for _, hook := range hooks[name] {
		_ = hook(data)
	}
}

func run_vhook(name string, data ...interface{}) interface{} {
	return vhooks[name](data...)
}

func run_vhook_noreturn(name string, data ...interface{}) {
	_ = vhooks[name](data...)
}

// Trying to get a teeny bit of type-safety where-ever possible, especially for such a critical set of hooks
func run_sshook(name string, data string) string {
	for _, hook := range sshooks[name] {
		data = hook(data)
	}
	return data
}

func run_pre_render_hook(name string, w http.ResponseWriter, r *http.Request, user *User, data interface{}) (halt bool) {
	// This hook runs on ALL pre_render hooks
	for _, hook := range pre_render_hooks["pre_render"] {
		if hook(w,r,user,data) {
			return true
		}
	}

	// The actual pre_render hook
	for _, hook := range pre_render_hooks[name] {
		if hook(w,r,user,data) {
			return true
		}
	}
	return false
}
