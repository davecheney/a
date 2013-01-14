package b

import "github.com/davecheney/a"

func B() bool {
	if a.A() {
		return true
	} 
	return false
}	
