package api

var Post = new(postApi)

type postApi struct {
}

//发帖
func (a *postApi) Create() {
}

//删帖
func (a *postApi) Delete() {
}

//赞
func (a *postApi) UpVote() {
}

//踩
func (a *postApi) DownVote() {
}
