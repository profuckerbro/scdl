package cmd

import (
	"github.com/profuckerbro/scdl/soundcloud"
)

func scdl(args []string) {
	url := args[0]

	if Find == true {

		soundcloud.Search(url)
		// exit
		return
	}
	// download song
	soundcloud.ExtractSong(url)

}
