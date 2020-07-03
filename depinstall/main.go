package depinstall

import (
	"github.com/munchkinhalfling/npm-clone/executil"
	"github.com/munchkinhalfling/npm-clone/prompt"
)

// Prompt user to install NPM dependencies
func InstallDeps(packageName string) {
	var questionRes = prompt.YesNo("Install npm dependencies now?", 'y')
	if questionRes {
		actuallyInstall(packageName)
	}
}

func actuallyInstall(packageName string) {
	executil.Run("cd", packageName)
	executil.Run("npm", "install")
	executil.Run("cd", "..")
}
