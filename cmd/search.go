package cmd

import (
	"fmt"
	"net/http"
	"net/url"

	dj "github.com/rchaganti/dadjoke-go"

	"github.com/spf13/cobra"
)

var (
	term  string
	page  int
	limit int
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search retrieves dad jokes based on a search term",
	Long:  "search is a command line tool for retrieving dad jokes from icanhazdadjoke.com based on a search term",
	Run: func(cmd *cobra.Command, args []string) {
		baseUrl, err := url.Parse("https://icanhazdadjoke.com")
		if err != nil {
			panic(err)
		}

		c := dj.Client{
			BaseUrl:    baseUrl,
			UserAgent:  "dj-go (https://github.com/rchaganti/dadjoke-go)",
			HttpClient: &http.Client{},
		}

		jokes, err := c.SearchDadJokes(term, page, limit)

		if err != nil {
			panic(err)
		}

		jokeResults := jokes.Results

		for i, j := range jokeResults {
			fmt.Printf("%d. %s\n", i+1, j.Joke)
		}
	},
}

func init() {
	searchCmd.Flags().StringVarP(&term, "term", "t", "", "search term")
	searchCmd.MarkFlagRequired("term")

	searchCmd.Flags().IntVarP(&page, "page", "p", 1, "page number to fetch")
	searchCmd.Flags().IntVarP(&limit, "limit", "l", 4, "number of jokes to fetch per page")

	rootCmd.AddCommand(searchCmd)
}
