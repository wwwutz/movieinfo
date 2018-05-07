package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/ryanbradynd05/go-tmdb"
	"github.com/urfave/cli"
	"io"
	"log"
	"net/http"
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

var TMDB_API string
var maxe int
var download bool

func downloadFile(URL string, filename string) error {
	fmt.Printf("### download(%s, %s)\n", URL, filename)
	if _, err := os.Stat(filename); err == nil {
		// path/to/whatever exists
		fmt.Println("### EXISTS: " + filename + " already exists. skipping")
	} else {
		// the path does not exist or some error occurred.

		response, err := http.Get(URL)
		if err != nil {
			return err
		}
		defer response.Body.Close()
		if response.StatusCode != http.StatusOK {
			return errors.New(response.Status)
		}

		var data bytes.Buffer

		_, err = io.Copy(&data, response.Body)
		if err != nil {
			return err
		}

		file, err := os.Create(filename) // Truncates if file already exists, be careful!
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		defer file.Close()

		len, err := file.Write(data.Bytes())

		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}
		fmt.Printf("\nLength: %d bytes", len)
		fmt.Printf("\nFile Name: %s", file.Name())
	}
	return nil
}

func tmdbURLfile(filename string, mID int) error {
	if _, err := os.Stat(filename); err == nil {
		// path/to/whatever exists
		fmt.Println("### EXISTS: " + filename + " already exists. skipping")
	} else {
		file, err := os.Create(filename) // Truncates if file already exists, be careful!
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("[InternetShortcut]\r\nURL=https://www.themoviedb.org/movie/%d\r\n", mID))

		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}
	}
	return nil
}

func tmdbMovie(name string) (*tmdb.Movie, error) {
	// Replace any .'s in the title with spaces
	name = strings.Replace(name, ".", " ", -1)
	name = strings.Replace(name, "_", " ", -1)

	db := tmdb.Init(TMDB_API)

	lookup, _ := db.SearchMovie(name, nil)

	mID := 0

	mIDs := make(map[int]int)

	for _, element := range lookup.Results {
		if mID == 0 {
			// keep first
			mID = element.ID
		}
		mIDs[mID] =+ 1
		fmt.Printf("---------- ID: %d\n", element.ID)
		fmt.Printf("OriginalTitle: %s\n", element.OriginalTitle)
		fmt.Printf("        Title: %s\n", element.Title)
		fmt.Printf("  ReleaseDate: %s\n", element.ReleaseDate)
		fmt.Printf("   PosterPath: %s\n", element.PosterPath)
		fmt.Printf(" BackdropPath: %s\n", element.BackdropPath)
		fmt.Printf("\nResults = %#v\n\n", element)

		if download {
			filename := fmt.Sprintf("%s-%s-%d", name, element.ReleaseDate, mID)
			downloadFile("https://image.tmdb.org/t/p/original"+element.PosterPath, filename+"-poster.jpg")
		}
		maxe -= 1
		if maxe == 0 {
			break
		}
	}

	if mID != 0 {
		res, _ := db.GetMovieImages(mID, nil)
		fmt.Printf("Images: %#v\n", res)
		return db.GetMovieInfo(mID, nil)
	}

	for mIDk, _ := range mIDs {
     tmdbURLfile(fmt.Sprintf("%s-%d.URL", name, mIDk),mIDk)
	}

	// Nothing found on TMdb
	return nil, errors.New("no TMdb match found when looking up movie")
}

func Movie(name string, year int) (MovieResult, error) {

	tmdbResult, err := tmdbMovie(name)
	fmt.Printf("\nresult = %#v\n\n", tmdbResult)
	if err == nil {

		// Parse release date
		year := 0000
		date, parseError := dateparse.ParseAny(tmdbResult.ReleaseDate)
		if parseError == nil {
			year = date.Year()
		}

		return MovieResult{
			Title:       tmdbResult.Title,
			ReleaseDate: tmdbResult.ReleaseDate,
			Year:        year,
			Genres:      tmdbResult.Genres,
		}, err
	}

	return MovieResult{Title: "N/A"}, err
}

func main() {
	app := cli.NewApp()
	app.Name = "movieinfo"
	app.Usage = "query tmdb.org to download backdrops, cover and more"
	app.UsageText = "movieinfo [movie]"
	app.Version = "1.0"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "download, d",
			Usage: "download images and metadata",
		},
		cli.BoolFlag{
			Name:  "force, f",
			Usage: "force overwrite",
		},
		cli.IntFlag{
			Name:  "year, y",
			Usage: "year",
		},
		cli.IntFlag{
			Name:  "max, m",
			Usage: "max. entries, 0 unlimited",
		},
		cli.StringFlag{
			Name:   "TMDB_API",
			Usage:  "tmdb.org API key",
			EnvVar: "TMDB_API",
		},
	}

	app.Action = func(c *cli.Context) error {
		name := c.Args().Get(0)
		year := c.Int("year")
		maxe = c.Int("max")
		download = c.Bool("download")
		TMDB_API = c.String("TMDB_API")
		fmt.Println("    name: ", name)
		fmt.Println("download: ", download)
		fmt.Println("   force: ", c.Bool("force"))
		fmt.Println("    year: ", year)
		fmt.Println("     max: ", maxe)
		fmt.Println("  apikey: ", TMDB_API)

		fmt.Println(Movie(name, year))

		return nil
	}

	app.Run(os.Args)
}
