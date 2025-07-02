var query = @"
IF NOT EXISTS (
    SELECT 1
    FROM dbo.LiveLocationSummary WITH (NOLOCK)
    WHERE serialNumber = @serialNumber
      AND tripNumber = @tripNumber
      AND EventTime > @eventTime
)
BEGIN
    IF EXISTS (
        SELECT 1
        FROM dbo.LiveLocationSummary WITH (NOLOCK)
        WHERE serialNumber = @serialNumber
          AND tripNumber = @tripNumber
    )
    BEGIN
        UPDATE dbo.LiveLocationSummary
        SET Latitude = @latitude,
            Longitude = @longitude,
            EventTime = @eventTime,
            Status = @status
        WHERE serialNumber = @serialNumber
          AND tripNumber = @tripNumber;
    END
    ELSE
    BEGIN
        INSERT INTO dbo.LiveLocationSummary (serialNumber, tripNumber, Latitude, Longitude, EventTime, Status)
        VALUES (@serialNumber, @tripNumber, @latitude, @longitude, @eventTime, @status);
    END
END";