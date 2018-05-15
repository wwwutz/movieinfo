# movieinfo
get ridiculous least information from tmdb

# warning

Do **not** take this source as go/golang reference. I'm a `perl` person. This is my first try to do a little bit more with golang. Based on `"golang string to int"` or `"golang write file"` web searches.
The `movieinfo` creates files, so it might overwrite yours.
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
    args:  total recall
   title:  total recall
download:  false
   force:  false
    year:  0
     max:  0
---------- ID: 64635
OriginalTitle: Total Recall
        Title: Total Recall
  ReleaseDate: 2012-08-02
   PosterPath: /tWBo7aZk3I1dLxmMj7ZJcN8uke5.jpg
 BackdropPath: /orFQbyZ6g7kPFaJXmgty0M88wJ0.jpg
---------- ID: 861
OriginalTitle: Total Recall
        Title: Total Recall
  ReleaseDate: 1990-06-01
   PosterPath: /ikYpJ0AjGBNnAYFnPJDUVIOcduR.jpg
 BackdropPath: /rPqCxVXBD89jeWMgJU3MeFA6GDV.jpg
---------- ID: 408340
OriginalTitle: Total Recall
        Title: Total Recall
  ReleaseDate:
   PosterPath:
 BackdropPath:
---------- ID: 180
OriginalTitle: Minority Report
        Title: Minority Report
  ReleaseDate: 2002-06-20
   PosterPath: /h3lpltSn7Rj1eYTPQO1lYGdw4Bz.jpg
 BackdropPath: /u8BvwuiiQ0uLFuXviKJU0cCHXIW.jpg
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
    args:  total recall
   title:  total recall
download:  true
   force:  false
    year:  0
     max:  0
---------- ID: 64635
OriginalTitle: Total Recall
        Title: Total Recall
  ReleaseDate: 2012-08-02
   PosterPath: /tWBo7aZk3I1dLxmMj7ZJcN8uke5.jpg
 BackdropPath: /orFQbyZ6g7kPFaJXmgty0M88wJ0.jpg
### download(https://image.tmdb.org/t/p/original/tWBo7aZk3I1dLxmMj7ZJcN8uke5.jpg, total recall-64635-2012-08-02-poster.jpg)
### download(https://image.tmdb.org/t/p/original/orFQbyZ6g7kPFaJXmgty0M88wJ0.jpg, total recall-64635-2012-08-02-backdrop.jpg)
---------- ID: 861
OriginalTitle: Total Recall
        Title: Total Recall
  ReleaseDate: 1990-06-01
   PosterPath: /ikYpJ0AjGBNnAYFnPJDUVIOcduR.jpg
 BackdropPath: /rPqCxVXBD89jeWMgJU3MeFA6GDV.jpg
### download(https://image.tmdb.org/t/p/original/ikYpJ0AjGBNnAYFnPJDUVIOcduR.jpg, total recall-861-1990-06-01-poster.jpg)
### download(https://image.tmdb.org/t/p/original/rPqCxVXBD89jeWMgJU3MeFA6GDV.jpg, total recall-861-1990-06-01-backdrop.jpg)
---------- ID: 408340
OriginalTitle: Total Recall
        Title: Total Recall
  ReleaseDate:
   PosterPath:
 BackdropPath:
---------- ID: 180
OriginalTitle: Minority Report
        Title: Minority Report
  ReleaseDate: 2002-06-20
   PosterPath: /h3lpltSn7Rj1eYTPQO1lYGdw4Bz.jpg
 BackdropPath: /u8BvwuiiQ0uLFuXviKJU0cCHXIW.jpg
