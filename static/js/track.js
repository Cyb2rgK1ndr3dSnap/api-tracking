document.addEventListener('DOMContentLoaded', function() {  

let statusMap = {};

// Función para obtener el texto del estado
function GetStatusText(status) {
    return statusMap[status]?.name || "";
}

// Función para obtener el color del estado
function GetStatusColor(status) {
    return statusMap[status]?.color || "";
}

async function FetchStatuses() {
    try {
        const response = await fetch("/status");
        if (!response.ok) {
            throw new Error("Error al obtener los estados");
        }

        const statusFilterElement = document.getElementById("status-filter");
        const statusDetailElement = document.getElementById("detail-status");
        const data = await response.json();
        // Poblar el mapa de estados
        data.statuses.forEach(status => {
            const option = document.createElement("option");
            option.value = status.id_status;
            option.textContent = status.name;
            statusFilterElement.appendChild(option);

            const detailOption = option.cloneNode(true);
            statusDetailElement.appendChild(detailOption);
            
            statusMap[status.id_status] = {
                name: status.name,
                color: status.color
            };
        });
    } catch (error) {
        console.error("Error al cargar los estados:", error);
    }
}

// Generar filas dinámicamente
const FetchShippings = async (status = "",pageSize = "10",pageNumber = "1") => {
    const response = await fetch(`/shipping?${status ? `&status=${status}` : ""}${pageSize ? `&size=${pageSize}` : ""}${pageNumber ? `&page=${pageNumber}` : ""}&web=true`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    });

    if (response.ok) {
        const result = await response.json();
        if(result != null){
            HtmlShippings(result);
        }else{
            document.getElementById('message').innerText = "No existen registros para la página"
        }
        document.getElementById('spinner').style.display = 'none';
    } else {
        const error = await response.json();
        console.log(`Error: ${error.error}`);
    }
};

HtmlShippings = async (data) => {

    const tableRows = document.getElementById("table-rows");
    tableRows.innerHTML = '';

    if(data.total_pages != null){
        const pageLimit = document.getElementById("page-limit");
        pageLimit.innerHTML = data.total_pages
    }

    data.data.forEach(item => {
        const last_update = new Date(item.last_update).toISOString().split("T")[0];
        const expiration_date = new Date(item.expiration_date).toISOString().split("T")[0];

        const row = document.createElement("div");
        row.classList.add("table-row");
        row.id = item.id_shipping
    
        row.style.backgroundColor = GetStatusColor(item.status)
    
        row.innerHTML = `
            <div class="column">${item.shipping_number}</div>
            <div class="column">${item.username}</div>
            <div class="column">${item.email}</div>
            <div class="column">${GetStatusText(item.status)}</div>
            <div class="column">${expiration_date}</div>
        `;

        row.addEventListener("click", () => showDetails(item));
        tableRows.appendChild(row);
    });
}

// Mostrar detalles en la tarjeta
const showDetails = (item) => {
    CleanInputs()
    
    document.getElementById("detail-item-edit").classList.remove("hidden");

    document.getElementById("detail-card").classList.remove("hidden");
    document.getElementById("toggle-close").classList.remove("hidden");
    document.getElementById("toggle-edit").classList.remove("hidden");
    document.getElementById("toggle-create").classList.add("hidden");

    const last_update = new Date(item.last_update).toISOString().split("T")[0];
    const expiration_date = new Date(item.expiration_date).toISOString().split("T")[0];

    document.getElementById("detail-id").value = item.id_shipping;
    document.getElementById("detail-shipping-number").value = item.shipping_number;
    document.getElementById("detail-username").value = item.username;
    document.getElementById("detail-email").value = item.email;
    document.getElementById("detail-quantity").value = item.quantity;
    document.getElementById("detail-weight").value = item.weight;
    document.getElementById("detail-amount").value = item.amount;
    document.getElementById("detail-status").value = item.status;
    document.getElementById("detail-updated").value = last_update;
    document.getElementById("detail-expiration").value = expiration_date;

    // Seleccionar el contenedor donde se añadirá el botón
    const buttonContainer = document.getElementById("detail-buttons");

    // Eliminar cualquier botón existente en el contenedor
    buttonContainer.innerHTML = "";

    // Crear un nuevo botón
    const editButton = document.createElement("button");
    editButton.id = "toggle-edit"
    editButton.textContent = "Editar";
    editButton.className = "edit-button";
    editButton.addEventListener("click", () => editShipping(item));

    const closeButton = document.createElement("button");
    closeButton.id = "toggle-close"
    closeButton.textContent = "Cerrar";
    closeButton.className = "edit-button";
    closeButton.addEventListener("click", () => closeShipping(item));

    buttonContainer.appendChild(closeButton);
    buttonContainer.appendChild(editButton);
};

const createDetails = () => {
    document.getElementById("detail-card").classList.remove("hidden");
    document.getElementById("toggle-create").classList.remove("hidden");
    
    document.getElementById("toggle-close").classList.add("hidden");
    document.getElementById("toggle-edit").classList.add("hidden");
    document.getElementById("detail-item-edit").classList.add("hidden");

    CleanInputs();
}

function updateRow(item) {
    const row = document.getElementById(item.id_shipping);
    
    if (row) {
        const columns = row.getElementsByClassName("column");

        columns[0].textContent = item.shipping_number;
        columns[1].textContent = item.username;
        columns[2].textContent = item.email;
        columns[3].textContent = GetStatusText(item.status);
        columns[4].textContent = new Date(item.expiration_date).toISOString().split("T")[0];

        row.style.backgroundColor = GetStatusColor(item.status)
    }
}

const createShipping = async () => {
    let expirationDate = document.getElementById("detail-expiration").value;

    // Verifica si la fecha no contiene 'T', agrega 'T' y la hora al final
    if (expirationDate && !expirationDate.includes('T')) {
        expirationDate += 'T00:00:00Z';  // Agregar la hora y zona horaria
    }

    const shippingData = {
        username: document.getElementById("detail-username").value,
        shipping_number: document.getElementById("detail-shipping-number").value,
        weight: parseFloat(document.getElementById("detail-weight").value),
        amount: parseFloat(document.getElementById("detail-amount").value),
        quantity: parseInt(document.getElementById("detail-quantity").value),
        expiration_date: expirationDate
    };

    try {
        // Realiza la solicitud POST al servidor
        const response = await fetch("/shipping", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(shippingData)
        });

        const data = await response.json();
        // Si la respuesta es exitosa (status 200-299)
        if (response.ok) {
            // Mostrar mensaje de éxito o lo que el servidor retorne
            document.getElementById('message').innerText = data.message
        } else {
            // Si hay un error, mostrar el mensaje de error
            document.getElementById('message').innerText = data.error
        }
    } catch (error) {
        // Manejo de errores en la solicitud
        console.error("Error en la solicitud:", error);
        alert("Error en la conexión. Intente nuevamente.");
    }
}

