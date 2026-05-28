document.getElementById('loginForm').addEventListener('submit', async function(e) {
    e.preventDefault(); // Empêche le rechargement de la page

    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    try {
        
        const response = await fetch('http://localhost:6767/users/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email, password })
        });

        const data = await response.json();

        if (response.ok) {
            localStorage.setItem('isLoggedIn', 'true');
            localStorage.setItem('user', JSON.stringify(data.user));
            
            
            alert("Ravi de vous revoir, " + data.user.username + " !");
            window.location.href = "profil"; 
        } else {
            alert(data.message || "Identifiants invalides.");
        }
    } catch (error) {
        console.error("Erreur connexion :", error);
        alert("Le serveur est injoignable.");
    }
});