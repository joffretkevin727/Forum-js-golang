-- ============================================================
-- SCHEMA SQL - Plateforme de forum
-- ============================================================

-- Table des utilisateurs (F-1, F-2)
CREATE TABLE users (
    id            SERIAL PRIMARY KEY,
    username      VARCHAR(30)  NOT NULL UNIQUE,
    email         VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(100) NOT NULL,
    role          VARCHAR(20)  NOT NULL DEFAULT 'user'
                  CHECK (role IN ('user', 'admin')),
    is_banned     BOOLEAN      NOT NULL DEFAULT FALSE,
    created_at    TIMESTAMP    NOT NULL DEFAULT NOW()
);

-- Table des tags / catégories (F-4, F-10)
CREATE TABLE tags (
    id   SERIAL PRIMARY KEY,
    name VARCHAR(30) NOT NULL UNIQUE
);

-- Table des topics (F-4)
CREATE TABLE topics (
    id         SERIAL PRIMARY KEY,
    title      VARCHAR(100) NOT NULL,
    body       VARCHAR(100) NOT NULL,
    status     VARCHAR(10)  NOT NULL DEFAULT 'open'
               CHECK (status IN ('open', 'closed', 'archived')),
    author_id  INTEGER      NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    -- Définition des clés étrangères
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Table de liaison topics <-> tags (F-4, F-10)
CREATE TABLE topic_tags (
    topic_id INTEGER NOT NULL,
    tag_id   INTEGER NOT NULL,
    PRIMARY KEY (topic_id, tag_id),
    -- Définition des clés étrangères
    FOREIGN KEY (topic_id) REFERENCES topics(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

-- Table des messages (F-5)
CREATE TABLE messages (
    id         SERIAL PRIMARY KEY,
    body       VARCHAR(500) NOT NULL,
    topic_id   INTEGER      NOT NULL,
    author_id  INTEGER      NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT NOW(),
    -- Définition des clés étrangères
    FOREIGN KEY (topic_id) REFERENCES topics(id) ON DELETE CASCADE,
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Table des likes / dislikes sur les messages (F-7)
CREATE TABLE message_votes (
    user_id    INTEGER  NOT NULL,
    message_id INTEGER  NOT NULL,
    vote       SMALLINT NOT NULL CHECK (vote IN (1, -1)),
    PRIMARY KEY (user_id, message_id),
    -- Définition des clés étrangères
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE
);

-- ============================================================
-- INDEX pour les performances
-- ============================================================

CREATE INDEX idx_topics_title ON topics (title);
CREATE INDEX idx_messages_topic_date ON messages (topic_id, created_at DESC);
CREATE INDEX idx_votes_message ON message_votes (message_id);
CREATE INDEX idx_topic_tags_tag ON topic_tags (tag_id);

-- ============================================================
-- VUE utilitaire : score de popularité des messages
-- ============================================================
CREATE VIEW message_scores AS
SELECT
    m.id AS message_id,
    m.topic_id,
    m.author_id,
    m.body,
    m.created_at,
    COALESCE(SUM(v.vote), 0) AS score
FROM messages m
LEFT JOIN message_votes v ON v.message_id = m.id
GROUP BY m.id, m.topic_id, m.author_id, m.body, m.created_at;