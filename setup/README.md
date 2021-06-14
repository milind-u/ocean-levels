# main.go
Main of the Go program for parsing and integrating the data from the coastlines and Google Maps Elevation API.

# coasts
Go package containing the code that parses and integrates that data of world coastlines and elevations.
Uses the Go package ```go-shp``` for the parsing ESRI Shapefiles in ```coastlines```.
If you intend to run this code, you must run the command ```go get github.com/jonas-p/go-shp```

# elevations.json
The integrated data of coordinates for world coastlines, and their elevations.<br>
Ouput of the Go program.

# coastlines
Dataset containing the coordinates of all the coastlines in the world, 
obtained from <a href="http://www.naturalearthdata.com/downloads/10m-physical-vectors/10m-coastline/" target="_blank">Natural Earth</a>
