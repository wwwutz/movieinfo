# movieinfo
get ridiculous least information from tmdb

The idea is to create a text file with movie meta information and download a poster and backdrop image into your current directory. The created text file has no distractive formatting and can be fed to `grep` or maybe  `find /movies -type f -name '*.txt' -print0 | xargs -0 grep -i schwarzenegger` . Which is a good thing.


```
$ movieinfo -d "total recall"
```
et voila
```
Imagining 'Total Recall'-735363-2001.URL
Minority Report-180-2002.URL
Minority Report-180-backdrop.jpg
Minority Report-180.jpg
Total Recall 2070 Machine Dreams-933701-1999.URL
Total Recall 2070 Machine Dreams-933701.jpg
Total Recall-597892-1987.URL
Total Recall-597892.jpg
Total Recall-64635-2012.URL
Total Recall-64635-backdrop.jpg
Total Recall-64635.jpg
Total Recall Die totale Erinnerung-861-1990.URL
Total Recall Die totale Erinnerung-861-backdrop.jpg
Total Recall Die totale Erinnerung-861.jpg
Total Recall(ed)-835797-2021.URL
Total Recall(ed)-835797.jpg

```
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
* `GOOS=windows GOARCH=amd64 go build` to create a windows `.exe`

# usage
Set up you API key. Either save it as environment variable `TMDB_API` or supply it on the commandline.

```
NAME:
   movieinfo - query tmdb.org to download backdrops, cover and more

USAGE:
   movieinfo [movie]

VERSION:
   0.4

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --download, -d               download images and metadata
   --verbose, --vv              whatever
   --year value, -y value       year (default: 0)
   --max value, -m value        max. entries, 0 unlimited (default: 0)
   --id value, -i value         tmdb movie ID (default: 0)
   --TMDB_API value             tmdb.org API key [$TMDB_API]
   --mvtoext value, --mv value  rename files to filename with this extension
   --removeartefacts, --rma     removes files with mID
   --help, -h                   show help
   --version, -v                print the version

```

### get info for "Total Recall"
 We'll find four versions of "Total Recall". The third ID 180 is obviously wrong.

```
$ movieinfo "total recall"
 arg[0]: total recall
 search: total recall
/*  1. ID: 861   */
        Title: Total Recall - Die totale Erinnerung
OriginalTitle: Total Recall
  ReleaseDate: 1990-06-01

/*  2. ID: 64635 */
        Title: Total Recall
  ReleaseDate: 2012-08-02

/*  3. ID: 597892 */
        Title: Total Recall
  ReleaseDate: 1987-01-01

/*  4. ID: 735363 */
        Title: Imagining 'Total Recall'
  ReleaseDate: 2001-01-01

/*  5. ID: 835797 */
        Title: Total Recall(ed)
  ReleaseDate: 2021-04-06

/*  6. ID: 933701 */
        Title: Total Recall 2070: Machine Dreams
  ReleaseDate: 1999-01-05

/*  7. ID: 180   */
        Title: Minority Report
  ReleaseDate: 2002-06-20


```

### download images for "Total Recall"

Let's just download all available poster and backdrop images and create a `.URL` file in the current directory.

