package qutils

type ArgsString []string

func (Args ArgsString) Default(index int, args ...string) string {
	if index >= 0 {
		if index < len(Args) {
			return Args[index]
		} else if index < len(args) {
			return Args[index]
		} else if len(args) == 1 {
			return args[0]
		}
	}
	return ""
}

type ArgsInt []int

func (Args ArgsInt) Default(index int, args ...int) int {
	if index >= 0 {
		if index < len(Args) {
			return Args[index]
		} else if index < len(args) {
			return Args[index]
		} else if len(args) == 1 {
			return args[0]
		}
	}
	return 0
}

type ArgsInt64 []int64

func (Args ArgsInt64) Default(index int, args ...int64) int64 {
	if index >= 0 {
		if index < len(Args) {
			return Args[index]
		} else if index < len(args) {
			return Args[index]
		} else if len(args) == 1 {
			return args[0]
		}
	}
	return 0
}

type ArgsFloat64 []float64

func (Args ArgsFloat64) Default(index int, args ...float64) float64 {
	if index >= 0 {
		if index < len(Args) {
			return Args[index]
		} else if index < len(args) {
			return Args[index]
		} else if len(args) == 1 {
			return args[0]
		}
	}
	return 0
}

type ArgsBool []bool

func (Args ArgsBool) Default(index int, args ...bool) bool {
	if index >= 0 {
		if index < len(Args) {
			return Args[index]
		} else if index < len(args) {
			return Args[index]
		} else if len(args) == 1 {
			return args[0]
		}
	}
	return false
}

type ArgsInterface []interface{}

func (Args ArgsInterface) Default(index int, args ...interface{}) interface{} {
	if index >= 0 {
		if index < len(Args) {
			return Args[index]
		} else if index < len(args) {
			return Args[index]
		} else if len(args) == 1 {
			return args[0]
		}
	}
	return nil
}

type ArgsArray [][]interface{}

func (Args ArgsArray) Default(index int, args ...[]interface{}) []interface{} {
	if index >= 0 {
		if index < len(Args) {
			return Args[index]
		} else if index < len(args) {
			return Args[index]
		} else if len(args) == 1 {
			return args[0]
		}
	}
	return []interface{}{}
}

type ArgsObject []map[string]interface{}

func (Args ArgsObject) Default(index int, args ...map[string]interface{}) map[string]interface{} {
	if index >= 0 {
		if index < len(Args) {
			return Args[index]
		} else if index < len(args) {
			return Args[index]
		} else if len(args) == 1 {
			return args[0]
		}
	}
	return map[string]interface{}{}
}
