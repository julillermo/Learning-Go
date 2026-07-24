package structs

import (
	"fmt"
	"math/rand"
	"reflect"
)

// Generally similar to JS-Objects and Python-Dictionaries.

func basicStructDeclaration() {

	type store struct {
		classification string
		owners         []string
		location       string
		coffee         []string
		pastries       []string
		landArea       float64
	}

	houseSweet := store{
		classification: "cafe",
		owners:         []string{"CG", "DC"},
		location:       "doon",
		coffee:         []string{"Espresso", "Latte", "Cappuccino"},
		pastries:       []string{"Croissant", "Muffin", "Scone"},
		landArea:       7107.0,
	}

	fmt.Println("reflect.TypeOf(houseSweet):", reflect.TypeOf(houseSweet))
	fmt.Println("houseSweet:", houseSweet)

	fmt.Println("")

	// nested struct
	type buildingSpace struct {
		owners []string
		stores []store
		rent   float64
	}

	georgeBuilding := buildingSpace{
		stores: []store{houseSweet},
		rent:   123.0,
	}

	// Assigning values within a struct (similar to JS-Obj & Python-Dict)
	georgeBuilding.owners = []string{"Christopher Nolan"}

	fmt.Println("reflect.TypeOf(georgeBuilding):", reflect.TypeOf(georgeBuilding))
	fmt.Println("georgeBuilding:", georgeBuilding)
}

func anonymousStructs() {
	someMusicAlbum := struct {
		producers        []string
		artists          []string
		genre            string
		parentalAdvisory bool
	}{
		producers:        []string{"McTypeMaster", "KeyboardSquire"},
		artists:          []string{"beep boop beep"},
		genre:            "techno",
		parentalAdvisory: false,
	}

	fmt.Println("reflect.TypeOf(anonymouseStruct):", reflect.TypeOf(someMusicAlbum))
	fmt.Println("anonymouseStruct someMusicAlbum:", someMusicAlbum)
}

func embeddedStructs() {
	// Difference from nested structs
	// - accessible at top level (kind of like class inheretance in a way)

	type jetpack struct {
		fuelCapacity float64
		socket       string
	}

	type robot struct {
		jetpack
		pilot string
	}

	superRobot := robot{
		pilot: "yo papa",
		jetpack: jetpack{
			fuelCapacity: 100.0,
			socket:       "compatible",
		},
	}

	fmt.Println("")
	fmt.Println("Embedded structs")
	fmt.Println("reflect.TypeOf(superRobot):", reflect.TypeOf(superRobot))
	fmt.Println("superRobot.pilot", superRobot.pilot)
	fmt.Println("superRobot.fuelCapacity", superRobot.fuelCapacity)
	fmt.Println("superRobot.socket", superRobot.socket)
}

/* Struct Methods */
// Kind of like how JS-Objects can carry with it functions you can call.
// However, these methods are declared AFTER the struct
// Note as well that this can't be declared inside of a function
type randomNumbers struct {
	numbers []int
}

func (rn randomNumbers) pickRandomNumber() int {
	randomIndex := rand.Intn(len(rn.numbers))
	return rn.numbers[randomIndex]
}

/* Optimizing for memory layout */
// - It's more efficient to place similar dataypes (even by bit)
//		one after the other.
// - I think this is because Go assigns 1 byte when the dtype isn't the exact same
// - This is moslty not notcieable, but can be applied for efficiency

func emptyStruct() (emptyStruct, someEmptyStructVar struct{}) {
	// empty structs are worth 0 bytes

	// anonymous empty struct{}
	emptyStruct = struct{}{}
	// named empty struct{}
	type someEmptyStruct struct{}
	someEmptyStructVar = someEmptyStruct{}

	return
}

func StructsMain() {
	fmt.Println("===== 5_Structs =====")

	basicStructDeclaration()
	anonymousStructs()
	embeddedStructs()

	fmt.Println("")
	fmt.Println("struct methods")
	rdmNumList := randomNumbers{
		numbers: []int{100, 200, 300, 400, 500},
	}
	fmt.Println("rdmNumList: ", rdmNumList)
	fmt.Println("rdmNumList.pickRandomNumber(): ", rdmNumList.pickRandomNumber())

	emptyStruct()
}