before:
```
$ ls -l
total 0

```
calling `movieinfo` command with `--download` option:
```
$ movieinfo -d "total recall"

 arg[0]: total recall
 search: total recall
/*  1. ID: 861   */
        Title: Total Recall - Die totale Erinnerung
OriginalTitle: Total Recall
  ReleaseDate: 1990-06-01

 URL: https://image.tmdb.org/t/p/original/fzXg7ysj9g3U3J5wOgE8KI3gtuH.jpg
file: Total Recall Die totale Erinnerung-861.jpg
 URL: https://image.tmdb.org/t/p/original/8ELKVrGIrTjkkULisVAa8qeDsTG.jpg
file: Total Recall Die totale Erinnerung-861-backdrop.jpg
/*  2. ID: 64635 */
        Title: Total Recall
  ReleaseDate: 2012-08-02

 URL: https://image.tmdb.org/t/p/original/6UUqnqTxDMBKeRx4CH4fBJFhhXf.jpg
file: Total Recall-64635.jpg
 URL: https://image.tmdb.org/t/p/original/uBHeAB2Ug9ELBzkMyls8CUjzn4i.jpg
file: Total Recall-64635-backdrop.jpg
/*  3. ID: 597892 */
        Title: Total Recall
  ReleaseDate: 1987-01-01

 URL: https://image.tmdb.org/t/p/original/3eVoUM8JA0fUpzGo065VNOa8m6k.jpg
file: Total Recall-597892.jpg
/*  4. ID: 735363 */
        Title: Imagining 'Total Recall'
  ReleaseDate: 2001-01-01

/*  5. ID: 835797 */
        Title: Total Recall(ed)
  ReleaseDate: 2021-04-06

 URL: https://image.tmdb.org/t/p/original/dOlb9WLetYE14cvHID8Oz8nFADu.jpg
file: Total Recall(ed)-835797.jpg
/*  6. ID: 933701 */
        Title: Total Recall 2070: Machine Dreams
  ReleaseDate: 1999-01-05

 URL: https://image.tmdb.org/t/p/original/yepxcOGRwchjR5cPJdRyouI9xbl.jpg
file: Total Recall 2070 Machine Dreams-933701.jpg
/*  7. ID: 180   */
        Title: Minority Report
  ReleaseDate: 2002-06-20

 URL: https://image.tmdb.org/t/p/original/gejqouERNdK59OneFBaOSHDUdUt.jpg
file: Minority Report-180.jpg
 URL: https://image.tmdb.org/t/p/original/qq4H9JfBKQ5DarMLI6lhUQjn9D7.jpg
file: Minority Report-180-backdrop.jpg

```

now we have:

```
$ ls -1
Imagining 'Total Recall'-735363-2001.URL
Minority Report-180-2002.URL
Minority Report-180-backdrop.jpg
Minority Report-180.jpg
Total Recall 2070 Machine Dreams-933701-1999.URL
Total Recall 2070 Machine Dreams-933701.jpg
Total Recall-597892-1987.URL
Total Recall-597892.jpg
Total Recall-64635-2012.URL
Total Recall-64635-backdrop.jpg
Total Recall-64635.jpg
Total Recall Die totale Erinnerung-861-1990.URL
Total Recall Die totale Erinnerung-861-backdrop.jpg
Total Recall Die totale Erinnerung-861.jpg
Total Recall(ed)-835797-2021.URL
Total Recall(ed)-835797.jpg

```

### sort out the mess

And now we can start sorting out. We're looking for the Arnie version via the `.URL` files
1. `Minority Report-180-2002.URL
` ...
1. `*-408340-*.URL
` points to ... well whatever
1. `Total Recall-64635-2012.URL
` looks like the remake
1. `Total Recall Die totale Erinnerung-861-1990.URL
` Yeah! That's the one. Go Arnie, pull that thing out of your nose

"Hmmm.. german ? Ah, that's what you meant with 'works for me' 8-)"

Let's get rid of the others (using filename completion via the tabulator key `[TAB]` or use a decent file manager.

Oh and try not to delete your movie files ;-)

### restrict number of results

use the `--max=n` option. This limits the output to `n` results. The default is zero (`0`) which means unlimited. Let's try one (`1`).

```
$ movieinfo -m 1 "total recall"
 arg[0]: total recall
 search: total recall
/*  1. ID: 861   */
        Title: Total Recall - Die totale Erinnerung
OriginalTitle: Total Recall
  ReleaseDate: 1990-06-01


```

If there would be more than one (`1`) hit, the result is random. The Arnie version might not be the first one to pop up. I'm using 4 or 5 for scripts (`--max=5`), in 99% of all cases the movie I was looking for pops up in the first results.

