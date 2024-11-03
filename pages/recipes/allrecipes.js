
let baseUrl = "http://localhost:8080/recipes/?filter_all=";

let searchValue
let foodType
let mealTime
let all = true

document.getElementById("toggleSwitch").addEventListener("change", function() {
  if (this.checked) {
      all= false 
  }else{
    all= true
  }
  getRecipes();
});


function createUrl(){
    searchValue= document.getElementById("searchInput").value
    foodType= document.getElementById("foodType").value
    mealTime= document.getElementById("mealTime").value
    baseUrl=  `${baseUrl}${all}`
    if (searchValue != ""){
      baseUrl = `${baseUrl}&filter_aproximation=${searchValue}`
    }
    if(foodType != "all"){
      baseUrl = `${baseUrl}&filter_type=${foodType}`
    }
    if(mealTime != "all"){
      baseUrl = `${baseUrl}&filter_moment=${mealTime}`
    }
    console.log(baseUrl)
}

function getRecipes(){
  createUrl();

  makeRequest(
    baseUrl,
    "GET",
    "",
    "application/json",
    true,
    successCreate,
    failed
  );
  baseUrl = "http://localhost:8080/recipes/?filter_all="
}

document.addEventListener("DOMContentLoaded", () => {
  getRecipes();
});

document.getElementById('foodType').onchange = () => {
  getRecipes();
}

document.getElementById('mealTime').onchange = () => {
  getRecipes();
}

const searchBar = document.getElementById('search-bar')

searchBar.addEventListener('submit', (e)=> {
    e.preventDefault()
    getRecipes();
})


function showRecipes(data) {
  const main = document.getElementById("recipe-list");
  main.innerHTML = '';
  data.forEach((recipe) => {
     // Crear el contenedor principal para cada receta
     let recipeContainer = document.createElement("div");
     recipeContainer.classList.add("recipe-container");
      recipeContainer.setAttribute("_id", recipe._id)
     // Nombre de la receta
     let recipeName = document.createElement("p");
     recipeName.textContent = recipe.recipe_name;
     recipeName.classList.add("big-font-size");
 
     // Momento de la receta
     let recipeMoment = document.createElement("p");
     recipeMoment.textContent = `Momento: ${recipe.recipe_moment}`;
     recipeMoment.classList.add("big-font-size");
 
     // Botón para mostrar más detalles
     let showMoreBtn = document.createElement("button");
     showMoreBtn.classList.add("btnS");
     let showMoreIcon= document.createElement("i")
     showMoreIcon.classList.add("fa-solid",  "fa-chevron-down")
     showMoreBtn.appendChild(showMoreIcon);

     let DeleteBtn = document.createElement("button");
     DeleteBtn.classList.add("btnS");
     let DeleteIcon= document.createElement("i")
     DeleteIcon.classList.add("fa-solid",  "fa-trash")
     DeleteBtn.appendChild(DeleteIcon);

     let UpdateBtn = document.createElement("button");
     UpdateBtn.classList.add("btnS");
     let UpdateIcon= document.createElement("i")
     UpdateIcon.classList.add("fa-solid",  "fa-pencil")
     UpdateBtn.appendChild(UpdateIcon);

     // Contenedor para la información adicional
     let recipeDetails = document.createElement("div");
     recipeDetails.classList.add("recipe-details");
     recipeDetails.style.display = "none"; // Ocultar detalles inicialmente
 
     // Información adicional de la receta
     let ingredients = document.createElement("p");
     ingredients.textContent = `Ingredients:  ${recipe.recipe_ingredients.map(ing => ` ${ing.Name}: ${ing.quantity}`).join(", ")}`;
     
     let description = document.createElement("p");
     description.textContent = `Descripction: ${recipe.recipe_description}`;
   
     // Agregar los detalles al contenedor de detalles
     recipeDetails.appendChild(ingredients);
     recipeDetails.appendChild(description);
 
     // Agregar el evento al botón para mostrar/ocultar detalles
     showMoreBtn.addEventListener("click", () => {
       if (recipeDetails.style.display === "none") {
         recipeDetails.style.display = "block";
         showMoreIcon.classList.remove("fa-chevron-down")
         showMoreIcon.classList.add("fa-chevron-up")
       } else {
         recipeDetails.style.display = "none";
         showMoreIcon.classList.remove("fa-chevron-up")
         showMoreIcon.classList.add("fa-chevron-down")
       }
     });
     DeleteBtn.addEventListener("click", () => {
      makeRequest(
        `http://localhost:8080/recipes/${recipe._id}`,
        "DELETE",
        "",
        "application/json",
        true,
        successDelete,
        failed
      );
     });
     UpdateBtn.addEventListener("click", () => {
        UpdateRecipe(recipe._id, recipe.recipe_name, recipe.recipe_description)
     });
     let recipeContainerFerst = document.createElement("div");
     recipeContainerFerst.classList.add("recipe-containerF");
     let buttonsContainer = document.createElement("div");
     buttonsContainer.classList.add("button-container");
      buttonsContainer.appendChild(showMoreBtn)
      buttonsContainer.appendChild(DeleteBtn)
      buttonsContainer.appendChild(UpdateBtn)
      if (!all){
        let CookBtn = document.createElement("button");
        CookBtn.classList.add("btnS");
        let CookIcon= document.createElement("i")
        CookIcon.classList.add("fa-solid",  "fa-utensils")
        CookBtn.appendChild(CookIcon);
        buttonsContainer.appendChild(CookBtn)
      }

     // Agregar los elementos al contenedor de la receta
     recipeContainer.appendChild(recipeName);
     recipeContainer.appendChild(recipeMoment);
     recipeContainer.appendChild(buttonsContainer);
     recipeContainerFerst.appendChild(recipeContainer)
     recipeContainerFerst.appendChild(recipeDetails);
    
    main.appendChild(recipeContainerFerst);
  });
}

