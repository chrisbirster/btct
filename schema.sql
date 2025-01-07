CREATE TABLE tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		description TEXT NOT NULL,
		complete BOOLEAN NOT NULL DEFAULT 0
	, owner_id TEXT NOT NULL DEFAULT 'default_user');
CREATE TABLE sqlite_sequence(name,seq);
