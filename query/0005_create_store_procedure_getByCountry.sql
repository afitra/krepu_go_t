DELIMITER //
CREATE PROCEDURE getByCountry(
    IN search_country VARCHAR(255),
    IN limit_count INT
)

BEGIN SELECT name, country FROM person
WHERE country LIKE CONCAT('%', search_country, '%');
END;
//
DELIMITER ;
