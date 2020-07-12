package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"enigma/constant"
	"enigma/roter"
)

func main() {
	//choose mode
	var (
		kek           bool
		input         string
		roter_setting string
	)
	flag.BoolVar(&kek, "kek", false, "Key encrypting key(repeat twice).")
	flag.StringVar(&input, "input", "", "Type string to encrypt/decrypt.")
	flag.StringVar(&roter_setting, "setting", "", "Setting of roters.")
	flag.Parse()
	if input == "" {
		log.Fatal("No string to encrypt/decrypt")
	} else if kek {
		if len(input) != 6 {
			log.Fatal("Kek must be repetition of three letters(six letters)")
		}
	} else {
		if roter_setting == "" {
			log.Fatal("Specify the setting of roters")
		} else if len(roter_setting) != 3 {
			log.Fatal("Setting must be three letters")
		}
	}

	//initialize roters
	roter1 := roter.Initialize(constant.RoterPattern1)
	roter2 := roter.Initialize(constant.RoterPattern2)
	roter3 := roter.Initialize(constant.RoterPattern3)

	//change the setting of the roters
	if !kek {
		setting_letters := strings.Split(roter_setting, "")
		fmt.Println(setting_letters)
	}

	if len(os.Args) < 2 {
		log.Fatal("Invalid args")
	}
	input_strs := strings.Split(os.Args[1], "")

	output_str := ""
	for _, char := range input_strs {

		//get encrypted char
		char_num := constant.StringConvertMap[char]
		encrypted_num := roter3.Values[roter2.Values[roter1.Values[char_num]]]

		//apply a reflector
		reflected_encrypted_num := constant.Reflector[encrypted_num]

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
		for key, value := range constant.StringConvertMap {
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
