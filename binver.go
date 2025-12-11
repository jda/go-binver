package binver

import (
	"runtime/debug"
)

var (
	NO_VERSION = "NO_VERSION"
)

func CanonVersion() string {
	version := ""
	commit := ""
	modified := false

	info, ok := debug.ReadBuildInfo()  
    if !ok {  
       return NO_VERSION
    }  

    for _, setting := range info.Settings {  
       switch setting.Key {  
       case "vcs.revision":  
          commit = setting.Value  
       case "vcs.modified":  
          modified = true  
       }  
    }

    if commit != "" {
    	version = commit
    } else {
    	return NO_VERSION
    }

    if modified {  
       version += "-tainted"  
    }

    return version
}