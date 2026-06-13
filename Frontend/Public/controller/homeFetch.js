let currentPage = 1;
const itemsPerPage = 10;
let allDiscussions = []; // Stockage global pour gérer la pagination côté front

// Fonction pour récupérer les topics depuis l'API Go
function fetchTopics() {
    fetch('http://localhost:6767/topics')
        .then(response => {
            if (!response.ok) {
                throw new Error(`Erreur serveur Go : ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            console.log("Topics reçus de Go :", data);
            allDiscussions = data || [];

            // On initialise la pagination avec le nombre total d'éléments reçus
            setupPagination(allDiscussions.length);
            // On affiche uniquement les éléments de la page courante
            displayCurrentPage();
        })
        .catch(error => {
            console.error("Impossible de joindre le serveur Go :", error);
            const container = document.getElementById('cards-container');
            if (container) {
                container.innerHTML = `<p class="error">Impossible de charger les sujets.</p>`;
            }
        });
}

// Fonction pour découper le tableau et afficher uniquement la page active
function displayCurrentPage() {
    const startIndex = (currentPage - 1) * itemsPerPage;
    const endIndex = startIndex + itemsPerPage;
    const paginatedItems = allDiscussions.slice(startIndex, endIndex);

    renderCards(paginatedItems);
}

// Gestion de la pagination graphique
function setupPagination(totalItems) {
    const container = document.getElementById('page-numbers-container');
    if (!container) return;
    container.innerHTML = "";

    const totalPages = Math.ceil(totalItems / itemsPerPage);

    for (let i = 1; i <= totalPages; i++) {
        const pageDiv = document.createElement('div');
        pageDiv.classList.add('page-number');
        pageDiv.textContent = i;

        if (i === currentPage) {
            pageDiv.classList.add('active-page');
        }

        pageDiv.addEventListener('click', () => {
            currentPage = i;
            console.log(`Chargement de la page : ${currentPage}`);

            setupPagination(totalItems); // Recalcule le style des boutons
            displayCurrentPage();        // Change les cartes affichées !
        });

        container.appendChild(pageDiv);
    }
}

// Génération des cartes HTML
function renderCards(discussions) {
    const container = document.getElementById('cards-container');
    if (!container) return;
    container.innerHTML = "";

    if (!discussions || discussions.length === 0) {
        container.innerHTML = "<p>Aucun sujet pour le moment.</p>";
        return;
    }

    discussions.forEach(item => {
        const tagsArray = item.tags || [];
        const tagsHTML = tagsArray.map(tag => `<p class="tag">${tag}</p>`).join('');

        let dateFormatee = "Date inconnue";
        if (item.date) {
            dateFormatee = new Date(item.date).toLocaleDateString('fr-FR', {
                day: 'numeric',
                month: 'short',
                hour: '2-digit',
                minute: '2-digit'
            });
        }

        // L'identifiant unique du topic venant de Go
        const topicId = item.id;

        const cardHTML = `
            <div class="card">
                <div class="header-card">
                    <div class="post-infos">
                        <p class="pseudo">${item.pseudo || 'Anonyme'}</p>
                        <p class="post-date">${dateFormatee}</p>
                    </div>
                    <div class="tags">
                        ${tagsHTML}
                    </div>
                </div>

                <div class="topic-title">
                    <p>${item.title || 'Sans titre'}</p>
                </div>

                <div class="topic-text">
                    <p>${item.text || item.body || ''}</p> 
                </div>

                <div class="line"></div>

                <div class="interactions-btn">
                    <div class="votes-btn">
                        <div class="up-vote">
                            <img src="assets/icons/upvote.svg" alt="Upvote Icon" class="icon-recipes">
                        </div>
                        <p>${item.upVotes || 0}</p>
                        <div class="width-spacer"></div>
                        <div class="down-vote">
                            <img src="assets/icons/downvote.svg" alt="Downvote Icon" class="icon-recipes">
                        </div>
                        <p>${item.downVotes || 0}</p>
                    </div>
                    
                    <a href="/topic_comments?topic_id=${topicId}" class="down-vote" style="text-decoration: none; color: inherit; display: flex; align-items: center;">
                        <p>Join Discussion</p>
                        <img src="assets/icons/arrowleft.svg" alt="Arrow Icon" class="icon-recipes">
                    </a>
                </div>
            </div>
        `;

        container.innerHTML += cardHTML;
    });
}

// Lancement au chargement de la page
document.addEventListener('DOMContentLoaded', () => {
    fetchTopics();
});