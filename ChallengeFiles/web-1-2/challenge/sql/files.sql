CREATE TABLE files (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  description TEXT NOT NULL,
  filename TEXT NOT NULL,
  owner TEXT NOT NULL,
  public INTEGER NOT NULL
);
