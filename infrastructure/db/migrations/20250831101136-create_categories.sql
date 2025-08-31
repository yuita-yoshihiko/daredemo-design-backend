
-- +migrate Up
CREATE TABLE IF NOT EXISTS categories (
  id BIGSERIAL,
  name TEXT NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  PRIMARY KEY (id),
  UNIQUE (name)
);

comment on table categories is 'デザイン情報カテゴリー';
comment on column categories.name is 'カテゴリー名';

-- +migrate Down
DROP TABLE IF EXISTS categories;
