CREATE TABLE "days" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "date" DATETIME NOT NULL
);

CREATE UNIQUE INDEX "days_date_idx" ON "days"("date");