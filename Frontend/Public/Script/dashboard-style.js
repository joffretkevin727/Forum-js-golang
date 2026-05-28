document.addEventListener("DOMContentLoaded", function() {
    // 1. On récupère l'URL actuelle (ex: reports.html?filter=messages)
    // On ne garde que le nom du fichier et les paramètres
    const currentUrl = window.location.pathname.split("/").pop() + window.location.search;

    // 2. On récupère tous les liens de la sidebar
    const menuLinks = document.querySelectorAll(".sidebar a");

    menuLinks.forEach(link => {
        // 3. On récupère la destination du lien (ex: reports.html?filter=messages)
        const linkHref = link.getAttribute("href");

        // 4. Si l'URL actuelle correspond exactement au lien, on ajoute 'active'
        if (currentUrl === linkHref) {
            link.classList.add("active");
        }
        
        // Optionnel : Gestion du Dashboard (si l'URL est vide ou juste le nom du fichier)
        if (currentUrl === "" || currentUrl === "dashboard.html") {
            if (linkHref === "dashboard.html") link.classList.add("active");
        }
    });
});