function UpdateRecipe(recipeId, name, description){
    let id= recipeId
    let model = document.createElement("div")
    model.classList.add("divUpdate")
    let exitbtn= document.createElement("button")
    exitbtn.classList.add("fa-solid","fa-xmark", "btnS")
    let title= document.createElement("h2")
    title.textContent= "Edit Recipe"
    let divName= document.createElement("div")
    divName.classList.add("inputRecipe")
    let pname= document.createElement("p")
    pname.classList.add("big-font-size")
    pname.textContent= " Recipe name: "
    let inputname= document.createElement("input")
    inputname.placeholder= "Enter new name"
    inputname.value= name
    divName.appendChild(pname)
    divName.appendChild(inputname)
    let divDes= document.createElement("div")
    divDes.classList.add("inputRecipe")
    let pdes= document.createElement("p")
    pdes.classList.add("big-font-size")
    pdes.textContent= "Description: "
    let inputdes= document.createElement("textarea")
    inputdes.classList.add("inputdes")
    inputdes.placeholder= "Enter new description"
    inputdes.rows= 5
    inputdes.cols= 30
    inputdes.value= description
    divDes.appendChild(pdes)
    divDes.appendChild(inputdes)
    let buttonEdit = document.createElement("button")
    buttonEdit.classList.add("fa-solid","fa-pencil", "edit-recipe")
    buttonEdit.textContent= "  Confirm   "

    model.appendChild(exitbtn)
    model.appendChild(title)
    model.appendChild(divName)
    model.appendChild(divDes)
    model.appendChild(buttonEdit)
    
    const modalOverlay= document.getElementById("modalOverlay")
    modalOverlay.appendChild(model)
    modalOverlay.showModal()
    buttonEdit.addEventListener("click", () =>{
      const data = {
        recipe_name: inputname.value,
        recipe_description: inputdes.value
      }
      makeRequest(
        `http://localhost:8080/recipes/${recipeId}`,
        "PUT",
        data,
        "application/json",
        true,
        successUpdate,
        failed
      );
    } )
   exitbtn.addEventListener("click", () => {
    modalClose();
});
}

function modalClose(){
    modalOverlay.close(); 
    modalOverlay.innerHTML= "";
}

function successCreate(response) {
  showRecipes(response)
  console.log("Éxito:", response);
}

function failed(response) {
  console.log("Falla:", response);
}

function successDelete(response) {
  getRecipes();
  console.log("Éxito:", response);
}

function successUpdate(response) {
  modalClose();
  getRecipes();
  console.log("Éxito:", response);
}


