package main


var(
	DefaultInitOpts = InitOptions{
		Args: []string{
			initCmd,
		},
	}
)

type InitOptions struct {
	Args []string
}

type InitOption func(*InitOptions)

func InitVersion(version string) InitOption{
	return func(options *InitOptions) {
		options.Args = append(options.Args, version)
	}
}

type cmd struct {
	projectName string
	opts InitOptions
}


func NewInitCmd(projectName string, setters ...InitOption) Cmd{
	opts := &DefaultInitOpts
	opts.Args = append(opts.Args, projectName)
	for _, setter := range setters{
		setter(opts)
	}
	return cmd{
		projectName: projectName,
		opts: *opts,
	}
}

func (cmd cmd) Exec() (err error) {
	err = mkdir(cmd.projectName)
	if err != nil{
		return
	}
	err = call(cmd.projectName, zos, cmd.opts.Args...)
	if err != nil{
		return
	}
	return
}
