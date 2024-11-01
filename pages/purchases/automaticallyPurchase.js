const baseUrlPurchases = "http://localhost:8080/purchases/?minimumList=true";
const baseUrlCreatePurchase = "http://localhost:8080/purchases/";
document.addEventListener('DOMContentLoaded', getMinimumList);

function getMinimumList() {
    makeRequest(
        baseUrlPurchases,
        "GET",
        "",
        "application/json",
        true,
        showMinimumList,
        failedGet
    )
    // let token = isUserLogged();
    // if (token == false) {
    //     window.location.href = "http://localhost:5500/pages/login/login.html";
    // }
    // const options = {
    //     method: "GET",
    //     headers: {
    //         "Content-Type": "application/json",
    //         Accept: "application/json",
    //         Authorization: `Bearer ${localStorage.getItem("authToken")}`,
    //     },
    // };

    // fetch(baseUrlPurchases, options)
    //     .then(response => response.json())
    //     .then(data => {
    //         if (!Array.isArray(data)) {
    //             console.error('La respuesta no es un arreglo:', data);
    //             throw new Error("La respuesta de la API no es un arreglo");
    //         }
    //         const foodTable = document.getElementById('dynamic-food-table');
    //         data.forEach(food => {
    //             const row = document.createElement('tr');
    //             row.innerHTML = `
    //                 <td class="food-name">${food.name}</td>
    //                 <td>U$D ${food.unit_price}</td>
    //                 <td>${food.current_quantity}</td>
    //                 <td>${food.minimum_quantity}</td>
    //                 <td class="quantity-to-buy">${food.minimum_quantity - food.current_quantity}</td>
    //               `;

    //             foodTable.appendChild(row);
    //         });
    //     })
    //     .catch(error => {
    //         console.error('Error al obtener datos:', error);
    //     });
}
function showMinimumList(data) {
    const foodTable = document.getElementById('dynamic-food-table');
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

    if (purchases.length === 0) {
        const dialog = document.createElement('div');
        dialog.className = 'dialog';
        dialog.textContent = 'There are not any products with current quantity below the minimum quantity. We will redirect you to home';
        const closeButton = document.createElement('button');
        closeButton.textContent = 'Close';
        closeButton.className = 'close-button';
        closeButton.style.pointerEvents = 'all';
        document.body.style.pointerEvents = 'none';
        document.body.scrollTop = 0;
        const icon = document.createElement('i');
        icon.className = 'fa fa-exclamation-triangle';
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
        return;
    }

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

