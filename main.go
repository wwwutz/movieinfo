package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/ryanbradynd05/go-tmdb"
	"github.com/urfave/cli"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MovieResult struct {
	Title       string
	ReleaseDate string
	Year        int
}

type MyMovieTxt struct {
	ID          int
	Title       string
	Year        int
	TagLine     string
	OverView    string
	ReleaseDate string
	ImdbId      string
}

var TMDB_API string
var maxe int
var download bool

func downloadFile(URL string, filename string) error {
	fmt.Printf("### download(%s, %s)\n", URL, filename)
	if _, err := os.Stat(filename); err == nil {
		// path/to/whatever exists
		fmt.Println("### EXISTS: " + filename + " already exists. skipping")
		return err
	}
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

	_, err = file.Write(data.Bytes())

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
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

func txtfile(filename string, data []byte) error {
	if _, err := os.Stat(filename); err == nil {
		// path/to/whatever exists
		fmt.Println("### EXISTS: " + filename + " already exists. skipping")
	} else {
		file, err := os.Create(filename) // Truncates if file already exists, be careful!
		if err != nil {
			log.Fatalf("failed creating file: %s", err)
		}
		defer file.Close()

		_, err = file.Write(data)

		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}
	}
	return nil
}

func tmdbMovieShort2txt(ms tmdb.MovieShort) ([]byte, error) {
	buf, err := json.MarshalIndent(ms, "", "---")
	if err != nil {
		log.Fatalf("MarshalIndent: %s", err)
	}
	return buf, err
}

func tmdbMovie(mID int, name string, argsyear int) (*tmdb.Movie, error) {

	var options = make(map[string]string)

	options["language"] = "de-DE"

	if argsyear != 0 {
		options["year"] = strconv.Itoa(argsyear)
	}

	db := tmdb.Init(TMDB_API)

	// no mID supplied: go search for a couple of movies
	if mID == 0 {
		lookup, _ := db.SearchMovie(name, options)
		if len(lookup.Results) == 1 {

		} else {

			for _, element := range lookup.Results {
				if mID == 0 {
					// keep first
					mID = element.ID
				}
				fmt.Printf("---------- ID: %d\n", element.ID)
				fmt.Printf("OriginalTitle: %s\n", element.OriginalTitle)
				fmt.Printf("        Title: %s\n", element.Title)
				fmt.Printf("  ReleaseDate: %s\n", element.ReleaseDate)
				// fmt.Printf("   PosterPath: %s\n", element.PosterPath)
				// fmt.Printf(" BackdropPath: %s\n", element.BackdropPath)
				//		fmt.Printf("\nResults = %#v\n\n", element)

				year := 0
				date, parseError := dateparse.ParseAny(element.ReleaseDate)
				if parseError == nil {
					year = date.Year()
				}

				if download {
					tmdbURLfile(fmt.Sprintf("%s-%d-%d.URL", name, element.ID, year), element.ID)
					filename := fmt.Sprintf("%s-%d-%d", name, element.ID, year)
					if element.PosterPath != "" {
						downloadFile("https://image.tmdb.org/t/p/original"+element.PosterPath, filename+"-poster.jpg")
					}
					if element.BackdropPath != "" {
						downloadFile("https://image.tmdb.org/t/p/original"+element.BackdropPath, filename+"-backdrop.jpg")
					}

					txt, err := tmdbMovieShort2txt(element)
					if err == nil {
						txtfile(fmt.Sprintf("%s-%d-%d.txt", name, element.ID, year), txt)
					}
				}
				maxe -= 1
				if maxe == 0 {
					break
				}
			}
		}
	}

	fmt.Printf("YEAH mID = %#v\n", mID)

	if mID != 0 {
		//		res, _ := db.GetMovieImages(mID, nil)
		//		fmt.Printf("Images: %#v\n", res)

		var m *tmdb.Movie
		m, err := db.GetMovieInfo(mID, options)
		fmt.Printf("tmdb.Movie: %#v\n", m)

		b, err := json.MarshalIndent(m, "", "---")
		if err != nil {
			fmt.Println("error:", err)
		}
		os.Stdout.Write(b)

		return m, err
	}

	// Nothing found on TMdb
	return nil, errors.New("no TMdb match found when looking up movie")
}

func mIDfromurlname(name string) int {
	// check if name fits as .URL file, return found mID or 0
	m := regexp.MustCompile(`-(\d+)-\d{4}.URL$`).FindStringSubmatch(name)
	mID := 0
	if len(m) > 0 {
		mi, err := strconv.Atoi(m[1])
		if err == nil && mi > 0 {
			mID = mi
		}
	}
	return mID
}

func cleanupname(name string) (string, int) {
	// " - 2015"  ==> year
	//  getyear := regexp.MustCompile(`\s-\s([12]\d\d\d)`)
	//	year,_ := strconv.Atoi(getyear.FindString(name))
	clname := name
	year := 0
	nameyear := regexp.MustCompile(`(.*?)\s-\s([12]\d\d\d)(.*)`).FindStringSubmatch(name)

	if len(nameyear) > 0 {
		clname = nameyear[1]
		year, _ = strconv.Atoi(nameyear[2])
	}
	//	re = regexp.MustCompile(`[_.\-]`)
	//	clname = re.ReplaceAllString(clname, ``)

	clname = regexp.MustCompile(`[_.\-]`).ReplaceAllString(clname, ` `)
	// trim everywhere
	clname = strings.Join(strings.Fields(clname), " ")
	return clname, year
}

func Movie(mID int, name string, year int) (MovieResult, error) {

	tmdbResult, err := tmdbMovie(mID, name, year)
	//	fmt.Printf("\nresult = %#v\n\n", tmdbResult)
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
		cli.IntFlag{
			Name:  "id, i",
			Usage: "tmdb movie ID",
		},
		cli.StringFlag{
			Name:   "TMDB_API",
			Usage:  "tmdb.org API key",
			EnvVar: "TMDB_API",
		},
	}

	app.Action = func(c *cli.Context) error {
		arg := c.Args().Get(0)

		forceyear := c.Int("year")
		title, year := cleanupname(arg)
		if forceyear > 0 {
			year = forceyear
		}
		// check if we supplied a complete .URL files as arg
		mID := c.Int("id")
		if mID == 0 {
			mID = mIDfromurlname(arg)
		}

		maxe = c.Int("max")
		download = c.Bool("download")
		TMDB_API = c.String("TMDB_API")
		fmt.Println("     arg: ", arg)
		fmt.Println("   title: ", title)
		fmt.Println("      id: ", mID)
		fmt.Println("download: ", download)
		fmt.Println("   force: ", c.Bool("force"))
		fmt.Println("    year: ", year)
		fmt.Println("     max: ", maxe)
		//		fmt.Println("  apikey: ", TMDB_API)

		Movie(mID, title, year)

		return nil
	}

	app.Run(os.Args)
}
