package maps

import "fmt"

/* Notes MAPS */
// - Maps are more akin to JS-Obj's (more so than structs?)
//		- It seems that structs were more closer to types that can have functions
// 		- In a way you can implement structs/interfaces using maps???
//		- You can't dynamically add a field after declaration for structs
// - Created using the make()  function

// - Structs access using `.` ex. user.Name
// - Maps access using `[key]` ex. ages["personA"]

func how2CreateAMap() {
	// key types are typically string
	//		structs can be used for keys, but I currently feel like I won't use this
	// value types are generally less restrictive

	// using make()
	fightingCharacterA := make(map[string]string)
	fightingCharacterA["lightAttack"] = "jab"
	fightingCharacterA["heavyAttack"] = "upper cut"
	fightingCharacterA["ultimate"] = "Super Duper Kaboom"

	fmt.Println("map using make(): ", fightingCharacterA)

	// Using inline
	fightingCharacterB := map[string]string{
		"lightAttack": "tap",
		"heavyAttack": "shove",
		"ultimate":    "Absolute Mega Best",
	}
	fmt.Println("map using inline: ", fightingCharacterB)

	// Map with struct
	type characterEquipment struct {
		rightHand string
		leftHand  string
	}

	fightingCharacterC := map[string]characterEquipment{
		"lightAttack": {
			rightHand: "fist",
			leftHand:  "fist",
		},
		"heavyAttack": {
			rightHand: "knife",
			leftHand:  "knife",
		},
		"ultimate": {
			rightHand: "energy",
			leftHand:  "energy",
		},
	}
	fmt.Println("map using struct as type: ", fightingCharacterC)

	fmt.Println("get length of keys of a map len(fightingCharacterB): ", len(fightingCharacterC))
}

func workingWithMaps() {
	someMap := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}

	fmt.Println("someMap contents", someMap)
	// Accessing a key-value pair (values can be assied through this)
	fmt.Println("access a key, someMap[key1]: ", someMap["key1"])

	delete(someMap, "key2")
	fmt.Println("delete using delete(someMap, key2): ", someMap)

	// check if a key exists
	_, ok1 := someMap["key1"]
	fmt.Println("check whether a key exists 'elem, ok := someMap[key1]': ", ok1)
	_, ok10 := someMap["key10"]
	fmt.Println("check whether a key exists 'elem, ok := someMap[key10]': ", ok10)
}

func nestedMaps() {
	nestedMap := map[string]map[string]string{
		"Tier1Key": {
			"Tier2Key": "nestedValue",
		},
	}
	fmt.Println("nestedMap example: ", nestedMap)
}

func MapsMain() {
	fmt.Println("===== 10_maps =====")

	how2CreateAMap()
	workingWithMaps()
	nestedMaps()
}