### supply tmdb.org ID via options

Once you have the tmdb.org ID of your movie you can restict the output and even get overview and cast in a `.txt` file

The command line option `-i <tmdb.org ID>` has highest priority and restricts output to this ID
```
$ movieinfo -d -i 64635
 arg[0]: total recall
 search: total recall
### START .txt
tmdbID:   64635
Title:    Total Recall
Tagline:  Was ist Wirklichkeit?
Release:  2012-08-02
Runtime:  1 h 58 min
Overview: Herzlich Willkommen bei Rekall, der Firma, die ihre Träume dank 
ihres Total Recall Verfahrens in reale Erinnerungen verwandeln kann. Für den 
Fabrikarbeiter Douglas Quaid klingt der Urlaub, den man nur in seiner Phantasie 
macht, nach der perfekten Lösung, seinem frustrierenden Leben zu entgehen. 
Obwohl er eine wundervolle Frau hat, die er liebt, könnten die Erinnerungen an 
ein Leben als Super-Agent genau das sein, was Douglas gerade braucht. Bei der 
Gedankenbefüllung geht jedoch etwas schief und Quaid wird plötzlich zu einem 
gejagten Mann. Auf der Flucht vor der Polizei verbündet sich Quaid mit einer 
Rebellenkämpferin. Zusammen wollen sie den Anführer jener 
Untergrundorganisation finden und Cohaagen aufhalten. Die Linie zwischen 
Fantasie und Realität beginnt zu verschwimmen und das Schicksal seiner Welt 
hängt am seidenen Faden, als Quaid herausfindet, was seine wahre Identität, 
seine wahre Liebe und sein wahres Schicksal ist.

-  1. Colin Farrell: Douglas Quaid / Hauser
-  2. Jessica Biel: Melina
-  3. Kate Beckinsale: Lori Quaid
-  4. Ethan Hawke: Carl Hauser (director’s cut)
-  5. Bill Nighy: Matthias
-  6. John Cho: McClane
-  7. Bryan Cranston: Cohaagen
-  8. Dylan Smith: Hammond
-  9. Cam Clarke: Terminal Announcer (voice)
- 10. Bokeem Woodbine: Harry
- 11. Will Yun Lee: Marek
- 12. Natalie Lisinska: Bohemian Nurse
- 13. Milton Barnes: Resistance Fighter
- 14. Stephen MacDonald: Slacker
- 15. James McGowan: Military Adjutant
- 16. Michael Therriault: Bank Clerk
- 17. Mishael Morgan: Rekall Receptionist
- 18. Linlyn Lue: Resistance Woman
- 19. Andrew Moodie: Factory Foreman
- 20. Kaitlyn Leeb: Three-Breasted Woman
- 21. Leo Guiyab: Hauser Cover Identities
- 22. Nykeem Provo: Hauser Cover Identities
- 23. Steve Byers: Henry Reed
- 24. Danny Waugh: Officer
- 25. Geoffrey Pounsett: Sentry Lieutentant
- 26. Jesse Bond: Lead Federal Police
- 27. Warren Belle: Security Sentry
- 28. Vincent Rother: Sentry
- 29. Matthew Nette: Sentry
- 30. Brooks Darnell: Stevens
- 31. Brett Donahue: Sentry Trooper
- 32. James Downing: Synth Captain
- 33. Simon Sinn: Murray
- 34. Lisa Chandler: Prostitute
- 35. Miranda Jade: Girl on Balcony
- 36. Shereen Airth: Red-Headed Lady
- 37. Philip Moran: Immigration Officer
- 38. Leigh Folsom Boyd: The Fall Announcer (voice)
- 39. Emily C. Chang: Newscaster
- 40. Clive Ashborn: Newscaster
- 41. Bill Coulter: Newscaster
- 42. Merella Fernandez: Newscaster
- 43. Alicia-Kay Markson: Newscaster
- 44. Brian Rodriguez: Newscaster
- 45. Bridget Hoffman: Chopper (voice)
- 46. Brian T. Delaney: ATC Dispatcher (voice)
- 47. J.J. Perry: Colony Police Officer (uncredited)

/  1. 52fe46e1c3a368484e0a8fb5 Colin Farrell
/  2. 52fe46e1c3a368484e0a8fc5 Jessica Biel
/  3. 52fe46e1c3a368484e0a8fbd Kate Beckinsale
/  4. 63d56861c15b550079fa8bd1 Ethan Hawke
/  5. 52fe46e1c3a368484e0a8fc1 Bill Nighy
/  6. 52fe46e1c3a368484e0a8fc9 John Cho
/  7. 52fe46e1c3a368484e0a8fb9 Bryan Cranston
/  8. 5e45792adb8a0000129b449a Dylan Smith
/  9. 5e457b9a3dd126001a5b8093 Cam Clarke
/ 10. 52fe46e1c3a368484e0a8fcd Bokeem Woodbine
/ 11. 52fe46e1c3a368484e0a8fe9 Will Yun Lee
/ 12. 52fe46e1c3a368484e0a905d Natalie Lisinska
/ 13. 52fe46e1c3a368484e0a9059 Milton Barnes
/ 14. 56aa6131c3a36872db006896 Stephen MacDonald
/ 15. 56aa5f96c3a36872e1007072 James McGowan
/ 16. 52fe46e1c3a368484e0a9051 Michael Therriault
/ 17. 56aa604892514154750055f8 Mishael Morgan
/ 18. 56aa645e92514159a5001f2f Linlyn Lue
/ 19. 56e5cc0fc3a3685aa0007ef9 Andrew Moodie
/ 20. 5700025092514167830029b0 Kaitlyn Leeb
/ 21. 5e45794941465c0018d3c1dc Leo Guiyab
/ 22. 5e45795bdb8a0000149b32c0 Nykeem Provo
/ 23. 52fe46e1c3a368484e0a8fed Steve Byers
/ 24. 593c51dcc3a3680f420123c0 Danny Waugh
/ 25. 5e457973db8a0000149b32fd Geoffrey Pounsett
/ 26. 52fe46e1c3a368484e0a8ff5 Jesse Bond
/ 27. 5c41f1fa9251416b02b2a1bd Warren Belle
/ 28. 5e457a589603310017f51bfe Vincent Rother
/ 29. 5e457a643dd12600185b5b2d Matthew Nette
/ 30. 52fe46e1c3a368484e0a8ff9 Brooks Darnell
/ 31. 5e457a7c9603310013f5bd23 Brett Donahue
/ 32. 5e457a8b41465c0016d3785c James Downing
/ 33. 5e457a989603310013f5bd65 Simon Sinn
/ 34. 52fe46e1c3a368484e0a9055 Lisa Chandler
/ 35. 5e457ab00c2710001384cadc Miranda Jade
/ 36. 5e457abf9603310013f5bda9 Shereen Airth
/ 37. 5e457b0a0c2710001a866ba5 Philip Moran
/ 38. 5e457b7141465c0016d37a11 Leigh Folsom Boyd
/ 39. 53313602c3a3686a780012b3 Emily C. Chang
/ 40. 5e457b1b3dd12600185b5c9a Clive Ashborn
/ 41. 5e457b3d3dd12600165bb278 Bill Coulter
/ 42. 5e457b4cdb8a0000149b35e4 Merella Fernandez
/ 43. 5e457b560c27100018858249 Alicia-Kay Markson
/ 44. 5e457b633dd12600185b5dc6 Brian Rodriguez
/ 45. 5e457ba7db8a0000149b3607 Bridget Hoffman
/ 46. 5e457b839603310017f51d25 Brian T. Delaney
/ 47. 604136dfd71fb4002b600d87 J.J. Perry
###  END  .txt
 URL: https://image.tmdb.org/t/p/original/6UUqnqTxDMBKeRx4CH4fBJFhhXf.jpg
file: Total Recall.jpg
 URL: https://image.tmdb.org/t/p/original/uBHeAB2Ug9ELBzkMyls8CUjzn4i.jpg
file: Total Recall-backdrop.jpg

$ ls -1
Total Recall-64635-2012.URL
Total Recall-backdrop.jpg
Total Recall.jpg
Total Recall.txt

```

