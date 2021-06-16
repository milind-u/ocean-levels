# main.go
Main of the Go program for parsing and integrating the data from the coastlines and Google Maps Elevation API.

# cln
One of my Go packages that contains collections. The only collection used in this project is ```TimeQueue```, which is used to ensure that requests are not being sent too fast too the Google Maps Elevation API, since it has a query frequency limit.

# util
Another one of my Go packages, that has various utility functions. The only file used in this project is ```util/util.go```.

# coasts
Go package specific to this project containing the code that parses and integrates that data of world coastlines and elevations.
Uses the Go package ```go-shp``` for the parsing ESRI Shapefiles in ```coastlines```.

# elevations.json
The integrated data of coordinates for world coastlines, and their elevations.<br>
Ouput of the Go program.

# coastlines
Dataset containing the coordinates of all the coastlines in the world, 
obtained from <a href="http://www.naturalearthdata.com/downloads/10m-physical-vectors/10m-coastline/" target="_blank">Natural Earth</a>

# Running the program
If you intend to run this setup program, you would have to move the directories ```coasts```, ```util```, and ```cln``` to ```$GOPATH/src```, and you would have to run the command ```go get github.com/jonas-p/go-shp``` to install the shapefile parsing library. You would also have to replace ```<API_KEY>``` in ```coasts/elevations.go``` with your Google Cloud Platform API Key that has Google Maps Elevation API enabled.
