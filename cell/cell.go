package cell

import (
    "fmt"
    "net/http"
    "appengine"
    "appengine/delay"
    "appengine/datastore"
    "time"
)

func init() {
    http.HandleFunc("/call", handler)
    http.HandleFunc("/_ah/stop", handlerStop)
}
var cellCallFunc = delay.Func("cellCall", cellCall)
type Status struct {
    Value int
}
func handler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    name := appengine.ModuleName(c)

    fmt.Fprintf(w, "Hello from cell %s",appengine.ModuleName(c))
    col := int(name[4]-'0');
    row := int(name[6]-'0');
    if row<3 {
        cellCallFunc.Call(c,col,row+1);
    }
    key := datastore.NewKey(c, "cell", appengine.ModuleName(c), 0, nil)
    _,err := datastore.Put(c,key,&Status{Value:2});
    if err!=nil {
        c.Warningf("Error saving:%v",err.Error())
    }
    time.Sleep(time.Second)
    _,err = datastore.Put(c,key,&Status{Value:1});
    if err!=nil {
        c.Warningf("Error saving:%v",err.Error())
    }
}

func cellCall(c appengine.Context,col int,row int) {
    time.Sleep(time.Millisecond*500)
    CallModule(fmt.Sprintf("cell%dx%d",col,row),c)
    //CallModule(fmt.Sprintf("cell%dx%d",col,row),c)
}


func handlerStop(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    key := datastore.NewKey(c, "cell", appengine.ModuleName(c), 0, nil)
    _,err := datastore.Put(c,key,&Status{Value:0});
    if err!=nil {
        c.Warningf("Error saving:%v",err.Error())
    }
    c.Infof("---Stopping");
}