### supply tmdb.org ID via filename.URL

Supplying a complete `movieinfo` generated `.URL` file will parse the filename for a potiential ID.

```
$ movieinfo "total recall-64635-2012.URL"
 arg[0]: total recall-64635-2012.URL
 search: total recall
### START .txt
tmdbID:   64635
Title:    Total Recall
Tagline:  Was ist Wirklichkeit?
Release:  2012-08-02
Runtime:  1 h 58 min
Overview: Herzlich Willkommen bei Rekall, der Firma, die ihre Träume dank ihres

-  1. Colin Farrell: Douglas Quaid / Hauser
-  2. Jessica Biel: Melina
-  3. Kate Beckinsale: Lori Quaid
-  4. Ethan Hawke: Carl Hauser (director’s cut)
-  5. Bill Nighy: Matthias
-  6. John Cho: McClane
-  7. Bryan Cranston: Cohaagen
-  8. Dylan Smith: Hammond
-  9. Cam Clarke: Terminal Announcer (voice)
- 10. Bokeem Woodbine: Harry
```
It'll try to do it's very best to extract the ID from a filename.

# The final result as a .txt file

Add `--download` to finally create a `.txt` file with all movie meta information we ( well ok, 'I' ) need

```
$ movieinfo -d -i 861
 arg[0]: 
 search: 
### START .txt
tmdbID:   861
Title:    Total Recall - Die totale Erinnerung
Tagline:  Mach Dich bereit für die Reise Deines Lebens
OTitle:   Total Recall
Release:  1990-06-01
Runtime:  1 h 53 min
Overview: In ferner Zukunft führt Bauarbeiter Douglas Quaid ein zufriedenes 
Leben mit seiner attraktiven Ehefrau Lori. Einzig seine immer wiederkehrenden 
Albträume vom Planeten Mars quälen ihn und so entschließt er sich zu einer 
virtuellen Reise auf den Roten Planeten. Doch bei der Erinnerungsimplantation 
geht etwas schief und Quaids Leben ändert sich radikal. Ist er wirklich 
derjenige, der er zu sein glaubt? Quaid begibt er sich auf die gefährliche 
Suche nach seiner wahren Identität.

-  1. Arnold Schwarzenegger: Douglas Quaid / Hauser
-  2. Rachel Ticotin: Melina
-  3. Sharon Stone: Lori
-  4. Ronny Cox: Vilos Cohaagen
-  5. Michael Ironside: Richter
-  6. Marshall Bell: George / Kuato
-  7. Mel Johnson Jr.: Benny
-  8. Michael Champion: Helm
-  9. Roy Brocksmith: Dr. Edgemar
- 10. Ray Baker: Bob McClane
- 11. Rosemary Dunsmore: Dr. Lull
- 12. David Knell: Ernie
- 13. Alexia Robinson: Tiffany
- 14. Dean Norris: Tony
- 15. Mark Carlton: Bartender
- 16. Debbie Lee Carrington: Thumbelina
- 17. Lycia Naff: Mary
- 18. Robert Costanzo: Harry
- 19. Michael LaGuardia: Stevens
- 20. Priscilla Allen: Fat Lady
- 21. Ken Strausbaugh: Immigration Officer
- 22. Marc Alaimo: Everett
- 23. Michael Gregory: Rebel Lieutenant
- 24. Ken Gildin: Hotel Clerk
- 25. Mickey Jones: Burly Miner
- 26. Parker Whitman: Martian Husband
- 27. Ellen Gollas: Martian Wife
- 28. Gloria Dorson: Woman in Phone Booth
- 29. Erika Carlsson: Miss Lonelyhearts
- 30. Benny Corral: Punk Cabbie
- 31. Bob Tzudiker: Doctor
- 32. Erik Cord: Lab Assistant
- 33. Frank Kopyc: Technician
- 34. Chuck Sloan: Scientist
- 35. Dave Nicolson: Scientist
- 36. Paula McClure: Newscaster
- 37. Rebecca Ruth: Reporter
- 38. Milt Tarver: Commercial Announcer
- 39. Roger Cudney: Agent
- 40. Monica Steuer: Mutant Mother
- 41. Sasha Rionda: Mutant Child
- 42. Linda Howell: Tennis Pro
- 43. Robert Picardo: Voice of Johnnycab (voice)
- 44. Kamala Lopez: Additional Voices (voice)
- 45. Morgan Lofting: Additional Voices (voice)
- 46. Patti Attar: Additional Voices (voice)
- 47. Bob Bergen: Additional Voices (voice)
- 48. Joe Unger: Additional Voices (voice)
- 49. Karlyn Michelson: Additional Voices (voice)
- 50. Joel Kramer: Harry's Henchman (uncredited)

/  1. 52fe4283c3a36847f8024e9b Arnold Schwarzenegger
/  2. 52fe4283c3a36847f8024ea3 Rachel Ticotin
/  3. 52fe4283c3a36847f8024e9f Sharon Stone
/  4. 52fe4283c3a36847f8024ea7 Ronny Cox
/  5. 52fe4283c3a36847f8024eab Michael Ironside
/  6. 52fe4283c3a36847f8024eaf Marshall Bell
/  7. 52fe4283c3a36847f8024eb3 Mel Johnson Jr.
/  8. 52fe4283c3a36847f8024eb7 Michael Champion
/  9. 52fe4283c3a36847f8024ebb Roy Brocksmith
/ 10. 52fe4283c3a36847f8024ebf Ray Baker
/ 11. 567b9b1f9251417de3001e50 Rosemary Dunsmore
/ 12. 578ce7fcc3a3682bb2011939 David Knell
/ 13. 578ce80ac3a3685b4400a435 Alexia Robinson
/ 14. 567b9b50c3a3684bcc001f93 Dean Norris
/ 15. 5e777625357c00001650a65b Mark Carlton
/ 16. 567b9bab9251417ddf001fb7 Debbie Lee Carrington
/ 17. 567b9bde9251417de5001e77 Lycia Naff
/ 18. 567b9c05c3a3684be3001e5f Robert Costanzo
/ 19. 5e77765bc8a2d4001720b224 Michael LaGuardia
/ 20. 5e777667d18b24001780229e Priscilla Allen
/ 21. 5e777676c8a2d4001920b2b9 Ken Strausbaugh
/ 22. 567b9c3ec3a3684bd3002322 Marc Alaimo
/ 23. 567b9c67c3a3684bcc001fcb Michael Gregory
/ 24. 5e7777512f3b170011528b85 Ken Gildin
/ 25. 567b9c9b9251417de9001f6f Mickey Jones
/ 26. 5e7777762f3b170011528b9f Parker Whitman
/ 27. 5e7777812f3b17001451fd46 Ellen Gollas
/ 28. 5e77778fcabfe400132366b0 Gloria Dorson
/ 29. 5e77779c2f3b170011528bb8 Erika Carlsson
/ 30. 5e7777b8357c00001151dba8 Benny Corral
/ 31. 5e77781f357c00001151dccf Bob Tzudiker
/ 32. 5e77783fb1f68d0012e56f83 Erik Cord
/ 33. 5e77787bd18b240017802a5e Frank Kopyc
/ 34. 5e7778cfc8a2d4001720b44e Chuck Sloan
/ 35. 5e7778dbd18b240013806605 Dave Nicolson
/ 36. 5e7778e8b1f68d0012e56ff2 Paula McClure
/ 37. 5e777905b1f68d0012e5701c Rebecca Ruth
/ 38. 5e777911cabfe400132367b1 Milt Tarver
/ 39. 5e77791da055ef00122e8a41 Roger Cudney
/ 40. 5e7779382f3b17001953560c Monica Steuer
/ 41. 5e777a67c8a2d4001520b898 Sasha Rionda
/ 42. 5e777b3ca055ef00122e8c2d Linda Howell
/ 43. 5e777b50c8a2d4001920b789 Robert Picardo
/ 44. 5e777b5ca055ef00142e6da4 Kamala Lopez
/ 45. 5e777b67d18b24001980242c Morgan Lofting
/ 46. 5e777b722f3b170019535b0d Patti Attar
/ 47. 5e777b80d18b240019802445 Bob Bergen
/ 48. 5e777ba42f3b170011528fa6 Joe Unger
/ 49. 5e777bb4357c00001151e17e Karlyn Michelson
/ 50. 602d2abb2cde98003e2dbe6a Joel Kramer
###  END  .txt
 URL: https://image.tmdb.org/t/p/original/fzXg7ysj9g3U3J5wOgE8KI3gtuH.jpg
file: Total Recall Die totale Erinnerung.jpg
 URL: https://image.tmdb.org/t/p/original/8ELKVrGIrTjkkULisVAa8qeDsTG.jpg
file: Total Recall Die totale Erinnerung-backdrop.jpg

$ ls -1
Total Recall Die totale Erinnerung-861-1990.URL
Total Recall Die totale Erinnerung-backdrop.jpg
Total Recall Die totale Erinnerung.jpg
Total Recall Die totale Erinnerung.txt

```

