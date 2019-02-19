package main


var(
	DefaultInitOpts = InitOptions{

	}
)

type InitOptions struct {

}

type InitOption func(*InitOptions)



type cmd struct {
	projectName string
	opts InitOptions
}


func NewInitCmd(projectName string, setters ...InitOption) Cmd{
	opts := &DefaultInitOpts
	for _, setter := range setters{
		setter(opts)
	}
	return cmd{
		projectName: projectName,
		opts: *opts,
	}
}

func (cmd cmd) Exec() error {
	mkdir(cmd.projectName, true)
	return nil
}
