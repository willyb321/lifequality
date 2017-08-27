package command

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/AlecAivazis/survey"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

// CmdIn is used like`in <place>`
func CmdIn(c *cli.Context) error {
	// Write your code here
	if string(c.Args().Get(0)) != "" {
		fmt.Println("Starting lookup for", c.Args().Get(0))
		searchLifeData(string(c.Args().Get(0)))
	} else {
		fmt.Println("Enter a city to look up.")
	}
	return nil
}

func selectCity(cities []string) string {
	city := ""
	prompt := &survey.Select{
		Message: "Choose matching City:",
		Options: cities,
	}
	survey.AskOne(prompt, &city, nil)
	return city
}

func searchLifeData(search string) {
	var URL *url.URL
	URL, err := url.Parse("https://api.teleport.org")
	URL.Path += "/api/cities/"
	parameters := url.Values{}
	parameters.Add("search", search)
	parameters.Add("embed", "city:search-results/city:item/city:urban_area/ua:scores")
	URL.RawQuery = parameters.Encode()
	response, err := http.Get(URL.String())
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var decData searchReturn
		err = json.Unmarshal(data, &decData)
		citiesData := decData
		cities := make([]string, 0)
		for _, v := range citiesData.Embedded.CitySearchResults {
			cities = append(cities, v.MatchingFullName)
		}
		city := selectCity(cities)
		if city != "" {
			for _, v := range citiesData.Embedded.CitySearchResults {
				if city == v.MatchingFullName {
					printLifeData(v.Embedded.CityItem.Embedded.CityUrbanArea.Embedded.UaScores.Categories,
						v.Embedded.CityItem.Embedded.CityUrbanArea.Embedded.UaScores.Summary,
						v.Embedded.CityItem.Embedded.CityUrbanArea.Embedded.UaScores.TeleportCityScore)
					break
				}
			}
		}
	}
}

func printLifeData(scores lifeScores, summary string, cityScore float64) {
	tableHeaders := []string{"Name", "Score"}
	var tableData [][]string
	for _, v := range scores {
		tableData = append(tableData, []string{v.Name, string(strconv.FormatFloat(v.ScoreOutOf10, 'f', 1, 64))})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(tableHeaders)
	table.SetFooter([]string{"Teleport Score", string(strconv.FormatFloat(cityScore, 'f', 1, 64))})
	table.AppendBulk(tableData)
	table.Render()
	// fmt.Printf("Summary:")
	// summary = strings.Replace(summary, "\n", "", -1)
	// summary = wordwrap.WrapString(summary, 3)
	// cityScore, _ = strconv.ParseFloat(cityScore, 64)
	// fmt.Println(summary)
}
