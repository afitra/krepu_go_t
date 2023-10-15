DELIMITER //
CREATE PROCEDURE getByName(
    IN search_name VARCHAR(255),
    IN limit_count INT
)

BEGIN SELECT name, country FROM person
WHERE name LIKE CONCAT('%', search_name, '%')
    LIMIT limit_count;
END;
//
DELIMITER ;