const editShipping = async (item) => {
    let expirationDate = document.getElementById("detail-expiration").value;

    // Verifica si la fecha no contiene 'T', agrega 'T' y la hora al final
    if (expirationDate && !expirationDate.includes('T')) {
        expirationDate += 'T00:00:00Z';  // Agregar la hora y zona horaria
    }

    const shippingData = {
        id_shipping: parseInt(document.getElementById("detail-id").value, 10),
        username: item.username !== document.getElementById("detail-username").value
        ? document.getElementById("detail-username").value
        : null,
        shipping_number: item.shipping_number !== document.getElementById("detail-shipping-number").value
        ? document.getElementById("detail-shipping-number").value
        : null,
        weight: item.weight !== parseFloat(document.getElementById("detail-weight").value)
        ? parseFloat(document.getElementById("detail-weight").value)
        : null,
        amount: item.amount !== parseFloat(document.getElementById("detail-amount").value)
        ? parseFloat(document.getElementById("detail-amount").value)
        : null,
        quantity: item.quantity !== parseInt(document.getElementById("detail-quantity").value, 10)
        ? parseInt(document.getElementById("detail-quantity").value, 10)
        : null,
        status: item.status !== parseInt(document.getElementById("detail-status").value, 10)
        ? parseInt(document.getElementById("detail-status").value, 10)
        : null,
        expiration_date: expirationDate !== item.expiration_date 
        ? expirationDate
        : null,
    };

    try {
        // Realiza la solicitud POST al servidor
        const response = await fetch("/shipping", {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(shippingData)
        });

        const data = await response.json();
        // Si la respuesta es exitosa (status 200-299)
        if (response.ok) {
            console.log(data)
            item.shipping_number = shippingData.shipping_number != null ? shippingData.shipping_number : item.shipping_number
            item.username = shippingData.username != null ? shippingData.username : item.username
            item.email = data.email != null ? data.email : item.email
            item.weight = shippingData.weight != null ? shippingData.weight : item.weight
            item.amount = shippingData.amount != null ? shippingData.amount : item.amount
            item.quantity = shippingData.quantity != null ? shippingData.quantity : item.quantity
            item.status = shippingData.status != null ? shippingData.status : item.status
            item.expiration_date = shippingData.expiration_date != null ? shippingData.expiration_date : item.expiration_date

            updateRow(item)

            document.getElementById('message').innerText = data.message;
            document.getElementById('message').classList.remove('hidden');
            document.getElementById("detail-card").classList.add("hidden");
            CleanInputs()
            setTimeout(() => {
                document.getElementById('message').classList.add('hidden');
            }, 10000);
        } else {
            // Si hay un error, mostrar el mensaje de error
            document.getElementById('message').innerText = data.error
        }
    } catch (error) {
        // Manejo de errores en la solicitud
        console.error("Error en la solicitud:", error);
        alert("Error en la conexión. Intente nuevamente.");
    }
}

