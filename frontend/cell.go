package frontend
import (
	"fmt"
	"appengine"
	"appengine/delay"
	"time"
)

var cellCallFunc = delay.Func("cellCall", cellCall)
func cellCall(c appengine.Context,col int,row int) {
	time.Sleep(time.Millisecond*500)
	CallModule(fmt.Sprintf("cell%dx%d",col,row),c)
}
