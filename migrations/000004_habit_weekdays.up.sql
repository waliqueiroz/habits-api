CREATE TABLE "habit_weekdays" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "habit_id" TEXT NOT NULL,
    "weekday" INTEGER NOT NULL,
    CONSTRAINT "habit_weekdays_habit_id_fkey" FOREIGN KEY ("habit_id") REFERENCES "habits" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE UNIQUE INDEX "habit_weekdays_habit_id_weekday_key" ON "habit_weekdays"("habit_id", "weekday");