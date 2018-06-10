# movieinfo
get ridiculous least information from tmdb

The idea is to create a text file with movie meta information and download a poster and backdrop image into your current directory. The created text file has no distractive formatting and can be fed to `grep` or maybe  `find /movies -type f -name '*.txt' -print0 | xargs -0 grep -i schwarzenegger` . Which is a good thing.


```
$ movieinfo -d "total recall"
```
et voila
```
Minority Report-180-2002.URL
Minority Report-180-backdrop.jpg
Minority Report-180.jpg
Total Recall-408340-0000.URL
Total Recall-64635-2012.URL
Total Recall-64635-backdrop.jpg
Total Recall-64635.jpg
Total Recall Die totale Erinnerung-861-1990.URL
Total Recall Die totale Erinnerung-861-backdrop.jpg
Total Recall Die totale Erinnerung-861.jpg

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
* `GOOS=wiondows GOARCH=amd64 go build` to create a windows `.exe`

# usage
Set up you API key. Either save it as environment variable `TMDB_API` or supply it on the commandline.

```
NAME:
   movieinfo - query tmdb.org to download backdrops, cover and more

USAGE:
   movieinfo [movie]

VERSION:
   0.3

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

/*  3. ID: 408340 */
        Title: Total Recall
  ReleaseDate: 

/*  4. ID: 180   */
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

 URL: https://image.tmdb.org/t/p/original/unjJqoBkzdUIA5Bi1rDdVHo0949.jpg
file: Total Recall Die totale Erinnerung-861.jpg
 URL: https://image.tmdb.org/t/p/original/rPqCxVXBD89jeWMgJU3MeFA6GDV.jpg
file: Total Recall Die totale Erinnerung-861-backdrop.jpg
/*  2. ID: 64635 */
        Title: Total Recall
  ReleaseDate: 2012-08-02

 URL: https://image.tmdb.org/t/p/original/4zgwx4HySRVjqSlmbrEKetJr5qo.jpg
file: Total Recall-64635.jpg
 URL: https://image.tmdb.org/t/p/original/orFQbyZ6g7kPFaJXmgty0M88wJ0.jpg
file: Total Recall-64635-backdrop.jpg
/*  3. ID: 408340 */
        Title: Total Recall
  ReleaseDate: 

/*  4. ID: 180   */
        Title: Minority Report
  ReleaseDate: 2002-06-20

 URL: https://image.tmdb.org/t/p/original/9niGbmFeaR27pu7cPuQQrStkLlt.jpg
file: Minority Report-180.jpg
 URL: https://image.tmdb.org/t/p/original/u8BvwuiiQ0uLFuXviKJU0cCHXIW.jpg
file: Minority Report-180-backdrop.jpg

```

now we have:

```
$ ls -1
Minority Report-180-2002.URL
Minority Report-180-backdrop.jpg
Minority Report-180.jpg
Total Recall-408340-0000.URL
Total Recall-64635-2012.URL
Total Recall-64635-backdrop.jpg
Total Recall-64635.jpg
Total Recall Die totale Erinnerung-861-1990.URL
Total Recall Die totale Erinnerung-861-backdrop.jpg
Total Recall Die totale Erinnerung-861.jpg

```

### sort out the mess

And now we can start sorting out. We're looking for the Arnie version via the `.URL` files
1. `Minority Report-180-2002.URL
` ...
1. `Total Recall-408340-0000.URL
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

