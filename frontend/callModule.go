package frontend
import (
	"appengine"
	"io/ioutil"
	"appengine/urlfetch"
)

func CallModule(module string,c appengine.Context) (string,error) {
	hostname, err := appengine.ModuleHostname(c,module, "", "")
	if err != nil {
		return "",err
	}
	url := "http://" + hostname + "/call"
	client := urlfetch.Client(c)
	resp, err := client.Get(url)
	if err != nil {
		return "",err
	}
	//fmt.Fprintf(w, "HTTP GET returned status %v", resp.Status)
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "",err
	}
	return string(contents),nil


}