package main

func mkdir(dir string, cd bool) (err error){
	err = call("mkdir", "-p", dir)
	if err != nil{
		return
	}
	if cd{
		err = call("cd", dir)
	}
	return
}
