# ocean-levels

Web app for climate change that will show the coasts in the world that would be submerged if global average ocean levels or temperature were to rise by a certain amount.

You can view the webapp at https://storage.googleapis.com/html_files_bucket/OceanLevels/index.html

# *.html, *.js, *.css, icon.jpg
Files for the webapp. If you intended to run this locally, you would have to change the path of ```elevations.json``` in ```index.html```

# setup
Golang code for parsing world coastline data and integrating it with elevation data fetched from Google Maps Elevation API, along with the coastline data itself.
Go program would not run right now because is not in proper Go directory structure.

# Google API Key
All instances of ```<API_KEY>``` in this repo must be replaced with your API key if you intend to run the code locally.
To get one, you must create a Google Cloud Project and enable billing and the necessary APIs (Google Maps JavaScript and Elevation), and then create an API Key in the Credentials section.<br>
```<API_KEY``` shows up in ```index.html``` and ```setup/coasts/elevations.go```
