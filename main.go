package main

import (
	"errors"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/ryanbradynd05/go-tmdb"
	"github.com/urfave/cli"
	"os"
	"strings"
)

type MovieResult struct {
	Title       string
	ReleaseDate string
	Year        int
	Genres      []struct {
		ID   int
		Name string
	}
}

// Returns the most likely TV result from the TMdb API
func tmdbMovie(name string) (*tmdb.Movie, error) {
	// Replace any .'s in the title with spaces
	name = strings.Replace(name, ".", " ", -1)
	name = strings.Replace(name, "_", " ", -1)

	db := tmdb.Init(os.Getenv("TMDB_API"))
	lookup, _ := db.SearchMovie(name, nil)

	mID := 0

	for _, element := range lookup.Results {
		if mID == 0 {
			// keep first
			mID = element.ID
		}
		fmt.Printf("---------- ID: %d\n", element.ID)
		fmt.Printf("OriginalTitle: %s\n", element.OriginalTitle)
		fmt.Printf("        Title: %s\n", element.Title)
		fmt.Printf("  ReleaseDate: %s\n", element.ReleaseDate)
		fmt.Printf("   PosterPath: %s\n", element.PosterPath)
		fmt.Printf(" BackdropPath: %s\n", element.BackdropPath)
	}

	if mID != 0 {
		return db.GetMovieInfo(mID, nil)
	}

	// Nothing found on TMdb
	return nil, errors.New("no TMdb match found when looking up movie")
}

func Movie(name string, year int) (MovieResult, error) {

	tmdbResult, err := tmdbMovie(name)

	if err == nil {

		// Parse release date
		date, parseError := dateparse.ParseAny(tmdbResult.ReleaseDate)

		if parseError == nil {
			return MovieResult{
				Title:       tmdbResult.Title,
				ReleaseDate: tmdbResult.ReleaseDate,
				Year:        date.Year(),
				Genres:      tmdbResult.Genres,
			}, err
		}

		return MovieResult{
			Title:       tmdbResult.Title,
			ReleaseDate: tmdbResult.ReleaseDate,
			Year:        0000,
			Genres:      tmdbResult.Genres,
		}, err
	}

	return MovieResult{Title: "NA"}, err
}

func main() {
	app := cli.NewApp()
	app.Name = "movieinfo"
	app.Usage = "movieinfo <movie>"
	app.UsageText = "usage text"
	app.Version = "1.0"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "force overwrite",
		},
		cli.IntFlag{
			Name:  "year, y",
			Usage: "year",
		},
	}

	app.Action = func(c *cli.Context) error {
		name := c.Args().Get(0)
		year := c.Int("year")
		fmt.Println(" name: ", name)
		fmt.Println("force: ", c.Bool("force"))
		fmt.Println(" year: ", year)

		fmt.Println(Movie(name, year))

		return nil
	}

	app.Run(os.Args)
}
