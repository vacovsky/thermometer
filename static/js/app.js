"use strict";

function CtoF(n) {
  return n * (9 / 5) + 32;
}

function getLastValue(lst) {
  return lst[lst.length - 1];
}

function populateTemplate(data) {}

var CHART_OPTIONS = {
  scales: {
    xAxes: [
      {
        type: "time",
        time: {
          unit: "day"
        }
      }
    ]
  }
};


function loadTempChart() {
  var Http = new XMLHttpRequest();
  var url = "/tempchartdata";
  Http.open("GET", url);
  Http.responseType = "json";
  Http.setRequestHeader("Content-Type", "application/json");
  Http.send();
  Http.onreadystatechange = e => {
    var chartData = Http.response;
    if (chartData != undefined) {
      //      populateTemplate();
      var config = {
        type: "line",
        data: {
          labels: chartData.Labels,
          datasets: [
            {
              label: "Ambient Air Temperature °F",
              fill: false,
              borderColor: window.chartColors.red,
              data: chartData.Data[0]
            },
            {
              label: "Ambient Humidity %",
              fill: false,
              borderColor: window.chartColors.orange,
              data: chartData.Data[1]
            }
          ]
        },
        options: {
          legend: {
            labels: {
              fontColor: "#bbb"
            }
          },
          elements: {
            line: {
              borderWidth: 2
            },

            point: {
              radius: 0
            }
          },
          aspectRatio: 2,
          responsive: true,
          title: {
            display: true,
            text: "Environmental"
          },
          tooltips: {
            mode: "index",
            intersect: false
          },
          hover: {
            mode: "nearest",
            intersect: true
          },
          scales: {
            xAxes: [
              {
                type: "time",
                display: true,
                ticks: {
                  fontColor: "#bbb"
                },
                scaleLabel: {
                  display: true
                }
              }
            ],
            yAxes: [
              {
                display: true,
                scaleLabel: {
                  display: true,
                  labelString: "°F / %"
                },
                ticks: {
                  beginAtZero: true,
                  min: 0,
                  max: 100,
                  fontColor: "#bbb",
                  stepSize: 20
                }
              }
            ]
          }
        }
      };

      var ctx = document.getElementById("temp-chart").getContext("2d");
      new Chart(ctx, config);

      weld(
        document.getElementById("#last-air-temp"),
        getLastValue(chartData.Data[0]) + " °F"
      );

      weld(
        document.getElementById("#last-humidity"),
        getLastValue(chartData.Data[1]) + " %"
      );
    }
  };
}

window.chartColors = {
  red: "rgb(255, 99, 132)",
  orange: "rgb(255, 159, 64)",
  yellow: "rgb(255, 205, 86)",
  green: "rgb(75, 192, 192)",
  blue: "rgb(54, 162, 235)",
  purple: "rgb(153, 102, 255)",
  grey: "rgb(201, 203, 207)"
};

function looper() {
  setTimeout(function() {
    loadTempChart();
    looper();
    console.log("Looping");
  }, 300000);
}

loadTempChart();
looper();
