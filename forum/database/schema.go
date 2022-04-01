package database

const (
	createUserTable = `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		nickname TEXT,
		email TEXT,
		created_at TEXT,
		password TEXT,
		firstname TEXT,
		lastname TEXT
	);`
	createSessionsTable = `CREATE TABLE IF NOT EXISTS sessions (
		id INTEGER PRIMARY KEY,
		uuid TEXT,
		user_id INTEGER,
		created_at TEXT
	);`
	createTagsTable = `CREATE TABLE IF NOT EXISTS tags (
		id INTEGER PRIMARY KEY,
		name TEXT
	);`
	createPostsTable = `CREATE TABLE IF NOT EXISTS posts (
		id INTEGER PRIMARY KEY,
		content TEXT,
		user_id INTEGER,
		created_at TEXT
	);`
	createCommentsTable = `CREATE TABLE IF NOT EXISTS comments (
		id INTEGER PRIMARY KEY,
		content TEXT,
		user_id INTEGER,
		created_at TEXT,
		post_id INTEGER
	);`
	createQuestionsTable = `CREATE TABLE IF NOT EXISTS questions (
		id INTEGER PRIMARY KEY,
		post_id INTEGER,
		title TEXT,
		views INTEGER
	);`
	createLikesTable = `CREATE TABLE IF NOT EXISTS likes (
		id INTEGER PRIMARY KEY,
		user_id INTEGER,
		post_id INTEGER,
		liked INTEGER
	);`
	createTagsQuestionsTable = `CREATE TABLE IF NOT EXISTS tags_questions (
		id INTEGER PRIMARY KEY,
		tag_id INTEGER,
		question_id INTEGER
	);`
	createQuestionsAnswersTable = `CREATE TABLE IF NOT EXISTS questions_answers (
		id INTEGER PRIMARY KEY,
		post_id INTEGER,
		question_id INTEGER
	);`
)
