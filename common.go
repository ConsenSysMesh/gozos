package main

func mkdir(dir string) (err error){
	return call(".","mkdir", "-p", dir)
}

