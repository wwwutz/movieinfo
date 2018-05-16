# movieinfo
get ridiculous least information from tmdb

# warning

Do **not** take this source as go/golang reference. I'm a `perl` person. This is my first try to do a little bit more with golang. Based on `"golang string to int"` or `"golang write file"` web searches.
The `movieinfo` command creates several files in your current directory, so it might overwrite yours.
It connects to the internet [[1]](https://www.youtube.com/watch?v=iDbyYGrswtg) to download data. Well, that's the idea.

You better have backup.
You have been warned.

[1] https://www.youtube.com/watch?v=iDbyYGrswtg

## prerequisites

* You will need a tmdb.org API key as seen on  https://www.themoviedb.org/documentation/api

1. goto https://www.themoviedb.org
1. create an account
1. get your API key

# features
* works for me. Fork it, change it: it'll work for you
* downloads images like poster, backdrop etc.
* creates unique filenames which group together on sort by name in your favourite filemanager like e.g. `/usr/bin/ls`
* is designed to be used as a command line tool
* due to `golang` it runs on
  * Synology DSM NAS
  * Debian laptop
  * Linux workstation
  * Windows (integrated into a [DirectoryOpus](https://www.gpsoft.com.au) button)

# repeat

* You will need a tmdb.org API key as seen on  https://www.themoviedb.org/documentation/api

## installation linux
* go get `go`
* `go get github.com/wwwutz/movieinfo`
* `cd $GOPATH/src/github.com/wwwutz/movieinfo`
* `go build`
* `cp movieinfo` somewhere in your `$PATH` if you prefer
* `GOOS=wiondows GOARCH=amd64 go build` to create a windows `.exe`

## usage examples

Set up you API key. Either save it as environment variable `TMDB_API` or supply it on the commandline.

### get info for "Total Recall"
 We'll find four versions of "Total Recall". The third ID 180 is obviously wrong.

```
$ ./movieinfo "total recall"
     arg:  total recall
  search:  total recall
      id:  0
download:  false
    year:  0
     max:  0
 verbose:  false

---------- ID: 64635
        Title: Total Recall
  ReleaseDate: 2012-08-02

---------- ID: 861
        Title: Total Recall - Die totale Erinnerung
OriginalTitle: Total Recall
  ReleaseDate: 1990-06-01

---------- ID: 408340
        Title: Total Recall
  ReleaseDate:

---------- ID: 180
        Title: Minority Report
  ReleaseDate: 2002-06-20
```

### download images for "Total Recall"

Let's just download all available poster and backdrop images and create a `.URL` file in the current directory.

before:
```
$ ls -l
total 8748
-rw-r--r-- 1 wwwutz wwwutz    1072 May  8 15:11 LICENSE
-rw-r--r-- 1 wwwutz wwwutz    1955 May  9 11:39 README.md
-rw-r--r-- 1 wwwutz wwwutz    5776 May  9 11:50 main.go
-rwxr-xr-x 1 wwwutz wwwutz 8941192 May  9 12:00 movieinfo
```
calling `movieinfo` command with `--download` option:
```
$ ./movieinfo -d "total recall"
     arg:  total recall
  search:  total recall
      id:  0
download:  true
    year:  0
     max:  0
 verbose:  false

---------- ID: 64635
        Title: Total Recall
  ReleaseDate: 2012-08-02

### download(https://image.tmdb.org/t/p/original/4zgwx4HySRVjqSlmbrEKetJr5qo.jpg, Total Recall-64635-2012-poster.jpg)
### download(https://image.tmdb.org/t/p/original/orFQbyZ6g7kPFaJXmgty0M88wJ0.jpg, Total Recall-64635-2012-backdrop.jpg)
---------- ID: 861
        Title: Total Recall - Die totale Erinnerung
OriginalTitle: Total Recall
  ReleaseDate: 1990-06-01

### download(https://image.tmdb.org/t/p/original/unjJqoBkzdUIA5Bi1rDdVHo0949.jpg, Total Recall Die totale Erinnerung-861-1990-poster.jpg)
### download(https://image.tmdb.org/t/p/original/rPqCxVXBD89jeWMgJU3MeFA6GDV.jpg, Total Recall Die totale Erinnerung-861-1990-backdrop.jpg)
---------- ID: 408340
        Title: Total Recall
  ReleaseDate:

---------- ID: 180
        Title: Minority Report
  ReleaseDate: 2002-06-20

### download(https://image.tmdb.org/t/p/original/9niGbmFeaR27pu7cPuQQrStkLlt.jpg, Minority Report-180-2002-poster.jpg)
### download(https://image.tmdb.org/t/p/original/u8BvwuiiQ0uLFuXviKJU0cCHXIW.jpg, Minority Report-180-2002-backdrop.jpg)
```

now we have:

```
$ ls -l
total 10324
-rw-r--r-- 1 wwwutz wwwutz    1072 May  8 15:11 LICENSE
-rw-r--r-- 1 wwwutz wwwutz  228686 May 16 14:52 Minority Report-180-2002-backdrop.jpg
-rw-r--r-- 1 wwwutz wwwutz  193067 May 16 14:52 Minority Report-180-2002-poster.jpg
-rw-r--r-- 1 wwwutz wwwutz      62 May 16 14:52 Minority Report-180-2002.URL
-rw-r--r-- 1 wwwutz wwwutz    7260 May 16 14:52 README.md
-rw-r--r-- 1 wwwutz wwwutz  148477 May 16 14:52 Total Recall Die totale Erinnerung-861-1990-backdrop.jpg
-rw-r--r-- 1 wwwutz wwwutz  134386 May 16 14:52 Total Recall Die totale Erinnerung-861-1990-poster.jpg
-rw-r--r-- 1 wwwutz wwwutz      62 May 16 14:52 Total Recall Die totale Erinnerung-861-1990.URL
-rw-r--r-- 1 wwwutz wwwutz      65 May 16 14:52 Total Recall-408340-0000.URL
-rw-r--r-- 1 wwwutz wwwutz  328461 May 16 14:52 Total Recall-64635-2012-backdrop.jpg
-rw-r--r-- 1 wwwutz wwwutz  466321 May 16 14:52 Total Recall-64635-2012-poster.jpg
-rw-r--r-- 1 wwwutz wwwutz      64 May 16 14:52 Total Recall-64635-2012.URL
-rw-r--r-- 1 wwwutz wwwutz   10458 May 16 14:50 main.go
-rwxr-xr-x 1 wwwutz wwwutz 9015305 May 16 14:50 movieinfo
```

### sort out the mess

And now we can start sorting out. We're looking for the Arnie version.
1. `Minority Report` ...
1. `Total Recall-408340-0000.URL` points to ... well whatever
1. `Total Recall-64635-2012*` looks like the remake
1. `Total Recall Die totale Erinnerung-861-1990*` Yeah! That's the one. Go Arnie, pull that thing out of your nose

"Hmmm.. german ? Ah, that's what you mean with 'works for me' 8-)"

Let's get rid of the others (using filename completion via the tabulator key `[TAB]` or use a decent file manager.

Oh and try not to delete your movie files ;-)

### restrict number of results

use the `--max=n` option. This limits the output to `n` results. The default is zero (0) which means unlimited. Let's try one (1).

```
$ ./movieinfo -m 1 "total recall"
     arg:  total recall
  search:  total recall
      id:  0
download:  false
    year:  0
     max:  1
 verbose:  false

---------- ID: 64635
        Title: Total Recall
  ReleaseDate: 2012-08-02
```

Which is, as we see, a bad thing. The Arnie version is not the first one to pop up. Anyways, I'm using 4 or 5 for scripts (`--max=5`), in 99% of all cases the movie I was looking for pops up in the first results.

### supply tmdb.org ID via options or filename

You know the tmdb.org ID of your movie ?

`-i <tmdb.org ID>` has highest priority and restricts output to this ID
```
$ ./movieinfo -i 64635
```
overrides `-ID-YYYY.URL` or `' - YYYY'` filename parsing


supplying a complete movie-info generated `.URL` file:

```
$ ./movieinfo "total recall-64635-2012.URL"
```

It'll try to do it's very best to extract the ID from a filename. But why should it ?

Well, it downloads and creates the most interesting `.txt` file only it you either supply the ID or there is only one result fopr your search.
