package main

type Zos interface {
	Init(projectName string, opts ...InitOption) error
}

type zos struct {

}

func NewZos() Zos{
	return &zos{

	}
}

func (zos zos) Init(projectName string, opts ...InitOption) error {
	return NewInitCmd(projectName, opts...).Exec()
}

