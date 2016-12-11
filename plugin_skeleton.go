package main

func init() {
	/*
		The UName field should match the name in the URL minus plugin_ and the file extension. The same name as the map index. Please choose a unique name which won't clash with any other plugins.
		The Name field is for the friendly name of the plugin shown to the end-user.
		The Author field is the author of this plugin. The one who created it.
		The URL field is for the URL pointing to the location where you can download this plugin.
		The Settings field points to the route for managing the settings for this plugin. Coming soon.
		The Active field should always be set to false in the init() function of a plugin. It's used internally by the software to determine whether an admin has enabled a plugin or not and whether to run it. This will be overwritten by the user's preference.
		The Type field is for the type of the plugin. This gets changed to "go" automatically and we would suggest leaving "".
		The Init field is for the initialisation handler which is called by the software to run this plugin. This expects a function. You should add your hooks, init logic, initial queries, etc. in said function.
	*/
	plugins["skeleton"] = Plugin{"skeleton","Skeleton","Azareal","","",false,"",init_test}
}

func init_test() {}