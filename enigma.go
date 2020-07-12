package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"enigma/roter"
	"enigma/util"
)

func main() {
	//initialize roters randomly
	line1 := make([]int, 0, len(util.RoterSeedMap))
	line2 := make([]int, 0, len(util.RoterSeedMap))
	line3 := make([]int, 0, len(util.RoterSeedMap))
	for _, element := range util.RoterSeedMap {
		line1 = append(line1, element)
	}
	for _, element := range util.RoterSeedMap {
		line2 = append(line1, element)
	}
	for _, element := range util.RoterSeedMap {
		line3 = append(line1, element)
	}
	//show lines of roters
	/*log.Println("line1: ", line1)
	log.Println("line2: ", line2)
	log.Println("line3: ", line3)*/
	roter1 := roter.Initialize(line1)
	roter2 := roter.Initialize(line2)
	roter3 := roter.Initialize(line3)

	if len(os.Args) != 2 {
		log.Fatal("Invalid args")
	}
	input_strs := strings.Split(os.Args[1], "")

	output_str := ""
	for _, char := range input_strs {

		//get encrypted char
		char_num := util.StringConvertMap[char]
		encrypted_num := roter3.Values[roter2.Values[roter1.Values[char_num]]]

		//apply a reflector
		reflected_encrypted_num := roter.Reflector[encrypted_num]

		//encrypt backwards
		for index, value := range roter3.Values {
			if value == reflected_encrypted_num {
				reflected_encrypted_num = index
				break
			}
		}
		for index, value := range roter2.Values {
			if value == reflected_encrypted_num {
				reflected_encrypted_num = index
				break
			}
		}
		for index, value := range roter1.Values {
			if value == reflected_encrypted_num {
				reflected_encrypted_num = index
				break
			}
		}

		//get char
		encrypted_char := ""
		for key, value := range util.StringConvertMap {
			if value == reflected_encrypted_num {
				encrypted_char = key
				break
			}
		}
		output_str += encrypted_char

		//rotate roters
		roter1.Plus()
		if roter1.IsOneRound {
			roter2.Plus()
			if roter2.IsOneRound {
				roter3.Plus()
			}
		}
	}
	fmt.Println(output_str)
}
