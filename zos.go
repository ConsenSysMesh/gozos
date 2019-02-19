package main



type Zos interface {
	Init(projectName string, opts ...InitOption) error
}

type zosWrapper struct {

}

func NewZos() Zos{
	return &zosWrapper{

	}
}

func (zos zosWrapper) Init(projectName string, opts ...InitOption) error {
	return NewInitCmd(projectName, opts...).Exec()
}

