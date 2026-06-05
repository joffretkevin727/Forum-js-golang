-- ============================================================
-- JEU DE DONNÉES DE TEST COHÉRENT
-- ============================================================

-- Tous les utilisateurs ci-dessous ont pour mot de passe : Password123!
INSERT INTO users (id, username, email, password_hash, is_banned)
VALUES (1, 'JeanSebastien', 'jean.seb@example.com', '$2a$10$wK9SBlVl7kR/y8W6n8fO/.wzCg7Fq6mEovqfbeX9LOnH4U2GepT7O', 0),
       (2, 'MarieCroissant', 'marie.c@example.com', '$2a$10$wK9SBlVl7kR/y8W6n8fO/.wzCg7Fq6mEovqfbeX9LOnH4U2GepT7O', 0),
       (3, 'PierreBoulanger', 'pierre.b@example.com', '$2a$10$wK9SBlVl7kR/y8W6n8fO/.wzCg7Fq6mEovqfbeX9LOnH4U2GepT7O',
        0),
       (4, 'TrollDuFour', 'troll.four@example.com', '$2a$10$wK9SBlVl7kR/y8W6n8fO/.wzCg7Fq6mEovqfbeX9LOnH4U2GepT7O', 1);
-- Utilisateur banni

-- Insertion des tags correspondants aux filtres de ton projet
INSERT INTO tags (id, name)
VALUES (1, 'Croissant'),
       (2, 'Cannelés'),
       (3, 'Gâteau Basque'),
       (4, 'Levain'),
       (5, 'Astuce');

-- Insertion des sujets de discussion (Topics)
INSERT INTO topics (id, title, body, status, author_id, pseudo, tags, like_count, dislike_count)
VALUES (1, 'Le secret du véritable Gâteau Basque',
        'Bonjour à tous ! Quelqu''un aurait la vraie recette traditionnelle du gâteau basque ? Notamment le secret pour que la crème pâtissière reste bien ferme après la cuisson au four.',
        'open', 1, 'JeanSebastien', '["Gâteau Basque", "Astuce"]', 2, 0),
       (2, 'Mon levain chef ne bulle plus, au secours !',
        'J''ai commencé un levain naturel il y a 4 jours. Il bullait bien au début mais depuis ce matin plus rien, une fine couche de liquide s''est formée sur le dessus. Est-il mort ?',
        'open', 2, 'MarieCroissant', '["Levain"]', 1, 1),
       (3, 'Réussir le feuilletage des croissants à coup sûr',
        'Après de nombreux essais ratés, j''ai enfin compris que la température du beurre de tournage doit être exactement la même que celle de la détrempe (environ 15°C). Voici mon guide complet...',
        'open', 3, 'PierreBoulanger', '["Croissant", "Astuce"]', 1, 0);

-- Association dans la table de liaison topic_tags
INSERT INTO topic_tags (topic_id, tag_id)
VALUES (1, 3), -- Topic 1 a le Tag 3 (Gâteau Basque)
       (1, 5), -- Topic 1 a le Tag 5 (Astuce)
       (2, 4), -- Topic 2 a le Tag 4 (Levain)
       (3, 1), -- Topic 3 a le Tag 1 (Croissant)
       (3, 5);
-- Topic 3 a le Tag 5 (Astuce)

-- Insertion des commentaires (Comments)
INSERT INTO comments (id, body, topic_id, author_id)
VALUES (1,
        'Il faut ajouter un peu de rhum ambré dans ta crème et surtout la laisser reposer une nuit entière au réfrigérateur avant de monter le gâteau !',
        1, 3),
       (2,
        'Le liquide au-dessus s''appelle du "hootch". C''est juste que ton levain a faim ! Enlève le liquide et nourris-le à nouveau avec de la farine de seigle.',
        2, 3),
       (3, 'N''importe quoi, le levain est mort, jette-le et achète de la levure chimique !', 2, 1),
       (4, 'Merci pour l''astuce des températures, mes croissants n''ont plus rien à voir !', 3, 2);

-- Simulation des likes et dislikes sur les topics
INSERT INTO liketopic (userid, topicid, `like`)
VALUES (2, 1, 1),  -- Marie aime le topic 1
       (3, 1, 1),  -- Pierre aime le topic 1
       (1, 2, -1), -- Jean met un dislike au topic 2
       (3, 2, 1),  -- Pierre met un like au topic 2
       (2, 3, 1);
-- Marie met un like au topic 3

-- Simulation des likes et dislikes sur les commentaires
INSERT INTO likecomment (userid, commentid, topicid, `like`)
VALUES (1, 1, 1, 1), -- Jean aime le commentaire de Pierre sur son topic
       (2, 2, 2, 1), -- Marie aime les conseils de Pierre sur le levain
       (3, 3, 2, -1); -- Pierre n''aime pas la fausse information de Jean