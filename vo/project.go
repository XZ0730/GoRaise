package vo

import (
	"Raising/model"
)

type Project struct {
	Pname      string   `json:"pname"`
	Info       string   `json:"info"`
	ImgUrl     DataList `json:"imgurl"`
	Pid        string   `json:"pid"`
	Accumulate float64  `json:"accumulate"`
	Ispass     bool     `json:"ispass"`
}
type Project_Img struct {
	Id       uint   `json:"id"`
	Img_url  string `json:"imgurl"`
	Img_name string `json:"imagename"`
	Pid      string `json:"pid"`
}
type Audit struct {
	Pname  string `json:"pname"`
	Info   string `json:"info"`
	Pid    string `json:"pid"`
	Uid    uint   `json:"uid"`
	IsPass bool   `json:"ispass"`
}

func BuildProject(project *model.Project, imgUrl interface{}, total uint) *Project {

	return &Project{
		Pid:   project.Pid,
		Pname: project.PName,
		Info:  project.Info,
		ImgUrl: DataList{
			Item:  imgUrl,
			Total: total,
		},
		Accumulate: project.Accumulate,
		Ispass:     false,
	}
}

//给用户看的简单项目列表

//只对用户自己开放的查看自己审核中的项目  每个用户可以查看自己的审核项目
