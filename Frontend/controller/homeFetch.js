let currentPage = 1;
const itemsPerPage = 10;

function setupPagination(totalItems) {
    const container = document.getElementById('page-numbers-container');
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

            setupPagination(totalItems);
        });

        container.appendChild(pageDiv);
    }
}

setupPagination(25);

const fakeData = [
    {
        pseudo: "Jean sebastien",
        date: "2 hours ago",
        tags: ["#Croissant", "#Cannelés"],
        title: "Want to know the secret of the gateau basque recipe ?",
        text: "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Esse, repudiandae, voluptates. Accusamus deserunt dignissimos.",
        upVotes: 20,
        downVotes: 25
    },
    {
        pseudo: "Marie Antoinette",
        date: "5 hours ago",
        tags: ["#Brioche", "#Boulangerie"],
        title: "The ultimate trick for a perfect puff pastry (pâte feuilletée)",
        text: "Dignissimos facilis iure libero quibusdam unde ut. Autem ex facilis hic id, natus nihil porro praesentium.",
        upVotes: 42,
        downVotes: 2
    }
];

function renderCards(discussions) {
    const container = document.getElementById('cards-container');
    container.innerHTML = "";

    discussions.forEach(item => {
        const tagsHTML = item.tags.map(tag => `<p class="tag">${tag}</p>`).join('');

        const cardHTML = `
            <div class="card">
                <div class="header-card">
                    <div class="post-infos">
                        <p class="pseudo">${item.pseudo}</p>
                        <p class="post-date">${item.date}</p>
                    </div>
                    <div class="tags">
                        ${tagsHTML}
                    </div>
                </div>

                <div class="topic-title">
                    <p>${item.title}</p>
                </div>

                <div class="topic-text">
                    <p>${item.text}</p>
                </div>

                <div class="line"></div>

                <div class="interactions-btn">
                    <div class="votes-btn">
                        <div class="up-vote">
                            <img src="assets/icons/upvote.svg" alt="Upvote Icon" class="icon-recipes">
                        </div>
                        <p>${item.upVotes}</p>
                        <div class="width-spacer"></div>
                        <div class="down-vote">
                            <img src="assets/icons/downvote.svg" alt="Downvote Icon" class="icon-recipes">
                        </div>
                        <p>${item.downVotes}</p>
                    </div>
                    <div class="down-vote">
                        <p>Join Discussion</p>
                        <img src="assets/icons/arrowleft.svg" alt="Arrow Icon" class="icon-recipes">
                    </div>
                </div>
            </div>
        `;

        container.innerHTML += cardHTML;
    });
}

renderCards(fakeData);