package soundcloud

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/fatih/color"
)

// Todo: Edge Cases
func TestGetArtwork(t *testing.T) {
	tables := []struct {
		url      string
		expected string
	}{
		{"https://soundcloud.com/uiceheidd/tell-me-you-love-me", "https://i1.sndcdn.com/artworks-DuzeporxapxyfgpP-PkGOzQ-t500x500.jpg"},
		{"https://soundcloud.com/uiceheidd/righteous", "https://i1.sndcdn.com/artworks-aoE3z9jDcBF2yf5p-f62P6Q-t500x500.jpg"},
	}
	for _, table := range tables {
		// request to user inputed SoundCloud URL
		resp, err := http.Get(table.url)
		if err != nil {
			log.Fatalln(err)
		}
		// response
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		result, _ := GetArtwork(body)
		expectedresult := table.expected
		if result != expectedresult {
			t.Error()
			red := color.New(color.FgRed).SprintFunc()
			fmt.Printf("%s GetTitle Failed: %s , expected %s got %s \n", red("[-]"), table.url, expectedresult, result)
		} else {
			green := color.New(color.FgGreen).SprintFunc()
			fmt.Printf("%s Passing: %s \n", green("[+]"), table.url)
		}
	}
}
