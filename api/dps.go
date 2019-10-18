package api

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/rgeoghegan/tabulate"
)

func insertFirst(slice []string, str string) []string {
	slice, slice[0] = append(slice[:1], slice[0:]...), str
	return slice
}

func remove(slice []string, pos int) []string {
	if pos >= len(slice) {
		return slice
	}
	return append(slice[:pos], slice[pos+1:]...)
}

//RunDps outputs processes in each containers
func RunDps(apiVersion string) {
	var (
		cli *client.Client
		err error
	)
	if apiVersion == "default" {
		cli, err = client.NewClientWithOpts(client.FromEnv)
	} else {
		cli, err = client.NewClientWithOpts(client.WithVersion(apiVersion))
	}
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})

	if err != nil {
		panic(err)
	}

	contents := make([][]string, 0)
	header := []string{"Image", "Container"}

	for _, container := range containers {
		processes, _ := cli.ContainerTop(ctx, container.ID, []string{"aux"})

		//header作成
		if len(header) == 2 {
			header = append(header, processes.Titles...)
		}

		for _, process := range processes.Processes {

			//イメージ名とコンテナ名の追加
			process = insertFirst(process, container.Names[0])
			process = insertFirst(process, container.Image)

			//Commandが長すぎる場合に省略
			if len(process[12]) > 20 {
				process[12] = process[12][:20] + "..."
			}

			//"Tty"の欄を削除
			process = remove(process, 8)

			contents = append(contents, process)
		}
	}

	if len(contents) == 0 {
		fmt.Println("No processes")
	} else {
		//"Tty"の欄を削除
		header = remove(header, 8)
		layout := &tabulate.Layout{Format: tabulate.SimpleFormat}
		layout.Headers = header
		tableText, err := tabulate.Tabulate(contents, layout)
		if err != nil {
			panic(err)
		}
		fmt.Println(tableText)
	}

}
