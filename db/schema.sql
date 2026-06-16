CREATE TABLE tasks (
  id   BIGINT  NOT NULL AUTO_INCREMENT PRIMARY KEY,
  title varchar(255)    NOT NULL,
  description text,
  status ENUM('todo', 'in-progress', 'done') DEFAULT 'todo' NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  completed_at TIMESTAMP DEFAULT NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