-  1. Colin Farrell: Doug Quaid/Carl Hauser
-  2. Kate Beckinsale: Lori Quaid
-  3. Jessica Biel: Melina
-  4. Bryan Cranston: Cohaagen
-  5. Bill Nighy: Matthias
-  6. John Cho: McClane
-  7. Bokeem Woodbine: Harry
-  8. Will Yun Lee: Marek
-  9. Steve Byers: Henry Reed
- 10. Currie Graham: Bergen
- 11. Jesse Bond: Lead Federal Police
- 12. Brooks Darnell: Stevens
- 13. Michael Therriault: Bank Clerk
- 14. Lisa Chandler: Prostitute
- 15. Milton Barnes: Resistance Fighter
- 16. Natalie Lisinska: Bohemian Nurse
- 17. Billy Choi: Street peddler
- 18. Emily Chang: Newscaster Lien Nguyen
- 19. James McGowan: Military Adjutant
- 20. Mishael Morgan: Rekall Receptionist
- 21. Stephen MacDonald: Slacker
- 22. Linlyn Lue: Resistance Woman
- 23. Andrew Moodie: Factory Foreman
- 24. Kaitlyn Wong: Three-Breasted Woman
- 25. Danny Waugh: Officer
- 26. Filip Watermann: Construction Worker (uncredited)

/  1. 52fe46e1c3a368484e0a8fb5 Colin Farrell
/  2. 52fe46e1c3a368484e0a8fbd Kate Beckinsale
/  3. 52fe46e1c3a368484e0a8fc5 Jessica Biel
/  4. 52fe46e1c3a368484e0a8fb9 Bryan Cranston
/  5. 52fe46e1c3a368484e0a8fc1 Bill Nighy
/  6. 52fe46e1c3a368484e0a8fc9 John Cho
/  7. 52fe46e1c3a368484e0a8fcd Bokeem Woodbine
/  8. 52fe46e1c3a368484e0a8fe9 Will Yun Lee
/  9. 52fe46e1c3a368484e0a8fed Steve Byers
/ 10. 52fe46e1c3a368484e0a8ff1 Currie Graham
/ 11. 52fe46e1c3a368484e0a8ff5 Jesse Bond
/ 12. 52fe46e1c3a368484e0a8ff9 Brooks Darnell
/ 13. 52fe46e1c3a368484e0a9051 Michael Therriault
/ 14. 52fe46e1c3a368484e0a9055 Lisa Chandler
/ 15. 52fe46e1c3a368484e0a9059 Milton Barnes
/ 16. 52fe46e1c3a368484e0a905d Natalie Lisinska
/ 17. 5322eb779251411f850049a7 Billy Choi
/ 18. 53313602c3a3686a780012b3 Emily Chang
/ 19. 56aa5f96c3a36872e1007072 James McGowan
/ 20. 56aa604892514154750055f8 Mishael Morgan
/ 21. 56aa6131c3a36872db006896 Stephen MacDonald
/ 22. 56aa645e92514159a5001f2f Linlyn Lue
/ 23. 56e5cc0fc3a3685aa0007ef9 Andrew Moodie
/ 24. 5700025092514167830029b0 Kaitlyn Wong
/ 25. 593c51dcc3a3680f420123c0 Danny Waugh
/ 26. 56db5cc6c3a3682dac0000d6 Filip Watermann
###  END  .txt
 URL: https://image.tmdb.org/t/p/original/4zgwx4HySRVjqSlmbrEKetJr5qo.jpg
file: Total Recall.jpg
 URL: https://image.tmdb.org/t/p/original/orFQbyZ6g7kPFaJXmgty0M88wJ0.jpg
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

-  1. Colin Farrell: Doug Quaid/Carl Hauser
-  2. Kate Beckinsale: Lori Quaid
-  3. Jessica Biel: Melina
-  4. Bryan Cranston: Cohaagen
-  5. Bill Nighy: Matthias
-  6. John Cho: McClane
-  7. Bokeem Woodbine: Harry
-  8. Will Yun Lee: Marek
-  9. Steve Byers: Henry Reed
- 10. Currie Graham: Bergen
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
Tagline:  Mach dich bereit für die Reise deines Lebens.
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

