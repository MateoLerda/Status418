const baseUrl = "http://localhost:8080/foods/";
const baseUrlCreatePurchase = "http://localhost:8080/purchases/";


document.addEventListener('DOMContentLoaded', getFoodsInPurchases);

function getFoodsInPurchases() {
    makeRequest(
        baseUrl,
        "GET",
        "",
        "application/json",
        true,
        showFoods,
        failedGet
    );
    // let token = isUserLogged();
    // if (token == false) {
    //     window.location.href = "http://127.0.0.1:5500/pages/login/login.html";
    // }

    // const options = {
    //     method: "GET",
    //     headers: {
    //         "Content-Type": "application/json",
    //         Accept: "application/json",
    //         Authorization: `Bearer ${localStorage.getItem("authToken")}`,
    //     },
    // };

    // fetch(baseUrl, options)
    //     .then(response => response.json())
    //     .then(data => {
    //         const foodTable = document.getElementById('dynamic-food-table');
    //         data.forEach(food => {
    //             const row = document.createElement('tr');
    //             row.setAttribute('food-code', food._id);
    //             row.innerHTML = `
    //                 <td>${food.name}</td>
    //                 <td>U$D ${food.unit_price}</td>
    //                 <td>${food.current_quantity}</td>
    //                 <td>${food.minimum_quantity}</td>
    //                 <td>
    //                     <button class="decrement"><i class="fa-solid fa-minus"></i></button>
    //                     <input type="number" value="0" size="2" readonly>
    //                     <button class="increment"><i class="fa-solid fa-plus"></i></button>

    //                 </td>
    //             `;
    //             const decrementButton = row.querySelector('.decrement');
    //             const input = row.querySelector('input');
    //             input.id = 'quantityToBuy';
    //             const incrementButton = row.querySelector('.increment');


    //             incrementButton.addEventListener('click', () => {
    //                 input.value = parseInt(input.value) + 1;
    //             });

    //             decrementButton.addEventListener('click', () => {
    //                 if (parseInt(input.value) > 0) {
    //                     input.value = parseInt(input.value) - 1;
    //                 }
    //             });

    //             foodTable.appendChild(row);
    //         });
    //     })
    //     .catch(error => {
    //         console.error('Error al obtener datos:', error);
    //     });
}

function failedGet(response) {
    console.log("Failed to get foods:", response);
}

function showFoods(data) {
    if (!data || data.length === 0) {
        alert('There are not any products with current quantity below the minimum quantity.');
        return;
    }
    const foodTable = document.getElementById('dynamic-food-table');
    data.forEach(food => {
        const row = document.createElement('tr');
        row.setAttribute('food-code', food._id);
        row.innerHTML = `
            <td id= "name">${food.name}</td>
            <td id="price">U$D ${food.unit_price}</td>
            <td>${food.current_quantity}</td>
            <td>${food.minimum_quantity}</td>
            <td>
                <button class="decrement"><i class="fa-solid fa-minus"></i></button>
                <input type="number" value="0" size="2" readonly>
                <button class="increment"><i class="fa-solid fa-plus"></i></button>
            </td>
        `;
        const decrementButton = row.querySelector('.decrement');
        const input = row.querySelector('input');
        input.id = 'quantityToBuy';
        const incrementButton = row.querySelector('.increment');
        incrementButton.addEventListener('click', () => {
            input.value = parseInt(input.value) + 1;
        });

        decrementButton.addEventListener('click', () => {
            if (parseInt(input.value) > 0) {
                input.value = parseInt(input.value) - 1;
            }
        });

        foodTable.appendChild(row);
    });
}

document.getElementById('btnManually').addEventListener('click', manuallyPurchase);

function manuallyPurchase() {
    const foodTable = document.getElementById('dynamic-food-table');
    const rows = foodTable.getElementsByTagName('tr');
    const selectedFoods = [];

    for (let row of rows) {
        const input = row.querySelector('input');
        const quantity = parseInt(input.value);
        const name = row.querySelector('#name').textContent;
        const price = row.querySelector('#price').textContent;
        const priceValue = parseFloat(price.replace('U$D ', ''));
        if (quantity > 0) {
            const foodCode = row.getAttribute('food-code');
            selectedFoods.push({ foodCode, name, quantity, priceValue });
        }
    }

    if (selectedFoods.length === 0) {
        const dialog = document.createElement('div');
        dialog.className = 'dialog';
        dialog.textContent = 'You have to select a food';
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
            window.location.reload();
        });

        dialog.appendChild(closeButton);
        document.body.appendChild(dialog);

        return;
    }


    let TotalCost = 0;
    const purchaseDto = {

        foods: selectedFoods.map(food => {
            TotalCost += food.priceValue * food.quantity;
            return {
                _id: food.foodCode,
                name: food.name,
                quantity: food.quantity,
            };
        }),
        total_cost: TotalCost
    };

    // Llama a makeRequest con el objeto JSON transformado
    makeRequest(
        baseUrlCreatePurchase,
        "POST",
        purchaseDto,
        "application/json",
        true,
        successPurchase,
        failedPurchase
    );

}



function successPurchase(response) {
    console.log('Purchase successful:', response);
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
   
}

function failedPurchase(response) {
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

}
