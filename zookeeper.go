package zookeeper

import (
	"fmt"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/samuel/go-zookeeper/zk"
)

var zkClient *zk.Conn

// Connect ...
func Connect(uri string) error {
	c, _, err := zk.Connect([]string{uri}, time.Second*30)
	if err != nil {
		fmt.Println("Error when connect to Zookeeper", uri, err)
		return err
	}

	fmt.Println(aurora.Green("*** CONNECTED TO ZOOKEEPER: " + uri))

	// Set client
	zkClient = c

	return nil
}




