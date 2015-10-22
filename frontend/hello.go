package frontend

import (
    "fmt"
    "net/http"
    "appengine"

    "appengine/datastore"
    "strings"
)

func init() {
    http.HandleFunc("/status", handlerStatus)
    http.HandleFunc("/call/", handler)
}


func handler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    var parts = strings.Split(r.URL.Path,"/");
    cX := int(parts[len(parts)-1][0]-'0');
    fmt.Fprintf(w,"Calling %d\n",cX);
    result,err := CallModule(fmt.Sprintf("cell%dx0",cX),c);
    if err!=nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return;
    }
    fmt.Fprintf(w, "Result:%v",result)

}
type Status struct  { Value int }
func handlerStatus(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    q := datastore.NewQuery("cell");
    var res []Status;
    keys,err := q.GetAll(c,&res);
    if err!=nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return;
    }
    fmt.Fprint(w, "[")
    for i,k := range(keys) {
        //if res[i].Value==1 {
        if i!=0 && i<len(keys) {
            fmt.Fprint(w,",")
        }
        fmt.Fprintf(w, `{"k":"%v","v":"%v"}`, k.String(), res[i].Value)
        //}
    }
    fmt.Fprint(w, "]")
}
