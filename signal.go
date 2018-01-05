package dbusUtil

import (
	"github.com/godbus/dbus"
	"encoding/json"
	"fmt"
	"time"
	"strconv"
	"github.com/ottojo/rclib"
)

func SignalToLogString(s *dbus.Signal) string {
	jsonBody, err := json.Marshal(s.Body)
	if err != nil {
		fmt.Println(err)
		jsonBody = []byte{}
	}
	return strconv.FormatInt(time.Now().UnixNano(), 10) + "," + s.Sender + "," + s.Name + "," + string(jsonBody) + "\n"
}

func ReceiveSignals(iface, path string, conn *dbus.Conn, c chan *dbus.Signal) {
	call := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='"+path+"',interface='"+iface+"'")
	if call.Err != nil {
		//TODO Error handling
		fmt.Println(call.Err)
	}
	conn.Signal(c)
}

func SendOnPackageReceived(p rclib.Package, conn *dbus.Conn, dbusInterface, dbusPath, dbusSignalName string) {
	jsonPackage, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	err = conn.Emit(dbus.ObjectPath(dbusPath), dbusInterface+"."+dbusSignalName, string(jsonPackage))
	if err != nil {
		fmt.Println(err)
	}
}
