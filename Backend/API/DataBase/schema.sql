-- ============================================================
-- SCHEMA SQL - Plateforme de forum (Version MySQL propre)
-- ============================================================

-- ========================
-- TABLE USERS
-- ========================
CREATE TABLE users (
    id            INT AUTO_INCREMENT PRIMARY KEY,
    username      VARCHAR(30)  NOT NULL UNIQUE,
    email         VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role          ENUM('user', 'admin', 'modo') NOT NULL DEFAULT 'user',
    is_banned     TINYINT(1) NOT NULL DEFAULT 0,
    created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

-- ========================
-- TABLE TAGS
-- ========================
CREATE TABLE tags (
    id   INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL UNIQUE
) ENGINE=InnoDB;

-- ========================
-- TABLE TOPICS
-- ========================
CREATE TABLE topics (
    id         INT AUTO_INCREMENT PRIMARY KEY,
    title      VARCHAR(100) NOT NULL,
    body       TEXT NOT NULL,
    status     ENUM('open', 'closed', 'archived') NOT NULL DEFAULT 'open',
    author_id  INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (author_id)
        REFERENCES users(id)
        ON DELETE CASCADE
) ENGINE=InnoDB;

-- ========================
-- TABLE TOPIC_TAGS
-- ========================
CREATE TABLE topic_tags (
    topic_id INT NOT NULL,
    tag_id   INT NOT NULL,

    PRIMARY KEY (topic_id, tag_id),

    FOREIGN KEY (topic_id)
        REFERENCES topics(id)
        ON DELETE CASCADE,

    FOREIGN KEY (tag_id)
        REFERENCES tags(id)
        ON DELETE CASCADE
) ENGINE=InnoDB;

-- ========================
-- TABLE MESSAGES
-- ========================
CREATE TABLE messages (
    id         INT AUTO_INCREMENT PRIMARY KEY,
    body       TEXT NOT NULL,
    topic_id   INT NOT NULL,
    author_id  INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (topic_id)
        REFERENCES topics(id)
        ON DELETE CASCADE,

    FOREIGN KEY (author_id)
        REFERENCES users(id)
        ON DELETE CASCADE
) ENGINE=InnoDB;

-- ========================
-- TABLE MESSAGE_VOTES
-- ========================
CREATE TABLE message_votes (
    user_id    INT NOT NULL,
    message_id INT NOT NULL,
    vote       TINYINT NOT NULL, -- +1 ou -1

    PRIMARY KEY (user_id, message_id),

    FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    FOREIGN KEY (message_id)
        REFERENCES messages(id)
        ON DELETE CASCADE
) ENGINE=InnoDB;

-- ENGINE=InnoDB sert a indiquer quel moteur se fais utiliser par Mysql