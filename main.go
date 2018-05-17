package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/araddon/dateparse"
	"github.com/ryanbradynd05/go-tmdb"
	"github.com/urfave/cli"
)

type MovieResult struct {
	Title       string
	ReleaseDate string
	Year        int
}

var TMDB_API string
var maxe int
var download bool
var verbose bool

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

	writefile(filename, data.Bytes())
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

func exists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		fmt.Println("### EXISTS: " + filename + " exists.")
		return true
	}
	return false
}

func writefile(filename string, data []byte) error {
	file, err := os.Create(filename) // Truncates if file already exists
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()

	_, err = file.Write(data)

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	return nil
}

func days(s int, n int) string {
	if s == 0 {
		return "0 sec"
	}

	var T [4]int

	T[0] += int(s / (60 * 60 * 24)) // days
	s -= T[0] * (60 * 60 * 24)
	T[1] += int(s / (60 * 60)) // hrs
	s -= T[1] * (60 * 60)
	T[2] += int(s / 60) // min
	s -= T[2] * 60
	T[3] += int(s) // min
	var L []string
	x := [4]string{"d", "h", "min", "sec"}
	j := 0
	for i := 0; i < len(x); i++ {
		y := T[i]
		if y != 0 {
			if n != 0 {
				n -= 1
				if n < 0 {
					break
				}
			}
			L = append(L, fmt.Sprintf("%d %s", y, x[i]))
			j += 1
		}
	}
	return strings.Join(L, " ")
}

func li(i int) int {
	n := 1
	if i >= 100000000 {
		n += 8
		i /= 100000000
	}
	if i >= 10000 {
		n += 4
		i /= 10000
	}
	if i >= 100 {
		n += 2
		i /= 100
	}
	if i >= 10 {
		n += 1
	}
	return n
}

func dumptmdbMovie(m *tmdb.Movie) error {
	b, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	return err
}

func panicon(err error, msg string) {
	if err != nil {
		_, file, no, ok := runtime.Caller(1)
		if ok {
			panic(fmt.Sprintf("// %s failed with err=%v\n// %s#%d\n", msg, err, file, no))
		} else {
         	panic(fmt.Sprintf("// %s failed with err=%v\n", msg, err))
		}
	}
}

func tmdbMovie2txt(tm tmdb.Movie) (string, error) {

	txt := fmt.Sprintf("tmdbID:   %d\n", tm.ID)
	txt += fmt.Sprintf("Title:    %s\n", tm.Title)
	if tm.Tagline != "" {
		txt += fmt.Sprintf("Tagline:  %s\n", tm.Tagline)
	}
	if tm.Title != tm.OriginalTitle {
		txt += fmt.Sprintf("OTitle:   %s\n", tm.OriginalTitle)
	}
	txt += fmt.Sprintf("Release:  %s\n", tm.ReleaseDate)
	txt += fmt.Sprintf("Runtime:  %s\n", days(int(tm.Runtime)*60, 0))
	txt += fmt.Sprintf("Overview: %s\n", tm.Overview)

	txt += "\n"
	l := li(len(tm.Credits.Cast))
	for i := range tm.Credits.Cast {
		txt += fmt.Sprintf("- %*d. %s: %s\n", l, i+1, tm.Credits.Cast[i].Name, tm.Credits.Cast[i].Character)
	}
	txt += "\n"
	for i := range tm.Credits.Cast {
		txt += fmt.Sprintf("/ %*d. %s %s\n", l, i+1, tm.Credits.Cast[i].CreditID, tm.Credits.Cast[i].Name)
	}
	/*
		txt += "\n"
		l = li(len(tm.Credits.Crew))
		for i := range tm.Credits.Crew {
			txt += fmt.Sprintf("= %*d. %s: %s / %s\n", l, i+1, tm.Credits.Crew[i].Department, tm.Credits.Crew[i].Name, tm.Credits.Crew[i].Job)
		}
	*/
	return txt, nil
}

