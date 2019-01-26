package core

import (
	"github.com/Xarvalus/go-networking/server/utils"
	"io/ioutil"
)

type Data struct {
	root string
	files []string
}

const DataPath = "rpc/core/data"

func InitData() *Data {
	files, err := ioutil.ReadDir(DataPath)
	utils.LogFatalError(err)

	var filesNames []string
	for _, file := range files {
		filesNames = append(filesNames, file.Name())
	}

	return &Data{
		root: DataPath,
		files: filesNames,
	}
}

// Includes data sources from 1st up to `CumulativelyUpTo`
type Args struct {
	CumulativelyUpTo int
}

func (data *Data) FetchFiles(args *Args, reply *string) error {
	channel := make(chan string)

	go data.readFiles(args, channel)

	for content := range channel {
		*reply += content
	}

	return nil
}

func (data *Data) readFiles(args *Args, channel chan string) {
	for _, file := range data.files[:args.CumulativelyUpTo] {
		content, err := ioutil.ReadFile(data.root + "/" + file)
		utils.LogFatalError(err)

		channel <- string(content)
	}

	close(channel)
}
