package main

import "sync"
import "encoding/json"

var group_update_mutex sync.Mutex

type GroupAdmin struct
{
	ID int
	Name string
	Rank string
	RankClass string
	CanEdit bool
	CanDelete bool
}

type Group struct
{
	ID int
	Name string
	Is_Mod bool
	Is_Admin bool
	Is_Banned bool
	Tag string
	Perms Perms
	PermissionsText []byte
	PluginPerms map[string]bool // Custom permissions defined by plugins. What if two plugins declare the same permission, but they handle them in incompatible ways? Very unlikely, we probably don't need to worry about this, the plugin authors should be aware of each other to some extent
	PluginPermsText []byte
	Forums []ForumPerms
	CanSee []int // The IDs of the forums this group can see
}

var group_create_mutex sync.Mutex
func create_group(group_name string, tag string, is_admin bool, is_mod bool, is_banned bool) (int, error) {
	var gid int
	err := group_entry_exists_stmt.QueryRow().Scan(&gid)
	if err != nil && err != ErrNoRows {
		return 0, err
	}
	if err != ErrNoRows {
		group_update_mutex.Lock()
		_, err = update_group_rank_stmt.Exec(is_admin, is_mod, is_banned, gid)
		if err != nil {
			return gid, err
		}
		_, err = update_group_stmt.Exec(group_name, tag, gid)
		if err != nil {
			return gid, err
		}

		groups[gid].Name = group_name
		groups[gid].Tag = tag
		groups[gid].Is_Banned = is_banned
		groups[gid].Is_Mod = is_mod
		groups[gid].Is_Admin = is_admin

		group_update_mutex.Unlock()
		return gid, nil
	}

	group_create_mutex.Lock()
	var permstr string = "{}"
	res, err := create_group_stmt.Exec(group_name, tag, is_admin, is_mod, is_banned, permstr)
	if err != nil {
		return 0, err
	}

	gid64, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	gid = int(gid64)

	var perms Perms = BlankPerms
	var blankForums []ForumPerms
	var blankIntList []int
	var plugin_perms map[string]bool = make(map[string]bool)
	var plugin_perms_bytes []byte = []byte("{}")
	if vhooks["create_group_preappend"] != nil {
		run_vhook("create_group_preappend", &plugin_perms, &plugin_perms_bytes)
	}

	groups = append(groups, Group{gid,group_name,is_mod,is_admin,is_banned,tag,perms,[]byte(permstr),plugin_perms,plugin_perms_bytes,blankForums,blankIntList})
	group_create_mutex.Unlock()

	// Generate the forum permissions based on the presets...
	fdata, err := fstore.GetAll()
	if err != nil {
		return 0, err
	}

	permupdate_mutex.Lock()
	for _, forum := range fdata {
		var thePreset string
		if is_admin {
			thePreset = "admins"
		} else if is_mod {
			thePreset = "staff"
		} else if is_banned {
			thePreset = "banned"
		} else {
			thePreset = "members"
		}

		permmap := preset_to_permmap(forum.Preset)
		permitem := permmap[thePreset]
		permitem.Overrides = true
		permstr, err := json.Marshal(permitem)
		if err != nil {
			return gid, err
		}
		perms := string(permstr)
		_, err = add_forum_perms_to_group_stmt.Exec(gid,forum.ID,forum.Preset,perms)
		if err != nil {
			return gid, err
		}

		err = rebuild_forum_permissions(forum.ID)
		if err != nil {
			return gid, err
		}
	}
	permupdate_mutex.Unlock()
	return gid, nil
}

func group_exists(gid int) bool {
	return (gid <= groupCapCount) && (gid > 0) && groups[gid].Name != ""
}