-  1. Arnold Schwarzenegger: Douglas Quaid/Hauser
-  2. Sharon Stone: Lori
-  3. Rachel Ticotin: Melina
-  4. Ronny Cox: Vilos Cohaagen
-  5. Michael Ironside: Richter
-  6. Marshall Bell: George/Kuato
-  7. Roy Brocksmith: Dr. Edgemar
-  8. Ray Baker: Bob McClane
-  9. Rosemary Dunsmore: Dr. Lull
- 10. Dean Norris: Tony
- 11. Debbie Lee Carrington: Thumbelina
- 12. Lycia Naff: Mary
- 13. Robert Costanzo: Harry
- 14. Marc Alaimo: Everett
- 15. Michael Gregory: Rebel Lieutenant
- 16. Mickey Jones: Burly Miner
- 17. Robert Picardo: Voice of Johnnycab (voice)
- 18. Michael Champion: Helm
- 19. Mel Johnson Jr.: Benny
- 20. David Knell: Ernie
- 21. Alexia Robinson: Tiffany

/  1. 52fe4283c3a36847f8024e9b Arnold Schwarzenegger
/  2. 52fe4283c3a36847f8024e9f Sharon Stone
/  3. 52fe4283c3a36847f8024ea3 Rachel Ticotin
/  4. 52fe4283c3a36847f8024ea7 Ronny Cox
/  5. 52fe4283c3a36847f8024eab Michael Ironside
/  6. 52fe4283c3a36847f8024eaf Marshall Bell
/  7. 52fe4283c3a36847f8024ebb Roy Brocksmith
/  8. 52fe4283c3a36847f8024ebf Ray Baker
/  9. 567b9b1f9251417de3001e50 Rosemary Dunsmore
/ 10. 567b9b50c3a3684bcc001f93 Dean Norris
/ 11. 567b9bab9251417ddf001fb7 Debbie Lee Carrington
/ 12. 567b9bde9251417de5001e77 Lycia Naff
/ 13. 567b9c05c3a3684be3001e5f Robert Costanzo
/ 14. 567b9c3ec3a3684bd3002322 Marc Alaimo
/ 15. 567b9c67c3a3684bcc001fcb Michael Gregory
/ 16. 567b9c9b9251417de9001f6f Mickey Jones
/ 17. 567b9cce9251417dda001e8d Robert Picardo
/ 18. 52fe4283c3a36847f8024eb7 Michael Champion
/ 19. 52fe4283c3a36847f8024eb3 Mel Johnson Jr.
/ 20. 578ce7fcc3a3682bb2011939 David Knell
/ 21. 578ce80ac3a3685b4400a435 Alexia Robinson
###  END  .txt
 URL: https://image.tmdb.org/t/p/original/unjJqoBkzdUIA5Bi1rDdVHo0949.jpg
file: Total Recall Die totale Erinnerung.jpg
 URL: https://image.tmdb.org/t/p/original/rPqCxVXBD89jeWMgJU3MeFA6GDV.jpg
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
ugly_m0VieNam3.avi
ugly_m0VieNam3.iso
ugly_m0VieNam3.URL

```
Now let them all be called `Nice-Name.*`

```
$ movieinfo --mvtoext .txt 3v3n_uglier.meta Nice-Name.txt ugly_m0VieNam3.avi ugly_m0VieNam3.iso ugly_m0VieNam3.URL

 filenames[0]: 3v3n_uglier.meta
 filenames[1]: Nice-Name.txt
 filenames[2]: ugly_m0VieNam3.avi
 filenames[3]: ugly_m0VieNam3.iso
 filenames[4]: ugly_m0VieNam3.URL
 mv 3v3n_uglier.meta Nice-Name.meta
 mv ugly_m0VieNam3.avi Nice-Name.avi
 mv ugly_m0VieNam3.iso Nice-Name.iso
 mv ugly_m0VieNam3.URL Nice-Name.URL

$ ls -1
Nice-Name.avi
Nice-Name.iso
Nice-Name.meta
Nice-Name.txt
Nice-Name.URL

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
ugly_m0VieNam3.iso
ugly_m0VieNam3.URL

```
Remember: It does *not* search in your folder. You must not supply more or less than one filename with the supplied extension. Captain Obvious.
