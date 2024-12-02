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
  const barChart = document.getElementById("momentBarChart");
  const pieChart = document.getElementById('momentPieChart')
  new Chart(barChart, {
    type: "bar",
    data: {
      labels: momentData.map((data) => data.moment),
      datasets: [
        {
          label: "Recipes by moment",
          data: momentData.map((data) => data.count),
          backgroundColor: [
            "rgba(42, 96, 73, 0.7)",
            "rgba(42, 96, 73, 0.7)",
            "rgba(42, 96, 73, 0.7)",
            "rgba(42, 96, 73, 0.7)",
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
      plugins: {
        title: {
          display: true,
          text: "Quantity of recipes by MOMENT",
          color: "#2a6049",
          font: {
            size: 18
          }
        },
      },
    },
  });

  new Chart(pieChart, {
    type: 'pie',
    data: {
      labels: momentData.map((data) => data.moment),
      datasets: [
        {
          label: 'Dataset 1',
          data: momentData.map((data) => data.count),
          backgroundColor: [
            'rgba(42, 96, 73)',
            'rgb(255, 99, 132)',
            'rgb(54, 162, 235)',
            'rgb(255, 205, 86)'
          ],
        }
      ]

    },
    options: {
      responsive: true,
      plugins: {
        legend: {
          position: 'top',
        },
        title: {
          display: true,
          text: 'Quantity of recipes by MOMENT',
          color: "#2a6049",
          font: {
            size: 18
          }
        }
      }
    },
  })
}

function fetchFoodTypeReport() {
  makeRequest(
    baseUrl + "foodtype",
    "GET",
    "",
    "application/json",
    true,
    makeFoodTypeReport,
    failedGet
  );
}

function makeFoodTypeReport(foodTypeData) {
  const barChart = document.getElementById("foodTypeBarChart");
  const pieChart = document.getElementById('foodTypePieChart')
  new Chart(barChart, {
    type: "bar",
    data: {
      labels: foodTypeData.map((data) => data.type),
      datasets: [
        {
          label: "Recipes by food type",
          data: foodTypeData.map((data) => data.count),
          backgroundColor: [
            "rgba(42, 96, 73, 0.7)",
            "rgba(42, 96, 73, 0.7)",
            "rgba(42, 96, 73, 0.7)",
            "rgba(42, 96, 73, 0.7)",
            "rgba(42, 96, 73, 0.7)"
          ],
          borderColor: [
            "rgba(42, 96, 73)",
            "rgba(42, 96, 73)",
            "rgba(42, 96, 73)",
            "rgba(42, 96, 73)",
            "rgba(42, 96, 73)"
            
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
      plugins: {
        title: {
          display: true,
          text: "Quantity of recipes by FOOD TYPE",
          color: "#2a6049",
          font: {
            size: 18
          }
        },
      },
    },
  });

  new Chart(pieChart, {
    type: 'pie',
    data: {
      labels: foodTypeData.map((data) => data.type),
      datasets: [
        {
          label: 'recipes by food type',
          data: foodTypeData.map((data) => data.count),
          backgroundColor: [
            'rgba(42, 96, 73)',
            'rgb(255, 99, 132)',
            'rgb(54, 162, 235)',
            'rgb(255, 205, 86)',
            'rgb(100, 110, 190)'
          ],
        }
      ]

    },
    options: {
      responsive: true,
      plugins: {
        legend: {
          position: 'top',
        },
        title: {
          display: true,
          text: 'Quantity of recipes by FOOD TYPE',
          color: "#2a6049",
          font: {
            size: 18
          }
        }
      }
    },
  })
}

function fetchCostReport() { }

function makeCostReport() { }

function showReports() {
  const userInfo = document.getElementById("user-info");
  const userMail = document.createElement("p");
  userMail.textContent = localStorage.getItem("user-mail");
  userMail.classList.add("green-color", "bold-words", "user-mail");
  userInfo.appendChild(userMail);
  fetchMomentReport();
  fetchFoodTypeReport();
  fetchCostReport();
}

function failedGet() {
  console.log("Failed to generate reports");
}
