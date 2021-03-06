package main

func init() {
	plugins["helloworld"] = NewPlugin("helloworld","Hello World","Azareal","http://github.com/Azareal","","","",init_helloworld,nil,deactivate_helloworld,nil,nil)
}

// init_helloworld is separate from init() as we don't want the plugin to run if the plugin is disabled
func init_helloworld() error {
	plugins["helloworld"].AddHook("rrow_assign", helloworld_reply)
	return nil
}

func deactivate_helloworld() {
	plugins["helloworld"].RemoveHook("rrow_assign", helloworld_reply)
}

func helloworld_reply(data interface{}) interface{} {
	reply := data.(*Reply)
	reply.Content = "Hello World!"
	reply.ContentHtml = "Hello World!"
	reply.Tag = "Auto"
	return nil
}
