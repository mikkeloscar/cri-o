package main

import (
	"os/exec"
	"os/user"
	"testing"

	"flag"

	"github.com/containers/storage"
	"github.com/urfave/cli"
)

func TestGetStore(t *testing.T) {
	// Make sure the tests are running as root
	failTestIfNotRoot(t)

	set := flag.NewFlagSet("test", 0)
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("root", "", "path to the root directory in which data, including images,  is stored")
	globalCtx := cli.NewContext(nil, globalSet, nil)
	command := cli.Command{Name: "imagesCommand"}
	c := cli.NewContext(nil, set, globalCtx)
	c.Command = command

	_, err := getStore(c)
	if err != nil {
		t.Error(err)
	}
}

func failTestIfNotRoot(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		t.Log("Could not determine user.  Running without root may cause tests to fail")
	} else if u.Uid != "0" {
		t.Fatal("tests will fail unless run as root")
	}
}

func getStoreForTests() (storage.Store, error) {
	set := flag.NewFlagSet("test", 0)
	globalSet := flag.NewFlagSet("test", 0)
	globalSet.String("root", "", "path to the root directory in which data, including images,  is stored")
	globalCtx := cli.NewContext(nil, globalSet, nil)
	command := cli.Command{Name: "testCommand"}
	c := cli.NewContext(nil, set, globalCtx)
	c.Command = command

	return getStore(c)
}

func pullTestImage(name string) error {
	cmd := exec.Command("crioctl", "image", "pull", name)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
