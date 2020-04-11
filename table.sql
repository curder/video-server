-- 用户表
CREATE TABLE IF NOT EXISTS users (
  id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
  login_name VARCHAR(255) UNIQUE NOT NULL,
  pwd TEXT NOT NULL
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 视频信息表
CREATE TABLE IF NOT EXISTS video_info (
  id VARCHAR(255) PRIMARY KEY NOT NULL,
  author_id INT,
  name TEXT,
  display_ctime TEXT,
  create_time DATETIME
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 评论表
CREATE TABLE IF NOT EXISTS comments (
  id VARCHAR(64) PRIMARY KEY NOT NULL,
  video_id VARCHAR(64),
  author_id INT,
  content TEXT,
  time DATETIME
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 会话表
CREATE TABLE IF NOT EXISTS sessions (
  session_id VARCHAR(255) PRIMARY KEY NOT NULL,
  TTL TINYTEXT,
  login_name VARCHAR(255)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;