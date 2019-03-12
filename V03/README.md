# CISCO DIGESTION SYSTEM

This is an excersize in a minimalized, *GoLang* server apps.   The objective was to deliver one (and only one) Executable file that has ZERO dependancies.  Put another way... a statically linked binary that fully contains the application.

## PLATFORM
The platform is *GoLang*. 
The DB is *BoltDB*.  
https://godoc.org/github.com/boltdb/bolt

the BoltDB is an *embedded JSON store*.  I choose Bolt for these reasons.

1. No Runtime Dependancy with Deliverable.
2. Bare minimum features so as to maximize use of resources. 
..* no need to authentication
..* no need to have more than one Writer.
..* emphasize reads.
..* no need to sophisticated indexing.
3. Would generate a separate BoltDB file that can be saved.
4. Perfect Time series solution

### TODO
* Dockerize.  Get to run in an alpne os. 
* ADD gRPC
* Create Test suite.
* Rearrange the Dir structure so the Source code is out of the way

* Possibly show this app running on a Ubuntu CORE OS. Distribute the App from a remote server automatically.
* Possibly demonstrate app on Amazon Core OS.  Just to see what it does.

* Add a Web portal. Probably vue.js ( another class ) 

* Improve the auto setup routine

### Alternative Systems
There are alternative systems using much of the same logic.  
One system listens for Messages from the CISCO call center directly thus does not read logs.


### MISSION
This app will consume Log data from a CISCO Call center,  import them into a BOLTDB and export them in a number of other formats that are simpler to understand, including a spread sheet format.
The app is designed to reside at the local call center NOC.  

There is another version of this system that will replicate the digested data to a cloud DataStore where multiple call centers can be consolidated.

This is part of my strategy to create an enterprice wide dash board and 


#### First time Run
The first time the app is run, sub directories will be created and a message will display on what to do next.  This is just to configure the system 

#### Second Run
At this point the system is configured at a Zero State, meaning there is no data.
The messages would have instructed you to copy all of the Log files into the /KKIn directory and run the app a second  time.

When the app finishes its run.
1. The KKOUT directory will be populated with a CONSOLIDATED JSON file for each Data type.
2. A KKErrors.xlsx file of errors will be created for inspection in an Excel Spread sheet.  This file should be migrated to a cloud repo for all managers to review.
3. The Data will be appended to the BOLTDB.  

The KKOUT and XLSX files are only what was added in the KKIn directoy. KKIn should be changed every X time period. 
The BoltDB will accumulate all records starting from the first run. The only way to reset the BoltDB is to physically delete it.  A new run of the app will just recreate the BoltDB and begin accumulation again.


## SAMPLE DATA
<root>/KKIN_cample_data has enough records to test the application.



# GOLANG Libraries.

The only non standard libraries in this system are...
go get github.com/360EntSecGroup-Skylar/excelize

go get -u github.com/asdine/storm

The compile line is...
go run main.go lib01.go lib02.go lib03.go lib10.go struct01.go


