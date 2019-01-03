package zabbix

import (
	"reflect"
	"testing"
)

const (
	testUsergroupName = "testUsergroupName"
)

func TestUsergroupCRUD(t *testing.T) {

	var z Zabbix

	/* Login */
	loginTest(&z, t)
	defer logoutTest(&z, t)

	/* Create and delete */
	ugCreatedIDs := testUsergroupCreate(t, z)
	defer testUsergroupDelete(t, z, ugCreatedIDs)

	/* Get */
	testUsergroupGet(t, z, ugCreatedIDs)
}

func testUsergroupCreate(t *testing.T, z Zabbix) []string {

	ugCreatedIDs, _, err := z.UsergroupCreate([]UsergroupObject{
		{
			Name: testUsergroupName,
		},
	})
	if err != nil {
		t.Fatal("Usergroup create error:", err)
	}

	if len(ugCreatedIDs) == 0 {
		t.Fatal("Usergroup create error: empty IDs array")
	}

	t.Logf("Usergroup create: success")

	return ugCreatedIDs
}

func testUsergroupDelete(t *testing.T, z Zabbix, ugCreatedIDs []string) []string {

	ugDeletedIDs, _, err := z.UsergroupDelete(ugCreatedIDs)
	if err != nil {
		t.Fatal("Usergroup delete error:", err)
	}

	if len(ugDeletedIDs) == 0 {
		t.Fatal("Usergroup delete error: empty IDs array")
	}

	if reflect.DeepEqual(ugDeletedIDs, ugCreatedIDs) == false {
		t.Fatal("Usergroup delete error: IDs arrays for created and deleted usergroup are mismatch")
	}

	t.Logf("Usergroup delete: success")

	return ugDeletedIDs
}

func testUsergroupGet(t *testing.T, z Zabbix, ugCreatedIDs []string) []UsergroupObject {

	ugObjects, _, err := z.UsergroupGet(UsergroupGetParams{
		UsrgrpIDs:    ugCreatedIDs,
		SelectRights: SelectExtendedOutput,
		SelectUsers:  SelectExtendedOutput,
		GetParameters: GetParameters{
			Filter: map[string]interface{}{
				"name": testUsergroupName,
			},
			Output: SelectExtendedOutput,
		},
	})

	if err != nil {
		t.Error("Usergroup get error:", err)
	} else {
		if len(ugObjects) == 0 {
			t.Error("Usergroup get error: unable to find created usergroup")
		} else {
			t.Logf("Usergroup get: success")
		}
	}

	return ugObjects
}