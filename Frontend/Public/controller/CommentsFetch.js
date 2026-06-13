document.addEventListener("DOMContentLoaded", () => {
    // CORRECTION : On extrait l'ID qui se trouve après le point d'interrogation (?id=...)
    const urlParams = new URLSearchParams(window.location.search);
    const topicId = urlParams.get('topic_id');

    // On vérifie que l'ID existe bien et qu'il s'agit d'un nombre valide
    if (!topicId || isNaN(topicId)) {
        console.error("Impossible de récupérer un ID de topic valide depuis l'URL de requêtes");
        const container = document.getElementById('comments-container');
        return;
    }

    console.log("ID du topic détecté :", topicId); // Petit log de contrôle pour toi

    // Lancement du chargement des commentaires avec le bon ID
    loadComments(topicId);
});

// Fonction qui fait l'appel API vers ton serveur Go (Port 6767)
function loadComments(topicId) {
    fetch(`http://localhost:6767/topiccomments?topic_id=${topicId}`)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Erreur serveur Go : ${response.status}`);
            }
            return response.json();
        })
        .then(comments => {
            console.log("Commentaires reçus de Go :", comments); // Regarde si ce log apparaît !
            renderComments(comments);
        })
        .catch(error => {
            console.error("Erreur lors de la récupération des commentaires :", error);
            const container = document.getElementById('comments-container');
            if (container) {
                container.innerHTML = `<p class="error">Impossible de charger les commentaires.</p>`;
            }
        });
}

// Fonction qui génère le HTML et l'injecte dans la page
function renderComments(comments) {
    const container = document.getElementById('comments-container');
    if (!container) return;

    // Si la liste est vide
    if (!comments || comments.length === 0) {
        container.innerHTML = '<p class="no-comments">Soyez le premier à laisser un commentaire !</p>';
        return;
    }

    container.innerHTML = ""; // On vide le message de chargement

    comments.forEach(comment => {
        // Formatage de la date SQL en format lisible (ex: 12 juin, 15:27)
        let dateFormatee = "Date inconnue";
        if (comment.created_at) {
            dateFormatee = new Date(comment.created_at).toLocaleDateString('fr-FR', {
                day: 'numeric',
                month: 'short',
                hour: '2-digit',
                minute: '2-digit'
            });
        }

        // Création du template HTML pour un commentaire
        const commentHTML = `
            <div class="comment-item" data-id="${comment.id}">
                <img src="Assets/images/user.svg" alt="User" class="comment-avatar">
                <div class="comment-content">
                    <div class="comment-meta">
                        <span class="comment-author">${comment.pseudo || 'Anonyme'}</span>
                        <span class="comment-time">${dateFormatee}</span>
                    </div>
                    <p class="comment-text">
                        ${comment.body}
                    </p>
                    <div class="comment-footer">
                        <span>Reply</span>
                        <span>Upvote (0)</span>
                    </div>
                </div>
            </div>
        `;

        // Injection dans le conteneur principal
        container.innerHTML += commentHTML;
    });
}