## unrelated feature: renaming files

`movieinfo --mvtoext` allows you to bulk rename files to the filename with the supplied extension.

Given the following files:
```

3v3n_uglier.meta
Nice-Name.txt
ugly_m0VieNam3.URL
ugly_m0VieNam3.avi
ugly_m0VieNam3.iso

```
Now let them all be called `Nice-Name.*`

```
$ movieinfo --mvtoext .txt 3v3n_uglier.meta Nice-Name.txt ugly_m0VieNam3.URL ugly_m0VieNam3.avi ugly_m0VieNam3.iso

 filenames[0]: 3v3n_uglier.meta
 filenames[1]: Nice-Name.txt
 filenames[2]: ugly_m0VieNam3.avi
 filenames[3]: ugly_m0VieNam3.iso
 filenames[4]: ugly_m0VieNam3.URL
 mv ugly_m0VieNam3.avi Nice-Name.avi
 mv ugly_m0VieNam3.iso Nice-Name.iso
 mv 3v3n_uglier.meta Nice-Name.meta
 mv ugly_m0VieNam3.URL Nice-Name.URL
 mv 3v3n_uglier.meta Nice-Name.meta

$ ls -1
Nice-Name.URL
Nice-Name.avi
Nice-Name.iso
Nice-Name.meta
Nice-Name.txt

```
Don't fall for it: It only uses the filenames you supply in the commandline:
```

$ movieinfo --mvtoext .txt *.{avi,iso}
 filenames[0]: ugly_m0VieNam3.avi
 filenames[1]: ugly_m0VieNam3.iso
# 2 != 1 : exit
 fail[1]:  no file with extension .txt supplied

```
Renaming one file:
```
$ movieinfo --mvtoext .txt ugly_m0VieNam3.avi Nice-Name.txt
 filenames[0]: ugly_m0VieNam3.avi
 filenames[1]: Nice-Name.txt
 mv ugly_m0VieNam3.avi Nice-Name.avi

$ ls -1
3v3n_uglier.meta
Nice-Name.avi
Nice-Name.txt
ugly_m0VieNam3.URL
ugly_m0VieNam3.iso

```
Remember: It does *not* search in your folder. You must not supply more or less than one filename with the supplied extension. Captain Obvious.