### download(https://image.tmdb.org/t/p/original/h3lpltSn7Rj1eYTPQO1lYGdw4Bz.jpg, total recall-180-2002-06-20-poster.jpg)
### download(https://image.tmdb.org/t/p/original/u8BvwuiiQ0uLFuXviKJU0cCHXIW.jpg, total recall-180-2002-06-20-backdrop.jpg)
```

now we have:

```
$ ls -l
total 10216
-rw-r--r-- 1 wwwutz wwwutz    1072 May  8 15:11 LICENSE
-rw-r--r-- 1 wwwutz wwwutz    1955 May  9 11:39 README.md
-rw-r--r-- 1 wwwutz wwwutz    5734 May  9 12:08 main.go
-rwxr-xr-x 1 wwwutz wwwutz 8941192 May  9 12:08 movieinfo
-rw-r--r-- 1 wwwutz wwwutz  228686 May  9 12:10 total recall-180-2002-06-20-backdrop.jpg
-rw-r--r-- 1 wwwutz wwwutz  283166 May  9 12:10 total recall-180-2002-06-20-poster.jpg
-rw-r--r-- 1 wwwutz wwwutz      62 May  9 12:10 total recall-180.URL
-rw-r--r-- 1 wwwutz wwwutz      65 May  9 12:10 total recall-408340.URL
-rw-r--r-- 1 wwwutz wwwutz  328461 May  9 12:10 total recall-64635-2012-08-02-backdrop.jpg
-rw-r--r-- 1 wwwutz wwwutz  298246 May  9 12:10 total recall-64635-2012-08-02-poster.jpg
-rw-r--r-- 1 wwwutz wwwutz      64 May  9 12:10 total recall-64635.URL
-rw-r--r-- 1 wwwutz wwwutz  148477 May  9 12:10 total recall-861-1990-06-01-backdrop.jpg
-rw-r--r-- 1 wwwutz wwwutz  184394 May  9 12:10 total recall-861-1990-06-01-poster.jpg
-rw-r--r-- 1 wwwutz wwwutz      62 May  9 12:10 total recall-861.URL
```

### sort out the mess

And now we can start sorting out. We're looking for the Arnie version.
1. `recall-180-2002-*` points to "Minority Report"
1. `recall-408340*` points to ... well whatever
1. `recall-64635-2012-*` looks like the remake
1. `recall-861-*` Yeah! That's the one. Go Arnie, pull that thing out of your nose

Let's get rid of the others (using filename completion via the tabulator key `[TAB]`. Press it , don't type:

```
$ rm tot[TAB]
total recall-180-2002-06-20-backdrop.jpg    total recall-408340.URL                     total recall-64635.URL                      total recall-861.URL
total recall-180-2002-06-20-poster.jpg      total recall-64635-2012-08-02-backdrop.jpg  total recall-861-1990-06-01-backdrop.jpg    
total recall-180.URL                        total recall-64635-2012-08-02-poster.jpg    total recall-861-1990-06-01-poster.jpg      
$ rm total\ recall-180-[TAB]
total recall-180-2002-06-20-backdrop.jpg  total recall-180-2002-06-20-poster.jpg    
$ rm total\ recall-180-2002-06-20-*
$
```

and so on. Or use a file manager.

### restrict number of results

use the `--max=n` option. This limits the output to `n` results. The default is zero (0) which means unlimited. Let's try one (1).

```
$ ./movieinfo -m 1 "total recall"
    args:  total recall
   title:  total recall
download:  false
   force:  false
    year:  0
     max:  1
---------- ID: 64635
OriginalTitle: Total Recall
        Title: Total Recall
  ReleaseDate: 2012-08-02
   PosterPath: /tWBo7aZk3I1dLxmMj7ZJcN8uke5.jpg
 BackdropPath: /orFQbyZ6g7kPFaJXmgty0M88wJ0.jpg
```

Which is, as we see, a bad thing. The Arnie version is not the first one to pop up. Anyways, I'm using 4 or 5 for scripts (`--max=5`), in 99% of all cases the movie I was looking for pops up in the first results.

### supply tmdb.org ID via options or filename

restricting search to a know ID can be achieved either by
```
$ ./movieinfo -i 64635
```

or by supplying a complete movie-info generated `.URL` file:

```
$ ./movieinfo "total recall-64635-2012.URL"
```
