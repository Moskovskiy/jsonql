// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 11: // ['\v','\v']
			return 1
		case r == 12: // ['\f','\f']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 33: // ['!','!']
			return 2
		case r == 34: // ['"','"']
			return 3
		case r == 37: // ['%','%']
			return 4
		case r == 38: // ['&','&']
			return 5
		case r == 39: // [''',''']
			return 6
		case r == 42: // ['*','*']
			return 7
		case r == 43: // ['+','+']
			return 8
		case r == 45: // ['-','-']
			return 9
		case r == 46: // ['.','.']
			return 10
		case r == 47: // ['/','/']
			return 11
		case r == 48: // ['0','0']
			return 12
		case 49 <= r && r <= 57: // ['1','9']
			return 13
		case r == 60: // ['<','<']
			return 14
		case r == 61: // ['=','=']
			return 15
		case r == 62: // ['>','>']
			return 16
		case 65 <= r && r <= 90: // ['A','Z']
			return 17
		case r == 91: // ['[','[']
			return 18
		case r == 93: // [']',']']
			return 19
		case r == 94: // ['^','^']
			return 20
		case r == 95: // ['_','_']
			return 21
		case 97 <= r && r <= 99: // ['a','c']
			return 17
		case r == 100: // ['d','d']
			return 22
		case r == 101: // ['e','e']
			return 17
		case r == 102: // ['f','f']
			return 23
		case 103 <= r && r <= 104: // ['g','h']
			return 17
		case r == 105: // ['i','i']
			return 24
		case 106 <= r && r <= 109: // ['j','m']
			return 17
		case r == 110: // ['n','n']
			return 25
		case 111 <= r && r <= 115: // ['o','s']
			return 17
		case r == 116: // ['t','t']
			return 26
		case 117 <= r && r <= 122: // ['u','z']
			return 17
		case r == 124: // ['|','|']
			return 27
		case r == 126: // ['~','~']
			return 28
		case r == 160: // [\u00a0,\u00a0]
			return 1
		case 8192 <= r && r <= 8202: // [\u2000,\u200a]
			return 1
		case r == 8239: // [\u202f,\u202f]
			return 1
		case r == 8287: // [\u205f,\u205f]
			return 1
		case r == 12288: // [\u3000,\u3000]
			return 1
		case r == 65279: // [\ufeff,\ufeff]
			return 1
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 29
		case r == 126: // ['~','~']
			return 30
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		case r == 92: // ['\','\']
			return 32
		default:
			return 3
		}
	},
	// S4
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		case r == 38: // ['&','&']
			return 33
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 34
		case r == 92: // ['\','\']
			return 35
		default:
			return 6
		}
	},
	// S7
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 36
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 37
		case 48 <= r && r <= 55: // ['0','7']
			return 38
		case 56 <= r && r <= 57: // ['8','9']
			return 39
		case r == 69: // ['E','E']
			return 40
		case r == 88: // ['X','X']
			return 41
		case r == 101: // ['e','e']
			return 40
		case r == 120: // ['x','x']
			return 41
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 37
		case 48 <= r && r <= 57: // ['0','9']
			return 13
		case r == 69: // ['E','E']
			return 40
		case r == 101: // ['e','e']
			return 40
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 42
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 43
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 44
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 122: // ['a','z']
			return 46
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 122: // ['a','z']
			return 46
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 100: // ['a','d']
			return 46
		case r == 101: // ['e','e']
			return 48
		case 102 <= r && r <= 122: // ['f','z']
			return 46
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case r == 97: // ['a','a']
			return 49
		case 98 <= r && r <= 122: // ['b','z']
			return 46
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 114: // ['a','r']
			return 46
		case r == 115: // ['s','s']
			return 50
		case 116 <= r && r <= 122: // ['t','z']
			return 46
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 110: // ['a','n']
			return 46
		case r == 111: // ['o','o']
			return 51
		case 112 <= r && r <= 116: // ['p','t']
			return 46
		case r == 117: // ['u','u']
			return 52
		case 118 <= r && r <= 122: // ['v','z']
			return 46
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 113: // ['a','q']
			return 46
		case r == 114: // ['r','r']
			return 53
		case 115 <= r && r <= 122: // ['s','z']
			return 46
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case r == 124: // ['|','|']
			return 54
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 55
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 56
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 3
		case r == 92: // ['\','\']
			return 57
		case r == 98: // ['b','b']
			return 57
		case r == 102: // ['f','f']
			return 57
		case r == 110: // ['n','n']
			return 57
		case r == 114: // ['r','r']
			return 57
		case r == 116: // ['t','t']
			return 57
		case r == 117: // ['u','u']
			return 58
		case r == 118: // ['v','v']
			return 57
		case r == 120: // ['x','x']
			return 59
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 6
		case r == 92: // ['\','\']
			return 60
		case r == 98: // ['b','b']
			return 60
		case r == 102: // ['f','f']
			return 60
		case r == 110: // ['n','n']
			return 60
		case r == 114: // ['r','r']
			return 60
		case r == 116: // ['t','t']
			return 60
		case r == 117: // ['u','u']
			return 61
		case r == 118: // ['v','v']
			return 60
		case r == 120: // ['x','x']
			return 62
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 36
		case r == 69: // ['E','E']
			return 63
		case r == 101: // ['e','e']
			return 63
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 64
		case r == 69: // ['E','E']
			return 65
		case r == 101: // ['e','e']
			return 65
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 37
		case 48 <= r && r <= 55: // ['0','7']
			return 38
		case 56 <= r && r <= 57: // ['8','9']
			return 39
		case r == 69: // ['E','E']
			return 40
		case r == 101: // ['e','e']
			return 40
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 37
		case 48 <= r && r <= 57: // ['0','9']
			return 39
		case r == 69: // ['E','E']
			return 40
		case r == 101: // ['e','e']
			return 40
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 66
		case r == 45: // ['-','-']
			return 66
		case 48 <= r && r <= 57: // ['0','9']
			return 67
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 68
		case 65 <= r && r <= 70: // ['A','F']
			return 69
		case 97 <= r && r <= 102: // ['a','f']
			return 69
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S43
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S44
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S45
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 122: // ['a','z']
			return 46
		}
		return NoState
	},
	// S46
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 122: // ['a','z']
			return 46
		}
		return NoState
	},
	// S47
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 122: // ['a','z']
			return 46
		}
		return NoState
	},
	// S48
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 101: // ['a','e']
			return 46
		case r == 102: // ['f','f']
			return 70
		case 103 <= r && r <= 122: // ['g','z']
			return 46
		}
		return NoState
	},
	// S49
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 107: // ['a','k']
			return 46
		case r == 108: // ['l','l']
			return 71
		case 109 <= r && r <= 122: // ['m','z']
			return 46
		}
		return NoState
	},
	// S50
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 109: // ['a','m']
			return 46
		case r == 110: // ['n','n']
			return 72
		case 111 <= r && r <= 122: // ['o','z']
			return 46
		}
		return NoState
	},
	// S51
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 115: // ['a','s']
			return 46
		case r == 116: // ['t','t']
			return 73
		case 117 <= r && r <= 122: // ['u','z']
			return 46
		}
		return NoState
	},
	// S52
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 107: // ['a','k']
			return 46
		case r == 108: // ['l','l']
			return 74
		case 109 <= r && r <= 122: // ['m','z']
			return 46
		}
		return NoState
	},
	// S53
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 116: // ['a','t']
			return 46
		case r == 117: // ['u','u']
			return 75
		case 118 <= r && r <= 122: // ['v','z']
			return 46
		}
		return NoState
	},
	// S54
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S55
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S56
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S57
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		case r == 92: // ['\','\']
			return 32
		default:
			return 3
		}
	},
	// S58
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 76
		case 65 <= r && r <= 70: // ['A','F']
			return 77
		case 97 <= r && r <= 102: // ['a','f']
			return 77
		}
		return NoState
	},
	// S59
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 78
		case 65 <= r && r <= 70: // ['A','F']
			return 79
		case 97 <= r && r <= 102: // ['a','f']
			return 79
		}
		return NoState
	},
	// S60
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 34
		case r == 92: // ['\','\']
			return 35
		default:
			return 6
		}
	},
	// S61
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 80
		case 65 <= r && r <= 70: // ['A','F']
			return 81
		case 97 <= r && r <= 102: // ['a','f']
			return 81
		}
		return NoState
	},
	// S62
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 82
		case 65 <= r && r <= 70: // ['A','F']
			return 83
		case 97 <= r && r <= 102: // ['a','f']
			return 83
		}
		return NoState
	},
	// S63
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 84
		case r == 45: // ['-','-']
			return 84
		case 48 <= r && r <= 57: // ['0','9']
			return 85
		}
		return NoState
	},
	// S64
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 64
		case r == 69: // ['E','E']
			return 65
		case r == 101: // ['e','e']
			return 65
		}
		return NoState
	},
	// S65
	func(r rune) int {
		switch {
		case r == 43: // ['+','+']
			return 86
		case r == 45: // ['-','-']
			return 86
		case 48 <= r && r <= 57: // ['0','9']
			return 87
		}
		return NoState
	},
	// S66
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 67
		}
		return NoState
	},
	// S67
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 67
		}
		return NoState
	},
	// S68
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 68
		case 65 <= r && r <= 70: // ['A','F']
			return 69
		case 97 <= r && r <= 102: // ['a','f']
			return 69
		}
		return NoState
	},
	// S69
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 68
		case 65 <= r && r <= 70: // ['A','F']
			return 69
		case 97 <= r && r <= 102: // ['a','f']
			return 69
		}
		return NoState
	},
	// S70
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 104: // ['a','h']
			return 46
		case r == 105: // ['i','i']
			return 88
		case 106 <= r && r <= 122: // ['j','z']
			return 46
		}
		return NoState
	},
	// S71
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 114: // ['a','r']
			return 46
		case r == 115: // ['s','s']
			return 89
		case 116 <= r && r <= 122: // ['t','z']
			return 46
		}
		return NoState
	},
	// S72
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 110: // ['a','n']
			return 46
		case r == 111: // ['o','o']
			return 90
		case 112 <= r && r <= 122: // ['p','z']
			return 46
		}
		return NoState
	},
	// S73
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 122: // ['a','z']
			return 46
		}
		return NoState
	},
	// S74
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 107: // ['a','k']
			return 46
		case r == 108: // ['l','l']
			return 91
		case 109 <= r && r <= 122: // ['m','z']
			return 46
		}
		return NoState
	},
	// S75
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 100: // ['a','d']
			return 46
		case r == 101: // ['e','e']
			return 92
		case 102 <= r && r <= 122: // ['f','z']
			return 46
		}
		return NoState
	},
	// S76
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 93
		case 65 <= r && r <= 70: // ['A','F']
			return 94
		case 97 <= r && r <= 102: // ['a','f']
			return 94
		}
		return NoState
	},
	// S77
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 93
		case 65 <= r && r <= 70: // ['A','F']
			return 94
		case 97 <= r && r <= 102: // ['a','f']
			return 94
		}
		return NoState
	},
	// S78
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 95
		case 65 <= r && r <= 70: // ['A','F']
			return 96
		case 97 <= r && r <= 102: // ['a','f']
			return 96
		}
		return NoState
	},
	// S79
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 95
		case 65 <= r && r <= 70: // ['A','F']
			return 96
		case 97 <= r && r <= 102: // ['a','f']
			return 96
		}
		return NoState
	},
	// S80
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 97
		case 65 <= r && r <= 70: // ['A','F']
			return 98
		case 97 <= r && r <= 102: // ['a','f']
			return 98
		}
		return NoState
	},
	// S81
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 97
		case 65 <= r && r <= 70: // ['A','F']
			return 98
		case 97 <= r && r <= 102: // ['a','f']
			return 98
		}
		return NoState
	},
	// S82
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 99
		case 65 <= r && r <= 70: // ['A','F']
			return 100
		case 97 <= r && r <= 102: // ['a','f']
			return 100
		}
		return NoState
	},
	// S83
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 99
		case 65 <= r && r <= 70: // ['A','F']
			return 100
		case 97 <= r && r <= 102: // ['a','f']
			return 100
		}
		return NoState
	},
	// S84
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 85
		}
		return NoState
	},
	// S85
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 85
		}
		return NoState
	},
	// S86
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 87
		}
		return NoState
	},
	// S87
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 87
		}
		return NoState
	},
	// S88
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 109: // ['a','m']
			return 46
		case r == 110: // ['n','n']
			return 101
		case 111 <= r && r <= 122: // ['o','z']
			return 46
		}
		return NoState
	},
	// S89
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 100: // ['a','d']
			return 46
		case r == 101: // ['e','e']
			return 102
		case 102 <= r && r <= 122: // ['f','z']
			return 46
		}
		return NoState
	},
	// S90
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 115: // ['a','s']
			return 46
		case r == 116: // ['t','t']
			return 103
		case 117 <= r && r <= 122: // ['u','z']
			return 46
		}
		return NoState
	},
	// S91
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 122: // ['a','z']
			return 46
		}
		return NoState
	},
	// S92
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 122: // ['a','z']
			return 46
		}
		return NoState
	},
	// S93
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 104
		case 65 <= r && r <= 70: // ['A','F']
			return 105
		case 97 <= r && r <= 102: // ['a','f']
			return 105
		}
		return NoState
	},
	// S94
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 104
		case 65 <= r && r <= 70: // ['A','F']
			return 105
		case 97 <= r && r <= 102: // ['a','f']
			return 105
		}
		return NoState
	},
	// S95
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		case r == 92: // ['\','\']
			return 32
		default:
			return 3
		}
	},
	// S96
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		case r == 92: // ['\','\']
			return 32
		default:
			return 3
		}
	},
	// S97
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 106
		case 65 <= r && r <= 70: // ['A','F']
			return 107
		case 97 <= r && r <= 102: // ['a','f']
			return 107
		}
		return NoState
	},
	// S98
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 106
		case 65 <= r && r <= 70: // ['A','F']
			return 107
		case 97 <= r && r <= 102: // ['a','f']
			return 107
		}
		return NoState
	},
	// S99
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 34
		case r == 92: // ['\','\']
			return 35
		default:
			return 6
		}
	},
	// S100
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 34
		case r == 92: // ['\','\']
			return 35
		default:
			return 6
		}
	},
	// S101
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 100: // ['a','d']
			return 46
		case r == 101: // ['e','e']
			return 108
		case 102 <= r && r <= 122: // ['f','z']
			return 46
		}
		return NoState
	},
	// S102
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 122: // ['a','z']
			return 46
		}
		return NoState
	},
	// S103
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 122: // ['a','z']
			return 46
		}
		return NoState
	},
	// S104
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 109
		case 65 <= r && r <= 70: // ['A','F']
			return 110
		case 97 <= r && r <= 102: // ['a','f']
			return 110
		}
		return NoState
	},
	// S105
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 109
		case 65 <= r && r <= 70: // ['A','F']
			return 110
		case 97 <= r && r <= 102: // ['a','f']
			return 110
		}
		return NoState
	},
	// S106
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 111
		case 65 <= r && r <= 70: // ['A','F']
			return 112
		case 97 <= r && r <= 102: // ['a','f']
			return 112
		}
		return NoState
	},
	// S107
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 111
		case 65 <= r && r <= 70: // ['A','F']
			return 112
		case 97 <= r && r <= 102: // ['a','f']
			return 112
		}
		return NoState
	},
	// S108
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 99: // ['a','c']
			return 46
		case r == 100: // ['d','d']
			return 113
		case 101 <= r && r <= 122: // ['e','z']
			return 46
		}
		return NoState
	},
	// S109
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		case r == 92: // ['\','\']
			return 32
		default:
			return 3
		}
	},
	// S110
	func(r rune) int {
		switch {
		case r == 34: // ['"','"']
			return 31
		case r == 92: // ['\','\']
			return 32
		default:
			return 3
		}
	},
	// S111
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 34
		case r == 92: // ['\','\']
			return 35
		default:
			return 6
		}
	},
	// S112
	func(r rune) int {
		switch {
		case r == 39: // [''',''']
			return 34
		case r == 92: // ['\','\']
			return 35
		default:
			return 6
		}
	},
	// S113
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 45
		case 65 <= r && r <= 90: // ['A','Z']
			return 46
		case r == 95: // ['_','_']
			return 47
		case 97 <= r && r <= 122: // ['a','z']
			return 46
		}
		return NoState
	},
}
