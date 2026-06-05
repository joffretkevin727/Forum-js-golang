const API_URL = 'http://localhost:6767';

document.addEventListener('DOMContentLoaded', () => {
    initFormSwitcher();
    initFormSubmissions();
});

function initFormSwitcher() {
    const tabSignIn = document.getElementById('tab-signin');
    const tabSignUp = document.getElementById('tab-signup');
    const signInForm = document.getElementById('signInForm');
    const signUpForm = document.getElementById('signUpForm');

    if (tabSignIn && tabSignUp) {
        tabSignIn.addEventListener('click', () => {
            tabSignUp.classList.remove('active');
            tabSignIn.classList.add('active');
            signUpForm.classList.remove('active');
            signInForm.classList.add('active');
        });

        tabSignUp.addEventListener('click', () => {
            tabSignIn.classList.remove('active');
            tabSignUp.classList.add('active');
            signInForm.classList.remove('active');
            signUpForm.classList.add('active');
        });
    }
}

function initFormSubmissions() {
    const signInForm = document.getElementById('signInForm');
    const signUpForm = document.getElementById('signUpForm');

    if (signInForm) {
        signInForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const email = document.getElementById('signInIdentifier').value;
            const password = document.getElementById('signInPassword').value;

            try {
                const response = await fetch(`${API_URL}/users/login`, {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ email, password })
                });
                const data = await response.json();

                // Single line comment: Stores user session object and explicit userId key in localStorage upon login.
                if (response.ok) {
                    localStorage.setItem('isLoggedIn', 'true');
                    localStorage.setItem('user', JSON.stringify(data.user));

                    // L'emplacement correct de l'enregistrement du userId est ICI
                    if (data.user && data.user.id) {
                        localStorage.setItem('userId', data.user.id);
                    }

                    alert(`Ravi de vous revoir, ${data.user.username} !`);
                    window.location.href = '/home';
                } else {
                    alert(data.message || 'Identifiants incorrects.');
                }
            } catch (error) {
                console.error('Erreur login:', error);
                alert('Le serveur API est injoignable.');
            }
        });
    }

    if (signUpForm) {
        signUpForm.addEventListener('submit', async (e) => {
            e.preventDefault();
            const email = document.getElementById('signUpEmail').value;
            const username = document.getElementById('signUpUsername').value;
            const password = document.getElementById('signUpPassword').value;

            try {
                const response = await fetch(`${API_URL}/users`, {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ username, email, password })
                });

                // Single line comment: Handles basic sign up response and toggles to the login tab layout view.
                if (response.ok) {
                    alert('Compte créé avec succès ! Vous pouvez maintenant vous connecter.');
                    document.getElementById('tab-signin').click();
                } else {
                    alert('Erreur lors de la création du compte (Email ou Username déjà pris).');
                }
            } catch (error) {
                console.error('Erreur inscription:', error);
                alert('Le serveur API est injoignable.');
            }
        });
    }
}