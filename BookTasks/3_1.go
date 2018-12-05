package main

import "fmt"

func main() {
	//makeAppend()
	//makeInt()
	//makeMapInt()
	//makeMapString()
	makeMapStringAgr()

}

func makeAppend() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func makeInt() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func makeMapInt() {
	var x map[string]int
	x = make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
}

func makeMapString() {
	//getElements := make(map[string]string)
	//getElements["H"] = "Hydrogen"
	//getElements["He"] = "Helium"
	//getElements["Li"] = "Lithium"
	//getElements["Be"] = "Beryllium"
	//getElements["B"] = "Boron"
	//getElements["C"] = "Carbon"
	//getElements["N"] = "Nitrogen"
	//getElements["O"] = "Oxygen"
	//getElements["F"] = "Fluorine"
	//getElements["Ne"] = "Neon"

	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	fmt.Println(elements["Li"])
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok)
	}
}

func makeMapStringAgr() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Ne"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
