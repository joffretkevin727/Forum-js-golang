const API_URL = 'http://localhost:6767';

document.addEventListener('DOMContentLoaded', () => {
    const signInForm = document.getElementById('signInForm');
    const signUpForm = document.getElementById('signUpForm');
    
    // Basculer l'affichage entre connexion et inscription
    document.getElementById('switchToSignUp').addEventListener('click', () => {
        signInForm.classList.remove('active');
        signUpForm.classList.add('active');
    });

    document.getElementById('switchToSignIn').addEventListener('click', () => {
        signUpForm.classList.remove('active');
        signInForm.classList.add('active');
    });

    // Gestion de la CONNEXION (Sign In)
    signInForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const email = document.getElementById('signInEmail').value;
        const password = document.getElementById('signInPassword').value;

        try {
            const response = await fetch(`${API_URL}/users/login`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ email, password })
            });
            const data = await response.json();

            if (response.ok) {
                localStorage.setItem('isLoggedIn', 'true');
                localStorage.setItem('user', JSON.stringify(data.user));
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

    // Gestion de la CRÉATION DE COMPTE (Sign Up)
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

            if (response.ok) {
                alert('Compte créé avec succès ! Vous pouvez maintenant vous connecter.');
                signUpForm.classList.remove('active');
                signInForm.classList.add('active');
            } else {
                alert('Erreur lors de la création du compte (Email ou Username déjà pris).');
            }
        } catch (error) {
            console.error('Erreur inscription:', error);
            alert('Le serveur API est injoignable.');
        }
    });
});