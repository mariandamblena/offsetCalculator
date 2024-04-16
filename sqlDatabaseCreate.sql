-- Crear tabla TipoSensor
CREATE TABLE TipoSensor (
    id INT PRIMARY KEY,
    Descripcion VARCHAR(255)
);
go
-- Crear tabla Sensor
CREATE TABLE Sensor (
    SerialNumber VARCHAR(50) PRIMARY KEY,
    TipoSensorId INT,
    FOREIGN KEY (TipoSensorId) REFERENCES TipoSensor(id)
);
go
-- Crear tabla Configuracion
CREATE TABLE Configuracion (
    SensorSerialNumber VARCHAR(50) PRIMARY KEY,
    TempCompensation BIT,
    HumidityCompensation BIT,
    PressureCompensation BIT,
    ModbusAddress INT,
    BaudRate INT,
    StopBits INT,
    CalibrationDate DATE,
    PurchaseLink VARCHAR(255),
    CsvConfigExtendedLink VARCHAR(255),
    FOREIGN KEY (SensorSerialNumber) REFERENCES Sensor(SerialNumber)
);
go
-- Crear tabla Datalog
CREATE TABLE Datalog (
    Timestamp DATETIME,
    Value FLOAT,
    SensorSerialNumber VARCHAR(50),
    FOREIGN KEY (SensorSerialNumber) REFERENCES Sensor(SerialNumber)
);
go