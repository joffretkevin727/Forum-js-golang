
const recipes = document.querySelectorAll('.col-side .recipe');

recipes.forEach(recipe => {
    recipe.addEventListener('click', () => {
        const currentActive = document.querySelector('.col-side .recipe.active');
        if (currentActive) {
            currentActive.classList.remove('active');
        }

        recipe.classList.add('active');
    });
});