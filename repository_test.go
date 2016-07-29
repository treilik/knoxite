/*
 * knoxite
 *     Copyright (c) 2016, Christian Muehlhaeuser <muesli@gmail.com>
 *
 *   For license see LICENSE.txt
 */

package knoxite

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateRepository(t *testing.T) {
	testPassword := "this_is_a_password"

	dir, err := ioutil.TempDir("", "knoxite")
	if err != nil {
		t.Errorf("Failed creating temporary dir for repository: %s", err)
		return
	}
	defer os.RemoveAll(dir)

	_, err = NewRepository(dir, testPassword)
	if err != nil {
		t.Errorf("Failed creating repository: %s", err)
	}

	_, err = OpenRepository(dir, testPassword)
	if err != nil {
		t.Errorf("Failed opening repository: %s", err)
	}
}

func TestCreateRepositoryError(t *testing.T) {
	testPassword := "this_is_a_password"

	_, err := NewRepository("invalidprotocol://foo", testPassword)
	if err == nil {
		t.Errorf("Creating repository should have failed, but didn't")
	}
}
