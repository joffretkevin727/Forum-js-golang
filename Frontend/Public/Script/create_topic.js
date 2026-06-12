// Single line comment: Manages tags behavior and enforces user authentication check before submitting the payload.
document.addEventListener("DOMContentLoaded", () => {
    const tagInput = document.getElementById("tag-input");
    const tagsContainer = document.getElementById("tags-container");
    const tagsHiddenInput = document.getElementById("tags-hidden-input");
    const datalist = document.getElementById("default-tags");
    const form = document.querySelector(".topic-creation");

    const defaultTags = Array.from(datalist.options).map(opt => opt.value.toLowerCase());
    let selectedTags = [];

    function updateTags() {
        tagsContainer.innerHTML = "";
        selectedTags.forEach((tag, index) => {
            const tagSpan = document.createElement("span");
            tagSpan.classList.add("tag");
            if (tag.type === "default") {
                tagSpan.classList.add("default-tag");
            } else {
                tagSpan.classList.add("created-tag");
            }
            tagSpan.innerHTML = `${tag.name} <span class="tag-remove" data-index="${index}">&times;</span>`;
            tagsContainer.appendChild(tagSpan);
        });
        tagsHiddenInput.value = JSON.stringify(selectedTags.map(t => t.name));
    }

    function addTag(value) {
        const cleanValue = value.trim();
        if (cleanValue === "") return;
        if (selectedTags.some(t => t.name.toLowerCase() === cleanValue.toLowerCase())) {
            tagInput.value = "";
            return;
        }
        const isDefault = defaultTags.includes(cleanValue.toLowerCase());
        const tagType = isDefault ? "default" : "created";
        selectedTags.push({ name: cleanValue, type: tagType });
        updateTags();
        tagInput.value = "";
    }

    tagInput.addEventListener("keydown", (e) => {
        if (e.key === "Enter") {
            e.preventDefault();
            addTag(tagInput.value);
        }
    });

    tagInput.addEventListener("input", () => {
        const options = Array.from(datalist.options).map(opt => opt.value);
        if (options.includes(tagInput.value)) {
            addTag(tagInput.value);
        }
    });

    tagsContainer.addEventListener("click", (e) => {
        if (e.target.classList.contains("tag-remove")) {
            const index = e.target.getAttribute("data-index");
            selectedTags.splice(index, 1);
            updateTags();
        }
    });

    // --- SOUBOUMISSION ET VÉRIFICATION DE CONNEXION ---
    if (form) {
        form.addEventListener("submit", async (e) => {
            e.preventDefault();

            // 1. CORRECTION : On récupère la chaîne et on la transforme en OBJET
            const userString = localStorage.getItem("user");
            if (!userString) {
                alert("Vous n'êtes pas connecté ! Vous devez avoir un compte et être identifié pour créer un sujet.");
                window.location.href = "http://localhost:6969/login";
                return;
            }
            const user = JSON.parse(userString); // Maintenant on a accès à user.id et user.username

            const formData = new FormData(form);

            // 2. CORRECTION : On ajoute le pseudo, l'ID et on renomme 'body' en 'text'
            const payload = {
                title: formData.get("titre"),
                text: formData.get("corps"),             // Changé 'body' en 'text' pour Go
                pseudo: user.username || "Anonyme",     // Ajouté pour le pseudo
                author_id: parseInt(user.id, 10),       // Ajouté pour l'ID (converti en entier)
                tags: JSON.parse(tagsHiddenInput.value || "[]")
            };

            try {
                // Adaptez l'URL vers votre API Go (port 6767)
                const response = await fetch("http://localhost:6767/topics", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                        "X-User-ID": user.id.toString() // ID propre envoyé en String dans le header
                    },
                    body: JSON.stringify(payload)
                });

                if (response.status === 401) {
                    alert("Session invalide ou expirée. Veuillez vous reconnecter.");
                    localStorage.removeItem("user"); // Nettoyage de la bonne clé
                    window.location.href = "http://localhost:6969/login";
                    return;
                }

                if (!response.ok) throw new Error("Erreur API Go");
                const data = await response.json();

                if (data && data.id) {
                    // Redirection finale vers votre serveur Express
                    window.location.href = `http://localhost:6969/topic/${data.id}`;
                }
            } catch (error) {
                console.error("Erreur:", error);
                alert("Impossible de créer le sujet.");
            }
        });
    }
});