const closeShipping = async (item) => {

}

document.getElementById("toggle-create").addEventListener("click", () => createShipping())

// Cerrar la tarjeta
document.getElementById("close-card").addEventListener("click", () => {
    document.getElementById("detail-item-edit").classList.remove("hidden");

    document.getElementById("detail-card").classList.add("hidden");
    document.getElementById("toggle-close").classList.add("hidden");
    document.getElementById("toggle-edit").classList.add("hidden");
    document.getElementById("toggle-create").classList.add("hidden");
});

document.getElementById("new-button").addEventListener("click", () => createDetails());

const CleanInputs = async () => {
    document.getElementById("detail-shipping-number").value = "";
    document.getElementById("detail-username").value = "";
    document.getElementById("detail-email").value = "";
    document.getElementById("detail-quantity").value = "";
    document.getElementById("detail-weight").value = "";
    document.getElementById("detail-amount").value = "";
    document.getElementById("detail-status").value = "";
    document.getElementById("detail-updated").value = "";
    document.getElementById("detail-expiration").value = "";
}

//Ejecutar fetch con los filtros
const FilterFetch = async () => {
    const statusElement = document.getElementById('status-filter');
    const status = statusElement.value;

    const pageSElement = document.getElementById('page-size');
    const pageS = pageSElement.value;

    const pageNElement = document.querySelector('.page-input');
    const pageN = pageNElement.value;

    document.getElementById('spinner').style.display = 'flex';
    FetchShippings(status, pageS, pageN)
}

// Filtrar por estado
document.getElementById("status-filter").addEventListener("change", () => {
    const input = document.querySelector('.page-input');
    input.value = 1;
    FilterFetch();
});

// Filtro de número de elementos
document.getElementById("page-size").addEventListener("change", () => {
    const page = document.querySelector('.page-input');
    page.value = 1;
    FilterFetch();
});

function goToPage(direction) {
    const input = document.querySelector('.page-input');
    let currentPage = parseInt(input.value, 10);
    const pageLimit = document.getElementById("page-limit");
    

    if (direction === 'prev' && currentPage > 1) {
        currentPage -= 1;
    } else if (direction === 'next' && currentPage < pageLimit.innerHTML) {
        currentPage += 1;
    }else{
        return;
    }

    input.value = currentPage;
    FilterFetch();
}

document.querySelector('.pagination-btn:first-child').addEventListener('click', () => goToPage('prev'));
document.querySelector('.pagination-btn:last-child').addEventListener('click', () => goToPage('next'));
document.querySelector('.page-input').addEventListener('input', (e) => FilterFetch());

FilterFetch();
FetchStatuses()
});