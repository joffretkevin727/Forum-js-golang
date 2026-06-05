-- ============================================================
-- SCHEMA SQL - Plateforme de forum (Version MySQL propre)
-- ============================================================

CREATE
DATABASE forum_db;
USE
forum_db;

-- ========================
-- TABLE USERS
-- ========================
CREATE TABLE users
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    username      VARCHAR(30)  NOT NULL UNIQUE,
    email         VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    is_banned     TINYINT(1) NOT NULL DEFAULT 0,
    created_at    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB;

-- ========================
-- TABLE TAGS
-- ========================
CREATE TABLE tags
(
    id   INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(30) NOT NULL UNIQUE
) ENGINE=InnoDB;

-- ========================
-- TABLE TOPICS
-- ========================
CREATE TABLE topics
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    title         VARCHAR(100) NOT NULL,
    body          TEXT         NOT NULL,
    status        ENUM('open', 'closed', 'archived') NOT NULL DEFAULT 'open',
    author_id     INT          NOT NULL,
    created_at    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    pseudo        VARCHAR(50)  NOT NULL DEFAULT 'Anonyme',
    tags          JSON NULL,
    like_count    INT          NOT NULL DEFAULT 0,
    dislike_count INT          NOT NULL DEFAULT 0,
    isLike        BOOLEAN               DEFAULT FALSE,

    FOREIGN KEY (author_id)
        REFERENCES users (id)
        ON DELETE CASCADE
) ENGINE=InnoDB;

-- ========================
-- TABLE TOPIC_TAGS
-- ========================
CREATE TABLE topic_tags
(
    topic_id INT NOT NULL,
    tag_id   INT NOT NULL,

    PRIMARY KEY (topic_id, tag_id),

    FOREIGN KEY (topic_id)
        REFERENCES topics (id)
        ON DELETE CASCADE,

    FOREIGN KEY (tag_id)
        REFERENCES tags (id)
        ON DELETE CASCADE
) ENGINE=InnoDB;

-- ========================
-- TABLE COMMENTS
-- ========================
CREATE TABLE comments
(
    id         INT AUTO_INCREMENT PRIMARY KEY,
    body       TEXT      NOT NULL,
    topic_id   INT       NOT NULL,
    author_id  INT       NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (topic_id)
        REFERENCES topics (id)
        ON DELETE CASCADE,

    FOREIGN KEY (author_id)
        REFERENCES users (id)
        ON DELETE CASCADE
) ENGINE=InnoDB;

-- ========================
-- TABLE LIKETOPIC
-- ========================
CREATE TABLE liketopic
(
    id      INT AUTO_INCREMENT,
    userid  INT     NOT NULL,
    topicid INT     NOT NULL,
    `like`  TINYINT NOT NULL CHECK (`like` IN (-1, 0, 1)),

    PRIMARY KEY (userid, topicid),
    KEY     id_idx (id),

    FOREIGN KEY (userid)
        REFERENCES users (id)
        ON DELETE CASCADE,

    FOREIGN KEY (topicid)
        REFERENCES topics (id)
        ON DELETE CASCADE
) ENGINE=InnoDB;

-- ========================
-- TABLE LIKECOMMENT
-- ========================
CREATE TABLE likecomment
(
    id        INT AUTO_INCREMENT,
    userid    INT     NOT NULL,
    commentid INT     NOT NULL,
    topicid   INT     NOT NULL,
    `like`    TINYINT NOT NULL CHECK (`like` IN (-1, 0, 1)),

    PRIMARY KEY (userid, commentid),
    KEY       id_idx (id),

    FOREIGN KEY (userid)
        REFERENCES users (id)
        ON DELETE CASCADE,

    FOREIGN KEY (commentid)
        REFERENCES comments (id)
        ON DELETE CASCADE,

    FOREIGN KEY (topicid)
        REFERENCES topics (id)
        ON DELETE CASCADE
) ENGINE=InnoDB;
