document.addEventListener('DOMContentLoaded', () => {
    // Initialise le système de basculement d'onglets
    initFormSwitcher();
});

function initFormSwitcher() {
    const tabSignIn = document.getElementById('tab-signin');
    const tabSignUp = document.getElementById('tab-signup');
    const signInForm = document.getElementById('signInForm');
    const signUpForm = document.getElementById('signUpForm');

    // Événement Clic sur l'onglet Sign In
    if (tabSignIn) {
        tabSignIn.addEventListener('click', () => {
            tabSignUp.classList.remove('active');
            tabSignIn.classList.add('active');

            signUpForm.classList.remove('active');
            signInForm.classList.add('active');
        });
    }

    // Événement Clic sur l'onglet Sign Up
    if (tabSignUp) {
        tabSignUp.addEventListener('click', () => {
            tabSignIn.classList.remove('active');
            tabSignUp.classList.add('active');

            signInForm.classList.remove('active');
            signUpForm.classList.add('active');
        });
    }
}