func tmdbMovie(mID int, search string, argsyear int) (*tmdb.Movie, error) {

	var options = make(map[string]string)

	options["language"] = "de-DE"

	if argsyear != 0 {
		options["year"] = strconv.Itoa(argsyear)
	}

	db := tmdb.Init(TMDB_API)

	// no mID supplied: go search for a couple of movies
	if mID == 0 {
		lookup, _ := db.SearchMovie(search, options)
		// one result: that's fine
		if len(lookup.Results) == 1 {
			mID = lookup.Results[0].ID
		} else {
			// more than one result:
			//  - download posters & backdrop
			//  - create minimal files to choose from
			// - do not download complete tmdb.Movie

			for _, element := range lookup.Results {
				fmt.Printf("---------- ID: %d\n", element.ID)
				fmt.Printf("        Title: %s\n", element.Title)
				if element.Title != element.OriginalTitle {
					fmt.Printf("OriginalTitle: %s\n", element.OriginalTitle)
				}
				fmt.Printf("  ReleaseDate: %s\n", element.ReleaseDate)
				// fmt.Printf("   PosterPath: %s\n", element.PosterPath)
				// fmt.Printf(" BackdropPath: %s\n", element.BackdropPath)
				//		fmt.Printf("\nResults = %#v\n\n", element)
				fmt.Printf("\n")
				year := 0
				date, parseError := dateparse.ParseAny(element.ReleaseDate)
				if parseError == nil {
					year = date.Year()
				}

				if download {
					// .URL, -{poster,backdrop}.jpg
					cleantitle, _ := cleanuptitle(element.Title)

					filename := fmt.Sprintf("%s-%d-%04d", cleantitle, element.ID, year)

					url := fmt.Sprintf("[InternetShortcut]\r\nURL=https://www.themoviedb.org/movie/%d\r\n", element.ID)
					writefile(filename+".URL", []byte(url))

					if element.PosterPath != "" {
						downloadFile("https://image.tmdb.org/t/p/original"+element.PosterPath, filename+"-poster.jpg")
					}
					if element.BackdropPath != "" {
						downloadFile("https://image.tmdb.org/t/p/original"+element.BackdropPath, filename+"-backdrop.jpg")
					}
				}
				maxe -= 1
				if maxe == 0 {
					break
				}
			}
		}
	}

	//	fmt.Printf("YEAH mID = %#v\n", mID)

	if mID != 0 {
		//		res, _ := db.GetMovieImages(mID, nil)
		//		fmt.Printf("Images: %#v\n", res)

		var m *tmdb.Movie
		options["append_to_response"] = "credits"
		m, err := db.GetMovieInfo(mID, options)
		panicon(err,"GetMovieInfo1")

		if verbose {
			dumptmdbMovie(m)
		}

		// no overview in our language ?
		if m.Overview == "" && options["language"] != m.OriginalLanguage {
			fmt.Printf("# Overview empty. Switching from %s to %s, retrying\n", options["language"], m.OriginalLanguage)
			options["language"] = m.OriginalLanguage
			m, err := db.GetMovieInfo(mID * -1, options)
			panicon(err,"GetMovieInfo2")
			if verbose {
				dumptmdbMovie(m)
			}
		}
		year := 0
		date, err := dateparse.ParseAny(m.ReleaseDate)
		if err == nil {
			year = date.Year()
		}

		txt, err := tmdbMovie2txt(*m)
		panicon(err,"tmdbMovie2txt")

		fmt.Printf("### START .txt\n%s###  END  .txt\n", txt)

		if download {
			// .txt, .URL, -{poster,backdrop}.jpg
			cleantitle, _ := cleanuptitle(m.Title)

			filename := fmt.Sprintf("%s-%d-%04d", cleantitle, mID, year)

			writefile(filename+".txt", []byte(txt))

			url := fmt.Sprintf("[InternetShortcut]\r\nURL=https://www.themoviedb.org/movie/%d\r\n", mID)
			writefile(filename+".URL", []byte(url))

			if m.PosterPath != "" {
				downloadFile("https://image.tmdb.org/t/p/original"+m.PosterPath, filename+"-poster.jpg")
			}
			if m.BackdropPath != "" {
				downloadFile("https://image.tmdb.org/t/p/original"+m.BackdropPath, filename+"-backdrop.jpg")
			}

		}

		return m, err
	}

	// Nothing found on TMdb
	return nil, errors.New("no TMdb match found when looking up movie")
}

func mIDfromurlname(name string) int {
	// check if name fits as .URL file, return found mID or 0
	m := regexp.MustCompile(`-(\d+)-\d{4}.\w{2,3}$`).FindStringSubmatch(name)
	mID := 0
	if len(m) > 0 {
		mi, err := strconv.Atoi(m[1])
		if err == nil && mi > 0 {
			mID = mi
		}
	}
	return mID
}

func cleanuptitle(name string) (string, int) {
	// " - 2015"  ==> year
	//  getyear := regexp.MustCompile(`\s-\s([12]\d\d\d)`)
	//	year,_ := strconv.Atoi(getyear.FindString(name))
	// total recall-861-1990.URL ==> total recall
	// get rid of bad chars
	clname := name
	year := 0
	nameyear := regexp.MustCompile(`(.*?)\s-\s([12]\d\d\d)(.*)`).FindStringSubmatch(name)

	if len(nameyear) > 0 {
		clname = nameyear[1]
		year, _ = strconv.Atoi(nameyear[2])
	}
	//	re = regexp.MustCompile(`[_.\-]`)
	//	clname = re.ReplaceAllString(clname, ``)
	clname = regexp.MustCompile(`\-\d+\-\d\d\d\d\.\w{2,3}$`).ReplaceAllString(clname, ``)
	clname = regexp.MustCompile(`\.\w{2,3}$`).ReplaceAllString(clname, ``)
	clname = regexp.MustCompile(`[_.\-:/\?\*]`).ReplaceAllString(clname, ` `)
	// trim everywhere
	clname = strings.Join(strings.Fields(clname), " ")
	return clname, year
}

func Movie(mID int, search string, year int) (MovieResult, error) {

	tmdbResult, err := tmdbMovie(mID, search, year)
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
			Name:  "verbose, vv",
			Usage: "whatever",
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
		search, year := cleanuptitle(arg)
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
		verbose = c.Bool("verbose")
		TMDB_API = c.String("TMDB_API")
		fmt.Println("     arg: ", arg)
		fmt.Println("  search: ", search)
		fmt.Println("      id: ", mID)
		fmt.Println("download: ", download)
		fmt.Println("    year: ", year)
		fmt.Println("     max: ", maxe)
		fmt.Println(" verbose: ", verbose)
		fmt.Print("\n")
		Movie(mID, search, year)

		return nil
	}

	app.Run(os.Args)
}
