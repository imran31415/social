CREATE TABLE social_user
(
    id BIGINT auto_increment PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    ## implement schema later
    profile JSON
);

CREATE TABLE social_post
(
    id BIGINT auto_increment PRIMARY KEY,
    user_id BIGINT NOT NULL,
    FOREIGN KEY  (user_id) REFERENCES social_user(id),
    ## implement schema later
    content JSON NOT NULL
);

CREATE TABLE social_feed (
    id BIGINT auto_increment PRIMARY KEY,
    owner_id BIGINT NOT NULL,
    FOREIGN KEY  (owner_id) REFERENCES social_user(id),
    post_id BIGINT NOT NULL,
    FOREIGN KEY  (post_id) REFERENCES social_post(id)
);

CREATE TABLE social_comment (
    id BIGINT auto_increment PRIMARY KEY,
    post_id BIGINT NOT NULL,
    FOREIGN KEY  (post_id) REFERENCES social_post(id),
    parent_comment_id BIGINT NOT NULL DEFAULT 0,
    # implement schema later
    content JSON NOT NULL
);