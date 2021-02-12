package api

var Reply = new(replyApi)

type replyApi struct {
}

func (a *replyApi) Create()   {}
func (a *replyApi) Delete()   {}
func (a *replyApi) UpVote()   {}
func (a *replyApi) DownVote() {}
