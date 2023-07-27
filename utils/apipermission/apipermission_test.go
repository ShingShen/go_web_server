package apipermission

import "testing"

func TestApiPermission(t *testing.T) {
	const name string = "authCheck"
	const errName string = "errName"

	apiPermission, err := ApiPermission(name)
	if err == nil {
		t.Logf("apiPermission passed: %v, %v", apiPermission, err)
	} else {
		t.Errorf("apiPermission failed: %v, %v", apiPermission, err)
	}

	apiPermissionNameErr, err := ApiPermission(errName)
	if err != nil {
		t.Logf("apiPermissionNameErr passed: %v, %v", apiPermissionNameErr, err)
	} else {
		t.Errorf("apiPermissionNameErr failed: %v, %v", apiPermissionNameErr, err)
	}
}
