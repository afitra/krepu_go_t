DELIMITER //
CREATE PROCEDURE getDataWithLimit(
    IN limit_count INT
)

BEGIN SELECT id, name, country FROM person LIMIT limit_count;

END;
//
DELIMITER ;
