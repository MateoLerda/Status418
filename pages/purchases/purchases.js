function getFoodsInPurchases() {
    fetch('https://localhost:8080/foods')
    .then(response => response.json())
    .then(data => {
        var foodTable = document.getElementById('dynamic-food-list');
        data.forEach(food => {
            var row = document.createElement('tr');
            row.innerHTML = `
            <td>${food.name}</td>
            <td>${food.price}</td>
            <td>${food.quantity}</td>
            <td>
                <input type="number" value="0" size="2" readonly>
                <button class="increment">+</button>
                <button class="decrement">-</button>
            </td>
            `;

            const input = row.querySelector('input');
            const incrementButton = row.querySelector('.increment');
            const decrementButton = row.querySelector('.decrement');

            incrementButton.addEventListener('click', () => {
                input.value = parseInt(input.value) + 1;
            });

            decrementButton.addEventListener('click', () => {
                if (parseInt(input.value) > 0) {
                    input.value = parseInt(input.value) - 1;
                }
            });
            foodTable.appendChild(fila);
        });
    })
    .catch(error => {
        console.error('Error al obtener datos:', error);
    });
}

document.addEventListener('DOMContentLoaded', getFoodsInPurchases());