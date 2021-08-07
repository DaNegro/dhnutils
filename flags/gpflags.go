package flags

import (
	"flag"
	"fmt"
)

// EntryArg :: Defining a struct type
type EntryArg struct {
	Action    int
	Timestamp int64
	Path      string
	Ext       string
}

// GenericEntryArg :: Defining a struct type
type GenericEntryArg struct {
	Action    int
	Mode      int
	Param1    string
	Param2    string
	Param3    string
	Param4    string
	Param5    string
	Param6    string
	Timestamp int64
}

//func Xmain() {
//	fmt.Println("Hello World!")
//}

func RealFlags() *EntryArg {
	result := new(EntryArg)

	// Register Int flag.
	count := flag.Int("count", 5, "count of iterations")
	action := flag.Int("action", 2, "action Id")

	ts := flag.Int64("ts", 1392642999, "Unix TS")

	// Register String flag.
	rootPath := flag.String("path", "C:\\tmp\\logs\\", "root path.")
	fileExt := flag.String("ext", ".java", "file ext.")

	// Parse the flags.
	flag.Parse()

	// Print the argument.s
	fmt.Println("Argument", *count)
	fmt.Println("path", *rootPath)
	fmt.Println("ts", *ts)

	//
	result.Action = *action
	result.Timestamp = *ts
	result.Path = *rootPath
	result.Ext = *fileExt

	return result
}

func GetGenericFlags() *GenericEntryArg {
	result := new(GenericEntryArg)

	// Register Int flag.
	action := flag.Int("action", 0, "action Id")
	mode := flag.Int("mode", 0, "mode Id")

	// Register Int64 flag.
	ts := flag.Int64("ts", 1392642999, "Unix Timestamp")

	// Register String flag.
	param1 := flag.String("Param1", "", "1st Param.")
	param2 := flag.String("Param2", "", "2nd Param.")
	param3 := flag.String("Param3", "", "3rd Param.")
	param4 := flag.String("Param4", "", "4th Param.")
	param5 := flag.String("Param5", "", "5th Param.")
	param6 := flag.String("Param6", "", "6th Param.")

	// Parse the flags.
	flag.Parse()

	result.Action = *action
	result.Mode = *mode
	result.Timestamp = *ts

	result.Param1 = *param1
	result.Param2 = *param2
	result.Param3 = *param3
	result.Param4 = *param4
	result.Param5 = *param5
	result.Param6 = *param6

	return result
}

func RealflagsTrain() *EntryArg {
	result := new(EntryArg)

	// Register Int flag.
	count := flag.Int("count", 5, "count of iterations")
	ts := flag.Int64("ts", 1392642999, "Unix TS")

	// Register String flag.
	rootPath := flag.String("path", "C:\\tmp\\ExtractIX\\", "root path.")
	fileExt := flag.String("ext", ".java", "file ext.")

	// Parse the flags.
	flag.Parse()

	// Print the argument.s
	fmt.Println("Argument", *count)
	fmt.Println("path", *rootPath)
	fmt.Println("ts", *ts)

	// Get int from the Int pointer.
	value := *count
	fmt.Println("Value -->> ", value)

	// Get String from the String pointer.
	value2 := *rootPath
	fmt.Println("Path -->> ", value2)

	// Get int from the Int pointer.
	value3 := *ts
	fmt.Println("ts -->> ", value3)
	//
	result.Timestamp = *ts
	result.Path = *rootPath
	result.Ext = *fileExt

	return result
}

func DemoFlags() {
	// Register Int flag.
	count := flag.Int("count", 5, "count of iterations")
	// Parse the flags.
	flag.Parse()

	// Print the argument.
	fmt.Println("Argument", *count)

	// Get int from the Int pointer.
	value := *count
	fmt.Println("Value", value)
}

