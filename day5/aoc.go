package main

import "fmt"

/////
// DAY 5
/////

const (
	// ModeParam = Parameters input
	ModeParam = 0
	// ModeImmed = Immediate input
	ModeImmed = 1
)

func getArg(mode int, n int, a codeArray) int {
	switch mode {
	case ModeParam:
		return a[a[n]]
	case ModeImmed:
		return a[n]
	}
	return 0
}

func parseInstruction(a codeArray, i int) (pos, v1, v2 int) {
	v1Mode := (a[i] % 1000) / 100
	v2Mode := (a[i] % 10000) / 1000
	posMode := (a[i] % 100000) / 10000

	v1 = getArg(v1Mode, i+1, a)
	v2 = getArg(v2Mode, i+2, a)

	switch posMode {
	case ModeParam:
		pos = a[i+3]
	case ModeImmed:
		pos = i + 3
	}

	return

}

func intCode(a codeArray, input int) (output int) {
	var skip int

	for i := 0; i < len(a); i += skip {
		opCode := a[i] % 100

		// Output validation
		if output != 0 && opCode != 99 {
			panic(fmt.Sprintf("Wrong output: %d  Opcode: %d", output, opCode))
		}

		switch opCode {
		case 1: // ADD
			pos, v1, v2 := parseInstruction(a, i)
			a[pos] = v1 + v2
			skip = 4
		case 2: // MULTIPLY
			pos, v1, v2 := parseInstruction(a, i)
			a[pos] = v1 * v2
			skip = 4
		case 3: // INPUT
			pos := a[i+1]
			a[pos] = input
			skip = 2
		case 4: // OUTPUT
			v1Mode := (a[i] % 1000) / 100
			output = getArg(v1Mode, i+1, a)
			skip = 2
		case 5: // JUMP-IF-TRUE
			_, v1, v2 := parseInstruction(a, i)
			if v1 != 0 {
				i = v2
				skip = 0
			} else {
				skip = 3
			}
		case 6: // JUMP-IF-FALSE
			_, v1, v2 := parseInstruction(a, i)
			if v1 == 0 {
				i = v2
				skip = 0
			} else {
				skip = 3
			}
		case 7: // LESS THAN
			pos, v1, v2 := parseInstruction(a, i)
			if v1 < v2 {
				a[pos] = 1
			} else {
				a[pos] = 0
			}
			skip = 4
		case 8: // EQUALS
			pos, v1, v2 := parseInstruction(a, i)
			if v1 == v2 {
				a[pos] = 1
			} else {
				a[pos] = 0
			}
			skip = 4
		case 99:
			return output
		default:
			panic(opCode)
		}
	}
	return
}

type codeArray []int

func main() {
	baseArray := codeArray{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1101, 91, 67, 225, 1102, 67, 36, 225, 1102, 21, 90, 225, 2, 13, 48, 224, 101, -819, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 7, 224, 224, 1, 223, 224, 223, 1101, 62, 9, 225, 1, 139, 22, 224, 101, -166, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 3, 224, 224, 1, 223, 224, 223, 102, 41, 195, 224, 101, -2870, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 1, 224, 224, 1, 224, 223, 223, 1101, 46, 60, 224, 101, -106, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 2, 224, 1, 224, 223, 223, 1001, 191, 32, 224, 101, -87, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 1, 224, 1, 223, 224, 223, 1101, 76, 90, 225, 1101, 15, 58, 225, 1102, 45, 42, 224, 101, -1890, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 101, 62, 143, 224, 101, -77, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 4, 224, 1, 224, 223, 223, 1101, 55, 54, 225, 1102, 70, 58, 225, 1002, 17, 80, 224, 101, -5360, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 3, 224, 1, 223, 224, 223, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1008, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 329, 1001, 223, 1, 223, 1108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 344, 101, 1, 223, 223, 107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 359, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 374, 1001, 223, 1, 223, 108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 389, 101, 1, 223, 223, 7, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 404, 1001, 223, 1, 223, 1108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 419, 101, 1, 223, 223, 1008, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 434, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 449, 1001, 223, 1, 223, 1007, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 464, 1001, 223, 1, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 479, 101, 1, 223, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 494, 1001, 223, 1, 223, 8, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 509, 101, 1, 223, 223, 1107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 524, 1001, 223, 1, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 539, 101, 1, 223, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 554, 101, 1, 223, 223, 1007, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 569, 101, 1, 223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 584, 101, 1, 223, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 599, 1001, 223, 1, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 614, 101, 1, 223, 223, 7, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 629, 1001, 223, 1, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 644, 101, 1, 223, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 659, 1001, 223, 1, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}

	a := make(codeArray, len(baseArray))
	b := make(codeArray, len(baseArray))
	copy(a, baseArray)
	copy(b, baseArray)

	fmt.Printf("The awnser to question one is: %d\n", intCode(a, 1))
	fmt.Printf("The awnser to question one is: %d\n", intCode(b, 5))
}
