document.addEventListener("DOMContentLoaded", () => {
    const tagInput = document.getElementById("tag-input");
    const tagsContainer = document.getElementById("tags-container");
    const tagsHiddenInput = document.getElementById("tags-hidden-input");
    const datalist = document.getElementById("default-tags");

    // Récupérer la liste des tags par défaut à partir du datalist
    const defaultTags = Array.from(datalist.options).map(opt => opt.value.toLowerCase());
    
    // Tableau pour stocker les tags sélectionnés (ex: [{name: "HTML", type: "default"}] )
    let selectedTags = [];

    // Fonction pour mettre à jour l'affichage et l'input caché
    function updateTags() {
        tagsContainer.innerHTML = "";
        selectedTags.forEach((tag, index) => {
            const tagSpan = document.createElement("span");
            tagSpan.classList.add("tag");
            
            // Applique la classe bleu ou vert selon l'origine du tag
            if (tag.type === "default") {
                tagSpan.classList.add("default-tag");
            } else {
                tagSpan.classList.add("created-tag");
            }
            
            tagSpan.innerHTML = `${tag.name} <span class="tag-remove" data-index="${index}">&times;</span>`;
            tagsContainer.appendChild(tagSpan);
        });

        // Met à jour l'input caché sous forme de chaîne de caractères (ex: "HTML,CSS,MonNouveauTag")
        tagsHiddenInput.value = selectedTags.map(t => t.name).join(",");
    }

    // Fonction pour ajouter un tag
    function addTag(value) {
        const cleanValue = value.trim();
        if (cleanValue === "") return;

        // Éviter les doublons
        if (selectedTags.some(t => t.name.toLowerCase() === cleanValue.toLowerCase())) {
            tagInput.value = "";
            return;
        }

        // Vérifier si le tag existe par défaut ou s'il est créé
        const isDefault = defaultTags.includes(cleanValue.toLowerCase());
        const tagType = isDefault ? "default" : "created";

        selectedTags.push({ name: cleanValue, type: tagType });
        updateTags();
        tagInput.value = ""; // Vide le champ
    }

    // Détecter la touche Entrée ou la sélection dans la liste
    tagInput.addEventListener("keydown", (e) => {
        if (e.key === "Enter") {
            e.preventDefault(); // Empêche de soumettre le formulaire entier
            addTag(tagInput.value);
        }
    });

    tagInput.addEventListener("input", () => {
        // Si l'utilisateur clique sur une option du datalist, l'ajout est instantané
        const options = Array.from(datalist.options).map(opt => opt.value);
        if (options.includes(tagInput.value)) {
            addTag(tagInput.value);
        }
    });

    // Supprimer un tag lors du clic sur le "x"
    tagsContainer.addEventListener("click", (e) => {
        if (e.target.classList.contains("tag-remove")) {
            const index = e.target.getAttribute("data-index");
            selectedTags.splice(index, 1);
            updateTags();
        }
    });
});