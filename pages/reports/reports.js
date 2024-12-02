const baseUrl = "http://localhost:8080/reports/";

document.addEventListener("DOMContentLoaded", showReports());

function fetchMomentReport() {
  makeRequest(
    baseUrl + "moment",
    "GET",
    "",
    "application/json",
    true,
    makeMomentReport,
    failedGet
  );
}

function makeMomentReport(momentData) {
  const chart = document.getElementById("momentChart");
  console.log(momentData);
  new Chart(chart, {
    type: "bar",
    data: {
      labels: momentData.map((data) => data.moment),
      datasets: [
        {
          label: "Recipes by moment",
          data: momentData.map((data) => data.count),
          backgroundColor: [
            "rgba(42, 96, 73, 0.2)",
            "rgba(42, 96, 73, 0.2)",
            "rgba(42, 96, 73, 0.2)",
            "rgba(42, 96, 73, 0.2)",
          ],
          borderColor: [
            "rgba(42, 96, 73)",
            "rgba(42, 96, 73)",
            "rgba(42, 96, 73)",
            "rgba(42, 96, 73)",
          ],
          borderWidth: 1,
        },
      ],
    },
    options: {
      scales: {
        y: {
          beginAtZero: true,
        },
      },
    },
  });
}

function fetchFoodTypeReport() {}
function fetchCostReport() {}

function showReports() {
  fetchMomentReport();
  fetchFoodTypeReport();
  fetchCostReport();
}

function failedGet() {
  console.log("Failed to generate reports");
}
