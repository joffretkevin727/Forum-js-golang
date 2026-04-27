const tabs = document.querySelectorAll('.row-start p');
const boxSignIn = document.querySelector('.box.active');
const boxSignUp = document.querySelector('.sign-up');

tabs.forEach((tab, index) => {
    tab.addEventListener('click', () => {
        document.querySelector('.row-start .active').classList.remove('active');
        tab.classList.add('active');

        if (index === 0) {
            boxSignIn.style.display = 'block';
            boxSignUp.style.display = 'none';
        } else {
            boxSignIn.style.display = 'none';
            boxSignUp.style.display = 'block';
        }
    });
});