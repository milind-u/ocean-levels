"use strict";

// There are too many elevation points so will only add every ELEVATION_INCR points
const ELEVATION_INCR = 50;
const ELEVATIONS_PATH = "setup/elevations.json";
const CIRCLE_COLOR = "#0000FF";
// Value of a change of one degree Celsius in Fahrenheit.
const C_TO_F = 9.0 / 5;
// Number of significant figures for floats
const FLOAT_PRECISION = 4;
const CURRENT_YEAR = new Date().getFullYear();
const MIN_YEAR = CURRENT_YEAR + 1;
// Start year of the sea level rise function
const START_YEAR = 2010;
//  Equilibrium time lag for sea level rise function
const EQ_TIME_LAG = 500;

/**
 * Different inputs taken in index.html, mapped to their id's
 * @enum {string}
 */
const Input = Object.freeze({
  YEAR: "year",
  DELTA_T_FAHRENHEIT: "delta-t-fahrenheit",
  DELTA_T_CELSIUS: "delta-t-celsius"
});

let map = null;
const locations = [];
let elevations = null;
fetchElevations();

function initYearInput() {
  const year = document.getElementById(Input.YEAR);
  year.min = MIN_YEAR;
  year.value = CURRENT_YEAR + 10;
  document.getElementById("legend-year").innerHTML = year.value;
}

function fetchElevations() {
  $.getJSON(ELEVATIONS_PATH, json => {
    elevations = Array.from(json);
    console.log(elevations.length, " elevations");
    if (map != null) {
      addCircles();
    }
  });
}

function initMap() {
  const mapElem = document.getElementById("map");
  map = new google.maps.Map(mapElem, {
    center: {lat: 0, lng: 0},
    zoom: 2,
    fullscreenControl: false
  });
  if (elevations != null) {
    addCircles();
  }
  const legend = document.getElementById("legend");
  map.controls[google.maps.ControlPosition.TOP_RIGHT].push(legend);
}

function addCircles() {
  for (let i = 0; i < elevations.length; i += ELEVATION_INCR) {
    const circle = new google.maps.Circle({
      center: elevations[i]["location"],
      radius: 1,
      editable: false,
      fillColor: CIRCLE_COLOR,
      strokeColor: CIRCLE_COLOR,
      fillOpacity: 0,
      strokeOpacity: 0,
      map
    });
    const location = {circle: circle, elevation: elevations[i]["elevation"]};
    displayLocationIfSubmerged(0, location);
    locations.push(location);
  }
}

function displaySubmergedLocations(input) {
  const result = roundInputToPrecision(input);

  console.log(input);
  let year, deltaSeaLevel, deltaT, deltaF;
  if (input === Input.YEAR) {
    deltaT = roundInputToPrecision(Input.DELTA_T_CELSIUS);
    deltaF = roundInputToPrecision(Input.DELTA_T_FAHRENHEIT);
  } else if (input === Input.DELTA_T_FAHRENHEIT) {
    deltaF = result;
    deltaT = deltaFToDeltaT(deltaF);
  } else { // DELTA_T_CELSIUS
    deltaT = result;
    deltaF = deltaTToDeltaF(deltaT);
  }
  year = Math.max(parseInt(document.getElementById(Input.YEAR).value), MIN_YEAR);
  deltaSeaLevel = getDeltaSeaLevel(deltaT, year);

  deltaSeaLevel = roundToPrecision(deltaSeaLevel);
  deltaT = roundToPrecision(deltaT);
  deltaF = roundToPrecision(deltaF);

  document.getElementById(Input.YEAR).value = year;
  document.getElementById(Input.DELTA_T_CELSIUS).value = deltaT;
  document.getElementById(Input.DELTA_T_FAHRENHEIT).value = deltaF;

  document.getElementById("delta-sea-level").innerHTML = deltaSeaLevel;
  document.getElementById("legend-year").innerHTML = year;

  for (const location of locations) {
    displayLocationIfSubmerged(deltaSeaLevel, location);
  }
}

function roundInputToPrecision(input) {
  return roundToPrecision(parseFloat(document.getElementById(input).value));
}

function roundToPrecision(f) {
  return parseFloat(f.toPrecision(FLOAT_PRECISION));
}

/**
 * Converts the change in temperature from fahrenheit to celsius
 * @param deltaF Change in temperature in fahrenheit
 * @return {number} Change in temperature in celsius
 */
function deltaFToDeltaT(deltaF) {
  return deltaF / C_TO_F;
}

/**
 * Converts the change in temperature from celsius to fahrenheit
 * @param deltaT Change in temperature in celsius
 * @return {number} Change in temperature in fahrenheit
 */
function deltaTToDeltaF(deltaT) {
  return deltaT * C_TO_F;
}

/**
 * Returns the change in sea level (m) given a certain change in temperature by a certain year (degrees Celsius).
 * Got this function from http://www.roperld.com/science/sealevelvstemperature.htm.
 * @param deltaT the change in temperature, in degrees Celsius
 * @param year the year by which the temperature increase occurs
 * @return {number} the change in sea level, in meters
 */
function getDeltaSeaLevel(deltaT, year) {
  let deltaSeaLevel = 0;
  const years_since_start = year - START_YEAR;
  const years_since_now = year - CURRENT_YEAR;
  const deltaTI = deltaT / years_since_now;

  for (let i = CURRENT_YEAR - START_YEAR; i < years_since_start; i++) {
    deltaSeaLevel += deltaTI * ((0.54 * (deltaTI * deltaTI)) + (0.39 * deltaTI) + 7.7) *
        Math.tanh((years_since_start - i) / (2 * EQ_TIME_LAG));
  }

  return deltaSeaLevel;
}

function displayLocationIfSubmerged(seaLevelRise, location) {
  // display all coasts that would be submerged given the sea level rise, and are not already below sea level
  const opacity = ((location.elevation < seaLevelRise) && (location.elevation >= 0) ? 1 : 0);
  location.circle.setOptions({fillOpacity: opacity, strokeOpacity: opacity});
}
