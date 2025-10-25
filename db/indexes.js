use database;

db.tasks.createIndex(
  { created_at: -1 },
  {
    name: "idx_created_at"
  }
);

db.tasks.createIndex(
  { status: 1 },
  {
    name: "idx_status"
  }
);

db.tasks.createIndex(
  { title: 1 },
  {
    name: "idx_tasks_title",
  }
);

db.tasks.createIndex(
  { status: 1, created_at: -1 },
  {
    name: "idx_tasks_status_createdAt_desc"
  }
);