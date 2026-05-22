-- ============================================================
-- JEU DE DONNÉES DE TEST (INSERTS)
-- ============================================================

-- 1. Insertion des utilisateurs (Les mots de passe fictifs devront correspondre à tes tests en Go)
INSERT INTO users (username, email, password_hash, role, is_banned) VALUES
('lucas', 'lucas@exemple.com', '$2a$10$Zm9ydW1nb2xhbmd0ZXN0cGFzc3dk', 'admin', 0),
('alice', 'alice@exemple.com', '$2a$10$Zm9ydW1nb2xhbmd0ZXN0cGFzc3dk', 'user', 0),
('bob_le_modo', 'bob@exemple.com', '$2a$10$Zm9ydW1nb2xhbmd0ZXN0cGFzc3dk', 'modo', 0),
('troll42', 'troll@exemple.com', '$2a$10$Zm9ydW1nb2xhbmd0ZXN0cGFzc3dk', 'user', 1); -- Utilisateur banni

-- 2. Insertion des tags
INSERT INTO tags (name) VALUES
('Golang'),
('MySQL'),
('MAMP'),
('Aide'),
('Discussions');

-- 3. Insertion des sujets (Topics)
INSERT INTO topics (title, body, status, author_id) VALUES
('Bienvenue sur le forum !', 'Ceci est le tout premier sujet pour saluer la communauté.', 'open', 1), -- Créé par lucas
('Problème de connexion MAMP et Go', 'Bonjour, mon API Go ne trouve pas MySQL sur le port 8889. Une idée ?', 'open', 2), -- Créé par alice
('Règles de modération', 'Merci de rester courtois sur ce forum de discussion.', 'archived', 3); -- Créé par bob_le_modo

-- 4. Liaison des tags aux sujets (Topic_Tags)
INSERT INTO topic_tags (topic_id, tag_id) VALUES
(1, 5), -- Sujet 1 a le tag 'Discussions'
(2, 1), -- Sujet 2 a le tag 'Golang'
(2, 2), -- Sujet 2 a le tag 'MySQL'
(2, 3), -- Sujet 2 a le tag 'MAMP'
(2, 4), -- Sujet 2 a le tag 'Aide'
(3, 5); -- Sujet 3 a le tag 'Discussions'

-- 5. Insertion des messages (Réponses)
INSERT INTO messages (body, topic_id, author_id) VALUES
('Super ! Content de voir ce projet démarrer.', 1, 2), -- Alice répond sur le Sujet 1
('Regarde si ton MAMP n''utilise pas plutôt le port 3306 dans les préférences.', 2, 1), -- Lucas répond sur le Sujet 2
('Ah merci Lucas ! C''était exactement ça, ça marche.', 2, 2); -- Alice répond sur le Sujet 2

-- 6. Insertion des votes sur les messages (Message_Votes)
INSERT INTO message_votes (user_id, message_id, vote) VALUES
(2, 2, 1),  -- Alice met un +1 au message de Lucas (id: 2)
(3, 2, 1);  -- Bob met aussi un +1 au message de Lucas (id: 2)