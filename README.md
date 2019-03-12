# NON SECRET GOLANG EXAMPLES
don't share as it includes keys and credentials



### V01  -GOLANG FIB CHALLENGE
a Fibanaci series solution.
Except I am doing it 3 different ways including..

1. Standard Loop Technique
2. Recursive Technique
3. GoLang Channels.
nothing secret



### V02  -GOLANG gRPC/ProtoBuf example building BLOCKCHAIN.
This is a GoLang Project. No Other Dependancies 
Example demonstrating.

1. GoLang 
2. ProtoBuf and gRPC in GO.
3. Very simple, BLOCKCHAIN solution.

This example has several purposes.
1. I have a project where I am replacing RESTful services with gRPC/PROTOBUF.  This was an experiment.
2. The other project is suppose to demonstrate an extremely fast Blockchain DB that can reside on an Embedded device.
3. I have clients that have large Public Web Exposure. IIS, LAMP and other web services. I do not believe that these Public facing portals can be adequetly protected from Hackers.  This is part of my strategy to demonstrate how I use GoLang,running as a servlet, as a middle wair system that has far fewer exploits than traditional web services.  
I am still working on instructions for how to compile the solution. 

### V03
*GOLANG  - BOLT.db*
This was a real prototype to scan error messages from a cisco call center system for FPL ( Florida Power and Light ).  3rd party systems, such as Call recording, were also being scanned.
The incoming data was usually in XML format.  This system created microservice executables that would digest these xml error files, convert to JSON, load into a BOLT db, and synch to cloud Google App Engine.  In addition, an excel report was automatically generated. Future versions would PUSH the Excel report to a git repo that managers could view.

### TICTOCK
*GOLANG - Go ROUTINES, CHANNELS, time package*
An experiment with golang time package as well as 
go routines and Channels.  A couple strategies. 

### V10
A Dockerized GoLang REST API with Postgres using Docker-Copose 
1. This project features goLang on a Scratch container.
2. A streamline means of developing golang on cloud.
The purpose of this project is to migrate a legacy system that was designed for SQL.  As oppose to trying to figure out how the dumps are working, this solution aims to just reproduce the Legacy system.
To be deployed to a Kubernetes or onprem system.
This project is missing some key parts..
Testing.

### GoLang Firebase
* The keys are fake *
A project that will use the Google Admin SDK to write and read records to the Firebase Cloud Platform.
TODO:  GoLang does not have a listener like Node.js so it will not respond to Add, Delete or update events.  But Firebase does have a Node.js listener that could relay these events to GoLang Procedures. This is not well documented, but I should add tis as an example.


#   projects to be added once I have them put together....

### F01
A flutter / Dart / gRPC solution ( Probably V02 client ).
Discuss the routine of compiling a Flutter ProtoFile.
< STAY TUNE >

### PTT02
*GOLANG - *
An excersize in importing GTFS Logisitcs data and doing analysis.
this project to be added latter as the deliverable
GTFS files are too large for GIT.  ( > 100 Megs)
Stay Tune.

