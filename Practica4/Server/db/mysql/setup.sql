/***CREATING ALL TABLES*/
CREATE TABLE Reporte(
        Carnet INT NOT NULL,
        Nombre VARCHAR(50) NULL,
        Curso VARCHAR(50) NULL,
        Fecha VARCHAR(100) NULL,
        Cuerpo VARCHAR(120) NULL
);

--   ENGINE = INNODB;

/* INSERT DATA */
INSERT INTO Reporte (Carnet, Nombre, Curso, Cuerpo)
VALUES (201588645, 'Tepokun', 'Redes 3', 'Ay no');
INSERT INTO Reporte (Carnet, Nombre, Curso, Cuerpo)
VALUES (201358845, 'Niu', 'Redes 3', 'Ay no');
INSERT INTO Reporte (Carnet, Nombre, Curso, Cuerpo)
VALUES (201555556, 'Macaco', 'Redes 3', 'Ay no');

-- DROP PROCEDURE IF EXISTS sp_GetEstudiante;
-- DELIMITER //
-- CREATE PROCEDURE sp_GetEstudiante()
--   BEGIN
--     SELECT * FROM DATOS_ESTUDIANTE;
--   END //
-- DELIMITER ;
-- /**Drop StoreProcedure**/
-- CALL sp_GetEstudiante();

Commit;
