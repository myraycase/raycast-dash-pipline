package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/antchfx/xmlquery"
)

const dashPiplineExec = "./dash-pipline"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("query is required")
		return
	}
	query := os.Args[1]

	items, err := execQuery(query)
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range items {
		fmt.Println(item.Title)
	}
}

func execQuery(query string) ([]*Item, error) {
	var sb strings.Builder
	cmd := exec.Command(dashPiplineExec, query)
	cmd.Stdout = &sb
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return parseItems(sb.String())
}

type Item struct {
	UID          string
	Title        string
	SubTitle     string
	Icon         string
	Quicklookurl string
}

func parseItems(itemsRaw string) ([]*Item, error) {
	r := strings.NewReader(itemsRaw)
	doc, err := xmlquery.Parse(r)
	if err != nil {
		return nil, err
	}

	items := make([]*Item, 0)

	xmlquery.FindEach(doc, "//item", func(_ int, node *xmlquery.Node) {
		subTitle := ""
		subTitles := node.Parent.SelectElements("subtitle")
		for _, t := range subTitles {
			if len(t.Attr) == 0 {
				subTitle = t.InnerText()
				break
			}
		}

		items = append(items, &Item{
			UID:          node.SelectAttr("uid"),
			Title:        node.SelectElement("title").InnerText(),
			SubTitle:     subTitle,
			Icon:         node.SelectElement("icon").InnerText(),
			Quicklookurl: node.SelectElement("quicklookurl").InnerText(),
		})
	})

	return items, nil
}
