SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

DROP TABLE IF EXISTS `user_coupon`;



create table user_coupon(
    id BIGINT(11) PRIMARY KEY AUTO_INCREMENT,
    user_id int(11) not null,
    coupon_id int(11) not null,
    used int(11) not null,
    create_time datetime DEFAULT current_timestamp,
    update_time datetime DEFAULT current_timestamp,
    UNIQUE key(user_id, coupon_id),
    key(coupon_id)
);