/*document.addEventListener('DOMContentLoaded', () => {
    const isLoggedIn = localStorage.getItem('isLoggedIn') === 'true';
    const user = JSON.parse(localStorage.getItem('user'));

    const profileLink = document.querySelector('.profile-icon');
    const startDiscussionBtn = document.querySelector('.new-discussion');

    if (isLoggedIn && user) {
        // Change le lien du profil pour aller sur la page profil et affiche le pseudo au survol
        profileLink.href = '/profil';
        profileLink.setAttribute('aria-label', `Profil de ${user.username}`);
        
        // Permet de cliquer sur "Start a discussion"
        startDiscussionBtn.style.opacity = '1';
        startDiscussionBtn.style.cursor = 'pointer';
    } else {
        // Si anonyme, redirige vers la page de login au clic
        profileLink.href = '/login';
        
        // Grise le bouton de création de topic car le middleware Go va bloquer la requête
        startDiscussionBtn.style.opacity = '0.5';
        startDiscussionBtn.style.cursor = 'not-allowed';
        startDiscussionBtn.addEventListener('click', (e) => {
            e.preventDefault();
            alert('Vous devez être connecté pour lancer une discussion.');
            window.location.href = '/login';
        });
    }
});*/