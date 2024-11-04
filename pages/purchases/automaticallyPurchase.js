const baseUrlPurchases = "http://localhost:8080/purchases/?minimumList=true";
const baseUrlCreatePurchase = "http://localhost:8080/purchases/";
document.addEventListener('DOMContentLoaded', getMinimumList);

function getMinimumList() {
    const userInfo = document.getElementById('user-info');
    const userMail = document.createElement('p');
    userMail.textContent = localStorage.getItem('user-mail');
    userMail.classList.add('green-color', 'bold-words', 'user-mail');
    userInfo.appendChild(userMail);
    makeRequest(
        baseUrlPurchases,
        "GET",
        "",
        "application/json",
        true,
        showMinimumList,
        failedGet
    )
  
}
function showMinimumList(data) {
    const foodTable = document.getElementById('dynamic-food-table');
    if (data.message) {
        showAlert('There are not any products with current quantity below the minimum quantity. We will redirect you to home')
        document.getElementById("alert-button").addEventListener(('click'), () => {
            window.location.href= "../home/home.html"
        })
    }
    data.forEach(food => {
        const row = document.createElement('tr');
        row.innerHTML = `
            <td class="food-name">${food.name}</td>
            <td>U$D ${food.unit_price}</td>
            <td>${food.current_quantity}</td>
            <td>${food.minimum_quantity}</td>
            <td class="quantity-to-buy">${food.minimum_quantity - food.current_quantity}</td>
          `;

        foodTable.appendChild(row);
    });
}

function failedGet(response) {
    console.log("Failed to get foods:", response);
}

document.getElementById('btnAutomatically').addEventListener('click', automaticallyPurchase);

function automaticallyPurchase() {
    const rows = document.querySelectorAll('#dynamic-food-table tr');
    const purchases = [];

    rows.forEach(row => {
        const foodName = row.querySelector('.food-name').textContent;
        const quantityToBuy = parseInt(row.querySelector('.quantity-to-buy').textContent, 10);

        if (quantityToBuy > 0) {
            purchases.push({
                name: foodName,
                quantity: quantityToBuy
            });
        }
    });

    const options = {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Accept: "application/json",
            Authorization: `Bearer ${localStorage.getItem("authToken")}`,
        },
        body: JSON.stringify({ purchases })
    };

    fetch(baseUrlCreatePurchase, options)
        .then(response => {
            if (!response.ok) {
                throw new Error('Error en la respuesta de la API');
            }
            return response.json();
        })
        .then(data => {
            const dialog = document.createElement('div');
            dialog.className = 'dialog';
            dialog.textContent = 'Purchase completed successfully, Thank you!';
            const closeButton = document.createElement('button');
            closeButton.textContent = 'Close';
            closeButton.className = 'close-button';
            closeButton.style.pointerEvents = 'all';
            document.body.style.pointerEvents = 'none';
            const icon = document.createElement('i');
            icon.className = 'fa fa-check-circle';
            icon.style.marginRight = '10px';
            dialog.insertBefore(icon, dialog.firstChild);
            document.body.style.overflow = 'hidden';
            closeButton.style.marginTop = '10px';
            closeButton.addEventListener('click', () => {
                document.body.removeChild(dialog);
                window.location = "http://localhost:5500/pages/home/home.html";
            });

            dialog.appendChild(closeButton);
            document.body.appendChild(dialog);
            console.log('Respuesta de la API:', data);
        })
        .catch(error => {
            console.error('Error al realizar la compra automática:', error);
            console.log('Failed to make purchase:', response);
            const dialog = document.createElement('div');
            dialog.className = 'dialog';
            dialog.textContent = 'Failed to make purchase, try again';
            const closeButton = document.createElement('button');
            closeButton.textContent = 'Close';
            closeButton.className = 'close-button';
            closeButton.style.pointerEvents = 'all';
            document.body.style.pointerEvents = 'none';
            const icon = document.createElement('i');
            icon.className = 'fa fa-times-circle';
            icon.style.marginRight = '10px';
            dialog.insertBefore(icon, dialog.firstChild);
            document.body.style.overflow = 'hidden';
            closeButton.style.marginTop = '10px';
            closeButton.addEventListener('click', () => {
                document.body.removeChild(dialog);
                window.location.reload();
            });

            dialog.appendChild(closeButton);
            document.body.appendChild(dialog);

        });
}

